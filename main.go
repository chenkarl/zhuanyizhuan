package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	info, err := getWeatherInfo("110000")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(info)
	place, err := getCityCode("北京市")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(place)
}

func getConfig() (conf map[string]interface{}, err error) {
	var ob interface{}
	filepath := "./config.json"
	data, err := ioutil.ReadFile(filepath)
	err = json.Unmarshal(data, &ob)
	conf = ob.(map[string]interface{})
	return
}
func getWeatherInfo(place string) (info interface{}, err error) {
	config, err := getConfig()
	if err != nil {
		return
	}
	url := config["weatherInfoURL"].(string) + "key=" + config["GaoDeKEY"].(string) + "&city=" + place
	log.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &info)
	return
}

// ErrorString implements Error's String method by returning itself.
type ErrorString string

func (e ErrorString) Error() string { return string(e) }

// NewError converts s to an ErrorString, which satisfies the Error interface.
func NewError(s string) error { return ErrorString(s) }

func getCityCode(name string) (place string, err error) {
	filepath := "./adcode_citycode.csv"
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
		if record[0] == name {
			place = record[1]
			return place, err
		}
	}
	return "", NewError("不存在该城市")
}
