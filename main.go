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
	for {
		if count == 1 {
			fmt.Println("Hello,Please Enter your name")
			fmt.Scanf(user, 2)
		}
		for scanner.Scan() {
			count++
			// input:=scanner.Text()
			option := 0
			if count > 1 {
				fmt.Println("Please choose one option \n 1. See Contacts \n 2. Search Contact \n 3. Enter New Contact \n 4. Delete Contact")
				fmt.Scanf("%d", &option)

				if option == 3 {
					done := pkg.SaveContacts()
					if !done {
						fmt.Println("An error occured Please try again later.")
					} else {
						fmt.Println("Success")
					}
					break
				}
				if option == 1 {
					done := pkg.SeeContacts()
					if !done {
						fmt.Println("An error occured Please try again later.")
					}
					break
				}
			}
			break
		}
	}
}
