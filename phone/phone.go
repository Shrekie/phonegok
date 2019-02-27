package phone

import "fmt"

var phoneBook []PhoneContact

type PhoneContact struct {
	Number string
	Name   string
}

func (p PhoneContact) Format() string {
	return fmt.Sprintf("number: %s name: %s", p.Number, p.Name)
}

func AddContact(p PhoneContact) {
	phoneBook = append(phoneBook, p)
}

func ContactByName(n string) PhoneContact {
	var result PhoneContact
	for _, v := range phoneBook {
		if v.Name == n {
			result = v
			break
		}
	}
	return result
}

func ContactByNumber(n string) PhoneContact {
	var result PhoneContact
	for _, v := range phoneBook {
		if v.Number == n {
			result = v
			break
		}
	}
	return result
}
