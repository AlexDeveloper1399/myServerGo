package main

import (
	"log"
	"net/http"
)

func main(){
http.HandleFunc("/",HandleRouteServer)
err := http.ListenAndServe(":9000",nil)
if err != nil{
	log.Fatal("ListenAndServe",err)
}

}
