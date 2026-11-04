package main

import (
	"errors"
	"strings"
)

var blacklist []string = []string{
	"sutarman",
	"andika",
	"fitri",
}

func FilterUser(usr User) (bool, error) {
	username := strings.ToLower(usr.username)
	gender := strings.ToLower(usr.gender)

	for _, v := range blacklist {
		if v == username {
			if username == "fitri" && usr.gender != "p" {
				return true, nil
			}
			return false, nil
		}
	}

	errGender := errors.New("Gender tidak valid")

	switch gender {
	case "p":
	case "l":
	default:
		return true, errGender
	}

	return true, nil
}
