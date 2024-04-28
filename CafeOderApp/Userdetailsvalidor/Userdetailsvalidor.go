package Userdetailsvalidor

func Userdetailsvalidation(firstname string, lastname string, city string, contactno string) (bool, bool, bool) {
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidMob := len(contactno) == 10
	isValidCity := len(city) >= 0
	return isValidName, isValidMob, isValidCity
}
