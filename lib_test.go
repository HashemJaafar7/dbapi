package go_database_api

import (
	"fmt"
	"testing"

	test "github.com/HashemJaafar7/go_test"
)

var db DB

func TestMain(m *testing.M) {
	act_err := Open(&db, "test")
	test.Test(false, true, "#v", act_err, nil)
	defer db.Close()

	m.Run()
}

func Test0(t *testing.T) {
	key := []byte{200, 98}
	exp_value := []byte{55}
	act_err := Update(db, key, exp_value)
	test.Test(false, true, "#v", act_err, nil)

	act_value, act_err := Get(db, key)
	test.Test(false, true, "#v", act_err, nil)
	test.Test(false, true, "#v", act_value, exp_value)

	key = []byte{8}
	act_value, act_err = Get(db, key)
	test.Test(false, true, "#v", act_err, fmt.Errorf("key [8] not found"))
	test.Test(false, true, "#v", act_value, nil)
}

func Test1(t *testing.T) {
	var values [][]byte
	var keys [][]byte

	act_err := View(db, func(key, value []byte) {
		values = append(values, value)
		keys = append(keys, key)
	})
	test.Test(false, true, "#v", act_err, nil)

	test.Debug("v", keys)
	test.Debug("v", values)
}

func Test2(t *testing.T) {
	key := []byte{1}
	exp_value := []byte{1}

	act_err := Delete(db, []byte{6})
	test.Test(false, true, "#v", act_err, nil)

	act_err = Delete(db, key)
	test.Test(false, true, "#v", act_err, nil)

	act_err = Add(db, key, exp_value)
	test.Test(false, true, "#v", act_err, nil)

	act_err = Add(db, key, exp_value)
	test.Test(false, true, "#v", act_err, fmt.Errorf("key [1] is used"))
}
