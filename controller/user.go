package controller

import (
	"TTMS/dao/redis"
	"TTMS/model/dto"
	"TTMS/pkg"
	"TTMS/pkg/resp"
	"TTMS/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

// @Summary 注册接口
// @Description 注册接口
// @Tags 全局接口
// @Accept json
// @Produc json
// @Param object body dto.UserSignUpReq true "请求参数"
// @Success 200 {object} resp.ResponseData "注册成功"
// @Router /signup [post]
func (*UserController) SignUpHandler(c *gin.Context) {
	var req *dto.UserSignUpReq
	if err := c.ShouldBindJSON(&req); err != nil { 
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}

	username := req.Username
	password := req.Password
	re_password := req.RePassword
	auth := req.Auth

	// 简单校验 // TODO 后续再加
	if len(username) == 0 || len(password) == 0 {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}

	if password != re_password {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return	
	}

	err := service.NewUserService().SignUp(username, password,auth)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return	
	}

	resp.ResponseSuccess(c, nil)
}

// @Summary 登录接口
// @Description 登录接口
// @Tags 全局接口
// @Accept json
// @Produc json
// @Param object body dto.UserLoginReq true "请求参数"
// @Success 200 {object} resp.ResponseData "登录响应信息"
// @Router /login [post]
func (*UserController) LoginHandler(c *gin.Context) {
	var req *dto.UserLoginReq
	if err := c.ShouldBindJSON(&req); err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}

	username := req.Username
	password := req.Password
	auth := req.Auth
	
	// 简单校验 // TODO 后续再加
	if len(username) == 0 || len(password) == 0 {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return	
	}

	user_id, err := service.NewUserService().Login(username, password, auth)
	if err!= nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
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

	// 将token存入redis中
	switch auth {
	case pkg.AuthAdmin:
		err = redis.SetAdminToken(token, int(user_id))
		if err!= nil {
			resp.ResponseError(c, resp.CodeError)
			return
		}	
	case pkg.AuthUser:
		err = redis.SetUserToken(token, int(user_id))
		if err!= nil {
			resp.ResponseError(c, resp.CodeError)
			return
		}
	default:
		resp.ResponseErrorWithMsg(c, resp.CodeError, "无权限")
		return
	}
	

	resp.ResponseSuccess(c, res)
}


func (*UserController) GetUserInfoHandler(c *gin.Context) {
	// 从上下文中获取用户ID和权限
	user_id := GetCurrentUserID(c)
	auth := GetCurrentUserAuthority(c)

	userinfo, err := service.NewUserService().GetUserInfo(user_id, auth)
	if err!= nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}

	resp.ResponseSuccess(c, userinfo)
}