package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// cards := newDeck()
	// cards.saveToFile("mycards")
	// cards := getDeckFromFile("mycards")
	// cards.print()

	// Struct
	var sudhanshu = person{
		firstName: "Sudhanshu",
		lastname:  "Patel",
		contactInfo: contactInfo{
			email:   "query.b2cs@gmail.com",
			zipcode: 214234,
		},
	}

	sudhanshu.updateFirstName("Sud")
	sudhanshu.print()

	fmt.Println("------------> MAP<-------------")
	colors := map[string]string{
		"red":   "0ff0000",
		"green": "0f2345",
	}
	for color, hex := range colors {
		fmt.Println(color, "->", hex)
	}

	// //var colors map[string]string = make(map[string]string)
	// colors := make(map[string]string)
	// colors["red"] = "234235"
	// fmt.Println(colors)
	// delete(colors, "red")
	// fmt.Println(colors)

	fmt.Println("--------------------->Interface<----------------------")
	eb := englishBot{}
	hb := hinglishBot{}

	printGreeting(eb)
	printGreeting(hb)

	fmt.Println("------------------> Rest Call <-------------------")
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error - ", err)
		os.Exit(1)
	}
	bs := make([]byte, 99999) // created byte slice of size 99999

	resultSize, _ := resp.Body.Read(bs) // bs should be big enough to capture all the data else data will missed
	fmt.Println("size of response ", resultSize)

	fmt.Println("Body ->>>>>>>", string(bs))
}
