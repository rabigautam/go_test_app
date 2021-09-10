package main

import (
	// "fmt"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"strings"

	// "strings"

	// "encoding/json"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

/**
*to load the env variable
 */

func envVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error occurs while loading .env")
	}
	return os.Getenv((key))
}

type Template struct {
	Template_id         string
	Template_name       string
	Template_create     string
	Template_update     string
	Template_html       string
	Template_text       string
	Is_removable        string
	Template_screenshot string
}
type Contact struct {
	Contact_id     string
	Contact_name   string
	Contact_create string
	Contact_update string
	Is_removable   string
}
type TemplateResponse struct {
	Result []Template
}
type ContactResponse struct {
	Result []Contact
}

func main() {

	baseUrl := envVariable("BASE_URL")
	token := envVariable(("AUTHENTICATION_TOKEN"))

	from := envVariable("MAIL")
	password := envVariable("PASWD")

	fmt.Println("from :"+from)

	templateUrl := baseUrl + "template/"
	body, err := doGetRequest(templateUrl, token)
	//request to given url

	if err != nil {
		log.Fatalf(err.Error())
	}
	var templateResponse TemplateResponse
	json.Unmarshal(body, &templateResponse)
	result := templateResponse.Result

	var required_template Template

	for _, value := range result {

		if value.Template_id == "17" {
			// Found!
			required_template = value

		}
	}

	contactUrl := baseUrl + "contact/?limit=11"
	body, err = doGetRequest(contactUrl, token)
	//request to given url

	if err != nil {
		log.Fatalf(err.Error())
	}
	var contactResponse ContactResponse
	json.Unmarshal(body, &contactResponse)
	resultContact := contactResponse.Result

	for key, contact := range resultContact {
		msg := strings.ReplaceAll(required_template.Template_text, "{{contact_name}}", contact.Contact_name)
		fmt.Println(key, "_", contact.Contact_create)

		sendMail(from, password, []string{contact.Contact_name}, msg)
	}

}

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
