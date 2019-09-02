package redisclusterutil

import (
	"fmt"
	"log"
	"time"

	redis "github.com/chasex/redis-go-cluster"
)

var cluster *redis.Cluster

//Init ...
func Init() {
	var err error
	cluster, err = redis.NewCluster(
		&redis.Options{
			StartNodes:   []string{"192.168.7.88:10000", "192.168.7.88:10002"},
			ConnTimeout:  100 * time.Millisecond,
			ReadTimeout:  100 * time.Millisecond,
			WriteTimeout: 150 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})

	if err != nil {
		panic("The Connection to the redis cluster couldnot be created. Hence do not let the server start")
	}
}

// //Lock ...
// func Lock(key, value string, timeoutMs int) (bool, error) {
// 	// r := pool.Get()
// 	// defer r.Close()

// 	cmd := redis.NewScript(1, lockScript)
// 	if res, err := cmd.Do(r, key, value, timeoutMs); err != nil {
// 		return false, err
// 	} else {
// 		return res == "OK", nil
// 	}
// }

//Ping ...
func Ping() {
	_, err := redis.String(cluster.Do("PING"))
	if err != nil {
		log.Printf("ERROR: fail ping redis conn: %s", err.Error())
		// os.Exit(1)
	}
}

//Set ...
func Set(key string, val string) error {

	_, err := cluster.Do("set", key, val)
	if err != nil {
		fmt.Printf("-set %s: %s\n", key, err.Error())
		return err
	}

	return nil
}

//Get ...
func Get(key string) (string, error) {

	value, err := redis.String(cluster.Do("GET", key))
	if err != nil {
		fmt.Printf("-get %s: %s\n", key, err.Error())
		return "", err
	}

	return value, nil
}
