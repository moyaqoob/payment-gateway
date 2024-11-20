package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fmt.Println("a basic app that provides a payment gateway")

	http.HandleFunc("/create-payment",handlerequest)
	http.HandleFunc("/",handleHealth)
	log.Println("Listening on localhost:4242")
	var err = http.ListenAndServe("localhost:4242",nil)
	if err!=nil{
		log.Fatal(err)
	}
}

func handlerequest(w http.ResponseWriter, r *http.Request){
	if r.Method=="POST"{
		response:=[]byte("server is uprunning successfully")
		_,err := w.Write(response)
		if err!=nil{
			fmt.Println("Something bad happen")
		}
	}
http.Error(w,"Method invalid",http.StatusMethodNotAllowed)	
}

func handleHealth(writer http.ResponseWriter,request *http.Request){
	response:=[]byte("server is running successfully less fncking gooo")

	_,err:=writer.Write(response)
	if err!=nil{
		fmt.Println(err)
	}
}