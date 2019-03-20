package amap

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type keyvalue map[string]interface{}

var (
	offset, page, count uint64
)

// GetURL ...
func GetURL(url string) (ret map[string]interface{}, err error) {
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
	err = json.Unmarshal(body, &ret)
	if ret["status"] != "1" {
		err = NewError(ret["info"].(string))
		return
	}
	return
}

// GetWeatherInfo 获取天气信息
func GetWeatherInfo(place string) (info interface{}, err error) {
	config, err := GetConfig()
	if err != nil {
		return
	}
	url := config["amapApiHost"].(string) + config["weatherInfoURL"].(string) + "key=" + config["GaoDeKEY"].(string) + "&city=" + place
	ret, err := GetURL(url)
	if err != nil {
		return
	}
	info = ret["lives"].(interface{})
	return
}

// GetPolygon 获取多边形经纬度之间城市
func GetPolygon(polygon string) (city []interface{}, err error) {
	offset = 20
	page = 1
	count = 21
	config, err := GetConfig()
	if err != nil {
		return
	}
	// https://restapi.amap.com/v3/place/polygon?polygon=116.460988,40.006919|116.48231,40.007381|116.47516,39.99713|116.472596,39.985227|116.45669,39.984989|116.460988,40.006919&keywords=kfc&output=xml&key=<用户的key>
	// 190103 直辖市级地名|190104地市级地名|190105区县级地名
	url := config["amapApiHost"].(string) + config["polygonURL"].(string) + "key=" + config["GaoDeKEY"].(string) + "&polygon=" + polygon + "&types=190103|190104&extensions=base"
	for offset*page < count {
		url = url + "&offset=" + strconv.FormatUint(offset, 10) + "&page=" + strconv.FormatUint(page, 10)
		ret, errURL := GetURL(url)
		if errURL != nil {
			return nil, errURL
		}
		var errPar error
		count, errPar = strconv.ParseUint(ret["count"].(string), 10, 64)
		if errPar != nil {
			return nil, errPar
		}
		page++
		city = append(city, ret["pois"].([]interface{})...)
	}
	return
}

// GetDistrict 获取城市信息
func GetDistrict(keywords string) (district map[string]interface{}, err error) {
	config, err := GetConfig()
	if err != nil {
		return
	}
	//https://restapi.amap.com/v3/config/district?keywords=北京&subdistrict=2&key=<用户的key>
	url := config["amapApiHost"].(string) + config["districtURL"].(string) + "key=" + config["GaoDeKEY"].(string) + "&keywords=" + keywords + "&subdistrict=1&extensions=base"
	ret, err := GetURL(url)
	if err != nil {
		return
	}
	district = ret["districts"].([]interface{})[0].(map[string]interface{})
	return
}
