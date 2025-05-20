package pkg

// 基于jwt的token身份验证：（token会涉及到access token和refresh token）
// 复习：jwt是一个三方库，提供了一系列方法
//（JWT全称JSON Web Token是一种跨域认证解决方案，属于一个开放的标准，它规定了一种Token 实现方式，目前多用于前后端分离项目和 OAuth2.0 业务场景下。）
// 用于生成token，解析token，加密token，解密token，生成token的payload，生成token的header，生成token的signature，生成token，验证token等。

import (
	"TTMS/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var Mysecret = []byte("jwt_secret")

// 自定义声明 : tag 的名称和数据库没有联系
// TODO 加一个权限字段，用于判断用户的权限（管理员/普通用户/场务员/运营经理/售票员/财务经理/会计）
type MyClaims struct {
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	Authority string `json:"authority"`
	jwt.RegisteredClaims
}

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

// 使用指定的secret生成返回token
func GenToken(userID int, username, authority string) (string, error) {
	// 在这里 用viper获取 ，确保配置文件已经被读取完
	TokenExpireDuration := int64(config.Conf.JwtConfig.Expire)
	// 创建一个我们自己的声明的数据
	c := MyClaims{
		UserID:    userID,
		Username:  username,
		Authority: authority,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "322Movie",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(TokenExpireDuration) * time.Second)),
		},
	}
	// 使用指定的签名算法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 返回tokenStr
	tokenStr, err := token.SignedString(Mysecret)
	return tokenStr, err
}

// 解析token
func ParseToken(tokenStr string) (*MyClaims, error) {
	// 解析token：mc接收解析token后其中包含的信息(claims)
	var mc = new(MyClaims)
	// 解析token：通过将tokenString解码到mc结构体中
	token, err := jwt.ParseWithClaims(tokenStr, mc, func(token *jwt.Token) (i interface{}, err error) {
		// 返回token的密钥
		return Mysecret, nil
	})
	if err != nil {
		// token解析失败
		return nil, err
	}
	if !token.Valid {
		// 校验token失败
		return nil, errors.New("invalid token")
	}
	// 校验token，token正确则返回负载的用户信息
	return mc, nil
}
