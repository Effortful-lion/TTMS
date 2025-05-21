package redis

import (
	"TTMS/config"
	"TTMS/pkg/common"
	"context"
	"fmt"
	"time"
)

// token:admin:1
// token:customer:2
// token:account:1
// ...

var (
	TokenPrefix = "token:"
	TokenAdminPrefix = "token:admin:"
	TokenStaffPrefix = "token:staff:"
	TokenManagerPrefix = "token:manager:"
	TokenTicketorPrefix = "token:ticketor:"
	TokenFinancePrefix = "token:finance:"
	TokenAccountPrefix = "token:account:"
	TokenCustomerPrefix = "token:customer:"
)

func SetToken(token string, userID int64, auth string) error {
	var TokenExpireDuration = config.Conf.JwtConfig.Expire
	var key string
	switch auth {
	case common.AuthAdmin:
		key = fmt.Sprintf("%s%d", TokenAdminPrefix, userID)
	case common.AuthStaff:
		key = fmt.Sprintf("%s%d", TokenStaffPrefix, userID)
	case common.AuthManager:
		key = fmt.Sprintf("%s%d", TokenManagerPrefix, userID)
	case common.AuthTicketor:
		key = fmt.Sprintf("%s%d", TokenTicketorPrefix, userID)
	case common.AuthFinance:
		key = fmt.Sprintf("%s%d", TokenFinancePrefix, userID)
	case common.AuthAccount:
		key = fmt.Sprintf("%s%d", TokenAccountPrefix, userID)
	case common.AuthUser:
		key = fmt.Sprintf("%s%d", TokenCustomerPrefix, userID)
	}
	return Rdb.Set(context.Background(), key, token, time.Duration(TokenExpireDuration) * time.Second).Err()
}

func GetToken(userID int64, auth string) (string, error) {
	var key string
	switch auth {
	case common.AuthAdmin:
		key = fmt.Sprintf("%s%d", TokenAdminPrefix, userID)
	case common.AuthStaff:
		key = fmt.Sprintf("%s%d", TokenStaffPrefix, userID)
	case common.AuthManager:
		key = fmt.Sprintf("%s%d", TokenManagerPrefix, userID)
	case common.AuthTicketor:
		key = fmt.Sprintf("%s%d", TokenTicketorPrefix, userID)
	case common.AuthFinance:
		key = fmt.Sprintf("%s%d", TokenFinancePrefix, userID)
	case common.AuthAccount:
		key = fmt.Sprintf("%s%d", TokenAccountPrefix, userID)
	case common.AuthUser:
		key = fmt.Sprintf("%s%d", TokenCustomerPrefix, userID)
	}	
	return Rdb.Get(context.Background(), key).Result()
}

