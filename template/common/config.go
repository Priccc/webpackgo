package common

import (
	// "encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	json "github.com/laktak/hjson-go"
)

func LoadConfig(file string) (IConfigInfo, error) {
	bs, err := GetFileContent(file)
	if err != nil {
		return nil, err
	}

	return ParseConfig(bs)
}

func ParseConfig(data []byte) (IConfigInfo, error) {
	var conf ConfigInfo
	err := json.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

type IConfigInfo interface {
	GetListeningPort() string
	Set(string, interface{}) error
	String() string
}

type ConfigInfo map[string]interface{}


func (conf ConfigInfo) String() string {
	r := ""
	for k, v := range conf {
		r = fmt.Sprintf("%s\r\n%s -> %s", r, k, v)
	}
	return r
}

func (conf ConfigInfo) Set(key string, val interface{}) error {
	if _, exists := conf[key]; exists {
		conf[key] = val
	} else {
		return errors.New("no such key")
	}
	return nil
}

func (conf ConfigInfo) GetListeningPort() string {
	if v, exists := conf["listeningPort"]; exists {
		return v.(string)
	}
	return ""
}

// get string from text file
func GetFileContent(file string) ([]byte, error) {
	if !IsFileExist(file) {
		return nil, os.ErrNotExist
	}
	b, e := ioutil.ReadFile(file)
	if e != nil {
		return nil, e
	}
	return b, nil
}

// exists returns whether the given file or directory exists or not
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
