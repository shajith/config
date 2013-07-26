package config

import (
	"io/ioutil"
	"launchpad.net/goyaml"
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

func (c Config) Get(path ...interface{}) (val interface{}) {
	var tmp interface{}
	tmp = map[interface{}]interface{}(c)
	for _, key := range path {
		if tmp == nil {
			val = nil
			return
		}

		switch tmp.(type) {
		case map[interface{}]interface{}:
			tmp = tmp.(map[interface{}]interface{})[key.(string)]
		case []interface{}:
			idx := key.(int)
			arr := tmp.([]interface{})
			if idx < len(arr) {
				tmp = arr[idx]
			} else {
				val = nil
				return
			}
		default:
			val = nil
			return
		}

	}
	val = tmp
	return
}
