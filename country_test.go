package gocountry

import (
	"fmt"
	"testing"
)

func TestTextSearch001(t *testing.T) {
	options := &Options{
		Full:      true, // search full country name
		CodeTwo:   true, // search for 2 character country code
		CodeThree: true, // search for 3 character country code
		Number:    true, //search for country number. (MUST BE 0 Padded)
	}

	countries := Search(options, "Example full countries: Benin, azerbaijan. Example three letter country code: BOL. Example two letter country code: BR Example number: 408")
	for _, country := range countries {
		fmt.Printf("Found country in text: %s, %s, %s, %s\n", country.Full(), country.CodeTwo(), country.CodeThree(), country.Number())
	}
}

func TestTextSearch002(t *testing.T) {
	countries := Search(&Options{Full: true}, "Example full countries: Benin, azerbaijan. Example three letter country code: BOL. Example two letter country code: BR Example number: 408")
	for _, country := range countries {
		fmt.Printf("Found country in text: %s, %s, %s, %s\n", country.Full(), country.CodeTwo(), country.CodeThree(), country.Number())
	}
}

func TestTextSearch003(t *testing.T) {
	countries := Search(&Options{CodeTwo: true}, "Example full countries: Benin, azerbaijan. Example three letter country code: BOL. Example two letter country code: BR Example number: 408")
	for _, country := range countries {
		fmt.Printf("Found country in text: %s, %s, %s, %s\n", country.Full(), country.CodeTwo(), country.CodeThree(), country.Number())
	}
}

func TestTextSearch004(t *testing.T) {
	countries := Search(&Options{CodeThree: true}, "Example full countries: Benin, azerbaijan. Example three letter country code: BOL. Example two letter country code: BR Example number: 408")
	for _, country := range countries {
		fmt.Printf("Found country in text: %s, %s, %s, %s\n", country.Full(), country.CodeTwo(), country.CodeThree(), country.Number())
	}
}

func TestTextSearch005(t *testing.T) {
	countries := Search(&Options{Number: true}, "Example full countries: Benin, azerbaijan. Example three letter country code: BOL. Example two letter country code: BR Example number: 408")
	for _, country := range countries {
		fmt.Printf("Found country in text: %s, %s, %s, %s\n", country.Full(), country.CodeTwo(), country.CodeThree(), country.Number())
	}
}

func TestTextSearch006(t *testing.T) {
	countries := Search(&Options{}, "Example full countries: Benin, azerbaijan. Example three letter country code: BOL. Example two letter country code: BR Example number: 408")
	for _, country := range countries {
		fmt.Printf("Found country in text: %s, %s, %s, %s\n", country.Full(), country.CodeTwo(), country.CodeThree(), country.Number())
	}
}

func TestNumberSearch001(t *testing.T) {
	country := FindByNumber(4)
	if country != nil {
		fmt.Printf("Found country by number: %s, %s, %s, %s\n", country.Full(), country.CodeTwo(), country.CodeThree(), country.Number())
	}
}
