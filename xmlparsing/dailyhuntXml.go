package xmlparsing

import (
	"encoding/xml"
	"fmt"
	"os"
)

//Address ...
type Address struct {
	City, State string
}

const (
	ver string = "1.2"
)

//Person ...
type Person struct {
	XMLName   xml.Name `xml:"person"`
	ID        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
	Height    float32  `xml:"height,omitempty"`
	Married   bool
	Address
	Comment string `xml:",comment"`
}

//AddressCode ...
type AddressCode struct {
	XMLName xml.Name `xml:"ADDRESS"`
	FROM    string   `xml:"FROM,attr"`
	TO      string   `xml:"TO,attr"`
	SEQ     int      `xml:"SEQ,attr"`
}

//USER ...
type USER struct {
	XMLName  xml.Name `xml:"USER"`
	USERNAME string   `xml:"USERNAME,attr"`
	PASSWORD string   `xml:"PASSWORD,attr"`
}

//SMS ...
type SMS struct {
	XMLName xml.Name `xml:"SMS"`
	Address AddressCode
	TEXT    string `xml:"TEXT,attr"`
	ID      int    `xml:"ID,attr"`
	NAI     int    `xml:"NAI,attr"`
	SPLIT   int    `xml:"SPLIT,attr"`
	CONCAT  int    `xml:"CONCAT,attr"`
}

//MESSAGE ...
type MESSAGE struct {
	XMLName xml.Name `xml:"MESSAGE"`
	VER     string   `xml:"VER,attr"`
	User    USER
	Sms     SMS
}

//SendMailXMLPreparation ...
func SendMailXMLPreparation(secretCode string) string {
	v := &MESSAGE{VER: ver}

	smsUserName := os.Getenv("SMS_USERNAME")
	smsPassword := os.Getenv("SMS_PASSWORD")
	smsMobileNo := os.Getenv("SMS_MOBILE_NUMBER")
	v.User = USER{USERNAME: smsUserName, PASSWORD: smsPassword}

	smsMessage := secretCode + " is the OTP to register your mobile number with dailyhunt for the rewards program"

	v.Sms = SMS{TEXT: smsMessage,
		ID: 1, NAI: 0, SPLIT: 0, CONCAT: 1}
	v.Sms.Address = AddressCode{FROM: "DLYHNT", TO: smsMobileNo, SEQ: 1}

	output, err := xml.Marshal(v)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}

	str2 := string(output)
	fmt.Println("xml: ", str2)

	return str2
}

//ExampleMarshalIndent ...
func ExampleMarshalIndent() {
	v := &Person{ID: 99, FirstName: "Arka", LastName: "Dutta", Age: 32}
	v.Comment = "Need more details"
	v.Address = Address{City: "Bengaluru", State: "Karnataka"}

	output, err := xml.Marshal(v)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}

	// dt2 := Person{}

	// xml.Unmarshal(output, &dt2)

	// os.Stdout.Write(output)
	str2 := string(output)
	fmt.Println("xml: ", str2)
	// fmt.Println(output)
}
