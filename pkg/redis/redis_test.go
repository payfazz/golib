package redis_test

import (
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
	connectionInfo := "localhost:6379"
	rc, err := redis.NewClient(connectionInfo, key)
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
	connectionInfo := "localhost:6379"
	rc, err := redis.NewClient(connectionInfo, key)
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
	connectionInfo := "localhost:6379"
	rc, err := redis.NewClient(connectionInfo, key)
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
