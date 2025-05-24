package mysql

import (
	"TTMS/model/do"

)

// --------------------------------Customer注册表-----------------------------------
type CustomerDao struct {
}

func NewCustomerDao() *CustomerDao {
	return &CustomerDao{}
}
func (ud *CustomerDao)SelectCustomerByID(id int64) (*do.Customer, error) {
	var customer do.Customer
	result := DB.Table("customer").Where("customer_id =?", id).First(&customer)
	if result.Error!= nil {
		return nil, result.Error
	}
	return &customer, nil
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

func (ud *CustomerDao)InsertCustomer(username, password string) (int64, error) {
	customer := do.Customer{
		CustomerName: username,
		CustomerPassword: password,
	}
	result := DB.Table("customer").Create(&customer)
	if result.Error!= nil {
		return 0, result.Error
	}	
	customer_id := customer.CustomerID
	return customer_id, nil
}

