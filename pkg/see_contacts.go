package pkg

import (
	"contact-management/database"
	"encoding/json"
	"fmt"
	"os"
)

func SeeContacts() bool {
	file, err := os.ReadFile("database/contacts.json")
	if err != nil {
		fmt.Println(err)
		return false
	}
	var contacts []database.Contacts

	if err = json.Unmarshal(file, &contacts); err != nil {
		fmt.Println(err)
		return false
	}
	for _, v := range contacts {
		fmt.Println(v.Name+" "+"%d", v.Phone, " ", v.Email)
	}
	return true
}
