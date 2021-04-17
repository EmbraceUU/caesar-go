package load

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Json struct {
	filepath string
}

func NewJson(filepath string) *Json {
	return &Json{
		filepath: filepath,
	}
}

func (j *Json) JsonMarshall(result interface{}) error {
	body, err := ioutil.ReadFile(j.filepath)
	if err != nil {
		return fmt.Errorf("read file err %v", err.Error())
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return fmt.Errorf("marshall err %v", err.Error())
	}
	return nil
}
