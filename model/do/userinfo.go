package do

//----------------------------------------UserInfo表-------------------------------------------------
type UserInfoGetter interface {
	GetUserID() int64
	GetUsername() string
}
 
type UserInfo struct {
	UserInfoID int64  `gorm:"column:user_info_id;type:int;primaryKey;autoIncrement"` // 用户ID，主键，自增
	Username   string `gorm:"column:username;type:varchar(100);not null;unique"`     // 用户名（唯一）

	// 关联
	Orders []Order
	Tickets []Ticket
}
func (u *UserInfo) GetUserID() int64 {
	return u.UserInfoID
}

func (u *UserInfo) GetUsername() string {
	return u.Username
}
//----------------------------------------AdminInfo表-------------------------------------------------
type AdminInfo struct {
	UserInfoID int64  `gorm:"column:user_info_id;type:int;primaryKey;autoIncrement"` // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一	
}
func (u *AdminInfo) GetUserID() int64 {
	return u.UserInfoID
}

func (u *AdminInfo) GetUsername() string {
	return u.Username
}
//----------------------------------------StaffInfo表----------------------------------------------
type StaffInfo struct {
	UserInfoID int64  `gorm:"column:user_info_id;type:int;primaryKey;autoIncrement"` // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一	
}
func (u *StaffInfo) GetUserID() int64 {
	return u.UserInfoID
}

func (u *StaffInfo) GetUsername() string {
	return u.Username
}
//----------------------------------------ManagerInfo表----------------------------------------------
type ManagerInfo struct {
	UserInfoID int64  `gorm:"column:user_info_id;type:int;primaryKey;autoIncrement"` // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique" json:"username,omitempty"` // 用户名，唯一
}
func (u *ManagerInfo) GetUserID() int64 {
	return u.UserInfoID
}

func (u *ManagerInfo) GetUsername() string {
	return u.Username
}
//----------------------------------------TicketorInfo表----------------------------------------------
type TicketorInfo struct {
	UserInfoID int64  `gorm:"column:user_info_id;type:int;primaryKey;autoIncrement"` // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一	
}
func (u *TicketorInfo) GetUserID() int64 {
	return u.UserInfoID
}

func (u *TicketorInfo) GetUsername() string {
	return u.Username
}
//----------------------------------------FinanceInfo表----------------------------------------------
type FinanceInfo struct {
	UserInfoID int64  `gorm:"column:user_info_id;type:int;primaryKey;autoIncrement"` // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一
}
func (u *FinanceInfo) GetUserID() int64 {
	return u.UserInfoID
}

func (u *FinanceInfo) GetUsername() string {
	return u.Username
}
//----------------------------------------AccountInfo表----------------------------------------------
type AccountInfo struct {
	UserInfoID int64  `gorm:"column:user_info_id;type:int;primaryKey;autoIncrement"` // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一
}
func (u *AccountInfo) GetUserID() int64 {
	return u.UserInfoID
}

func (u *AccountInfo) GetUsername() string {
	return u.Username
}


// 用户基本信息表
// type UserBasic struct {
// }

