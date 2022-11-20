package main

import "fmt"

func main() {

	lat := 35.62235961
	lng := 139.7273458

	lonW := lng - (lat * 0.000046038) - lng*0.000083043 + 0.010040
	latW := lat - (lat * 0.00010695) + lng*0.000017464 + 0.0046017

	fmt.Println(lonW, latW)
}
