package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var tempStruct Data //The struct that will be constructed


func main() {
	http.HandleFunc("/data", dataHandler)
	http.HandleFunc("/submit", submitHandler)

	//Initialization of the map used in tempStruct
	tempStruct.CopyAndPaste = make(map[string]bool)
	tempStruct.CopyAndPaste["inputEmail"] = false
	tempStruct.CopyAndPaste["inputCardNumber"] = false
	tempStruct.CopyAndPaste["inputCVV"] = false

	fmt.Println("Server now running on localhost:8080")
	fmt.Println(`Try running: curl -X POST -d '{"WebsiteUrl":"test123"}' http://localhost:8080/data ...`)
	fmt.Println("You can now use the form at '../server/index.html' \n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


type Dimension struct {
	Width  string	`json:"width"`
	Height string	`json:"height"`
}

type Data struct {
	WebsiteUrl         string	`json:"websiteurl,omitempty"`
	SessionId          string	`json:"sessionid,omitempty"`
	ResizeFrom         Dimension	`json:"resizefrom,omitempty"`
	ResizeTo           Dimension	`json:"resizeto,omitempty"`
	CopyAndPaste       map[string]bool	`json:"copyandpaste,omitempty"` // map[fieldId]true
	FormCompletionTime int	`json:"formcompletiontime,omitempty"` // Seconds
}

type reqData struct {
	EventType					 string `json:"eventtype,omitempty"`
	WebsiteUrl         string	`json:"websiteurl,omitempty"`
	SessionId          string	`json:"sessionid,omitempty"`
	ResizeFrom         Dimension	`json:"resizefrom,omitempty"`
	ResizeTo           Dimension	`json:"resizeto,omitempty"`
	FormId						 string			`json:"formid,omitempty"`
	Pasted						 bool				`json:"pasted,omitempty"`
	FormCompletionTime int	`json:"formcompletiontime,omitempty"` // Seconds
}





func dataHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	//log.Printf("Request received: %s, %s", body, err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to read body"))
		return
	}

	req := &reqData{}

	if err = json.Unmarshal(body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to unmarshal JSON request"))
		return
	}
	//log.Printf("route /data called with request: %+v", req)

	//Initialization of the struct under construction
	tempStruct.WebsiteUrl = req.WebsiteUrl
	tempStruct.SessionId = req.SessionId

	//construction of the struct with the JSON received
	switch os := req.EventType; os {
	case "resize":
		log.Printf("last event: %s", os)
		tempStruct.ResizeFrom = req.ResizeFrom
		tempStruct.ResizeTo = req.ResizeTo

	case "copyAndPaste":
		log.Printf("last event: %s", os)
		tempStruct.CopyAndPaste[req.FormId] = req.Pasted
	}

	log.Printf("State of the struct under construction:\n \x1B[33m %+v \n \x1B[0m", tempStruct)
	w.WriteHeader(http.StatusOK)
}




func submitHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	//log.Printf("Request received: %s, %s", body, err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to read body"))
		return
	}

	req := &reqData{}

	if err = json.Unmarshal(body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to unmarshal JSON request"))
		return
	}

	//Initialization of the struct under construction
	tempStruct.WebsiteUrl = req.WebsiteUrl
	tempStruct.SessionId = req.SessionId

	//construction of the struct with the JSON received
	//We only consider the timeTaken event
	switch os := req.EventType; os {
	case "timeTaken":
		log.Printf("last event: %s", os)
		tempStruct.FormCompletionTime = req.FormCompletionTime
	}
	log.Printf("\x1B[32m Button pressed! Here is the struct:\n %+v \n \x1B[0m" , tempStruct)

	w.WriteHeader(http.StatusOK)
}
