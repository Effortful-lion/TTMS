package mysql

import (
	"TTMS/model/do"

	"gorm.io/gorm"
)

type BaseUserDao interface {
	GetUserLoginByID(id int64) (do.UserTypeGetter, error)
	GetUserLoginByUsername(username string) (do.UserTypeGetter, error)
	InsertUserLogin(username, password string) error
	GetUserInfoByID(id int64) (do.UserInfoGetter, error)
}
// --------------------------------User注册表-----------------------------------
type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (*UserDao) GetUserLoginByID(id int64) (do.UserTypeGetter, error) {
	var user do.UserLogin
	err := DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (*UserDao) GetUserLoginByUsername(username string) (do.UserTypeGetter, error) {
	var user do.UserLogin
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 忽略记录不存在的错误，返回 nil 和 nil
			return nil, nil
		}
		// 处理其他错误
		return nil, err
	}
	return &user, nil
}

func (*UserDao) InsertUserLogin(username, password string) error {
	user := do.UserLogin{Username: username, Password: password}
	err := DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (*UserDao) GetUserInfoByID(id int64) (do.UserInfoGetter, error) {
	var user do.UserInfo
	err := DB.First(&user, id).Error	
	if err!= nil {
		return nil, err	
	}
	return &user, nil
}

// --------------------------------Admin注册表-----------------------------------
type AdminDao struct {
}

func NewAdminDao() *AdminDao {
	return &AdminDao{}
}

func (*AdminDao) GetUserLoginByID(id int64) (do.UserTypeGetter, error) {
	var user do.AdminLogin
	err := DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (*AdminDao) GetUserLoginByUsername(username string) (do.UserTypeGetter, error) {
	var user do.AdminLogin
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 忽略记录不存在的错误，返回 nil 和 nil
			return nil, nil
		}
		// 处理其他错误
		return nil, err
	}
	return &user, nil
}

func (*AdminDao) InsertUserLogin(username, password string) error {
	user := do.AdminLogin{Username: username, Password: password}
	err := DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (*AdminDao) GetUserInfoByID(id int64) (do.UserInfoGetter, error) {
	var user do.AdminInfo
	err := DB.First(&user, id).Error	
	if err!= nil {
		return nil, err	
	}
	return &user, nil
}
//--------------------------------Staff注册表-----------------------------------
type StaffDao struct {
}

func NewStaffDao() *StaffDao {
	return &StaffDao{}
}

func (*StaffDao) GetUserLoginByID(id int64) (do.UserTypeGetter, error) {
	var user do.StaffLogin
	err := DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (*StaffDao) GetUserLoginByUsername(username string) (do.UserTypeGetter, error) {
	var user do.StaffLogin
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 忽略记录不存在的错误，返回 nil 和 nil
			return nil, nil
		}
		// 处理其他错误
		return nil, err
	}
	return &user, nil
}

func (*StaffDao) InsertUserLogin(username, password string) error {
	user := do.StaffLogin{Username: username, Password: password}
	err := DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (*StaffDao) GetUserInfoByID(id int64) (do.UserInfoGetter, error) {
	var user do.StaffInfo
	err := DB.First(&user, id).Error	
	if err!= nil {
		return nil, err	
	}
	return &user, nil
}
//--------------------------------Manager注册表-----------------------------------
type ManagerDao struct {
}

func NewManagerDao() *ManagerDao {
	return &ManagerDao{}
}

func (*ManagerDao) GetUserLoginByID(id int64) (do.UserTypeGetter, error) {
	var user do.ManagerLogin
	err := DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (*ManagerDao) GetUserLoginByUsername(username string) (do.UserTypeGetter, error) {
	var user do.ManagerLogin
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 忽略记录不存在的错误，返回 nil 和 nil
			return nil, nil
		}
		// 处理其他错误
		return nil, err
	}
	return &user, nil
}

func (*ManagerDao) InsertUserLogin(username, password string) error {
	user := do.ManagerLogin{Username: username, Password: password}
	err := DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (*ManagerDao) GetUserInfoByID(id int64) (do.UserInfoGetter, error) {
	var user do.ManagerInfo
	err := DB.First(&user, id).Error	
	if err!= nil {
		return nil, err	
	}
	return &user, nil
}
//--------------------------------Customer注册表-----------------------------------


