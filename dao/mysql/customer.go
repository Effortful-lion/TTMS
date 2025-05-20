package mysql

import (
	"TTMS/model/do"

)

// --------------------------------User注册表-----------------------------------
type CustomerDao struct {
}

func NewCustomerDao() *CustomerDao {
	return &CustomerDao{}
}

func (ud *CustomerDao)SelectCustomerByUsername(username string) (*do.Customer, error) {
	var customer do.Customer
	result := DB.Table("customer").Where("customer_name = ?", username).First(&customer)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, nil
		}
		return nil, result.Error
	}
	return &customer, nil
}

func (ud *CustomerDao)InsertCustomer(username, password string) error {
	customer := do.Customer{
		CustomerName: username,
		CustomerPassword: password,
	}
	result := DB.Table("customer").Create(&customer)
	if result.Error!= nil {
		return result.Error
	}	
	return nil
}

