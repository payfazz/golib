package redis

import (
	"encoding/json"
	"errors"
	"time"

	goredis "github.com/go-redis/redis"
	e "github.com/payfazz/golib/pkg/errors"
)

// Client ...
type Client struct {
	rc *goredis.Client
}

// Set , put data to redis, with seconds duration, when 0 seconds given, it will not expire
func (c *Client) Set(key string, data interface{}, seconds int64) error {
	var payload interface{}
	// k := reflect.ValueOf(data).Kind()
	jsonBytes, errJSON := json.Marshal(data) // try to marshal it to json, when failed set payload as it is.
	if errJSON != nil {
		return errJSON
	}
	payload = string(jsonBytes)
	cmd := c.rc.Set(key, payload, time.Duration(seconds)*time.Second)
	ok, err := cmd.Result()
	if err != nil {
		return err
	}
	if ok != "OK" {
		return e.Service("REDIS_ERR", "redis", errors.New("failed saving data to redis"))
	}
	return nil
}

// Get , get data from redis with given key and assign in to output
func (c *Client) Get(key string, output interface{}) error {
	cmd := c.rc.Get(key)
	payload, err := cmd.Result()
	if err != nil {
		return err
	}
	errUnmarshal := json.Unmarshal([]byte(payload), output)
	if errUnmarshal != nil {
		output = payload
	}
	return nil
}

// Del ...
func (c *Client) Del(key string) error {
	return c.rc.Del(key).Err()
}

// NewClient , get new redis client
func NewClient(address, password string, db int, name string) (*Client, error) {
	options := &goredis.Options{
		Addr: address,
		DB:   db,
		OnConnect: func(con *goredis.Conn) error {
			_, err := con.ClientSetName(name).Result()
			if err != nil {
				return err
			}
			return nil
		},
	}
	if password != "" {
		options.Password = password
	}

	rc := goredis.NewClient(options)
	c := &Client{
		rc: rc,
	}
	return c, nil
}
