package filestore

import (
	"Todo-Cli-With-Golang/entity"
	"fmt"
	"os"
	"strings"
	json2 "encoding/json"

)

type FileStore struct {
	filePath string
}

func New(path string) FileStore{
	return  FileStore{filePath: path}
}

func (f FileStore) Save(u entity.User) {
	f.writeFileUser(u)
}


func(f FileStore) loadUserFromStorage()[]entity.User {
	return  f.Load()

}


func (f FileStore) Load() []entity.User {
	var uStore []entity.User
	data, err := os.ReadFile(f.filePath)
	if err != nil {
		fmt.Println("cannot read file:", err)
		return nil
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	for _, line := range lines {
		var user entity.User
		err := json2.Unmarshal([]byte(line), &user)
		if err != nil {
			fmt.Println("error parsing user json:", err)

			continue
		}
		uStore = append(uStore, user)
		//fmt.Printf("user: %+v\n", user)
	}
	return uStore
}

func (f FileStore) writeFileUser(user entity.User) {
	var file *os.File
	//با مقدادیر زیر ما یک فایل ایجاد،اگر ایجاد شده بود اپند یا به ان چیزی اضافه می کنیم و ان را می خوانیم و اطلاعات یوزر خود را در این فایل سیو می کنیم
	file, err := os.OpenFile(f.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("can,t create or open file", err)

		return
	}
	defer file.Close()
	var data []byte

	data, err = json2.Marshal(user)
	if err != nil {
		fmt.Println("error to sava data format in the json", err)

		return
	}
	data = append(data, '\n')

	file.Write(data)

}