package phone

import "fmt"

var phoneBook []PhoneContact

type PhoneContact struct {
	number int
	name   string
}

func (p PhoneContact) Format() string {
	return fmt.Sprintf("number: %d name: %s", p.number, p.name)
}

func AddContact(p PhoneContact) {
	phoneBook = append(phoneBook, p)
}

func GetContact(n string) PhoneContact {
	var result PhoneContact
	for _, v := range phoneBook {
		if v.name == n {
			result = v
			break
		}
	}
	return result
}
