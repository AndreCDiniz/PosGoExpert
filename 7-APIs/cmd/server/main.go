package main

import "github.com/AndreCDiniz/PosGoExpert/7-APIs/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
