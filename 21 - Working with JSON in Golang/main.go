package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseame"`
	Price    int
	Patform  string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"ReactJs", 299, "OakshadeEducation", "abc123", []string{"web-dev", "js"}},
		{"MERN", 399, "OakshadeEducation", "bcd123", []string{"web-dev", "js"}},
		{"Angular", 299, "OakshadeEducation", "xyz123", nil},
	}

	// Package this data as JSON Data:
	// finalJson, err := json.Marshal(courses)
	finalJson, err := json.MarshalIndent(courses, "", "\t") // Cleaner output
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
    [    
		{
                "courseame": "ReactJs",
                "Price": 299,
                "website": "OakshadeEducation",
                "tags": [
                        "web-dev",
                        "js"
                ]
        },
        {
                "courseame": "MERN",
                "Price": 399,
                "website": "OakshadeEducation",
                "tags": [
                        "web-dev",
                        "js"
                ]
        },
        {
                "courseame": "Angular",
                "Price": 299,
                "website": "OakshadeEducation"
        }
	]
	`)

	// var onlinecourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON Valid")
		// json.Unmarshal(jsonDataFromWeb, &onlinecourse)
		// fmt.Printf("%#v\n", onlinecourse)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	// Some cases where we just want to add some data to key value. We don't use struct in this case

	var myOnlineData interface{}
	err := json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	if err != nil {
		fmt.Println("Error: ", err)

	} else {
		courses := myOnlineData.([]interface{})
		for _, course := range courses {
			courseMap := course.(map[string]interface{})
			fmt.Printf("%#v\n", courseMap)
		}
	}
}
