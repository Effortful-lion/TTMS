package dto


type UserSignUpReq struct {
	Username string `json:"username" binding:"required"` // 用户名，必填
	Password string `json:"password" binding:"required"` // 密码，必填	
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 确认密码，必填且必须与Password相同
	Auth		string `json:"auth" binding:"required,oneof=admin user staff manager ticketor finance account"`
}

type UserLoginReq struct {
	Username string `json:"username" binding:"required"` // 用户名，必填
	Password string `json:"password" binding:"required"` // 密码，必填
	Auth 	 string `json:"auth"` // 权限，选填
}

type UserLoginResp struct {
	UserID int64  `json:"user_id"`  // 用户ID
	Auth   string `json:"auth"`		// 用户权限
	Token string `json:"token"` 	// 登录成功后返回的token
}

type UserInfoResp struct {
	UserID int64  `json:"user_id"`  // 用户ID
	Username string `json:"username"` // 用户名
	Auth   string `json:"auth"`		// 用户权限
}
