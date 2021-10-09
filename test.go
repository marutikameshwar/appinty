package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	Id       int   `json:"Id"`
	Username string `json:"Username"`
	Emailid  string `json:"Emailid"`
	Password string `json:"Password"`

}

type UserPost struct{
	Id int `json:"Id"`
	UserId int `json:"UserId"`
	Caption string `json:"Caption"`
	Image string `json:"Image"`
	TimeStamp string`json:"TimeStamp"`

}

var Users []User
var UserPosts []UserPost


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func userdata(){
	Users=[]User{
		User{Id: 100,Username: "maruti",Emailid: "maruti@gmail.com",Password: "123"},
		User{Id: 200,Username: "john",Emailid: "john@yahoo.com",Password: "567"},
	}

}

func userpost(){
	UserPosts=[]UserPost{
		{Id:1,Image: "https://miro.medium.com/max/1200/1*mk1-6aYaf_Bes1E3Imhc0A.jpeg",Caption: "babay yodha",TimeStamp: "10/34/56" ,UserId: 100},
		{Id: 2,Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSlrZqTCInyg6RfYC7Ape20o-EWP1EN_A8fOA&usqp=CAU",Caption: "cute puppy",TimeStamp: "10/34/56",UserId: 200},
	}
}

func returnAllUsers(w http.ResponseWriter, r *http.Request){
	var U User
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields() // catch unwanted fields

	// anonymous struct type: handy for one-time use
	err := d.Decode(&U)
	if err != nil {
		// bad JSON or unrecognized json field
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// got the input we expected: no more, no less
	log.Println(U)
}

func returnAllPosts(w http.ResponseWriter, r *http.Request){
	var U UserPost
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields() // catch unwanted fields

	// anonymous struct type: handy for one-time use
	err := d.Decode(&U)
	if err != nil {
		// bad JSON or unrecognized json field
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// got the input we expected: no more, no less
	log.Println(U.Image)
}

func returnuserbyid(w http.ResponseWriter, r *http.Request){
	var url=r.URL.Path
	var arr=strings.Split(url,"/")
	//fmt.Println(arr[2])
	var temp =-1
	for i:=0;i< len(Users);i++{
		if strconv.Itoa(Users[i].Id)==arr[2] {
			fmt.Println(Users[i].Id)
			temp=i
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "  ")
			encoder.Encode(Users[temp])
			break
		}
	}
	if(temp!=-1){
		fmt.Fprintf(w, "No data found")
	}

}

func returnpostbyid(w http.ResponseWriter, r *http.Request){
	var url=r.URL.Path
	var arr=strings.Split(url,"/")
	//fmt.Println(arr[2])
	var temp =-1
	for i:=0;i< len(UserPosts);i++{
		if strconv.Itoa(UserPosts[i].Id)==arr[2] {
			fmt.Println(UserPosts[i].Id)
			temp=i
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "  ")
			encoder.Encode(UserPosts[temp])
			break
		}
	}
	if(temp!=-1){
		fmt.Fprintf(w, "No data found")
	}

}

func returnallusererposts(w http.ResponseWriter, r *http.Request){
	var url=r.URL.Path
	var arr=strings.Split(url,"/")
	//fmt.Println(arr[3])
	var temp=-1
	var logic=false
	for i:=0;i<len(UserPosts);i++{
		if strconv.Itoa(UserPosts[i].UserId)==arr[3]{
			fmt.Println(UserPosts[i].Id)
			temp=i
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "  ")
			encoder.Encode(UserPosts[temp])
			logic=true
		}
	}
	if logic==true{
		fmt.Fprintf(w, "invalid user/no posts by user")
	}

}
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users",returnAllUsers)
	http.HandleFunc("/users/",returnuserbyid)
	http.HandleFunc("/posts",returnAllPosts)
    http.HandleFunc("/posts/",returnpostbyid)
	http.HandleFunc("/posts/users/",returnallusererposts)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	userdata()
	userpost()
	handleRequests()
}