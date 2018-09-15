package main

import (

	"fmt"
	"net/http"
	"html"

)


func main() {
	// Listen on port 9200
	http.HandleFunc("/", handler)
	fmt.Println("# Listening on 9200")
	http.ListenAndServe(":9200", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}


func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}

}
