package model

import (
	"github.com/jinzhu/gorm"
	"simplelist/res/cache"
	"strconv"
)

type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:Uid"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"title"`
	Status    int    `gorm:"default:0"`
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64 `gorm:"default:"0`
}

func (Task *Task) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(Task.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (Task *Task) AddView() {
	cache.RedisClient.Incr(cache.TaskViewKey(Task.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Task.ID)))
}
