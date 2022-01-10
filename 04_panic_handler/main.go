package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	date, ex := GetNowDate()
	fmt.Println(date, ex)

	dateTheError, excep := GetNowDateRetError()
	if excep != nil {
		fmt.Println(excep)
	} else {
		fmt.Println(dateTheError)
	}

	fmt.Println("end..")
}

func GetNowDate() (date string, e error) {
	defer func() {
		if ex := recover(); ex != nil {
			fmt.Println(ex)
		}
	}()

	now := time.Now().UTC()
	date = now.String()
	date = string(date)

	endDate, e := time.Parse("2006-01-02", "21-02-24")
	if endDate.Before(now) {
		panic("panic generated")
	}

	return date, e
}

func GetNowDateRetError() (date string, e error) {
	defer func() {
		if ex := recover(); ex != nil {
			fmt.Println(ex)
		}
	}()

	now := time.Now().UTC()
	date = now.String()
	date = string(date)

	endDate, e := time.Parse("2006-01-02", "21-02-24")
	if endDate.Before(now) {
		e = errors.New("THE DATE WRONG")
	} else {
		e = nil
	}

	return date, e
}
