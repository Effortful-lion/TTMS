package service

import (
	"TTMS/dao/mysql"
	"TTMS/model/dto"
	"TTMS/pkg/common"
	"errors"
)

type EmployService struct {
}

func NewEmployService() *EmployService {
	return &EmployService{}
}

func (u *EmployService) GetUserInfo(id int64) (data *dto.UserInfoResp, err error) {
	userdao := mysql.NewEmployDao()
	user, err := userdao.SelectEmployByID(id)
	if err!= nil {
		return nil, errors.New("查询数据库失败")
	}
	if user == nil {
		return nil, errors.New("用户不存在")	
	}
	var res dto.UserInfoResp
	res.UserID = user.EmployID
	res.Username = user.EmployName
	// 根据员工id查询用户角色
	userroledao := mysql.NewUserRoleDao()
	roleName, err := userroledao.SelectRoleByUserID(id)
	if err!= nil {
		return nil, errors.New("查询数据库失败")
	}
	res.Auth = roleName
	return &res, nil
}

func (u *EmployService) Login(username, password, auth string) (id int64, err error) {
	userdao := mysql.NewEmployDao()

	// 从数据库中查询用户名是否存在
	user, err := userdao.SelectEmployByUsername(username)
	if err != nil {
		return 0, errors.New("查询数据库失败")
	}
	if user == nil {
		return 0, errors.New("用户名不存在")
	}
	// 如果存在，比较密码是否正确
	if user.EmployPassword != password {
		
		return 0, errors.New("用户名或密码错误")
	}
	// 检查该用户的角色是否相符
	user_id := user.EmployID
	// 根据员工id查询用户角色
	userroledao := mysql.NewUserRoleDao()
	roleName, err := userroledao.SelectRoleByUserID(user_id)
	// 比较
	if err!= nil {
		return 0, errors.New("查询数据库失败")	
	}
	if roleName != auth {
		return 0, errors.New("权限错误,用户不存在")
	}

	return user.EmployID, nil
}

func (u *EmployService) SignUp(username, password, auth string) (err error) {
	userdao := mysql.NewEmployDao()
	user, err := userdao.SelectEmployByUsername(username)
	if err != nil {
		return errors.New("查询数据库失败")
	}
	if user != nil {
		return errors.New("用户名已存在")
	}

	// 员工注册，返回员工ID
	var employ_id int64
	if employ_id, err = userdao.InsertEmploy(username, password); err!= nil {
		return errors.New("注册失败")	
	}

	// 权限注册
	userroledao := mysql.NewUserRoleDao()
	role_id := 0
	switch auth {
	case common.AuthAdmin:
		role_id = common.AuthAdminID
	case common.AuthManager:
		role_id = common.AuthManagerID
	case common.AuthStaff:
		role_id = common.AuthStaffID
	case common.AuthFinance:
		role_id = common.AuthFinanceID
	case common.AuthTicketor:
		role_id = common.AuthTicketorID
	case common.AuthAccount:
		role_id = common.AuthAccountID
	default:
		return errors.New("权限错误")
	}
	err = userroledao.InsertUserRole(employ_id, int64(role_id))
	if err!= nil {
		return errors.New("权限注册失败")	
	}
	return nil
}