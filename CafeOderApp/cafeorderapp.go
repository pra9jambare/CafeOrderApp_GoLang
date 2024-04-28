package main

import (
	"CafeOrderApp/Userdetailsvalidor"
	"fmt"
)

var firstname string
var lastname string
var contactno string
var city string
var selectmenu string
var retry string

func main() {
	for {
		fmt.Println("Welcome To Cafe Filmy Flavour")
		UserDetails()
		ShowUserDetails()
		isValidName, isValidMob, isValidCity := Userdetailsvalidor.Userdetailsvalidation(firstname, lastname, city, contactno)
		println(isValidName, isValidMob, isValidCity)
		if isValidName == true && isValidMob == true && isValidCity == true {
			mainmenu(selectmenu)
			retrymenue(retry)
			fmt.Scan(&retry)
			if retry == "Yes" || retry == "Y" || retry == "yes" || retry == "y" {
				continue
			} else {
				break
			}
		} else {
			user_input_reply(isValidName, isValidMob, isValidCity)
			retrymenue(retry)
			fmt.Scan(&retry)
			if retry == "Yes" || retry == "Y" || retry == "yes" || retry == "y" {
				continue
			} else {
				break
			}
		}
	}
}
func UserDetails() {
	fmt.Println("Provide Detials About You")
	fmt.Println("First Name : ")
	fmt.Scan(&firstname)
	fmt.Println("Last Name : ")
	fmt.Scan(&lastname)
	fmt.Println("Contact No : ")
	fmt.Scan(&contactno)
	fmt.Println("City : ")
	fmt.Scan(&city)
	fmt.Println(" ")
}

func ShowUserDetails() {
	fmt.Println("")
	fmt.Println("Please Verify Your Details")
	fmt.Println("First Name : ", firstname)
	fmt.Println("Last Name : ", lastname)
	fmt.Println("Contact No : ", contactno)
	fmt.Println("City : ", city)
	fmt.Println(" ")
}

func mainmenu(selectmenu string) {

	fmt.Println("Please Select")
	fmt.Println("Press 1. for Menu Card")
	fmt.Println("Press 2. for Order")
	fmt.Println("Press 3. for Complaint")
	fmt.Println("Press 4.for About Us")
	fmt.Println("Press 5. for User Visit Details")
	fmt.Println(" ")
	fmt.Scan(&selectmenu)

	switch selectmenu {
	case "1":
		fmt.Println(" ")
		fmt.Println("Hello To Menucard")
		fmt.Println("Press 1. Pizza")
		fmt.Println("Press 2. Burger")
		fmt.Println("Press 3. Fries")
		fmt.Println(" ")
		var menu_select string
		fmt.Scan(&menu_select)
		switch menu_select {
		case "1":
			fmt.Println("Veg Pizza - 160rs")
			fmt.Println("Chicken Pizza - 2000rs")
			fmt.Println("Suprime Veg Pizza - 260rs")
			fmt.Println("Suprime Chicken Pizza - 300rs")
			fmt.Println(" ")
		case "2":
			fmt.Println("Veg Burger - 90rs")
			fmt.Println("Chicken Burger - 130rs")
			fmt.Println("Suprime Veg Burger - 150rs")
			fmt.Println("Suprime Chicken Burger - 200rs")
			fmt.Println(" ")
		case "3":
			fmt.Println("Salted Fries - 100rs")
			fmt.Println("Masala Fries - 150rs")
			fmt.Println("Chilly Fries - 200rs")
			fmt.Println(" ")
		default:
			fmt.Println(" ")
			fmt.Println("Wrong Input Please Retry")
			fmt.Println(" ")

		}
	case "2":
		fmt.Println("Hello To Order")
	case "3":
		var orderid uint
		var issue string
		fmt.Println("Hello To Complaint Section")
		fmt.Println(" ")
		ShowUserDetails()
		fmt.Println(" ")
		fmt.Println("Please enter the order-id")
		fmt.Scan(&orderid)
		fmt.Println("Please enter your issue")
		fmt.Scan(&issue)
		fmt.Println("Hello", firstname, "Complaint is sucessfully raised for order id", orderid)
		fmt.Println(" ")

	case "4":
		fmt.Println(" ")
		fmt.Println("About Us ")
		fmt.Println("Thank you for visiting us , Please find the details below ")
		fmt.Println("Owner: Pranav Jambare ")
		fmt.Println("Founded : 2022 ")
		fmt.Println("Category : Food ")
		fmt.Println(" ")

	case "5":
		fmt.Println(" ")
		fmt.Println("User Visit Details ")
		fmt.Println("User Visit Details - Menu Card")

	default:
		fmt.Println(" ")
		fmt.Println("You have select invalid menu")
		fmt.Println(" ")
	}
}

func retrymenue(retry string) {
	fmt.Println(" ")
	fmt.Println("Do you want still want to continue - Yes / No")
}

func userdetailsvalidation(firstname string, lastname string, city string, contactno string) (bool, bool, bool) {
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidMob := len(contactno) == 10
	isValidCity := len(city) >= 0
	return isValidName, isValidMob, isValidCity
}

func user_input_reply(isValidName bool, isValidMob bool, isValidCity bool) {
	if !isValidName {
		fmt.Println("You have written a short name")
	}
	if !isValidMob {
		fmt.Println("You have written a short mobile, please enter 10 digit")
	}
	if !isValidCity {
		fmt.Println("You have written incorrect City")
	}
}

// 	uservisitedmenucard = append(uservisitedmenucard, firstname+" "+lastname)
// 	var firstnames []string
// 	for _, uservistedmenue := range uservisitedmenucard {
// 		var names = strings.Fields(uservistedmenue)
// 		firstnames = append(firstnames, names[0])
// 	}
// 	return firstnames
// }
