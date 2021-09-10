package main

import (
	"fmt"
	"log"
)

func main() {

	baseUrl := envVariable("BASE_URL")
	token := envVariable(("AUTHENTICATION_TOKEN"))


	required_template, err := getTemplate(baseUrl, token)

	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(required_template, "required template")

	required_contact, err := getContact(baseUrl, token)
	fmt.Println(required_contact,"required_contact")

	if err != nil {
		log.Fatalln(err.Error())
	}

	mergeTemplateWithContact(required_template,required_contact)
	

}
