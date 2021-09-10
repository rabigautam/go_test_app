package main

import (
	"encoding/json"
	"log"
)

func getContact(baseUrl string, token string) (ContactResponse, error) {

	contactUrl := baseUrl + "contact/?limit=11"
	body, err := doGetRequest(contactUrl, token)
	//request to given url

	if err != nil {
		log.Fatalf(err.Error())
		return ContactResponse{}, err
	}
	var contactResponse ContactResponse
	json.Unmarshal(body, &contactResponse)
	var required_contact ContactResponse
	
	required_contact = contactResponse

	return required_contact, nil
}
