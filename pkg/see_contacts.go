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
	// maxName:=0
	// maxEmail:=0
	// maxPhone:=0

	// for _,v:= range contacts{
	// 	if len(v.Name)>maxName{
	// 		maxName=len(v.Name)
	// 	}
	// 	if len(v.Email)>maxEmail{
	// 		maxEmail=len(v.Email)
	// 	}
	// 	if len(v.Phone)>maxPhone{
	// 		maxName=len(v.Name)
	// 	}
	// }
	if len(contacts)==0{
		fmt.Print("Nothing to display...")
		return true
	}
	for _, v := range contacts {
		fmt.Printf("%s        +254%d          %s\n",v.Name,v.Phone, v.Email)
	}
	return true
}
