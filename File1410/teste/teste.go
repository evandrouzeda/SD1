package main

import (
	comands "SD1/File1410/pkg"
	"fmt"
)

//comands "SD1/File1410/pkg"
type msg struct {
	Name string `json:"name"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

func main() {
	msg := comands.LOGIN("s[0]", "s[1]", "s[2]")
	fmt.Println(msg)
	/*
		pay := msg{"Payload01", "evandro", "uzeda"}

		/* The json marshal function returns a byte array and an error
		b, err := json.Marshal(pay)

		/* check the error
		if err != nil {
			fmt.Println("There was an error :[")
			return
		}

		/*convert the byte array to a string and print it

		s1 := string(b)
		fmt.Println(s1)*/
}
