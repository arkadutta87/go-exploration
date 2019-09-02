package restclient

import (
	"bytes"
	"fmt"
	"net/http"
)

const (
	smsEndpoint string = "https://api.equence.in/xmlcasting"
)

//SendSms ...
func SendSms(smsBody string) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", smsEndpoint, bytes.NewBuffer([]byte(smsBody)))
	if err != nil {
		fmt.Println(err)
	}

	// you can then set the Header here
	// I think the content-type should be "text/xml" like json...
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Accept", "application/json")
	// now POST it
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}
