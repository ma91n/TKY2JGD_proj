package main

import (
	"fmt"
	"github.com/everystreet/go-proj/v8/proj"
	"github.com/golang/geo/s1"
)

func main() {
	wgsLng, wgsLat := Tky2Wgs(128542740/float64(60*60*256), 32706756/float64(60*60*256))
	fmt.Printf("%f %f\n", wgsLng, wgsLat) // 139.474598 35.492354
}

func Tky2Wgs(lng, lat float64) (float64, float64) {
	coord := proj.LP{
		Lng: s1.Angle(lng),
		Lat: s1.Angle(lat),
	}

	err := proj.CRSToCRS(
		"EPSG:4301", // 日本測地系(TOKYO) 緯度経度
		"EPSG:4326", // 世界測地系(WGS84) 緯度経度
		func(pj proj.Projection) {
			proj.TransformForward(pj, &coord)
		})
	if err != nil {
		panic(err) // サンプルコードなのでpanicにしていますが、errorを戻り値にした方が良いです
	}
	return float64(coord.Lng), float64(coord.Lat)
}
