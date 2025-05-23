package controller

import (
	"TTMS/model/dto"
	"TTMS/pkg/common"
	"TTMS/pkg"
	"TTMS/pkg/resp"
	"TTMS/service"
	"TTMS/dao/redis"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) GetUserInfoHandler(c *gin.Context) {
	// 从上下文中获取用户ID和权限
	user_id := GetCurrentUserID(c)
	auth := GetCurrentUserAuthority(c)
	if auth == "" {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	var err error
	var userinfo *dto.UserInfoResp
	if auth == common.AuthUser {
		// 如果是普通用户，就走普通用户的获取用户信息流程
		userinfo, err = service.NewCustomerService().GetUserInfo(user_id)	
		if err!= nil {
			resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
			return	
		}
	}else{
		// 其他就走员工获取用户信息流程
		userinfo, err = service.NewEmployService().GetUserInfo(user_id)
		if err!= nil {
			resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
			return	
		}
	}

	resp.ResponseSuccess(c, userinfo)
}

func (uc *UserController) LoginHandler(c *gin.Context) {
	var req *dto.UserLoginReq
	if err := c.ShouldBindJSON(&req); err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}

	username := req.Username
	password := req.Password
	auth := req.Auth
	
	if len(username) == 0 || len(password) == 0 {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return	
	}

	var user_id int64
	var err error
	// 如果登录的是普通用户，就走普通用户的登录流程
	if auth == common.AuthUser {
		user_id, err = service.NewCustomerService().Login(username, password)
		if err!= nil {
			resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
			return
		}
	}else{
		// 其他就走员工登录流程
		user_id, err = service.NewEmployService().Login(username, password, auth)
		if err!= nil {
			resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
			return
		}
	}
	var res = make(map[string]any)
	
	// 生成token
	token, err := pkg.GenToken(int(user_id), username, auth)
	if err != nil {
		resp.ResponseError(c, resp.CodeError)
		return	
	}
	res["user_id"] = user_id
	res["token"] = token

	if err := redis.SetToken(token, user_id, auth); err != nil {
		resp.ResponseError(c, resp.CodeError)
		return
	}

	resp.ResponseSuccess(c, res)
}

func (uc *UserController) SignUpHandler(c *gin.Context) {
	var req *dto.UserSignUpReq
	if err := c.ShouldBindJSON(&req); err != nil { 
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}

	username := req.Username
	password := req.Password
	re_password := req.RePassword
	auth := req.Auth

	if len(username) == 0 || len(password) == 0 {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}

	if password != re_password {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return	
	}

	// 如果注册的是普通用户，就走普通用户的注册流程
	if auth == common.AuthUser {
		err := service.NewCustomerService().SignUp(username, password)
		if err != nil {
			resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
			return	
		}
	}else {
		// 其他就走员工注册流程，还要加入权限设置
		err := service.NewEmployService().SignUp(username, password, auth)
		if err!= nil {
			resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
			return
		}
	}
	resp.ResponseSuccess(c, nil)
}