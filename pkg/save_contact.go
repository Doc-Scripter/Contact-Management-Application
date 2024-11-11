package pkg

import (
	"contact-management/database"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func SaveContacts() bool {
	name := ""
	email := ""
	numbers := 0
	fmt.Println("Enter Name")
	fmt.Scanf("%s", &name)
	fmt.Println("Enter Email")
	fmt.Scanf("%s", &email)
	fmt.Println("Enter Phone Numeber")
	fmt.Scanf("%d", numbers)
	var contacts database.Contacts
	file, err := os.ReadFile("database/contacts.json")
	if err != nil {
		log.Fatal(err)
	}
	var existing []database.Contacts
	err = json.Unmarshal(file, &existing)
	if err != nil {
		fmt.Println(err)
		return false
	}
	contacts.Name = name
	contacts.Email = email
	contacts.Phone = numbers

	existing = append(existing, contacts)
	fmt.Println(existing)
	saved, err := json.Marshal(existing)
	if err != nil {
		fmt.Println(err)
		return false
	}
	os.WriteFile("database/contacts.json", saved, 0o777)

	return true
}
