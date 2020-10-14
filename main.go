package main

import (
	"clean-api/api"
)

func main(){
	server := api.HTTPServer{}
	server.Listen()
}
