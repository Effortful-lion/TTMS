package redis

import (
	"TTMS/dao/mysql"
	"TTMS/pkg/common"
	"context"
	"fmt"
	"time"
)

// 初始化的时候，给角色和权限设置好

var role = []string{common.AuthAdmin, common.AuthUser, common.AuthStaff, common.AuthManager, common.AuthTicketor, common.AuthFinance, common.AuthAccount}

var resource_role = map[string][]string{
	common.ResourceUserinfo:{common.AuthAdmin, common.AuthUser, common.AuthStaff, common.AuthManager, common.AuthTicketor, common.AuthFinance, common.AuthAccount},
	common.ResourceManage:{common.AuthAdmin},
	common.ResourceSale:{common.AuthAdmin, common.AuthUser},
	common.ResourceCollection:{common.AuthAdmin},
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
	fmt.Println("初始化资源权限成功")
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

// TODO 定时同步角色-资源权限表
func SyncRoleResource() error {
	tk := time.NewTicker(1 * time.Hour)
	for range tk.C {
		// 遍历所有资源，查询所有角色
		// eg: resource:userinfo, resource:manage, resource:sale, resource:collection
		keys, err := Rdb.Keys(context.Background(), "resource:*").Result()
		if err!= nil {
			return err
		}
		// keys: []string{"resource:userinfo", "resource:manage", "resource:sale", "resource:collection"}
		// 需要的是 []string{"userinfo", "manage", "sale", "collection"}
		for _, v := range keys {
			// 分割字符串，获取资源名
			resource := v[8:] // 从第8个字符开始，获取资源名
			// 插入资源名
			err := mysql.NewResourceDao().InsertResource(resource)
			if err!= nil {
				return err	
			}
			// 获取资源的角色列表
			roles, err := Rdb.SMembers(context.Background(), "resource:"+resource).Result()
			if err!= nil {
				return err
			}
			// roles: []string{"admin", "user", "staff", "manager", "ticketor", "finance", "account"}
			// 遍历角色列表，设置角色-资源权限表
			for _, role := range roles {
				// 调用 mysql 的函数，设置角色-资源权限表
				role_id := common.GetRoleID(role)
				resource_id, err := mysql.NewResourceDao().SelectResourceByName(resource)
				if err!= nil {
					return err
				}
				err = mysql.NewUserRoleDao().SyncRoleResource(int64(role_id), resource_id)
				if err!= nil {
					return err
				}
			}	
		}
		
	}
	
	return nil
}