package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"./amap"
)

var (
	City, Distance string
)

func main() {
	fmt.Printf("请输入输入查询的城市与距离（km）：")
	_, err := fmt.Scanln(&City, &Distance)
	if err != nil {
		log.Fatalln(err)
	}
	district, err := amap.GetDistrict(City)
	if err != nil {
		log.Fatalln(err)
	}
	center := district["center"].(string)
	centerArr := strings.Split(center, ",")
	lon, err := strconv.ParseFloat(centerArr[0], 64)
	lat, err := strconv.ParseFloat(centerArr[1], 64)
	dis, err := strconv.ParseFloat(Distance, 64)
	if err != nil {
		log.Fatalln(err)
	}
	ploy := amap.GetLocation(lon, lat, dis)
	log.Println(ploy)
	city, err := amap.GetPolygon(ploy)
	if err != nil {
		log.Fatalln(err)
	}
	for _, value := range city {
		name := value.(map[string]interface{})["name"]
		place, err := amap.GetCityCode(name.(string))
		if err != nil {
			log.Fatalln(err)
			continue
		}
		info, err := amap.GetWeatherInfo(place)
		if err != nil {
			log.Fatalln(err)
			continue
		}
		if !strings.Contains(info.([]interface{})[0].(map[string]interface{})["weather"].(string), "雨") {
			log.Println("可", name, info)
		}
		if strings.Contains(info.([]interface{})[0].(map[string]interface{})["weather"].(string), "雨") {
			log.Println("不可", name, info)
		}
	}
}
