package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

const (
	// 用于Token的前缀
	TokenUser = "token_user"
	TokenAdmin = "token_admin"
	TokenStaff = "token_staff"
	TokenManager = "token_manager"
	TokenTicketor = "token_ticketor"
	TokenFinance = "token_fiance"
	TokenAccount = "token_account"
)
// --------------------------------------User---------------------------------------
// 将生成的token存储到redis中
func SetUserToken(token string, ID int) error {
	var TokenExpireDuration = viper.GetInt("auth.expire")
	return Rdb.Set(context.Background(),"token_user" + fmt.Sprint(ID),token,time.Hour * time.Duration(TokenExpireDuration)).Err()
}

// 通过用户ID得到用户token
func GetUserToken(ID int) (token string,err error){
	return Rdb.Get(context.Background(),"token_user" + fmt.Sprint(ID)).Result()
}

// // 通过token得到用户ID
// func GetUserID() (ID int,err error){
// 	token, _ := GetUserToken(ID)
// 	// 去掉前缀
// 	useridstr,found := strings.CutPrefix(token,"token_user")
// 	if !found {
// 		return 0,errors.New("token error")
// 	}
// 	return  strconv.Atoi(useridstr)
// }

// // 删除用户的登录token
// func DelUserToken(userID int) error {
// 	key := "token_user" + strconv.Itoa(userID)
// 	result := Rdb.Del(context.Background(),key)
// 	if result.Err() != nil {
// 		return result.Err()
// 	}
// 	return nil
// }

// // 删除用户的 忘记密码的修改 token
// func DelUserToken2(email string) error {
// 	key := "token" + email
// 	result := Rdb.Del(context.Background(),key)
// 	if result.Err() != nil {
// 		return result.Err()
// 	}
// 	return nil
// }

//-----------------------------------------Admin--------------------------------------------
// 将生成的token存储到redis中
func SetAdminToken(token string, ID int) error {
	var TokenExpireDuration = viper.GetInt("auth.expire")
	return Rdb.Set(context.Background(),TokenAdmin + fmt.Sprint(ID),token,time.Hour * time.Duration(TokenExpireDuration)).Err()
}

// 通过用户ID得到用户token
func GetAdminToken(ID int) (token string,err error){
	return Rdb.Get(context.Background(),TokenAdmin + fmt.Sprint(ID)).Result()
}