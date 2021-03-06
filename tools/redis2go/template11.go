package main

const template11 string = `/// -------------------------------------------------------------------------------
/// THIS FILE IS ORIGINALLY GENERATED BY redis2go.exe.
/// PLEASE DO NOT MODIFY THIS FILE.
/// -------------------------------------------------------------------------------

package {{packagename}}

import (
	"errors"
	{{fmt}}
	{{import_struct_format}}
	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
	"github.com/gomodule/redigo/redis"
)

// {{classname}} : 代表 1 个 redis 对象
type {{classname}} struct {
	Key {{key_type}}
	{{fields_def}}

	dirtyDataIn{{classname}} map[string]interface{}
	dirtyDataForStructFiledIn{{classname}} map[string]interface{}
	isLoadIn{{classname}} bool
	dbKeyIn{{classname}} string
	ddbNameIn{{classname}} string
	expireIn{{classname}} uint
}

// New{{classname}} : New{{classname}} 的构造函数
func New{{classname}}(dbName string, key {{key_type}}) *{{classname}} {
	return &{{classname}} {
		Key: key,
		ddbNameIn{{classname}}: dbName,
		dbKeyIn{{classname}}: {{func_dbkey}},
		dirtyDataIn{{classname}}: make(map[string]interface{}),
		dirtyDataForStructFiledIn{{classname}}: make(map[string]interface{}),
	}
}

// HasKey : 是否存在 KEY
//          返回值，若访问数据库失败返回-1；若 key 存在返回 1 ，否则返回 0 。
func (obj{{classname}} *{{classname}}) HasKey() (int, error) {
	db := go_redis_orm.GetDB(obj{{classname}}.ddbNameIn{{classname}})
	val, err := redis.Int(db.Do("EXISTS", obj{{classname}}.dbKeyIn{{classname}}))
	if err != nil {
		return -1, err
	}
	return val, nil
}

// Load : 从 redis 加载数据
func (obj{{classname}} *{{classname}}) Load() error {
	if obj{{classname}}.isLoadIn{{classname}} == true {
		return errors.New("alreay load")
	}
	db := go_redis_orm.GetDB(obj{{classname}}.ddbNameIn{{classname}})
	val, err := redis.Values(db.Do("HGETALL", obj{{classname}}.dbKeyIn{{classname}}))
	if err != nil {
		return err
	}
	if len(val) == 0 {
		return go_redis_orm.ERR_ISNOT_EXIST_KEY
	}
	var data struct {
		{{fields_def_db}}
	}
	if err := redis.ScanStruct(val, &data); err != nil {
		return err
	}
	{{fields_init}}
	obj{{classname}}.isLoadIn{{classname}} = true
	return nil
}

// Save : 保存数据到 redis
func (obj{{classname}} *{{classname}}) Save() error {
	if len(obj{{classname}}.dirtyDataIn{{classname}}) == 0 && len(obj{{classname}}.dirtyDataForStructFiledIn{{classname}}) == 0 {
		return nil
	}
	for k := range(obj{{classname}}.dirtyDataForStructFiledIn{{classname}}) {
		_ = k
		{{fields_save}}
	}
	db := go_redis_orm.GetDB(obj{{classname}}.ddbNameIn{{classname}})
	if _, err := db.Do("HMSET", redis.Args{}.Add(obj{{classname}}.dbKeyIn{{classname}}).AddFlat(obj{{classname}}.dirtyDataIn{{classname}})...); err != nil {
		return err
	}
	if obj{{classname}}.expireIn{{classname}} != 0 {
		if _, err := db.Do("EXPIRE", obj{{classname}}.dbKeyIn{{classname}}, obj{{classname}}.expireIn{{classname}}); err != nil {
			return err
		}
	}
	obj{{classname}}.dirtyDataIn{{classname}} = make(map[string]interface{})
	obj{{classname}}.dirtyDataForStructFiledIn{{classname}} = make(map[string]interface{})
	return nil
}

// Delete : 从 redis 删除数据
func (obj{{classname}} *{{classname}}) Delete() error {
	db := go_redis_orm.GetDB(obj{{classname}}.ddbNameIn{{classname}})
	_, err := db.Do("DEL", obj{{classname}}.dbKeyIn{{classname}})
	if err == nil {
		obj{{classname}}.isLoadIn{{classname}} = false
		obj{{classname}}.dirtyDataIn{{classname}} = make(map[string]interface{})
		obj{{classname}}.dirtyDataForStructFiledIn{{classname}} = make(map[string]interface{})
	}
	return err
}

// IsLoad : 是否已经从 redis 导入数据
func (obj{{classname}} *{{classname}}) IsLoad() bool {
	return obj{{classname}}.isLoadIn{{classname}}
}

// Expire : 向 redis 设置该对象过期时间
func (obj{{classname}} *{{classname}}) Expire(v uint) {
	obj{{classname}}.expireIn{{classname}} = v
}

// DirtyData : 获取该对象目前已脏的数据
func (obj{{classname}} *{{classname}}) DirtyData() (map[string]interface{}, error) {
	for k := range(obj{{classname}}.dirtyDataForStructFiledIn{{classname}}) {
		_ = k
		{{fields_save2}}
	}
	data := make(map[string]interface{})
	for k, v := range(obj{{classname}}.dirtyDataIn{{classname}}) {
		data[k] = v
	}
	obj{{classname}}.dirtyDataIn{{classname}} = make(map[string]interface{})
	obj{{classname}}.dirtyDataForStructFiledIn{{classname}} = make(map[string]interface{})
	return data, nil
}

// Save2 : 保存数据到 redis 的第 2 种方法
func (obj{{classname}} *{{classname}}) Save2(dirtyData map[string]interface{}) error {
	if len(dirtyData) == 0 {
		return nil
	}
	db := go_redis_orm.GetDB(obj{{classname}}.ddbNameIn{{classname}})
	if _, err := db.Do("HMSET", redis.Args{}.Add(obj{{classname}}.dbKeyIn{{classname}}).AddFlat(dirtyData)...); err != nil {
		return err
	}
	if obj{{classname}}.expireIn{{classname}} != 0 {
		if _, err := db.Do("EXPIRE", obj{{classname}}.dbKeyIn{{classname}}, obj{{classname}}.expireIn{{classname}}); err != nil {
			return err
		}
	}
	return nil
}

{{func_get}}

{{func_set}}`

const getFuncString = `// Get{{field_name_upper}} : 获取字段值
func (obj{{classname}} *{{classname}}) Get{{field_name_upper}}() {{field_type}} {
	return obj{{classname}}.{{field_name_lower}}
}`

const setFuncString = `// Set{{field_name_upper}} : 设置字段值
func (obj{{classname}} *{{classname}}) Set{{field_name_upper}}(value {{field_type}}) {
	obj{{classname}}.{{field_name_lower}} = value
	obj{{classname}}.dirtyDataIn{{classname}}["{{field_name_lower_all}}"] = value
}`

const setFuncStringFieldstring = `// Set{{field_name_upper}} : 设置字段值
func (obj{{classname}} *{{classname}}) Set{{field_name_upper}}(value {{field_type}}) {
	obj{{classname}}.{{field_name_lower}} = value
	obj{{classname}}.dirtyDataIn{{classname}}["{{field_name_lower_all}}"] = string([]byte(value))
}`

const setFuncStringFieldbyte = `// Set{{field_name_upper}} : 设置字段值
func (obj{{classname}} *{{classname}}) Set{{field_name_upper}}(value {{field_type}}) {
	obj{{classname}}.{{field_name_lower}} = value
	tmp := make([]byte, len(value))
	copy(tmp, value)
	obj{{classname}}.dirtyDataIn{{classname}}["{{field_name_lower_all}}"] = tmp
}`

const getFuncStringForStructFiled = `// Get{{field_name_upper}} : 获取字段值
func (obj{{classname}} *{{classname}}) Get{{field_name_upper}}(mutable bool) *{{field_type}} {
	if mutable {
		obj{{classname}}.dirtyDataForStructFiledIn{{classname}}["{{field_name_lower_all}}"] = nil
	}
	return &obj{{classname}}.{{field_name_lower}}
}`

const getFuncStringSave = `if k == "{{field_name_lower_all}}" {
	data, err := {{struct_format}}.Marshal(&obj{{classname}}.{{field_name_lower}})
	if err != nil {
		return err
	}
	obj{{classname}}.dirtyDataIn{{classname}}["{{field_name_lower_all}}"] = data
}`

const getFuncStringSave2 = `if k == "{{field_name_lower_all}}" {
	data, err := {{struct_format}}.Marshal(&obj{{classname}}.{{field_name_lower}})
	if err != nil {
		return nil, err
	}
	obj{{classname}}.dirtyDataIn{{classname}}["{{field_name_lower_all}}"] = data
}`

const dbkeyFuncStringInt = `"{{classname}}:" + fmt.Sprintf("%d", key)`

const dbkeyFuncStringStr = `"{{classname}}:" + key`
