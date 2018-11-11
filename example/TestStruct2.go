/// -------------------------------------------------------------------------------
/// THIS FILE IS ORIGINALLY GENERATED BY redis2go.exe.
/// PLEASE DO NOT MODIFY THIS FILE.
/// -------------------------------------------------------------------------------
package main

import (
	"errors"
	"fmt"
	"strconv"

	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
	"github.com/gomodule/redigo/redis"
)

type TestStruct2 struct {
	Key    int32
	values map[int32]*TestStruct2Item

	__dirtyData map[int32]int
	__isLoad    bool
	__dbKey     string
	__dbName    string
	__expire    uint
}

func NewTestStruct2(dbName string, key int32) *TestStruct2 {
	return &TestStruct2{
		Key:         key,
		values:      make(map[int32]*TestStruct2Item),
		__dbName:    dbName,
		__dbKey:     "TestStruct2:" + fmt.Sprintf("%d", key),
		__dirtyData: make(map[int32]int),
	}
}

// 若访问数据库失败返回-1；若 key 存在返回 1 ，否则返回 0 。
func (this *TestStruct2) HasKey() (int, error) {
	db := go_redis_orm.GetDB(this.__dbName)
	val, err := redis.Int(db.Do("EXISTS", this.__dbKey))
	if err != nil {
		return -1, err
	}
	return val, nil
}

func (this *TestStruct2) Load() error {
	if this.__isLoad == true {
		return errors.New("alreay load!")
	}
	db := go_redis_orm.GetDB(this.__dbName)
	val, err := redis.Values(db.Do("HGETALL", this.__dbKey))
	if err != nil {
		return err
	}
	if len(val) == 0 {
		return go_redis_orm.ERR_ISNOT_EXIST_KEY
	}
	for i := 0; i < len(val); i += 2 {
		temp := string(val[i].([]byte))
		tempUint64, err := strconv.ParseUint(temp, 10, 64)
		subKey := int32(tempUint64)
		if err != nil {
			return err
		}
		item := NewTestStruct2Item(subKey, this)
		err = item.Unmarshal(val[i+1].([]byte))
		if err != nil {
			return err
		}
		this.values[subKey] = item
	}
	this.__isLoad = true
	return nil
}

func (this *TestStruct2) Save() error {
	if len(this.__dirtyData) == 0 {
		return nil
	}
	tempData := make(map[int32][]byte)
	for k, _ := range this.__dirtyData {
		if item, ok := this.values[k]; ok {
			var err error
			tempData[k], err = item.Marshal()
			if err != nil {
				return err
			}
		}
	}
	db := go_redis_orm.GetDB(this.__dbName)
	if _, err := db.Do("HMSET", redis.Args{}.Add(this.__dbKey).AddFlat(tempData)...); err != nil {
		return err
	}
	if this.__expire != 0 {
		if _, err := db.Do("EXPIRE", this.__dbKey, this.__expire); err != nil {
			return err
		}
	}
	this.__dirtyData = make(map[int32]int)
	return nil
}

func (this *TestStruct2) Delete() error {
	db := go_redis_orm.GetDB(this.__dbName)
	_, err := db.Do("DEL", this.__dbKey)
	if err == nil {
		this.__isLoad = false
		this.__dirtyData = make(map[int32]int)
	}
	return err
}

func (this *TestStruct2) NewItem(subKey int32) *TestStruct2Item {
	item := NewTestStruct2Item(subKey, this)
	this.values[subKey] = item
	this.__dirtyData[subKey] = 1
	return item
}

func (this *TestStruct2) DeleteItem(subKey int32) error {
	if _, ok := this.values[subKey]; ok {
		db := go_redis_orm.GetDB(this.__dbName)
		_, err := db.Do("HDEL", this.__dbKey, subKey)
		if err != nil {
			return err
		}
		delete(this.values, subKey)
		if _, ok := this.__dirtyData[subKey]; ok {
			delete(this.__dirtyData, subKey)
		}
	}
	return nil
}

func (this *TestStruct2) GetItem(subKey int32) *TestStruct2Item {
	if item, ok := this.values[subKey]; ok {
		return item
	}
	return nil
}

func (this *TestStruct2) GetItems() []*TestStruct2Item {
	var ret []*TestStruct2Item
	for _, v := range this.values {
		ret = append(ret, v)
	}
	return ret
}

func (this *TestStruct2) IsLoad() bool {
	return this.__isLoad
}

func (this *TestStruct2) Expire(v uint) {
	this.__expire = v
}
