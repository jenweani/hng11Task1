package main

import (
	"hng11task1/api"
	"hng11task1/configs"
)

func main(){
	configs.Load()
	srv := api.NewServer(8888, api.BuildRoutesHandler())
	srv.Listen()
}