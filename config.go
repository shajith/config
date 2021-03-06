package config

import (
	"io/ioutil"
	"launchpad.net/goyaml"
	"errors"
	"fmt"
)

type Config map[interface{}]interface{}

func New(path string) Config {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var m map[interface{}]interface{}

	goyaml.Unmarshal(b, &m)
	return m
}

func makeError(path []interface{}) error {
	return errors.New(fmt.Sprintf("config: couldn't get path '%v'", path))
}

func (c Config) Get(path ...interface{}) (val interface{}, err error) {
	var tmp interface{}
	tmp = map[interface{}]interface{}(c)
	for _, key := range path {
		switch tmp.(type) {
		case map[interface{}]interface{}:
			tmp = tmp.(map[interface{}]interface{})[key.(string)]
		case []interface{}:
			idx := key.(int)
			arr := tmp.([]interface{})
			if idx < len(arr) {
				tmp = arr[idx]
			} else {
				return nil, makeError(path)
			}
		default:
			return nil, makeError(path)
		}

	}
	if(tmp != nil) {
		return tmp, nil
	} else {
		return nil, makeError(path)
	}
}
