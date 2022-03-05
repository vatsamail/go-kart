package main

import (
	"fmt"
	"log"

	"example.com/adv_greetings"
	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(quote.Go())
	message, err := greetings.Hello("Pushparaj")
	if err == nil {
		fmt.Println(message)
	}

	log.SetPrefix("Greetings Log: ")
	log.SetFlags(0)
	fmt.Println("What's next")

	new_message, new_err := adv_greetings.Hi("Gopal")
	fmt.Println("Write something")
	fmt.Println(new_message)
	fmt.Println(new_err)

	message, err = greetings.Hello("James")
	if err == nil {
		fmt.Println(message)
	}

	names := []string{"Amar", "Akbar", "Antony"}
	messages, err := adv_greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Printing Array")
	fmt.Println(messages)

	// final error
	messagex, errx := greetings.Hello("")
	if errx != nil {
		log.Fatal(errx)
	}
	fmt.Sprintf("Nil message: %v", messagex)
}
