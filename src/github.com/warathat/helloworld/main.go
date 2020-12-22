package main

import (
	bnkgopackage "github.com/chaiyarin/bnk-gopackage" //เพิ่ม package bnk เข้ามา
	"github.com/warathat/school"
)

func main() {
	println("Hello World")

	println("School Name", school.SchoolName)
	println("School Address:", school.GetSchoolAddress())

	println("Member Name : ", bnkgopackage.GetFullNameCherprang())
	// นำ package bnk มาใช้งาน
}
