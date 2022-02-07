package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	measurementDuration := flag.String("duration", "day", "measurements duration")
	flag.Parse()
	var url string = "http://localhost:8000/get?period=" + *measurementDuration
	fmt.Println("request = %s", url)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
