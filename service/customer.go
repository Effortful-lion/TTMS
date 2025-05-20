package service

import (
	"TTMS/dao/mysql"
	"errors"
)

type CustomerService struct {
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (u *CustomerService) SignUp(username, password string) (err error) {
	// 定义接口变量
	userdao := mysql.NewCustomerDao()
	// 统一处理用户注册逻辑
	user, err := userdao.SelectCustomerByUsername(username)
	if err != nil {
		return errors.New("查询数据库失败")
	}
	if user != nil {
		return errors.New("用户名已存在")
	}
	if err = userdao.InsertCustomer(username, password); err != nil {
		return errors.New("注册失败")
	}
	return nil
}

func (u *CustomerService) Login(username, password string) (id int64, err error) {
	userdao := mysql.NewCustomerDao()

	// 从数据库中查询用户名是否存在
	user, err := userdao.SelectCustomerByUsername(username)
	if err != nil {
		return 0, errors.New("查询数据库失败")
	}
	if user == nil {
		return 0, errors.New("用户名不存在")
	}
	// 如果存在，比较密码是否正确
	if user.CustomerPassword == password {
		return user.CustomerID, nil
	}
	return 0, errors.New("用户名或密码错误")
}