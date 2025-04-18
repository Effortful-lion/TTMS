package do

type UserTypeGetter interface {
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
//----------------------------------------UserInfo表-------------------------------------------------
type UserInfoGetter interface {
	GetUserID() int64
	GetUsername() string
}
 
type UserInfo struct {
	UserID   int64  `gorm:"column:user_id;type:int;primaryKey;autoIncrement"`  // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一
	// 其他用户信息字段
}
func (u *UserInfo) GetUserID() int64 {
	return u.UserID
}

func (u *UserInfo) GetUsername() string {
	return u.Username
}
//----------------------------------------AdminInfo表-------------------------------------------------
type AdminInfo struct {
	UserID   int64  `gorm:"column:user_id;type:int;primaryKey;autoIncrement"`  // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一	
}
func (u *AdminInfo) GetUserID() int64 {
	return u.UserID
}

func (u *AdminInfo) GetUsername() string {
	return u.Username
}
//----------------------------------------StaffInfo表----------------------------------------------
type StaffInfo struct {
	UserID   int64  `gorm:"column:user_id;type:int;primaryKey;autoIncrement"`  // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一	
}
func (u *StaffInfo) GetUserID() int64 {
	return u.UserID
}

func (u *StaffInfo) GetUsername() string {
	return u.Username
}
//----------------------------------------ManagerInfo表----------------------------------------------
type ManagerInfo struct {
	UserID   int64  `gorm:"column:user_id;type:int;primaryKey;autoIncrement" json:"user_id,omitempty"`   // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique" json:"username,omitempty"` // 用户名，唯一
}
func (u *ManagerInfo) GetUserID() int64 {
	return u.UserID
}

func (u *ManagerInfo) GetUsername() string {
	return u.Username
}
//----------------------------------------TicketorInfo表----------------------------------------------
type TicketorInfo struct {
	UserID   int64  `gorm:"column:user_id;type:int;primaryKey;autoIncrement"`  // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一	
}
func (u *TicketorInfo) GetUserID() int64 {
	return u.UserID
}

func (u *TicketorInfo) GetUsername() string {
	return u.Username
}
//----------------------------------------FinanceInfo表----------------------------------------------
type FinanceInfo struct {
	UserID   int64  `gorm:"column:user_id;type:int;primaryKey;autoIncrement"`  // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一
}
func (u *FinanceInfo) GetUserID() int64 {
	return u.UserID
}

func (u *FinanceInfo) GetUsername() string {
	return u.Username
}
//----------------------------------------AccountInfo表----------------------------------------------
type AccountInfo struct {
	UserID   int64  `gorm:"column:user_id;type:int;primaryKey;autoIncrement"`  // 用户ID，主键，自增
	Username string `gorm:"column:username;type:varchar(100);not null;unique"` // 用户名，唯一
}
func (u *AccountInfo) GetUserID() int64 {
	return u.UserID
}

func (u *AccountInfo) GetUsername() string {
	return u.Username
}


// 用户基本信息表
// type UserBasic struct {
// }

