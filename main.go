package main

import (
	"fmt"
	"time"
)

type Printable interface {
	DebugPrintable() string
}

type Address struct {
	postalCode int
	ward       string
}

type UserProfile struct {
	name    string
	age     int
	address Address
}

func (profile UserProfile) description() string {
	return profile.name + ", " + string(profile.age)
}

func (profile UserProfile) DebugPrintable() string {
	return profile.name
}

func (profile *UserProfile) set(name string) {
	profile.name = name
}

func main() {
	hiraganaList := []string{"あ", "い", "う", "い", "え"}
	hiraganaList = append(hiraganaList, "か")

	for index, hiragana := range hiraganaList {
		switch hiragana {
		case "あ":
			fmt.Println("This is the first letter of my first name", hiragana)
		default:
			fmt.Println(index, hiragana)
		}
	}

	information := make(map[string]string)
	information["title"] = "Surprise"
	information["subtitle"] = "This is Aymen"

	fmt.Println(information)

	aymenAddress := Address{
		postalCode: 111111,
		ward:       "Nerima",
	}

	aymenProfile := UserProfile{
		name:    "Aymen",
		age:     23,
		address: aymenAddress,
	}

	aymenProfile.name = "Aymen Yasine"
	fmt.Println(aymenProfile.description())

	aymenProfile.set("Pepitoo")
	fmt.Println(aymenProfile.description())

	go printEverySecond()

	fmt.Println("End")
}

func printEverySecond() {
	for range time.Tick(time.Second) {
		fmt.Println("every second")
	}
}
