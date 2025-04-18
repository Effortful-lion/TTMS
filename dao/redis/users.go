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
	fmt.Println("TokenExpireDuration:",TokenExpireDuration)
	return Rdb.Set(context.Background(),"token_user" + fmt.Sprint(ID),token,time.Second * time.Duration(TokenExpireDuration)).Err()
}

// 通过用户ID得到用户token
func GetUserToken(ID int) (token string,err error){
	return Rdb.Get(context.Background(),"token_user" + fmt.Sprint(ID)).Result()
}

// // 删除用户的登录token
// func DelUserToken(userID int) error {
// 	key := "token_user" + strconv.Itoa(userID)
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
	return Rdb.Set(context.Background(),TokenAdmin + fmt.Sprint(ID),token,time.Second * time.Duration(TokenExpireDuration)).Err()
}

// 通过用户ID得到用户token
func GetAdminToken(ID int) (token string,err error){
	return Rdb.Get(context.Background(),TokenAdmin + fmt.Sprint(ID)).Result()
}

//----------------------------------------Manager-------------------------------------------
// 将生成的token存储到redis中
func SetManagerToken(token string, ID int) error {
	var TokenExpireDuration = viper.GetInt("auth.expire")
	return Rdb.Set(context.Background(),TokenManager + fmt.Sprint(ID),token,time.Second * time.Duration(TokenExpireDuration)).Err()
}

// 通过用户ID得到用户token
func GetManagerToken(ID int) (token string,err error){
	return Rdb.Get(context.Background(),TokenManager + fmt.Sprint(ID)).Result()
}

//----------------------------------------Staff-------------------------------------------
// 将生成的token存储到redis中
func SetStaffToken(token string, ID int) error {
	var TokenExpireDuration = viper.GetInt("auth.expire")
	return Rdb.Set(context.Background(),TokenStaff + fmt.Sprint(ID),token,time.Second * time.Duration(TokenExpireDuration)).Err()
}

// 通过用户ID得到用户token
func GetStaffToken(ID int) (token string,err error){
	return Rdb.Get(context.Background(),TokenStaff + fmt.Sprint(ID)).Result()
}
//----------------------------------------Finance-------------------------------------------
// 将生成的token存储到redis中
func SetFinanceToken(token string, ID int) error {
	var TokenExpireDuration = viper.GetInt("auth.expire")
	return Rdb.Set(context.Background(),TokenFinance + fmt.Sprint(ID),token,time.Second * time.Duration(TokenExpireDuration)).Err()
}

// 通过用户ID得到用户token
func GetFinanceToken(ID int) (token string,err error){
	return Rdb.Get(context.Background(),TokenFinance + fmt.Sprint(ID)).Result()
}
//----------------------------------------Account-------------------------------------------
// 将生成的token存储到redis中
func SetAccountToken(token string, ID int) error {
	var TokenExpireDuration = viper.GetInt("auth.expire")
	return Rdb.Set(context.Background(),TokenAccount + fmt.Sprint(ID),token,time.Second * time.Duration(TokenExpireDuration)).Err()
}

// 通过用户ID得到用户token
func GetAccountToken(ID int) (token string,err error){
	return Rdb.Get(context.Background(),TokenAccount + fmt.Sprint(ID)).Result()
}
//----------------------------------------Ticketor-------------------------------------------
// 将生成的token存储到redis中
func SetTicketorToken(token string, ID int) error {
	var TokenExpireDuration = viper.GetInt("auth.expire")
	return Rdb.Set(context.Background(),TokenTicketor + fmt.Sprint(ID),token,time.Second * time.Duration(TokenExpireDuration)).Err()
}

// 通过用户ID得到用户token
func GetTicketorToken(ID int) (token string,err error){
	return Rdb.Get(context.Background(),TokenTicketor + fmt.Sprint(ID)).Result()
}
