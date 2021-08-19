package ssh

import (
	"io/ioutil"
	"os"
	"strings"
)

type fromFileManager struct {
	list []string
}

func (f *fromFileManager) GetList() []string {
	return f.list
}

func NewManagerFromFile(file string) (Manager, error) {
	data, err := ioutil.ReadFile(file)
	if os.IsNotExist(err) {
		err = ioutil.WriteFile(file, []byte{},0755)
		if err != nil {
			return nil, err
		}
		return &fromFileManager{list: []string{}}, nil
	} else if err != nil {
		return nil, err
	}
	return &fromFileManager{list: strings.Split(string(data), "\n")}, nil
}
