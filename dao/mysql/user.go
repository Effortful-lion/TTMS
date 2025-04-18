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

func (ud *UserDao) SelectUserLoginByID(id int64) (do.UserTypeGetter, error) {
	auth := ud.auth
	//var user do.UserLogin
	switch auth {
	case pkg.AuthAdmin:
		user := do.AdminLogin{}
		err := DB.Where("user_id = ?", id).First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthUser:
		user := do.UserLogin{}
		err := DB.Where("user_id = ?", id).First(&user).Error
		if err!= nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthStaff:
		user := do.StaffLogin{}
		err := DB.Where("user_id = ?", id).First(&user).Error
		if err!= nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthManager:
		user := do.ManagerLogin{}
		err := DB.Where("user_id = ?", id).First(&user).Error
		if err!= nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthFinance:
		user := do.FinanceLogin{}
		err := DB.Where("user_id = ?", id).First(&user).Error
		if err!= nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthTicketor:
		user := do.TicketorLogin{}
		err := DB.Where("user_id = ?", id).First(&user).Error
		if err!= nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthAccount:
		user := do.AccountLogin{}
		err := DB.Where("user_id = ?", id).First(&user).Error
		if err!= nil {
			return nil, err
		}
		return &user, nil
	default:
		return nil, errors.New("权限错误")
	}
}

func (ud *UserDao) SelectUserLoginByUsername(username string) (do.UserTypeGetter, error) {
	auth := ud.auth
	switch auth {
	case pkg.AuthAdmin:
		user := do.AdminLogin{}
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
	case pkg.AuthUser:
		user := do.UserLogin{}
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
	case pkg.AuthStaff:
		user := do.StaffLogin{}
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
	case pkg.AuthManager:
		user := do.ManagerLogin{}
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
	case pkg.AuthFinance:
		user := do.FinanceLogin{}
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
	case pkg.AuthTicketor:
		user := do.TicketorLogin{}
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
	case pkg.AuthAccount:
		user := do.AccountLogin{}
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
	default:
		return nil, errors.New("权限错误")
	}
}

func (ud *UserDao) InsertUserLogin(username, password string) error {
	auth := ud.auth
	switch auth {
	case pkg.AuthAdmin:
		user := do.AdminLogin{Username: username, Password: password}
		// 插入注册表
		err := DB.Create(&user).Error
		if err != nil {
			return err
		}
		// 插入信息表
		admin := do.AdminInfo{Username: username}
		err = DB.Create(&admin).Error
		if err!= nil {
			return err
		}
	case pkg.AuthUser:
		user := do.UserLogin{Username: username, Password: password}
		err := DB.Create(&user).Error
		if err != nil {
			return err
		}
		// 插入信息表
		userinfo := do.UserInfo{Username: username}
		err = DB.Create(&userinfo).Error
		if err!= nil {
			return err
		}
	case pkg.AuthStaff:
		user := do.StaffLogin{Username: username, Password: password}
		err := DB.Create(&user).Error
		if err != nil {
			return err
		}
		// 插入信息表
		staff := do.StaffInfo{Username: username}
		err = DB.Create(&staff).Error
		if err!= nil {
			return err
		}
	case pkg.AuthManager:
		user := do.ManagerLogin{Username: username, Password: password}
		err := DB.Create(&user).Error
		if err != nil {
			return err
		}
		// 插入信息表
		manager := do.ManagerInfo{Username: username}
		err = DB.Create(&manager).Error
		if err!= nil {
			return err
		}
	case pkg.AuthFinance:
		user := do.FinanceLogin{Username: username, Password: password}
		err := DB.Create(&user).Error
		if err != nil {
			return err
		}
		// 插入信息表
		finance := do.FinanceInfo{Username: username}
		err = DB.Create(&finance).Error
		if err!= nil {
			return err
		}
	case pkg.AuthTicketor:
		user := do.TicketorLogin{Username: username, Password: password}
		err := DB.Create(&user).Error
		if err != nil {
			return err
		}
		// 插入信息表
		ticketor := do.TicketorInfo{Username: username}
		err = DB.Create(&ticketor).Error
		if err!= nil {
			return err
		}
	case pkg.AuthAccount:
		user := do.AccountLogin{Username: username, Password: password}
		err := DB.Create(&user).Error
		if err != nil {
			return err
		}
		// 插入信息表
		account := do.AccountInfo{Username: username}
		err = DB.Create(&account).Error
		if err!= nil {
			return err
		}
	default:
		return errors.New("权限错误")
	}
	
	return nil
}

func (ud *UserDao) SelectUserInfoByID(id int64) (do.UserInfoGetter, error) {
	auth := ud.auth
	switch auth {
	case pkg.AuthAdmin:
		user := do.AdminInfo{}
		err := DB.Where("user_id = ?", id).First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthUser:
		user := do.UserInfo{}
		err := DB.Where("user_id = ?", id).First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthStaff:
		user := do.StaffInfo{}
		err := DB.Where("user_id = ?", id).First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthManager:
		user := do.ManagerInfo{}
		err := DB.Where("user_id = ?", id).First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthFinance:
		user := do.FinanceInfo{}
		err := DB.Where("user_id = ?", id).First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthTicketor:
		user := do.TicketorInfo{}
		err := DB.Where("user_id = ?", id).First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	case pkg.AuthAccount:
		user := do.AccountInfo{}
		err := DB.Where("user_id = ?", id).First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	default:
		return nil, errors.New("权限错误")
	}
}