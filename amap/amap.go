package amap

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// GetWeatherInfo 获取天气信息
func GetWeatherInfo(place string) (info interface{}, err error) {
	config, err := GetConfig()
	if err != nil {
		return
	}
	url := config["amapApiHost"].(string) + config["weatherInfoURL"].(string) + "key=" + config["GaoDeKEY"].(string) + "&city=" + place
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

// GetPolygon 获取多边形经纬度之间城市
func GetPolygon(polygon string) (city interface{}, err error) {
	config, err := GetConfig()
	if err != nil {
		return
	}
	url := config["amapApiHost"].(string) + config["polygonURL"].(string) + "key=" + config["GaoDeKEY"].(string) + "&polygon=" + polygon + "&types=190103|190104&extensions=base"
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
	err = json.Unmarshal(body, &city)
	return
}
