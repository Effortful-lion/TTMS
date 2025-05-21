package redis

import (
	"TTMS/pkg/common"
	"context"
)

// 初始化的时候，给角色和权限设置好

var role = []string{common.AuthAdmin, common.AuthUser, common.AuthStaff, common.AuthManager, common.AuthTicketor, common.AuthFinance, common.AuthAccount}

var resource_role = map[string][]string{
	common.ResourceUserinfo:{common.AuthAdmin, common.AuthUser, common.AuthStaff, common.AuthManager, common.AuthTicketor, common.AuthFinance, common.AuthAccount},
	common.ResourceManage:{common.AuthAdmin},
	common.ResourceSale:{common.AuthAdmin, common.AuthTicketor, common.AuthUser},
}

func InitRedisResource() error {
	// 初始化角色
	err := SetRole(role)
	if err!= nil {
		return err
	}
	// 初始化资源权限
	err = SetResourceRole(resource_role)
	if err!= nil {
		return err
	}
	return nil
}

// 设置角色
func SetRole(role []string) error {
	for _, v := range role {
		// role 列表
		err := Rdb.SAdd(context.Background(),"role", v).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func SetResourceRole(resource_role map[string][]string) error {
	for k, v := range resource_role {
		// 资源列表: k: 资源名，v: 角色列表
		// SADD resource:k v1 v2 v3
		err := Rdb.SAdd(context.Background(),"resource:"+k, v).Err()
		if err!= nil {
			return err	
		}
	}	
	return nil
}

func SetResourceIsMember(resource, role string) bool {
	key := "resource:" + resource
	return Rdb.SIsMember(context.Background(), key, role).Val()
}