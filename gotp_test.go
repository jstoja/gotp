package gotp

import "testing"

func TestHotp(t *testing.T) {
	password := "MayTheFBeWithyou"
	hotpPassword := Hotp(password, 42)
	hotpExpected := "861699"
	if hotpPassword != hotpPassword {
		t.Error("The htop should be %s but is %s", hotpPassword, hotpExpected)
	}
}
