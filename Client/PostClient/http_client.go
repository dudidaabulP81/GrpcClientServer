package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please provide client id as a program parameter")
		os.Exit(1)
	}

	clientId := os.Args[1]
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++{
		func() {
			currentMeasurement := rand.Intn(45)
			secTillNextMeasurement := rand.Intn(5) + 1
			formData := url.Values{}
			formData.Add("id", clientId)
			formData.Add("temperature", strconv.Itoa(currentMeasurement))
			time.Sleep(time.Duration(secTillNextMeasurement) * time.Second)
			resp, err := http.PostForm("http://localhost:8000/set", formData)
			if err != nil {
				log.Fatalln(err)
			}
			defer resp.Body.Close()
		}()
	}
}
