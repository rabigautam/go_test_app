package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"os"
)


func doGetRequest(url string, token string) ([]byte, error) {
	//request to given url
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}

	//add headers to request

	req.Header.Add("Authorization", token)
	req.Header.Add("Content_type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return body, err
}

// Sending Email Using Smtp in Golang

// Main function
func sendMail(from string, password string, toList []string, msg string) {

	host := "smtp.gmail.com"

	// Its the default port of smtp server
	port := "587"

	body := []byte(msg)

	// to act as username.
	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	// handling the errors
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully sent mail to all user in toList")
}
