package do

type UserGetter interface {
	GetPassword() string
	GetUserID() int64
}

//----------------------------------------UserLogin表-------------------------------------------------
// 用户登录注册表
type UserLogin struct {
	UserID   int64  `gorm:"column:user_id;type:int;primaryKey;autoIncrement"`  // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一
	Password string `gorm:"column:password;type:varchar(100);not null"`        // 密码
}

func (u *UserLogin) GetPassword() string {
	return u.Password
}

func (u *UserLogin) GetUserID() int64 {
	return u.UserID
}

//----------------------------------------AdminLogin表-------------------------------------------------
type AdminLogin struct {
	UserID   int64  `gorm:"column:user_id;type:int;primaryKey;autoIncrement"`  // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一
	Password string `gorm:"column:password;type:varchar(100);not null"`        // 密码
}

func (a *AdminLogin) GetPassword() string {
	return a.Password
}

func (a *AdminLogin) GetUserID() int64 {
	return a.UserID	
}

//----------------------------------------StaffLogin表-------------------------------------------------
type StaffLogin struct {
	UserID   int64  `gorm:"column:user_id;type:int;primaryKey;autoIncrement"`  // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一
	Password string `gorm:"column:password;type:varchar(100);not null"`        // 密码
}

func (a *StaffLogin) GetPassword() string {
	return a.Password
}

func (a *StaffLogin) GetUserID() int64 {
	return a.UserID	
}
//----------------------------------------ManagerLogin表-------------------------------------------------
type ManagerLogin struct {
	UserID   int64  `gorm:"column:user_id;type:int;primaryKey;autoIncrement"`  // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一
	Password string `gorm:"column:password;type:varchar(100);not null"`        // 密码
}

func (a *ManagerLogin) GetPassword() string {
	return a.Password
}

func (a *ManagerLogin) GetUserID() int64 {
	return a.UserID	
}
//----------------------------------------TicketorLogin表-------------------------------------------------
type TicketorLogin struct {
	UserID   int64  `gorm:"column:user_id;type:int;primaryKey;autoIncrement"`  // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一
	Password string `gorm:"column:password;type:varchar(100);not null"`        // 密码
}

func (a *TicketorLogin) GetPassword() string {
	return a.Password
}

func (a *TicketorLogin) GetUserID() int64 {
	return a.UserID	
}
//----------------------------------------FinanceLogin表-------------------------------------------------
type FinanceLogin struct {
	UserID   int64  `gorm:"column:user_id;type:int;primaryKey;autoIncrement"`  // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一
	Password string `gorm:"column:password;type:varchar(100);not null"`        // 密码
}

func (a *FinanceLogin) GetPassword() string {
	return a.Password
}

func (a *FinanceLogin) GetUserID() int64 {
	return a.UserID	
}
//----------------------------------------AccountLogin表-------------------------------------------------
type AccountLogin struct {
	UserID   int64  `gorm:"column:user_id;type:int;primaryKey;autoIncrement"`  // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一
	Password string `gorm:"column:password;type:varchar(100);not null"`        // 密码
}

func (a *AccountLogin) GetPassword() string {
	return a.Password
}

func (a *AccountLogin) GetUserID() int64 {
	return a.UserID	
}
