package service

import (
	"blog-go/cache"
	"blog-go/models"
	"github.com/go-redis/redis"
)

//得到一篇文章的详情
func GetOneRecord(recordId uint64) (*models.Record, error) {
	//get from cache
	record, err := cache.GetOneRecordCache(recordId)
	if err == redis.Nil || err != nil {
		//get from mysql
		record, errSel := models.GetOne(int(recordId))
		if errSel != nil {
			return nil, errSel
		} else {
			//set cache
			errSet := cache.SetOneRecordCache(recordId, record)
			if errSet != nil {
				return nil, errSet
			} else {
				return record, errSel
			}
		}
	} else {
		return record, err
	}
}
