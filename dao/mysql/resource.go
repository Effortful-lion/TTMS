package mysql

import "TTMS/model/do"

type ResourceDao struct {
}

func NewResourceDao() *ResourceDao {
	return &ResourceDao{}
}

func (rd *ResourceDao) InsertResource(resourceName string) error {
	// 插入前先检查是否存在同名资源
	var count int64
	if err := DB.Model(&do.Resource{}).Where("resource_name = ?", resourceName).Count(&count).Error; err != nil {
		return err
	}
	return DB.Create(&do.Resource{ResourceName: resourceName}).Error
}

func (rd *ResourceDao) SelectResourceByName(resourceName string) (int64, error) {
	var resource do.Resource
	if err := DB.Where("resource_name = ?", resourceName).First(&resource).Error; err != nil {
		return 0, err
	}	
	return resource.ResourceID, nil
}