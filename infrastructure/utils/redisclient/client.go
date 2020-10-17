package redisclient

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

const (
	defaultHost = "localhost:6379"
)

var (
	address       string
	password      string
	expiredTime   int
	redisInstance *redis.Client
	mu            sync.Mutex
)

// PingRedis ...
func PingRedis(hostPortAddr string, expTime string, passwd string) (*redis.Client, error) {

	address = hostPortAddr
	if address == "" {
		address = defaultHost
	}

	var err error
	expiredTime, err = strconv.Atoi(expTime)
	if err != nil {
		expiredTime = 5
	}

	password = passwd

	client := connector()
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, err
}

// check redisclient connection
func connector() *redis.Client {
	if redisInstance == nil {
		mu.Lock()
		defer mu.Unlock()

		if redisInstance == nil {
			return initRedis()
		}
	}

	return redisInstance
}

func initRedis() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
	redisInstance = client
	for i := 1; i <= 10; i++ {
		_, err := client.Ping().Result()
		if err == nil {
			break
		}
		logrus.Errorf("redisclient error: %s", err)
		logrus.Warn("Error connect to redisclient, reconnecting to : ", address, " trial: ", i)
		time.Sleep(3000 * time.Millisecond)

		if i == 10 {
			logrus.Fatal("System force stop because redisclient cannot connect")
		}
	}

	return redisInstance
}

func CreateNewKey(key string, val interface{}, ttl ...int) error {
	client := connector()

	var expiredTimeInt int
	var err error

	if len(ttl) < 1 {
		expiredTimeInt = expiredTime
	} else {
		expiredTimeInt = ttl[0]
	}

	durationTimeout := time.Duration(expiredTimeInt) * time.Second

	err = client.Set(key, val, durationTimeout).Err()

	return err
}

func ExtendTTLRedis(keys string, duration int) bool {
	client := connector()

	durationTimeout := time.Duration(duration) * time.Second
	res, err := client.Expire(keys, durationTimeout).Result()

	if err != nil {
		return false
	}

	return res
}

func IsKeyExist(key string) (status bool) {
	client := connector()
	_, err := client.Get(key).Result()
	status = false
	if err == redis.Nil {
		status = false
		fmt.Println("Keys does not exist")
	} else if err != nil {
		status = false
		panic(err)
	} else {
		status = true
	}
	return
}

func GetContentByKey(key string) (interface{}, bool) {
	client := connector()
	val, err := client.Get(key).Result()
	if err == redis.Nil {
		return "", false
	}
	return val, true
}

func DeleteKeys(keys ...string) int64 {
	client := connector()
	n, err := client.Del(keys...).Result()
	if err != nil {
		return 0
	}
	return n
}

// UpdateKey is updating the key value without resetting the TTL
func UpdateKey(key string, newValue string) error {

	// eval "local ttl = redisclient.call('ttl', ARGV[1]) if ttl > 0 then return redisclient.call('SETEX', ARGV[1], ttl, ARGV[2]) end" 0 key 987
	client := connector()
	script := "local ttl = redisclient.call('ttl', KEYS[1]) if ttl > 0 then return redisclient.call('SETEX', KEYS[1], ttl, ARGV[1]) end"

	_, err := client.Eval(script, []string{key}, []string{newValue}).Result()

	return err
}

func GetTTLByEval(key string) (int64, error) {
	client := connector()
	script := "local ttl = redisclient.call('ttl', KEYS[1]) if ttl > 0 then return redisclient.call('ttl', KEYS[1]) end"

	res, err := client.Eval(script, []string{key}).Result()

	if err != nil {
		return -1, err
	}

	return res.(int64), err
}
