package cache

import (
	"blog-go/dao/redisd"
	"blog-go/models"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

//token的过期时长
const ArticleDuration = time.Minute * 5

//cache的名字
func getRecordCacheName(recordId uint64) string {
	return "record_" + strconv.FormatUint(recordId, 10)
}

//从cache得到一篇文章
func GetOneRecordCache(recordId uint64) (*models.Record, error) {
	key := getRecordCacheName(recordId)
	val, err := redisd.RDB.Get(key).Result()

	if err == redis.Nil || err != nil {
		return nil, err
	} else {
		record := models.Record{}
		if err := json.Unmarshal([]byte(val), &record); err != nil {
			//t.Error(target)
			return nil, err
		}
		return &record, nil
	}
}

//向cache保存一篇文章
func SetOneRecordCache(recordId uint64, record *models.Record) error {
	key := getRecordCacheName(recordId)
	content, err := json.Marshal(record)
	if err != nil {
		fmt.Println(err)
		return err
	}
	errSet := redisd.RDB.Set(key, content, ArticleDuration).Err()
	if errSet != nil {
		return errSet
	}
	return nil
}
