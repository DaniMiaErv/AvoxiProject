package main

import (
	CheckIp "github.com/DaniMiaErv/AvoxiCheckIP"
)

func main() {

	testIp := "1"
	countryList := []string{"United States", "Germany", "Japan"}

	println(CheckIp.CheckIPWhiteList(testIp, countryList))

}
