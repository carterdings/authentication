package logic

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	mrand "math/rand"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/carterdings/authentication/entity"
	"github.com/carterdings/authentication/repo/dao"
	"github.com/carterdings/authentication/repo/errs"
	"github.com/carterdings/authentication/repo/log"
	"github.com/patrickmn/go-cache"
)

// AuthService service interface
type AuthService interface {
	CreateUser(req *entity.UserReq) error
	DeleteUser(req *entity.UserReq) error
	CreateRole(req *entity.RoleReq) error
	DeleteRole(req *entity.RoleReq) error
	AddRoleToUser(req *entity.UserRoleReq) error
	Authenticate(req *entity.UserRoleReq) (*entity.UserRoleRsp, error)
	Invalidate(req *entity.UserRoleReq) error
	CheckRole(req *entity.UserRoleReq) (*entity.UserRoleRsp, error)
	AllRoles(req *entity.UserRoleReq) (*entity.UserRoleRsp, error)
}

// NewServie new service
func NewServie(privKey, pubKey []byte) AuthService {
	return &service{
		privKey: privKey,
		pubKey:  pubKey,
	}
}

// service service
type service struct {
	privKey []byte
	pubKey  []byte

	userLock sync.Mutex
	roleLock sync.Mutex
	bindLock sync.Mutex
}

// CreateUser create user
func (s *service) CreateUser(req *entity.UserReq) error {
	salt := make([]byte, entity.SaltLen)
	randBytes(salt)

	pwd := append([]byte(req.Password), salt...)
	user := &entity.User{
		UserName:   req.UserName,
		Salt:       salt,
		Password:   hashData(pwd),
		CreateTime: time.Now().Unix(),
	}
	s.userLock.Lock()
	_, ok := dao.Get(user.Key())
	if ok {
		// user exists
		s.userLock.Unlock()
		log.Errorf("User %s exists", user.Key())
		return errs.New(entity.ErrCodeUserExists, "User exists")
	}
	// create user
	dao.Set(user.Key(), user, cache.NoExpiration)
	// free lock
	s.userLock.Unlock()
	return nil
}

// DeleteUser delete user
func (s *service) DeleteUser(req *entity.UserReq) error {
	user := &entity.User{
		UserName: req.UserName,
	}
	s.userLock.Lock()
	_, ok := dao.Get(user.Key())
	if !ok {
		// user not exists
		s.userLock.Unlock()
		log.Errorf("User %s not exist", user.Key())
		return errs.New(entity.ErrCodeUserNotExist, "User not exist")
	}
	dao.Delete(user.Key())
	// free lock
	s.userLock.Unlock()
	return nil
}

// CreateRole create role
func (s *service) CreateRole(req *entity.RoleReq) error {
	role := &entity.Role{
		RoleName:   req.RoleName,
		CreateTime: time.Now().Unix(),
	}
	s.roleLock.Lock()
	_, ok := dao.Get(role.Key())
	if ok {
		// role exists
		s.roleLock.Unlock()
		log.Errorf("Role %s exists", role.Key())
		return errs.New(entity.ErrCodeRoleExists, "Role exists")
	}
	// create user
	dao.Set(role.Key(), role, cache.NoExpiration)
	// free lock
	s.roleLock.Unlock()
	return nil
}

// DeleteRole delete role
func (s *service) DeleteRole(req *entity.RoleReq) error {
	role := &entity.Role{
		RoleName: req.RoleName,
	}
	s.roleLock.Lock()
	_, ok := dao.Get(role.Key())
	if !ok {
		// role not exist
		s.roleLock.Unlock()
		log.Errorf("Role %s not exist", role.Key())
		return errs.New(entity.ErrCodeRoleNotExist, "Role not exist")
	}
	dao.Delete(role.Key())
	// free lock
	s.roleLock.Unlock()
	return nil
}

// AddRoleToUser add role to user
func (s *service) AddRoleToUser(req *entity.UserRoleReq) error {
	user := &entity.User{
		UserName: req.UserName,
	}
	_, ok := dao.Get(user.Key())
	if !ok {
		// user not exists
		log.Errorf("User %s not exist", user.Key())
		return errs.New(entity.ErrCodeUserNotExist, "User not exist")
	}
	role := &entity.Role{
		RoleName: req.RoleName,
	}
	_, ok = dao.Get(role.Key())
	if !ok {
		// role not exist
		log.Errorf("Role %s not exist", role.Key())
		return errs.New(entity.ErrCodeRoleNotExist, "Role not exist")
	}

	roles := make(map[string]bool)
	s.bindLock.Lock()
	ur, ok := dao.Get(user.UserRolesKey())
	if ok {
		// find roles
		roles = ur.(map[string]bool)
	}
	roles[req.RoleName] = true
	dao.Set(user.UserRolesKey(), roles, cache.NoExpiration)
	s.bindLock.Unlock()

	dao.Set(req.Key(), nil, cache.NoExpiration)
	return nil
}

// Authenticate authenticate
func (s *service) Authenticate(req *entity.UserRoleReq) (*entity.UserRoleRsp, error) {
	rsp := &entity.UserRoleRsp{}
	user := &entity.User{
		UserName: req.UserName,
	}
	// query user
	u, ok := dao.Get(user.Key())
	if !ok {
		return rsp, errs.New(entity.ErrCodeUserNotExist, "User not exist")
	}
	usr, ok := u.(*entity.User)
	if !ok {
		return rsp, errs.New(entity.ErrCodeUserNotExist, "User not exist")
	}
	pwd := append([]byte(req.Password), usr.Salt...)
	// check password
	digest := hashData(pwd)
	if !bytes.Equal(digest, usr.Password) {
		return rsp, errs.New(entity.ErrCodeInvalidPassword, "Invalid password")
	}
	buf := make([]byte, 16)
	randBytes(buf)
	random := base64.RawURLEncoding.EncodeToString(buf)
	// generate token
	tok := &entity.Token{
		UserName:  req.UserName,
		Password:  digest,
		Rand:      random,
		Timestamp: time.Now().Unix(),
		Expire:    entity.TokenExpire,
	}
	token, err := genToken(tok, s.privKey)
	if err != nil {
		return rsp, errs.Newf(entity.ErrCodeGenToken, "generate token: %v", err)
	}
	rsp.Token = token
	dao.Set(token, nil, time.Duration(entity.TokenExpire)*time.Second)
	return rsp, nil
}

func (s *service) checkToken(req *entity.UserRoleReq) error {
	user := &entity.User{
		UserName: req.UserName,
	}
	// query user
	u, ok := dao.Get(user.Key())
	if !ok {
		return errs.New(entity.ErrCodeUserNotExist, "User not exist")
	}
	usr, ok := u.(*entity.User)
	if !ok {
		return errs.New(entity.ErrCodeUserNotExist, "User not exist")
	}
	// check token
	err := checkToken(req.Token, usr, s.pubKey)
	if err != nil {
		return err
	}
	// check token if invalidated
	_, ok = dao.Get(req.Token)
	if !ok {
		return errs.New(entity.ErrCodeInvalidToken, "Invalid token: invalidated")
	}
	return nil
}

// Invalidate invalidate
func (s *service) Invalidate(req *entity.UserRoleReq) error {
	err := s.checkToken(req)
	if err != nil {
		return err
	}
	// invalidate
	dao.Delete(req.Token)
	return nil
}

// CheckRole check role
func (s *service) CheckRole(req *entity.UserRoleReq) (*entity.UserRoleRsp, error) {
	rsp := &entity.UserRoleRsp{}
	err := s.checkToken(req)
	if err != nil {
		return rsp, err
	}
	_, ok := dao.Get(req.Key())
	if !ok {
		// role not match user
		return rsp, nil
	}
	rsp.CheckResult = true
	return rsp, nil
}

// AllRoles all roles
func (s *service) AllRoles(req *entity.UserRoleReq) (*entity.UserRoleRsp, error) {
	rsp := &entity.UserRoleRsp{}
	err := s.checkToken(req)
	if err != nil {
		return rsp, err
	}
	user := &entity.User{
		UserName: req.UserName,
	}
	ur, ok := dao.Get(user.UserRolesKey())
	if !ok {
		// no roles
		return rsp, nil
	}
	roles, ok := ur.(map[string]bool)
	if !ok {
		return rsp, nil
	}
	rlist := make([]string, 0)
	for k, _ := range roles {
		rlist = append(rlist, k)
	}
	rsp.Roles = rlist
	return rsp, nil
}

func genToken(tok *entity.Token, privKey []byte) (string, error) {
	data := append([]byte(fmt.Sprintf("u=%s&r=%s&t=%d&e=%d", tok.UserName, tok.Rand, tok.Timestamp, tok.Expire)),
		tok.Password...)
	sign, err := RsaSign(data, privKey)
	if err != nil {
		return "", err
	}
	token := base64.RawURLEncoding.EncodeToString(sign)
	v := url.Values{}
	v.Set("user", tok.UserName)
	v.Set("rand", tok.Rand)
	v.Set("ts", strconv.FormatInt(tok.Timestamp, 10))
	v.Set("expire", strconv.FormatInt(tok.Expire, 10))
	v.Set("token", token)
	return v.Encode(), nil
}

func checkToken(token string, user *entity.User, pubKey []byte) error {
	v, err := url.ParseQuery(token)
	if err != nil {
		return errs.Newf(entity.ErrCodeInvalidToken, "Invalid token: %v", err)
	}
	ts, _ := strconv.ParseInt(v.Get("ts"), 10, 64)
	expire, _ := strconv.ParseInt(v.Get("expire"), 10, 64)
	now := time.Now().Unix()
	if now-ts > expire {
		return errs.New(entity.ErrCodeExpiredToken, "Expired token")
	}
	tok := &entity.Token{
		UserName:  v.Get("user"),
		Password:  user.Password,
		Rand:      v.Get("rand"),
		Timestamp: ts,
		Expire:    expire,
	}
	data := append([]byte(fmt.Sprintf("u=%s&r=%s&t=%d&e=%d", tok.UserName, tok.Rand, tok.Timestamp, tok.Expire)),
		tok.Password...)
	tokenStr := v.Get("token")
	sign, err := base64.RawURLEncoding.DecodeString(tokenStr)
	if err != nil {
		return errs.Newf(entity.ErrCodeInvalidToken, "Invalid token: %v", err)
	}
	err = RsaVerify(data, sign, pubKey)
	if err != nil {
		return errs.Newf(entity.ErrCodeInvalidToken, "Invalid token: %v", err)
	}
	return nil
}

func randBytes(buf []byte) {
	_, err := rand.Read(buf)
	if err != nil {
		// It always returns len(p) and a nil error.
		mrand.Read(buf)
	}
}
