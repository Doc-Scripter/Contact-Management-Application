package main

import (
	"bufio"
	"contact-management/pkg"
	"fmt"
	"os"
)

func main() {
	args := os.Stdin
	scanner := bufio.NewScanner(args)
	count := 1
	user := ""
	if count == 1 {
		fmt.Println("Hello,Please Enter your name")
		fmt.Scanf(user, 2)
		count++
	}
	for scanner.Scan() {
		option := 0
		for count > 1 {
			fmt.Print("Please choose one option \n 1. See Contacts \n 2. Search Contact \n 3. Enter New Contact \n 4. Delete Contact \n 5. Exit\n\n")
			fmt.Scanf("%d", &option)

			if option == 3 {
				done := pkg.SaveContacts()
				if !done {
					fmt.Println("An error occured Please try again later.")
				} else {
					fmt.Print("Success \n\n")
				}
				count = 2
			}
			if option == 1 {
				done := pkg.SeeContacts()
				if !done {
					fmt.Println("An error occured Please try again later.")
				} else {
					fmt.Print("Success \n\n")
				}
				count = 2
			}
			if option == 4 {
				done := pkg.DeleteContact()
				if !done {
					fmt.Print("An error occured Please try again later.\n\n")
				}else{
					fmt.Print("Success \n\n")

				}
			}
			if option == 5 {
				fmt.Println("Have a good day ,come again next time:)")
				return
			}
		}
	}
}
