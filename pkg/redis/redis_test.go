package redis_test

import (
	"strings"
	"testing"
	"time"

	"github.com/payfazz/golib/pkg/redis"
)

type Payload struct {
	ID    int       `json:"id"`
	Name  string    `json:"name"`
	Date  time.Time `json:"date"`
	State bool      `json:"state"`
}

func TestSetGetString(t *testing.T) {
	key := "pkg:redis:test:get-set-string"
	expected := string(time.Now().UnixNano())
	address := "localhost:6379"
	password := ""
	db := 0
	rc, err := redis.NewClient(address, password, db, key)
	if err != nil {
		t.Error(err)
	}
	err = rc.Set(key, expected, 10)
	if err != nil {
		t.Error(err)
	}

	actual := ""
	err = rc.Get(key, &actual)

	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf(`expected value : '%s', got '%s'`, expected, actual)
	}
}

func TestSetGetNumber(t *testing.T) {
	key := "pkg:redis:test:get-set-number"
	expected := time.Now().UnixNano()
	address := "localhost:6379"
	password := ""
	db := 0
	rc, err := redis.NewClient(address, password, db, key)
	if err != nil {
		t.Error(err)
	}
	err = rc.Set(key, expected, 10)
	if err != nil {
		t.Error(err)
	}

	actual := int64(0)
	err = rc.Get(key, &actual)

	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf(`expected value : '%v', got '%v'`, expected, actual)
	}
}

func TestSetGetStruct(t *testing.T) {
	key := "pkg:redis:test:get-set-struct"
	expected := &Payload{ID: 10}
	address := "localhost:6379"
	password := ""
	db := 0
	rc, err := redis.NewClient(address, password, db, key)
	if err != nil {
		t.Error(err)
	}
	err = rc.Set(key, expected, 10)
	if err != nil {
		t.Error(err)
	}

	actual := &Payload{}
	err = rc.Get(key, actual)

	if err != nil {
		t.Error(err)
	}
	if expected.ID != actual.ID {
		t.Errorf(`expected ID : '%v', got '%v'`, expected.ID, actual.ID)
	}
	if expected.Name != actual.Name {
		t.Errorf(`expected Name : '%v', got '%v'`, expected.Name, actual.Name)
	}
	if expected.Date != actual.Date {
		t.Errorf(`expected Date : '%v', got '%v'`, expected.Date, actual.Date)
	}
	if expected.State != actual.State {
		t.Errorf(`expected State : '%v', got '%v'`, expected.State, actual.State)
	}
}

func TestSetGetHash(t *testing.T) {
	key := "pkg:redis:test:get-set-hash"
	field := "testHash"
	expected := string(time.Now().UnixNano())
	address := "localhost:6379"
	password := ""
	db := 0
	rc, err := redis.NewClient(address, password, db, key)
	if err != nil {
		t.Error(err)
	}
	err = rc.HSet(key, field, expected)
	if err != nil {
		t.Error(err)
	}
	err = rc.Exp(key, 30)
	if err != nil {
		t.Error(err)
	}

	isExists, err := rc.HExists(key, field)
	if err != nil {
		t.Error(err)
	}
	if !isExists {
		t.Errorf(`expected failed`)
	}

	actual, err := rc.HGet(key, field)
	if err != nil {
		t.Error(err)
	}
	if expected != strings.Trim(actual, "\"") {
		t.Errorf(`expected value : '%s', got '%s'`, expected, actual)
	}

	err = rc.HDel(key, field)
	if err != nil {
		t.Error(err)
	}
	data, err := rc.HVals(key)
	if err != nil {
		t.Error(err)
	}
	if len(data) > 0 {
		t.Errorf(`expected value : empty, got '%s'`, data)
	}

}

func TestGetSetSet(t *testing.T) {
	key := "pkg:redis:test:get-set-set"
	expected := string(time.Now().UnixNano())
	address := "localhost:6379"
	password := ""
	db := 0
	rc, err := redis.NewClient(address, password, db, key)
	if err != nil {
		t.Error(err)
	}
	err = rc.SADD(key, expected)
	if err != nil {
		t.Error(err)
	}
	err = rc.Exp(key, 30)
	if err != nil {
		t.Error(err)
	}
	isMember, err := rc.SISMEMBER(key, expected)
	if err != nil {
		t.Error(err)
	}
	if !isMember {
		t.Error(`expected failed`)
	}

	err = rc.SREM(key, expected)
	if err != nil {
		t.Error(`expected failed`)
	}
	data, err := rc.SMEMBER(key)
	if err != nil {
		t.Error(`expected failed`)
	}
	if len(data) != 0 {
		t.Errorf(`expected value : empty, got '%s'`, data)
	}
}
