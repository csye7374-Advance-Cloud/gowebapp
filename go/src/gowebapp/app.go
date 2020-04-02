package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
)

func getCurrentTime(response http.ResponseWriter, request *http.Request) {
	currentTime := time.Now()
	fmt.Fprintf(response, currentTime.String())
}

func getTimeInZones(response http.ResponseWriter, request *http.Request){
	queryValues := request.URL.Query()
	timeInZones := make(map[string]string)
	currentTime := time.Now()
	for k, v := range queryValues {
        	fmt.Println("key:", k, "value:", v[0])
			location, err := time.LoadLocation(strings.ToUpper(v[0]))
			fmt.Println("Location : ",location);
    		if err != nil {
        		fmt.Println(err)
				timeInZones[v[0]]=err.Error()
    		}else{
				timeInZones[location.String()]=currentTime.In(location).String()
		  		fmt.Println("ZONE : ", location, " Time : ", currentTime.In(location))
			}
	}
	json.NewEncoder(response).Encode(timeInZones)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getCurrentTime)
	router.HandleFunc("/time", getTimeInZones)
	fmt.Println("Go Server Started at port 8000..")
	log.Fatal(http.ListenAndServe(":8000", router))
}
