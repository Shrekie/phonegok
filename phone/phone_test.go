package phone

import "testing"

func TestPhoneFormat(t *testing.T) {
	phoneContact := PhoneContact{"96115677", "David Thomas"}
	response := phoneContact.Format()
	expected := "number: 96115677 name: David Thomas"
	if response != expected {
		t.Errorf("%q was not equal to %q", response, expected)
	}
}

func TestPhoneContactByName(t *testing.T) {
	phoneContact := PhoneContact{"96115677", "David Thomas"}
	AddContact(phoneContact)
	response := ContactByName(phoneContact.Name)
	expected := "96115677"
	if response.Number != expected {
		t.Errorf("%q was not equal to %q", response, expected)
	}
}

func TestPhoneContactByNumber(t *testing.T) {
	phoneContact := PhoneContact{"96115677", "David Thomas"}
	AddContact(phoneContact)
	response := ContactByNumber(phoneContact.Number)
	expected := "David Thomas"
	if response.Name != expected {
		t.Errorf("%q was not equal to %q", response, expected)
	}
}
