package main

import (
	"encoding/json"
	"fmt"
)

// มีการ Import Encoding JSON เข้ามา

type customer struct { // การประกาศโครงสร้าง struct // มีการปรับ Field ให้ขึ้นต้นตัวใหญ่
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func main() {
	// cus := customer{
	// 	Firstname: "Warathat",
	// 	Lastname:  "Pant",
	// 	Code:      111111,
	// 	Phone:     "0812345678",
	// } // การกำหนดค่าเริ่มต้น ให้ customer struct
	// fmt.Println(cus)

	// fmt.Println(cus.Firstname)
	// fmt.Println(cus.Lastname)
	// fmt.Println(cus.Code)
	// fmt.Println(cus.Phone)

	// Struck with Array
	// customerLists := [3]customer{}
	// customerLists[0] = customer{
	// 	Firstname: "W1",
	// 	Lastname:  "P1",
	// 	Code:      1111111,
	// 	Phone:     "0812345678",
	// }
	// customerLists[1] = customer{
	// 	Firstname: "W2",
	// 	Lastname:  "P2",
	// 	Code:      222222,
	// 	Phone:     "0812345679",
	// }
	// customerLists[2] = customer{
	// 	Firstname: "W3",
	// 	Lastname:  "P3",
	// 	Code:      333333,
	// 	Phone:     "0812345671",
	// }
	// fmt.Println(customerLists)

	// Struck with Slice
	// customerLists := []customer{}
	// customer1 := customer{
	// 	Firstname: "W1",
	// 	Lastname:  "P1",
	// 	Code:      1111111,
	// 	Phone:     "0812345678",
	// }
	// customer2 := customer{
	// 	Firstname: "W2",
	// 	Lastname:  "P2",
	// 	Code:      222222,
	// 	Phone:     "0812345679",
	// }
	// customer3 := customer{
	// 	Firstname: "W3",
	// 	Lastname:  "P3",
	// 	Code:      333333,
	// 	Phone:     "0812345671",
	// }
	// customerLists = append(customerLists, customer1)
	// customerLists = append(customerLists, customer2)
	// customerLists = append(customerLists, customer3)
	// fmt.Println(customerLists)

	// To JSON
	customerLists := []customer{}
	customer1 := customer{
		Firstname: "W1",
		Lastname:  "P1",
		Code:      1111111,
		Phone:     "0812345678",
	}
	customer2 := customer{
		Firstname: "W2",
		Lastname:  "P2",
		Code:      222222,
		Phone:     "0812345679",
	}
	customer3 := customer{
		Firstname: "W3",
		Lastname:  "P3",
		Code:      333333,
		Phone:     "0812345671",
	}
	customerLists = append(customerLists, customer1)
	customerLists = append(customerLists, customer2)
	customerLists = append(customerLists, customer3)

	customerListsJson, err := json.Marshal(customerLists)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(customerListsJson))
}
