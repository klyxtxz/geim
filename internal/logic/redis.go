package logic

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	ctx context.Context
	rdb *redis.Client
}

func Initredis(conf string) (redisObj *Redis) {
	re := new(Redis)
	re.ctx = context.Background()
	re.rdb = redis.NewClient(&redis.Options{
		Addr:     conf,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	log.Print("connect redis success:", conf)
	return re
}

func (re *Redis) put(token, id, room string) error {
	err := re.rdb.HSet(re.ctx, token, "id", id, "room", room).Err()
	log.Print("redis put:", id, ",", room)
	return err

}
func (re *Redis) get(token string) (map[string]string, error) {
	hash, err := re.rdb.HGetAll(re.ctx, token).Result()
	log.Print("redis get: token:", token, " hash:", hash)
	return hash, err

}
