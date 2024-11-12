package pkg

import (
	"contact-management/database"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func DeleteContact()bool{
  name:=""
  fmt.Println("Enter full or part of name to delete")
  fmt.Scanf("%s", &name)

  var contacts []database.Contacts
 file,err:=os.ReadFile("database/contacts.json")

 if err!=nil{
  fmt.Println(err)
  return false
 }
 json.Unmarshal(file,&contacts)
 var updatedContacts []database.Contacts
 var contact database.Contacts
 for _,v:=range contacts{
  if strings.Contains(v.Name,name){
    continue
  }
  contact.Name=v.Name
  contact.Email=v.Email
  contact.Phone=v.Phone
  updatedContacts=append(updatedContacts, contact)
 }

newFile,err:=json.Marshal(updatedContacts)
if err!=nil{
  fmt.Println(err)
  return false
}

os.WriteFile("database/contacts.json",newFile,0o777)
  return true
}