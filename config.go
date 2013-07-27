package config

import (
	"io/ioutil"
	"launchpad.net/goyaml"
	"errors"
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
				return nil, errors.New("Couldn't fetch path")
			}
		default:
			return nil, errors.New("Couldn't fetch path")
		}

	}
	if(tmp != nil) {
		return tmp, nil
	} else {
		return nil, errors.New("Couldn't fetch path")
	}
}
