package service

import (
	"TTMS/dao/mysql"
	"TTMS/model/dto"
	"errors"
)

type UserService struct {
	auth string
}

func NewUserService(auth string) *UserService {
	return &UserService{auth: auth}
}

func (u *UserService) SignUp(username, password string) (err error) {
	auth := u.auth
	// 定义接口变量
	userdao := mysql.NewUserDao(auth)
	// 统一处理用户注册逻辑
	user, err := userdao.SelectUserLoginByUsername(username)
	if err != nil {
		return errors.New("查询数据库失败")
	}
	if user != nil {
		return errors.New("用户名已存在")
	}
	if err = userdao.InsertUserLogin(username, password); err != nil {
		return errors.New("注册失败")
	}

	return nil
}

func (u *UserService) Login(username, password string) (data int64, err error) {
	auth := u.auth
	// 定义接口变量
	userdao := mysql.NewUserDao(auth)

	// 从数据库中查询用户名是否存在
	user, err := userdao.SelectUserLoginByUsername(username)
	if err != nil {
		return 0, errors.New("查询数据库失败")
	}
	if user == nil {
		return 0, errors.New("用户名不存在")
	}
	// 如果存在，比较密码是否正确
	if user.GetPassword() == password {
		return user.GetUserID(), nil
	}
	return 0, errors.New("用户名或密码错误")
}

func (u *UserService) GetUserInfo(user_id int64) (data *dto.UserInfoResp, err error) {
	auth := u.auth
	// 定义接口变量
	userdao := mysql.NewUserDao(auth)
	// 从数据库中查询用户信息
	user, err := userdao.SelectUserInfoByID(user_id)
	if err!= nil {
		return nil, errors.New("查询数据库失败")
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	userinfo := &dto.UserInfoResp{
		UserID: user.GetUserID(),
		Username: user.GetUsername(),
		Auth: u.auth,	
	}

	return userinfo, nil
}
