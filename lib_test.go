package dbapi

import (
	"testing"

	"github.com/HashemJaafar7/testutils"
)

func fTest[t any](actual t, expected t) {
	testutils.Test(true, false, true, 10, "v", actual, expected)
}

var db DB

func TestMain(m *testing.M) {
	act_err := Open(&db, "test")
	fTest(act_err, nil)
	db.DropAll()
	defer db.Close()

	m.Run()
}

func Test0(t *testing.T) {
	key := []byte{200, 98}
	exp_value := []byte{55}
	act_err := Update(db, key, exp_value)
	fTest(act_err, nil)

	act_value, act_err := Get(db, key)
	fTest(act_err, nil)
	fTest(act_value, exp_value)

	key = []byte{8}
	act_value, act_err = Get(db, key)
	fTest(act_err.Error(), "ErrKeyNotFound : key [8] not found")
	fTest(act_value, nil)
}

func Test1(t *testing.T) {
	var values [][]byte
	var keys [][]byte

	act_err := View(db, func(key, value []byte) {
		values = append(values, value)
		keys = append(keys, key)
	})
	fTest(act_err, nil)

	testutils.Debug("v", keys)
	testutils.Debug("v", values)
}

func Test2(t *testing.T) {
	key := []byte{1}
	expValue := []byte{1}

	{
		actErr := Delete(db, []byte{6})
		fTest(actErr, nil)
	}
	{
		actErr := Delete(db, key)
		fTest(actErr, nil)
	}
	{
		actErr := Add(db, key, expValue)
		fTest(actErr, nil)
	}
	{
		actErr := Add(db, key, expValue)
		fTest(actErr.Error(), "ErrKeyIsUsed : key [1] is used")
	}
}
