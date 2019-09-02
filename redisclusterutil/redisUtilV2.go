package redisclusterutil

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
)

const (
	lockScript = `
		return redis.call('SET', KEYS[1], ARGV[1], 'NX', 'PX', ARGV[2])
	`
	unlockScript = `
		if redis.call("get",KEYS[1]) == ARGV[1] then
		    return redis.call("del",KEYS[1])
		else
		    return 0
		end
	`
)

var clusterV2 redisc.Cluster

//InitV2 ..
func InitV2() {

	clusterV2 = redisc.Cluster{
		StartupNodes: []string{"192.168.7.88:10000", "192.168.7.88:10002"},
		DialOptions:  []redis.DialOption{redis.DialConnectTimeout(5 * time.Second)},
		CreatePool:   createPool,
	}

	// initialize its mapping
	if err := clusterV2.Refresh(); err != nil {
		log.Fatalf("Refresh failed: %v", err)
	}

}

//Lock ...
func Lock(key, value string, timeoutMs int) (bool, error) {

	// create a script that takes 2 keys and 2 values, and returns 1
	script := redis.NewScript(1, lockScript)
	fmt.Println("The script was created")

	conn := clusterV2.Get()
	defer conn.Close()
	// bind it to the right node for the required keys, ahead of time
	if err := redisc.BindConn(conn, key); err != nil {
		log.Fatalf("BindConn failed: %v", err)
	}

	// cmd := redis.NewScript(1, lockScript)
	if res, err := script.Do(conn, key, value, timeoutMs); err != nil {
		return false, err
	} else {
		return res == "OK", nil
	}
}

// Unlock attempts to remove the lock on a key so long as the value matches.
// If the lock cannot be removed, either because the key has already expired or
// because the value was incorrect, an error will be returned.
func Unlock(key, value string) error {
	conn := clusterV2.Get()
	defer conn.Close()

	// bind it to the right node for the required keys, ahead of time
	if err := redisc.BindConn(conn, key); err != nil {
		log.Fatalf("BindConn failed: %v", err)
	}

	cmd := redis.NewScript(1, unlockScript)
	if res, err := redis.Int(cmd.Do(conn, key, value)); err != nil {
		return err
	} else if res != 1 {
		return errors.New("Unlock failed, key or secret incorrect")
	}

	// Success
	return nil
}

func createPool(addr string, opts ...redis.DialOption) (*redis.Pool, error) {
	return &redis.Pool{
		MaxIdle:     5,
		MaxActive:   10,
		IdleTimeout: time.Minute,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, opts...)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}, nil
}

//PingV2 ...
func PingV2() {

	conn := clusterV2.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		log.Printf("ERROR: fail ping redis conn: %s", err.Error())
		// os.Exit(1)
	}
}

//SetV2 ...
func SetV2(key string, val string) error {

	conn := clusterV2.Get()
	defer conn.Close()
	_, err := conn.Do("set", key, val)
	if err != nil {
		fmt.Printf("-set %s: %s\n", key, err.Error())
		return err
	}

	return nil
}

//GetV2 ...
func GetV2(key string) (string, error) {

	conn := clusterV2.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		fmt.Printf("-get %s: %s\n", key, err.Error())
		return "", err
	}

	return value, nil
}

//Close ...
func Close() error {
	return clusterV2.Close()
}
