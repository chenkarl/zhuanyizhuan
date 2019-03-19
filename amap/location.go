package amap

import "strconv"

var (
	// EE, ES, WS, WW, WN, EN 六点坐标
	EE, ES, WS, WW, WN, EN string
)

// GetLocation lon 经度 lat 纬度 dis 距离 km
// 暂时只支持北半球与东经
func GetLocation(lon float64, lat float64, dis float64) (polygon string) {
	EELon := strconv.FormatFloat(lon+dis/100, 'f', 6, 64)
	EELat := strconv.FormatFloat(lat, 'f', 6, 64)
	EE = EELon + "," + EELat

	ESLon := strconv.FormatFloat(lon+dis/200, 'f', 6, 64)
	ESLat := strconv.FormatFloat(lat-dis/115.6, 'f', 6, 64)
	ES = ESLon + "," + ESLat

	WSLon := strconv.FormatFloat(lon-dis/200, 'f', 6, 64)
	WSLat := strconv.FormatFloat(lat-dis/115.6, 'f', 6, 64)
	WS = WSLon + "," + WSLat

	WWLon := strconv.FormatFloat(lon-dis/100, 'f', 6, 64)
	WWLat := strconv.FormatFloat(lat, 'f', 6, 64)
	WW = WWLon + "," + WWLat

	ENLon := strconv.FormatFloat(lon+dis/200, 'f', 6, 64)
	ENLat := strconv.FormatFloat(lat+dis/115.6, 'f', 6, 64)
	EN = ENLon + "," + ENLat

	WNLon := strconv.FormatFloat(lon-dis/200, 'f', 6, 64)
	WNLat := strconv.FormatFloat(lat+dis/115.6, 'f', 6, 64)
	WN = WNLon + "," + WNLat

	return EE + "|" + ES + "|" + WS + "|" + WW + "|" + WN + "|" + EN
}
