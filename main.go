package main

import (
		"log"
		"net/http"
		"demo_gorm/configs/routes"
		"os"
		"io"
)



func main(){
	f, err := os.OpenFile("/Users/unbxd/go-workspace/src/demo_gorm/log/development.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	mw := io.MultiWriter(f, os.Stdout)
	log.SetOutput(mw)
	router := routes.NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}
