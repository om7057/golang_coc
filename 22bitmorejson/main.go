package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //This is an alias
	Price    int
	Platfrom string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LCO.in", "abc1@!23", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LCO.in", "aqc1@!23", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 99, "LCO.in", "abqw1@!23", nil},
	}
	//package this data as JSON data

	// finalJson, err := json.Marshal(lcoCourses)
	// finalJson, err := json.MarshalIndent(lcoCourses, "lco", "\t")
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "website": "LCO.in",
                "tags": ["web-dev","js"] 
    }

	`)

	var lcoCourse course 

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON ")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Not JSON")
	}

	//Extracting as key value pairs 

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Println("%#v\n", myOnlineData)

	for k, v:= range myOnlineData {
		fmt.Printf("key is %v and value is %v and type of data is %T\n",k , v, v)
	}

}
