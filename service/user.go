package service

import (
	"TTMS/dao/mysql"
	"TTMS/pkg"
	"errors"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) SignUp(username, password, auth string) (err error) {
	// 定义接口变量
	var userdao mysql.BaseUserDao

	// 根据权限类型赋值
	switch auth {
	case pkg.AuthUser:
		userdao = mysql.NewUserDao()
	case pkg.AuthAdmin:
		userdao = mysql.NewAdminDao()
	case pkg.AuthAccount:
		userdao = mysql.NewAccountDao()
	default:
		return errors.New("权限错误")
	}

	// 统一处理用户注册逻辑
	user, err := userdao.GetUserLoginByUsername(username)
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

func (*UserService) Login(username, password, auth string) (data int64, err error) {
	// 定义接口变量
	var userdao mysql.BaseUserDao

	// 根据权限类型赋值
	switch auth {
	case pkg.AuthUser:
		userdao = mysql.NewUserDao()
	case pkg.AuthAdmin:
		userdao = mysql.NewAdminDao()
	default:
		return 0, errors.New("权限错误")
	}

	// 从数据库中查询用户名是否存在
	user, err := userdao.GetUserLoginByUsername(username)
	if err != nil {
		return 0, errors.New("查询数据库失败")
	}
	// 如果存在，比较密码是否正确
	if user != nil && user.GetPassword() == password {
		return user.GetUserID(), nil
	}
	return 0, errors.New("用户名或密码错误")
}

func (*UserService) GetUserInfo(user_id int64, auth string) (data any, err error) {
	// 定义接口变量
	var userdao mysql.BaseUserDao	
	// 根据权限类型赋值
	switch auth {
	case pkg.AuthUser:
		userdao = mysql.NewUserDao()
	case pkg.AuthAdmin:
		userdao = mysql.NewAdminDao()
	default:
		return nil, errors.New("权限错误")		
	}

	// 从数据库中查询用户信息
	user, err := userdao.GetUserInfoByID(user_id)
	if err!= nil {
		return nil, errors.New("查询数据库失败")
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	return user, nil
}
