package phone

import "testing"

func TestPhoneFormatReturn(t *testing.T) {
	phoneContact := PhoneContact{96115677, "David Thomas"}
	response := phoneContact.Format()
	expected := "number: 96115677 name: David Thomas"
	if response != expected {
		t.Errorf("%q was not equal to %q", response, expected)
	}
}

func TestPhoneContactRetrieve(t *testing.T) {
	phoneContact := PhoneContact{96115677, "David Thomas"}
	AddContact(phoneContact)
	response := GetContact(phoneContact.name)
	expected := 96115677
	if response.number != expected {
		t.Errorf("%q was not equal to %q", response, expected)
	}
}
