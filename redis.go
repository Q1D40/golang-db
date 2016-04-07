package db

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

type Redis struct {
}

func NewRedis() *Redis {
	return &Redis{}
}

func (this *Redis) conn() (redis.Conn, error) {
	redisconn := beego.AppConfig.String("redisconn")
	redispassword := beego.AppConfig.String("redispassword")
	c, err := redis.Dial("tcp", redisconn)
	if err != nil {
		c.Close()
		return c, err
	}
	if redispassword == "" {
		return c, err
	}
	_, err = c.Do("AUTH", redispassword)
	if err != nil {
		c.Close()
		return c, err
	}
	return c, err
}

func (this *Redis) Put(k string, v string) error {
	c, err := this.conn()
	defer c.Close()
	if err != nil {
		return err
	}
	_, err = c.Do("SET", k, v)
	if err != nil {
		return err
	}
	return nil
}

func (this *Redis) Setex(k string, ttl int64, v string) error {
	c, err := this.conn()
	defer c.Close()
	if err != nil {
		return err
	}
	_, err = c.Do("SETEX", k, ttl, v)
	if err != nil {
		return err
	}
	return nil
}

func (this *Redis) Get(k string) (interface{}, error) {
	c, err := this.conn()
	defer c.Close()
	if err != nil {
		return "", err
	}
	v, err := redis.String(c.Do("GET", k))
	if err != nil {
		return v, err
	}
	return v, nil
}

func (this *Redis) Lpush(k string, v string) error {
	c, err := this.conn()
	defer c.Close()
	if err != nil {
		return err
	}
	_, err = c.Do("lpush", k, v)
	if err != nil {
		return err
	}
	return nil
}
