package database

import (
	"github.com/teamssix/cf/pkg/util/pubutil"
)

func InsertOSSCache(OSSCache []pubutil.OSSCache) {
	DeleteOSSCache(OSSCache[0].AccessKeyId)
	CacheDb.Create(&OSSCache)
}

func DeleteOSSCache(AccessKeyId string) {
	var OSSCache []pubutil.OSSCache
	CacheDb.Where("access_key_id = ?", AccessKeyId).Delete(&OSSCache)
}

func SelectOSSCache(provider string) []pubutil.OSSCache {
	var OSSCache []pubutil.OSSCache
	AccessKeyId := SelectConfigInUse(provider).AccessKeyId
	CacheDb.Where("access_key_id = ?", AccessKeyId).Find(&OSSCache)
	return OSSCache
}
