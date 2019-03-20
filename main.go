package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/chenkarl/zhuanyizhuan/amap"
)

type PlaceWeather struct {
	place   string
	weather string
}
type PlaceWeatherArr struct {
	placeweather []PlaceWeather
}

var (
	city, distance string
)

func main() {
	RunServer()
	// fmt.Printf("请输入输入查询的城市与距离（km）：")
	// _, err := fmt.Scanln(&City, &Distance)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 获取参数
	r.ParseForm()
	if len(r.Form["city"]) > 0 {
		city = r.Form["city"][0]
	}
	if len(r.Form["distance"]) > 0 {
		distance = r.Form["distance"][0]
	}
	// 获取区域信息
	district, err := amap.GetDistrict(city)
	if err != nil {
		log.Fatalln(err)
	}
	// 获取该区域的经纬度
	center := district["center"].(string)
	centerArr := strings.Split(center, ",")
	lon, err := strconv.ParseFloat(centerArr[0], 64)
	lat, err := strconv.ParseFloat(centerArr[1], 64)
	dis, err := strconv.ParseFloat(distance, 64)
	if err != nil {
		log.Fatalln(err)
	}
	// 获取中心点距离内的城市列表
	ploy := amap.GetLocation(lon, lat, dis)
	city, err := amap.GetPolygon(ploy)
	if err != nil {
		log.Fatalln(err)
	}
	// 获取列表内城市的天气信息
	var pwArr []PlaceWeather
	for _, value := range city {
		name := value.(map[string]interface{})["name"]
		place, err := amap.GetCityCode(name.(string))
		if err != nil {
			log.Fatalln(err)
			continue
		}
		infoArr, err := amap.GetWeatherInfo(place)
		info := infoArr.([]interface{})[0]
		if err != nil {
			log.Fatalln(err)
			continue
		}
		var pw PlaceWeather
		if !strings.Contains(info.(map[string]interface{})["weather"].(string), "雨") {
			pw.place = name.(string)
			pw.weather = info.(map[string]interface{})["weather"].(string)

			// fmt.Fprint(w, "可", name, info)
			// log.Println("可", name, info)
		}
		if strings.Contains(info.(map[string]interface{})["weather"].(string), "雨") {
			pw.place = name.(string)
			pw.weather = info.(map[string]interface{})["weather"].(string)
		}
		pwArr = append(pwArr, pw)
	}
	w.Header().Set("Content-Type", "application/json")
	log.Println(pwArr)
	j, err := json.Marshal(pwArr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(j))
	fmt.Fprint(w, string(j))
	//w.Write(string(j))
}
func RunServer() {
	handler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/hello", &handler)
	server.ListenAndServe()
}
