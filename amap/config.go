package amap

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

// ErrorString implements Error's String method by returning itself.
type ErrorString string

func (e ErrorString) Error() string { return string(e) }

// NewError converts s to an ErrorString, which satisfies the Error interface.
func NewError(s string) error { return ErrorString(s) }

// GetConfig get config
func GetConfig() (conf map[string]interface{}, err error) {
	var ob interface{}
	filepath := "./config/config.json"
	data, err := ioutil.ReadFile(filepath)
	err = json.Unmarshal(data, &ob)
	conf = ob.(map[string]interface{})
	return
}

// GetCityCode get citycode
func GetCityCode(name string) (place string, err error) {
	filepath := "./data/adcode_citycode.csv"
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	reader := csv.NewReader(strings.NewReader(string(data[:])))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if strings.Contains(record[0], name) {
			place = record[1]
			return place, err
		}
	}
	return "", NewError("不存在该城市")
}
