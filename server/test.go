package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


var ck http.Cookie


func main() {

	http.HandleFunc("/data", dataHandler)
	http.HandleFunc("/submit", submitHandler)

	//The cookie initialization
	expires := time.Now().AddDate(0, 0, 1)
	ck = http.Cookie{Name: "DataCookie", Value: "", Expires: expires, HttpOnly: true}

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

type reqData struct { //the struct of the request sent by the frontend
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

	//Unmarshall of the received cookie in tempDataStruct, in order to modify it to send it back
	tempDataStruct := Data{}
	tempDataStruct.CopyAndPaste = make(map[string]bool) //Initialization of the CP map in the struct
	json.Unmarshal([]byte(ck.Value), &tempDataStruct)
	log.Printf("cookie unmarshalled:", tempDataStruct)

	tempDataStruct.WebsiteUrl = req.WebsiteUrl
	tempDataStruct.SessionId = req.SessionId

	//construction of the struct with the JSON request received
	switch os := req.EventType; os {
	case "resize":
		log.Printf("last event: %s", os)

		tempDataStruct.ResizeFrom = req.ResizeFrom
		tempDataStruct.ResizeTo = req.ResizeTo

		log.Printf("cookie unmarshalled and modified:", tempDataStruct)

	case "copyAndPaste":
		log.Printf("last event: %s", os)

		tempDataStruct.CopyAndPaste[req.FormId] = req.Pasted

		log.Printf("cookie unmarshalled and modified:", tempDataStruct)
	}

	//Store the new value of the cookie
	tempFinal, _ := json.Marshal(&tempDataStruct)
	ck.Value = string(tempFinal)
	// write the cookie to response
	log.Printf("cookie sent:", ck)
	http.SetCookie(w, &ck)


	log.Printf("State of the struct under construction:\n \x1B[33m %+v \n \x1B[0m", ck.Value)


	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "null")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

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

	//Unmarshall of the received cookie in tempDataStruct, in order to modify it to send it back
	tempDataStruct := Data{}
	tempDataStruct.CopyAndPaste = make(map[string]bool) //Initialization of the CP map in the struct
	json.Unmarshal([]byte(ck.Value), &tempDataStruct)
	log.Printf("cookie unmarshalled:", tempDataStruct)

	tempDataStruct.WebsiteUrl = req.WebsiteUrl
	tempDataStruct.SessionId = req.SessionId

	//construction of the struct with the JSON received
	//We only consider the timeTaken event
	switch os := req.EventType; os {
	case "timeTaken":
		log.Printf("last event: %s", os)

		tempDataStruct.FormCompletionTime = req.FormCompletionTime

		log.Printf("cookie unmarshalled and modified:", tempDataStruct)
	}

	//Store the new value of the cookie
	tempFinal, _ := json.Marshal(&tempDataStruct)
	ck.Value = string(tempFinal)
	// write the cookie to response
	log.Printf("cookie sent:", ck)
	http.SetCookie(w, &ck)

	log.Printf("\x1B[32m Button pressed! Here is the final struct:\n %+v \n \x1B[0m" , ck.Value)


	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "null")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.WriteHeader(http.StatusOK)
}
