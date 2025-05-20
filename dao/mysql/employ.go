package mysql

import "TTMS/model/do"

type EmployDao struct {
}

func NewEmployDao() *EmployDao {
	return &EmployDao{}
}

func (ud *EmployDao) SelectEmployByUsername(username string) (*do.Employ, error) {
	var employ do.Employ
	result := DB.Table("employ").Where("employ_name = ?", username).First(&employ)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, nil
		}
		return nil, result.Error
	}
	return &employ, nil
}

func (ud *EmployDao) InsertEmploy(username, password string) (id int64, err error) {
	employ := do.Employ{
		EmployName:     username,
		EmployPassword: password,
	}
	result := DB.Table("employ").Create(&employ)
	if result.Error!= nil {
		return 0, result.Error
	}
	return employ.EmployID, nil
}