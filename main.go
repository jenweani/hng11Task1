package main

import (
	"hng11task1/api"
	// "hng11task1/configs"
)

func main(){
	// configs.Load()
	srv := api.NewServer(8080, api.BuildRoutesHandler())
	srv.Listen()
}