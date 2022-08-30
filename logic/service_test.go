package logic

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey"
	"github.com/carterdings/authentication/entity"
	"github.com/carterdings/authentication/repo/dao"
	"github.com/carterdings/authentication/repo/errs"
	"github.com/stretchr/testify/assert"
)

// Test_AuthService_CreateUser ...
func Test_AuthService_CreateUser(t *testing.T) {
	patches := patchDaoGet(nil)
	patches = patchDaoSet(patches)
	patches = patchDaoDelete(patches)
	defer patches.Reset()

	privKey, pubKey, err := GenRsaKey()
	assert.Nil(t, err)
	s := NewServie(privKey, pubKey)

	tests := []struct {
		name string
		req  *entity.UserReq
		err  error
	}{
		{"test_1", &entity.UserReq{
			UserName: "carter",
			Password: "pwd",
		}, nil},
		{"test_2", &entity.UserReq{
			UserName: "cat",
			Password: "pwd",
		}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.CreateUser(tt.req)
			assert.Equal(t, tt.err, err)
		})
	}
}

// Test_AuthService_DeleteUser ..
func Test_AuthService_DeleteUser(t *testing.T) {
	patches := patchDaoGet(nil)
	patches = patchDaoSet(patches)
	patches = patchDaoDelete(patches)
	defer patches.Reset()

	privKey, pubKey, err := GenRsaKey()
	assert.Nil(t, err)
	s := NewServie(privKey, pubKey)

	tests := []struct {
		name string
		req  *entity.UserReq
		err  error
	}{
		{"test_1", &entity.UserReq{
			UserName: "carter",
			Password: "pwd",
		}, nil},
		{"test_2", &entity.UserReq{
			UserName: "cat",
			Password: "pwd",
		}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteUser(tt.req)
			assert.Equal(t, tt.err, err)
		})
	}
}

// Test_AuthService_CreateRole ...
func Test_AuthService_CreateRole(t *testing.T) {
	patches := patchDaoGet(nil)
	patches = patchDaoSet(patches)
	patches = patchDaoDelete(patches)
	defer patches.Reset()

	privKey, pubKey, err := GenRsaKey()
	assert.Nil(t, err)
	s := NewServie(privKey, pubKey)

	tests := []struct {
		name string
		req  *entity.RoleReq
		err  error
	}{
		{"test_1", &entity.RoleReq{
			RoleName: "root",
		}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.CreateRole(tt.req)
			assert.Equal(t, tt.err, err)
		})
	}
}

// Test_AuthService_DeleteRole ...
func Test_AuthService_DeleteRole(t *testing.T) {
	patches := patchDaoGet(nil)
	patches = patchDaoSet(patches)
	patches = patchDaoDelete(patches)
	defer patches.Reset()

	privKey, pubKey, err := GenRsaKey()
	assert.Nil(t, err)
	s := NewServie(privKey, pubKey)

	tests := []struct {
		name string
		req  *entity.RoleReq
		err  error
	}{
		{"test_1", &entity.RoleReq{
			RoleName: "root",
		}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteRole(tt.req)
			assert.Equal(t, tt.err, err)
		})
	}
}

// Test_AuthService_AddRoleToUser ...
func Test_AuthService_AddRoleToUser(t *testing.T) {
	patches := patchDaoGet(nil)
	patches = patchDaoSet(patches)
	patches = patchDaoDelete(patches)
	defer patches.Reset()

	privKey, pubKey, err := GenRsaKey()
	assert.Nil(t, err)
	s := NewServie(privKey, pubKey)

	tests := []struct {
		name string
		req  *entity.UserRoleReq
		err  error
	}{
		{"test_1", &entity.UserRoleReq{
			UserName: "cat",
			Password: "pwd",
			RoleName: "root",
			Token:    "token",
		},
			errs.New(entity.ErrCodeUserNotExist, "User not exist")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.AddRoleToUser(tt.req)
			assert.Equal(t, tt.err, err)
		})
	}
}

// Test_AuthService_Authenticate ...
func Test_AuthService_Authenticate(t *testing.T) {
	patches := patchDaoGet(nil)
	patches = patchDaoSet(patches)
	patches = patchDaoDelete(patches)
	defer patches.Reset()

	privKey, pubKey, err := GenRsaKey()
	assert.Nil(t, err)
	s := NewServie(privKey, pubKey)

	tests := []struct {
		name string
		req  *entity.UserRoleReq
		err  error
	}{
		{"test_1", &entity.UserRoleReq{
			UserName: "cat",
			Password: "pwd",
			RoleName: "root",
			Token:    "token",
		},
			errs.New(entity.ErrCodeUserNotExist, "User not exist")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := s.Authenticate(tt.req)
			assert.Equal(t, tt.err, err)
			t.Logf("Authenticate req: %+v, rsp: %+v", tt.req, rsp)
		})
	}
}

// Test_AuthService_Invalidate ...
func Test_AuthService_Invalidate(t *testing.T) {
	patches := patchDaoGet(nil)
	patches = patchDaoSet(patches)
	patches = patchDaoDelete(patches)
	defer patches.Reset()

	privKey, pubKey, err := GenRsaKey()
	assert.Nil(t, err)
	s := NewServie(privKey, pubKey)

	tests := []struct {
		name string
		req  *entity.UserRoleReq
		err  error
	}{
		{"test_1", &entity.UserRoleReq{
			UserName: "cat",
			Password: "pwd",
			RoleName: "root",
			Token:    "token",
		},
			errs.New(entity.ErrCodeUserNotExist, "User not exist")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.Invalidate(tt.req)
			assert.Equal(t, tt.err, err)
		})
	}
}

// Test_AuthService_CheckRole ...
func Test_AuthService_CheckRole(t *testing.T) {
	patches := patchDaoGet(nil)
	patches = patchDaoSet(patches)
	patches = patchDaoDelete(patches)
	defer patches.Reset()

	privKey, pubKey, err := GenRsaKey()
	assert.Nil(t, err)
	s := NewServie(privKey, pubKey)

	tests := []struct {
		name string
		req  *entity.UserRoleReq
		err  error
	}{
		{"test_1", &entity.UserRoleReq{
			UserName: "cat",
			Password: "pwd",
			RoleName: "root",
			Token:    "token",
		},
			errs.New(entity.ErrCodeUserNotExist, "User not exist")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := s.CheckRole(tt.req)
			assert.Equal(t, tt.err, err)
			t.Logf("CheckRole req: %+v, rsp: %+v", tt.req, rsp)
		})
	}
}

// Test_AuthService_AllRoles ...
func Test_AuthService_AllRoles(t *testing.T) {
	patches := patchDaoGet(nil)
	patches = patchDaoSet(patches)
	patches = patchDaoDelete(patches)
	defer patches.Reset()

	privKey, pubKey, err := GenRsaKey()
	assert.Nil(t, err)
	s := NewServie(privKey, pubKey)

	tests := []struct {
		name string
		req  *entity.UserRoleReq
		err  error
	}{
		{"test_1", &entity.UserRoleReq{
			UserName: "cat",
			Password: "pwd",
			RoleName: "root",
			Token:    "token",
		},
			errs.New(entity.ErrCodeUserNotExist, "User not exist")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := s.AllRoles(tt.req)
			assert.Equal(t, tt.err, err)
			t.Logf("AllRoles req: %+v, rsp: %+v", tt.req, rsp)
		})
	}
}

// Test_Token ...
func Test_Token(t *testing.T) {
	privKey, pubKey, err := GenRsaKey()
	assert.Nil(t, err)

	tests := []struct {
		name string
		tk   *entity.Token
		err  error
	}{
		{"test_1", &entity.Token{
			UserName:  "cat",
			Password:  []byte("pwd"),
			Rand:      "rand",
			Timestamp: 123456,
			Expire:    7200,
		}, errs.New(entity.ErrCodeExpiredToken, "Expired token")},
		{"test_2", &entity.Token{
			UserName:  "cat",
			Password:  []byte("pwd"),
			Rand:      "rand",
			Timestamp: time.Now().Unix(),
			Expire:    7200,
		}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := genToken(tt.tk, privKey)
			assert.Nil(t, err)
			t.Logf("token: %s", token)

			user := &entity.User{
				Password: tt.tk.Password,
			}
			err = checkToken(token, user, pubKey)
			assert.Equal(t, tt.err, err)
		})
	}

}

func patchDaoGet(patches *gomonkey.Patches) *gomonkey.Patches {
	if patches == nil {
		patches = gomonkey.NewPatches()
	}
	return patches.ApplyFunc(dao.Get, func(key string) (interface{}, bool) {
		fmt.Println("key: ", key)
		switch {
		case key == "user_cat":
			return &entity.User{
				UserName: "cat",
			}, true
		case strings.HasPrefix(key, "user_roles_"):
			return map[string]bool{"root": true}, true
		case strings.HasPrefix(key, "role_root"):
			return &entity.Role{RoleName: "root"}, true
		}
		return nil, false
	})
}

func patchDaoSet(patches *gomonkey.Patches) *gomonkey.Patches {
	if patches == nil {
		patches = gomonkey.NewPatches()
	}
	return patches.ApplyFunc(dao.Set, func(key string, val interface{}, d time.Duration) {})
}

func patchDaoDelete(patches *gomonkey.Patches) *gomonkey.Patches {
	if patches == nil {
		patches = gomonkey.NewPatches()
	}
	return patches.ApplyFunc(dao.Delete, func(key string) {})
}
