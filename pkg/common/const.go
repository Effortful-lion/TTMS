package common



// 定义角色
const (
	AuthAdmin    = "admin"    // 系统管理员
	AuthUser     = "user"     // 普通用户
	AuthStaff    = "staff"    // 场务员
	AuthManager  = "manager"  // 运营经理
	AuthTicketor = "ticketor" // 售票员
	AuthFinance  = "finance"  // 财务经理
	AuthAccount  = "account"  // 会计
)

// 定义角色id
const (
	AuthAdminID    = 1    // 系统管理员
	AuthUserID     = 7    // 普通用户
	AuthStaffID    = 2    // 场务员
	AuthManagerID  = 3    // 运营经理
	AuthTicketorID = 4    // 售票员
	AuthFinanceID  = 5    // 财务经理
	AuthAccountID  = 6    // 会计
)

// 定义访问资源
const (
	ResourceManage = "manage" // 管理
	ResourceUserinfo = "userinfo" // 用户信息
	ResourceSale = "sale"		// 售票或买票
)