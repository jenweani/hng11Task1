package main

import(
	"hng11task1/api"
)

func main(){
	srv := api.NewServer(8888, api.BuildRoutesHandler())
	srv.Listen()
}