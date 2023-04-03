package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
)

var configFile string
var defaultValues map[string]json.RawMessage
var currentValues map[string]json.RawMessage

func init() {
	defaultValues = make(map[string]json.RawMessage)
	currentValues = make(map[string]json.RawMessage)
}

func SetConfigFile(path string) error {
	configFile = path
	if v, err := parseConfigFile(configFile); err != nil {
		currentValues = v
		return err
	} else {
		currentValues = v
		return nil
	}
}

func GetConfigFile() string {
	return configFile
}

func parseConfigFile(conf string) (map[string]json.RawMessage, error) {
	if content, err := ioutil.ReadFile(conf); err != nil {
		return nil, errors.New("failed to read config file:" + err.Error())
	} else {
		var currentV map[string]json.RawMessage
		if err := json.Unmarshal(content, &currentV); err != nil {
			return nil, errors.New("failed to parse config file:" + err.Error())
		}
		return currentV, nil
	}
}

// GeneDefaultConfig 从所有注册的配置中， 生成一份默认的配置
func GeneDefaultConfig() ([]byte, error) {
	content, err := json.MarshalIndent(defaultValues, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to get default settings:%s", err)
	}
	return content, nil
}

// RegConfig 注册配置项和默认值
func RegConfig(key string, val interface{}) error {
	c, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("failed to reg config with key %s:%s", key, err)
	}
	_, exists := defaultValues[key]
	if exists {
		return fmt.Errorf("failed to reg config with key :%s:%s", key, err)
	}
	defaultValues[key] = c
	return nil
}

// GetConfig 调用者自己负责val的结构, 不支持嵌套的struct
func GetConfig(key string, val interface{}) error {
	confDefault, defaultExists := defaultValues[key]
	if !defaultExists {
		return fmt.Errorf("get config error for key:%s, not reg yet", key)
	}

	confCurrent, currentExists := currentValues[key]

	v := confDefault

	if currentExists {
		t := reflect.TypeOf(val).Elem()
		if t.Kind() != reflect.Struct {
			v = confCurrent
		} else {
			v = mergeValueStruct(confDefault, confCurrent)
		}
	}

	err := json.Unmarshal(v, val)
	if err != nil {
		return fmt.Errorf("failed to parse config for key %s:%s", key, err)
	}
	return nil
}

func mergeValueStruct(_default, _current json.RawMessage) json.RawMessage {
	_d := map[string]json.RawMessage{}
	_c := map[string]json.RawMessage{}
	// won't error
	_ = json.Unmarshal(_default, &_d)
	_ = json.Unmarshal(_current, &_c)

	_v := map[string]json.RawMessage{}
	for k, v := range _d {
		val, exists := _c[k]
		if !exists {
			val = v
		}
		_v[k] = val
	}

	val, _ := json.Marshal(_v)

	return val
}
