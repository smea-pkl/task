package main

import "os"

func InsertData(usr string) bool {
	filename := "data.txt"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return false
	}

	defer file.Close()

	file.WriteString(usr)

	return true
}