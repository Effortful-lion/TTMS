package mysql

import "TTMS/model/do"

type UserRoleDao struct {
}

func NewUserRoleDao() *UserRoleDao {
	return &UserRoleDao{}
}

func (ud *UserRoleDao)InsertUserRole(employ_id int64, roleID int64) error {
	return DB.Create(&do.UserRole{EmployID: employ_id, RoleID: roleID}).Error
}

func (ud *UserRoleDao)SelectRoleByUserID(user_id int64) (string, error) {
	var role do.Role
	if err := DB.Model(&do.UserRole{}).Select("role.role_name").Joins("JOIN role ON user_role.role_id = role.role_id").Where("user_role.employ_id = ?", user_id).First(&role).Error; err != nil {
		return "", err
	}
	return role.RoleName, nil
}