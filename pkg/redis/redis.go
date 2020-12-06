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

// Exp set redis key expiration time (in seconds)
func (c *Client) Exp(key string, seconds int64) error {
	cmd := c.rc.Expire(key, time.Duration(seconds)*time.Second)
	count, err := cmd.Result()
	if err != nil {
		return err
	}
	if !count {
		return e.Service("REDIS_ERR", "redis", errors.New("failed add data to set"))
	}
	return nil
}

// HSet insert one record to redis hash
func (c *Client) HSet(key, field string, data interface{}) error {
	var payload interface{}
	jsonBytes, errJSON := json.Marshal(data)
	if errJSON != nil {
		return errJSON
	}
	payload = string(jsonBytes)
	cmd := c.rc.HSet(key, field, payload)
	count, err := cmd.Result()
	if err != nil {
		return err
	}
	if !count {
		return e.Service("REDIS_ERR", "redis", errors.New("failed add data to hash"))
	}
	return nil
}

// HDel remove one record from redis hash
func (c *Client) HDel(key, field string) error {

	cmd := c.rc.HDel(key, field)
	count, err := cmd.Result()
	if err != nil {
		return err
	}
	if count == 0 {
		return e.Service("REDIS_ERR", "redis", errors.New("failed remove data from hash"))
	}
	return nil
}

// HExists check if data is a member of the set
func (c *Client) HExists(key, field string) (bool, error) {
	cmd := c.rc.HExists(key, field)
	count, err := cmd.Result()
	if err != nil {
		return false, err
	}
	if !count {
		return false, nil
	}
	return true, nil
}

// HVals get all hash record, output as a slice
func (c *Client) HVals(key string) ([]string, error) {
	cmd := c.rc.HVals(key)
	data, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// HGet get one hash record
func (c *Client) HGet(key, field string) (string, error) {
	cmd := c.rc.HGet(key, field)
	data, err := cmd.Result()
	if err != nil {
		return "", err
	}
	return data, nil
}

// SADD insert one record to redis set
func (c *Client) SADD(key string, data interface{}) error {
	var payload interface{}
	// k := reflect.ValueOf(data).Kind()
	jsonBytes, errJSON := json.Marshal(data) // try to marshal it to json, when failed set payload as it is.
	if errJSON != nil {
		return errJSON
	}
	payload = string(jsonBytes)
	cmd := c.rc.SAdd(key, payload)
	count, err := cmd.Result()
	if err != nil {
		return err
	}
	if count == 0 {
		return e.Service("REDIS_ERR", "redis", errors.New("failed add data to set"))
	}
	return nil
}

// SREM remove one record from redis set
func (c *Client) SREM(key string, data interface{}) error {
	var payload interface{}
	// k := reflect.ValueOf(data).Kind()
	jsonBytes, errJSON := json.Marshal(data) // try to marshal it to json, when failed set payload as it is.
	if errJSON != nil {
		return errJSON
	}
	payload = string(jsonBytes)
	cmd := c.rc.SRem(key, payload)
	count, err := cmd.Result()
	if err != nil {
		return err
	}
	if count == 0 {
		return e.Service("REDIS_ERR", "redis", errors.New("failed add data to set"))
	}
	return nil
}

// SISMEMBER check if data is a member of the set
func (c *Client) SISMEMBER(key string, data interface{}) (bool, error) {
	var payload interface{}
	// k := reflect.ValueOf(data).Kind()
	jsonBytes, errJSON := json.Marshal(data) // try to marshal it to json, when failed set payload as it is.
	if errJSON != nil {
		return false, errJSON
	}
	payload = string(jsonBytes)
	cmd := c.rc.SIsMember(key, payload)
	count, err := cmd.Result()
	if err != nil {
		return false, err
	}
	if !count {
		return false, nil
	}
	return true, nil
}

// SMEMBER get all Set record, output as a slice
func (c *Client) SMEMBER(key string) ([]string, error) {

	cmd := c.rc.SMembers(key)
	data, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}