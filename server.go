package main

import (
	"fmt"
	"net/http"
	"strings"
)

func HandleRouteServer(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(
		r.Form,
		r.URL.Path,
		r.URL.Scheme,
		r.URL.Host )
		for k,v := range r.Form{
			fmt.Println("key",k)
			fmt.Println("value",strings.Join(v,""))
		}
		fmt.Fprintf(w,"hello World!")
}
