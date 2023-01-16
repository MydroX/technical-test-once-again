package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Data struct {
	Users []map[interface{}]interface{} `yaml:"users"`
}

func main() {
	yfile, err := os.ReadFile("users.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var data Data

	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		log.Fatal(err)
	}

	var emailList []string

	for _, user := range data.Users {

		var firstname string
		var lastname string

		for i, name := range user {
			if fmt.Sprintf("%v", i) == "firstname" {
				firstname = fmt.Sprintf("%v", name)
			} else if fmt.Sprintf("%v", i) == "lastname" {
				lastname = fmt.Sprintf("%v", name)
			} else {
				log.Fatalf("unable to read names")
			}
		}

		if firstname == "" || lastname == "" {
			continue
		}

		emailList = append(emailList, fmt.Sprintf("%s.%s@example.com", firstname, lastname))
	}

	fmt.Println(emailList)
}
