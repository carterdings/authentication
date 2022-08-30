package main

import (
	"net/http"

	"github.com/carterdings/authentication/entity"
	"github.com/carterdings/authentication/logic"
	"github.com/carterdings/authentication/repo/errs"
	"github.com/gin-gonic/gin"
)

// Service service
type Service struct {
	logic.AuthService
}

// CreateUser create user handler
func (s *Service) CreateUser(c *gin.Context) {
	req := &entity.UserReq{}
	rsp := &entity.CommRsp{}
	var err error

	defer func() {
		if err != nil {
			rsp.Code = errs.ErrCode(err)
			rsp.Msg = errs.ErrMsg(err)
		}
		c.PureJSON(http.StatusOK, gin.H{
			"code": rsp.Code,
			"msg":  rsp.Msg,
		})
	}()

	err = c.BindJSON(req)
	if err != nil {
		err = errs.New(entity.ErrCodeInvalidParam, err.Error())
		return
	}
	err = s.AuthService.CreateUser(req)
	if err != nil {
		return
	}
}

// DeleteUser delete user handler
func (s *Service) DeleteUser(c *gin.Context) {
	req := &entity.UserReq{}
	rsp := &entity.CommRsp{}
	var err error

	defer func() {
		if err != nil {
			rsp.Code = errs.ErrCode(err)
			rsp.Msg = errs.ErrMsg(err)
		}
		c.PureJSON(http.StatusOK, gin.H{
			"code": rsp.Code,
			"msg":  rsp.Msg,
		})
	}()

	err = c.BindJSON(req)
	if err != nil {
		err = errs.New(entity.ErrCodeInvalidParam, err.Error())
		return
	}
	err = s.AuthService.DeleteUser(req)
	if err != nil {
		return
	}
}

// CreateRole create role handler
func (s *Service) CreateRole(c *gin.Context) {
	req := &entity.RoleReq{}
	rsp := &entity.CommRsp{}
	var err error

	defer func() {
		if err != nil {
			rsp.Code = errs.ErrCode(err)
			rsp.Msg = errs.ErrMsg(err)
		}
		c.PureJSON(http.StatusOK, gin.H{
			"code": rsp.Code,
			"msg":  rsp.Msg,
		})
	}()

	err = c.BindJSON(req)
	if err != nil {
		err = errs.New(entity.ErrCodeInvalidParam, err.Error())
		return
	}
	err = s.AuthService.CreateRole(req)
	if err != nil {
		return
	}
}

// DeleteRole delete role handler
func (s *Service) DeleteRole(c *gin.Context) {
	req := &entity.RoleReq{}
	rsp := &entity.CommRsp{}
	var err error

	defer func() {
		if err != nil {
			rsp.Code = errs.ErrCode(err)
			rsp.Msg = errs.ErrMsg(err)
		}
		c.PureJSON(http.StatusOK, gin.H{
			"code": rsp.Code,
			"msg":  rsp.Msg,
		})
	}()

	err = c.BindJSON(req)
	if err != nil {
		err = errs.New(entity.ErrCodeInvalidParam, err.Error())
		return
	}
	err = s.AuthService.DeleteRole(req)
	if err != nil {
		return
	}
}

// AddRoleToUser add role to user
func (s *Service) AddRoleToUser(c *gin.Context) {
	req := &entity.UserRoleReq{}
	rsp := &entity.CommRsp{}
	var err error

	defer func() {
		if err != nil {
			rsp.Code = errs.ErrCode(err)
			rsp.Msg = errs.ErrMsg(err)
		}
		c.PureJSON(http.StatusOK, gin.H{
			"code": rsp.Code,
			"msg":  rsp.Msg,
		})
	}()

	err = c.BindJSON(req)
	if err != nil {
		err = errs.New(entity.ErrCodeInvalidParam, err.Error())
		return
	}
	err = s.AuthService.AddRoleToUser(req)
	if err != nil {
		return
	}
}

// Authenticate authenticate
func (s *Service) Authenticate(c *gin.Context) {
	req := &entity.UserRoleReq{}
	rsp := &entity.UserRoleRsp{}
	var err error

	defer func() {
		if err != nil {
			rsp.Code = errs.ErrCode(err)
			rsp.Msg = errs.ErrMsg(err)
		}
		c.PureJSON(http.StatusOK, gin.H{
			"code":  rsp.Code,
			"msg":   rsp.Msg,
			"token": rsp.Token,
		})
	}()

	err = c.BindJSON(req)
	if err != nil {
		err = errs.New(entity.ErrCodeInvalidParam, err.Error())
		return
	}
	rsp, err = s.AuthService.Authenticate(req)
	if err != nil {
		return
	}
}

// Invalidate invalidate
func (s *Service) Invalidate(c *gin.Context) {
	req := &entity.UserRoleReq{}
	rsp := &entity.UserRoleRsp{}
	var err error

	defer func() {
		if err != nil {
			rsp.Code = errs.ErrCode(err)
			rsp.Msg = errs.ErrMsg(err)
		}
		c.PureJSON(http.StatusOK, gin.H{
			"code": rsp.Code,
			"msg":  rsp.Msg,
		})
	}()

	err = c.BindJSON(req)
	if err != nil {
		err = errs.New(entity.ErrCodeInvalidParam, err.Error())
		return
	}
	err = s.AuthService.Invalidate(req)
	if err != nil {
		return
	}
}

// CheckRole check role
func (s *Service) CheckRole(c *gin.Context) {
	req := &entity.UserRoleReq{}
	rsp := &entity.UserRoleRsp{}
	var err error

	defer func() {
		if err != nil {
			rsp.Code = errs.ErrCode(err)
			rsp.Msg = errs.ErrMsg(err)
		}
		c.PureJSON(http.StatusOK, gin.H{
			"code":         rsp.Code,
			"msg":          rsp.Msg,
			"check_result": rsp.CheckResult,
		})
	}()

	err = c.BindJSON(req)
	if err != nil {
		err = errs.New(entity.ErrCodeInvalidParam, err.Error())
		return
	}
	rsp, err = s.AuthService.CheckRole(req)
	if err != nil {
		return
	}
}

// AllRoles all roles
func (s *Service) AllRoles(c *gin.Context) {
	req := &entity.UserRoleReq{}
	rsp := &entity.UserRoleRsp{}
	var err error

	defer func() {
		if err != nil {
			rsp.Code = errs.ErrCode(err)
			rsp.Msg = errs.ErrMsg(err)
		}
		c.PureJSON(http.StatusOK, gin.H{
			"code":  rsp.Code,
			"msg":   rsp.Msg,
			"roles": rsp.Roles,
		})
	}()

	err = c.BindJSON(req)
	if err != nil {
		err = errs.New(entity.ErrCodeInvalidParam, err.Error())
		return
	}
	rsp, err = s.AuthService.AllRoles(req)
	if err != nil {
		return
	}
}
