package mysql

import (
	"TTMS/model/do"
	"TTMS/pkg"
	"errors"

	"gorm.io/gorm"
)

// --------------------------------User注册表-----------------------------------
type UserDao struct {
	auth string
}

func NewUserDao(auth string) *UserDao {
	return &UserDao{auth: auth}
}

func (ud *UserDao) GetUserLoginByID(id int64) (do.UserTypeGetter, error) {
	auth := ud.auth
	//var user do.UserLogin
	switch auth {
	case pkg.AuthAdmin:
		user := do.AdminLogin{UserID: id}
		err := DB.First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthUser:
		user := do.UserLogin{UserID: id}
		err := DB.First(&user).Error
		if err!= nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthStaff:
		user := do.StaffLogin{UserID: id}
		err := DB.First(&user).Error
		if err!= nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthManager:
		user := do.ManagerLogin{UserID: id}
		err := DB.First(&user).Error
		if err!= nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthFinance:
		user := do.FinanceLogin{UserID: id}
		err := DB.First(&user).Error
		if err!= nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthTicketor:
		user := do.TicketorLogin{UserID: id}
		err := DB.First(&user).Error
		if err!= nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthAccount:
		user := do.AccountLogin{UserID: id}
		err := DB.First(&user).Error
		if err!= nil {
			return nil, err
		}
		return &user, nil
	default:
		return nil, errors.New("权限错误")
	}
}

func (ud *UserDao) GetUserLoginByUsername(username string) (do.UserTypeGetter, error) {
	auth := ud.auth
	switch auth {
	case pkg.AuthAdmin:
		user := do.AdminLogin{Username: username}
		err := DB.First(&user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// 忽略记录不存在的错误，返回 nil 和 nil
				return nil, nil
			}
			// 处理其他错误
			return nil, err
		}
		return &user, nil
	case pkg.AuthUser:
		user := do.UserLogin{Username: username}
		err := DB.First(&user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// 忽略记录不存在的错误，返回 nil 和 nil
				return nil, nil
			}
			// 处理其他错误
			return nil, err
		}
		return &user, nil
	case pkg.AuthStaff:
		user := do.StaffLogin{Username: username}
		err := DB.First(&user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// 忽略记录不存在的错误，返回 nil 和 nil
				return nil, nil
			}
			// 处理其他错误
			return nil, err
		}
		return &user, nil
	case pkg.AuthManager:
		user := do.ManagerLogin{Username: username}
		err := DB.First(&user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// 忽略记录不存在的错误，返回 nil 和 nil
				return nil, nil
			}
			// 处理其他错误
			return nil, err
		}
		return &user, nil
	case pkg.AuthFinance:
		user := do.FinanceLogin{Username: username}
		err := DB.First(&user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// 忽略记录不存在的错误，返回 nil 和 nil
				return nil, nil
			}
			// 处理其他错误
			return nil, err
		}
		return &user, nil
	case pkg.AuthTicketor:
		user := do.TicketorLogin{Username: username}
		err := DB.First(&user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// 忽略记录不存在的错误，返回 nil 和 nil
				return nil, nil
			}
			// 处理其他错误
			return nil, err
		}
		return &user, nil
	case pkg.AuthAccount:
		user := do.AccountLogin{Username: username}
		err := DB.First(&user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// 忽略记录不存在的错误，返回 nil 和 nil
				return nil, nil
			}
			// 处理其他错误
			return nil, err
		}
		return &user, nil
	default:
		return nil, errors.New("权限错误")
	}
}

func (ud *UserDao) InsertUserLogin(username, password string) error {
	auth := ud.auth
	switch auth {
	case pkg.AuthAdmin:
		user := do.AdminLogin{Username: username, Password: password}
		err := DB.Create(&user).Error
		if err != nil {
			return err
		}
	case pkg.AuthUser:
		user := do.UserLogin{Username: username, Password: password}
		err := DB.First(&user).Error
		if err!= nil {
			return err
		}
	case pkg.AuthStaff:
		user := do.StaffLogin{Username: username, Password: password}
		err := DB.First(&user).Error
		if err != nil {
			return err
		}
	case pkg.AuthManager:
		user := do.ManagerLogin{Username: username, Password: password}
		err := DB.First(&user).Error
		if err != nil {
			return err
		}
	case pkg.AuthFinance:
		user := do.FinanceLogin{Username: username, Password: password}
		err := DB.First(&user).Error
		if err != nil {
			return err
		}
	case pkg.AuthTicketor:
		user := do.TicketorLogin{Username: username, Password: password}
		err := DB.First(&user).Error
		if err != nil {
			return err
		}
	case pkg.AuthAccount:
		user := do.AccountLogin{Username: username, Password: password}
		err := DB.First(&user).Error
		if err != nil {
			return err
		}
	default:
		return errors.New("权限错误")
	}
	
	return nil
}

func (ud *UserDao) GetUserInfoByID(id int64) (do.UserInfoGetter, error) {
	
	auth := ud.auth
	switch auth {
	case pkg.AuthAdmin:
		user := do.AdminLogin{UserID: id}
		err := DB.First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthUser:
		user := do.UserLogin{UserID: id}
		err := DB.First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthStaff:
		user := do.StaffLogin{UserID: id}
		err := DB.First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthManager:
		user := do.ManagerLogin{UserID: id}
		err := DB.First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthFinance:
		user := do.FinanceLogin{UserID: id}
		err := DB.First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthTicketor:
		user := do.TicketorLogin{UserID: id}
		err := DB.First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthAccount:
		user := do.AccountLogin{UserID: id}
		err := DB.First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	default:
		return nil, errors.New("权限错误")
	}
}