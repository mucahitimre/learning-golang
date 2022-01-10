package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var member = GetInterfacedMember()
	fullName := member.GetFullName()
	fmt.Printf("Full Name: %s \n", fullName)

	age := member.GetAge()
	fmt.Printf("Age: %d \n", age)

	// pointer olan methodlar interface'e tanımlanamaz
	stMember := GetStructMember()
	address := stMember.GetOrAddAddress("USA")
	fmt.Printf("Address: %s \n", address)

	fmt.Println("end..")
}

func GetInterfacedMember() IMember {
	return Member{
		Name:      "Joey",
		Surname:   "Tribbiani",
		Address:   "",
		BirthYear: 1967,
	}
}

func GetStructMember() Member {
	return Member{
		Name:      "Joey",
		Surname:   "Tribbiani",
		Address:   "",
		BirthYear: 1967,
	}
}

type Member struct {
	Name      string `json:name`
	Surname   string `json:surname`
	Address   string `json:address`
	BirthYear int    `json:birthyear`
}

func (member Member) GetFullName() string {

	return strings.Join([]string{member.Name, member.Surname}, " ")
}

func (member Member) GetAge() int {
	nowYear := time.Now().Year()
	return nowYear - member.BirthYear
}

func (member *Member) GetOrAddAddress(address string) string {
	if member.Address == "" {
		member.Address = address
	}

	return member.Address
}

type IMember interface {
	GetFullName() string
	GetAge() int
}
