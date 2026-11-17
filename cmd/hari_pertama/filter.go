package main

import (
	"strings"
)

var blacklist []string = []string{
	"sutarman",
	"fitri",
	"asep",
}

func FilterUser(usr User) (bool, error) {
	username := strings.ToLower(usr.Username)
	gender := strings.ToLower(usr.Gender)

	for _, v := range blacklist {
		if v == username {
			return false, nil
		}
	}

	switch gender {
	case "p":
	case "l":
	default:
		return false, nil
	}

	return true, nil
}
