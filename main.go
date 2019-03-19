package main

import (
	"fmt"
	"log"
)

var (
	City, Distance string
)

func main() {
	fmt.Printf("请输入输入查询的城市：")
	_, err := fmt.Scanln(&City, &Distance)
	if err != nil {
		log.Fatalln(err)
	}

	// place, err := amap.GetCityCode("北京市")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// info, err := amap.GetWeatherInfo(place)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// city, err := amap.GetPolygon("116,39|117,40")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}
