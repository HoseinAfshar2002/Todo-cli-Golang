package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	ID         int
	Title      string
	Date       string
	CategoryId int
	DoneTask   bool
	UserId     int
}
type Category struct {
	ID     int
	Title  string
	Color  string
	UserId int
}

var userStorage []User
var authUser *User
var taskStorage []Task
var categoryStorage []Category

const userStoragePath = "user.txt"

func main() {
	loadUserStorage()

	command := flag.String("command", "no command", "Run Command")
	flag.Parse()

	// ایجاد یک حلققه برای اجرای کامند های ما
	for {
		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter the new  command")
		scanner.Scan()
		//متن وارد شده رو درون کامند می ریزیم
		*command = scanner.Text()

	}

}
func runCommand(command string) {
	if command != "register-user" && command != "exit" && authUser == nil {

		loginUser()

		if authUser == nil {
			return
		}

	}
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "list-task":
		listTask()

	case "login-user":
		loginUser()
	case "exit":
		os.Exit(0)

	default:
		fmt.Println("command is not valid", command)

	}
}
func createTask() {

	scanner := bufio.NewScanner(os.Stdin)
	var name, date, category string

	fmt.Println("please enter the task name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the task date")
	scanner.Scan()
	date = scanner.Text()

	fmt.Println("please enter the task category id")
	scanner.Scan()
	category = scanner.Text()

	//در خط زیر کتگوری ای دی رو از استرینگ به عدد تبدیل می کنیم

	categoryId, err := strconv.Atoi(category)
	if err != nil {
		fmt.Println("category is not valid", err)

		return
	}
	//چک می کنیم که کتگوری ای دی وجود داشته باشد و یوزر ای همان یور دای دی باشد که دراد دان را فراخانی می کند
	isFound := false
	for _, c := range categoryStorage {
		if c.ID == categoryId && c.UserId == authUser.ID {
			isFound = true

			break
		}
	}

	if !isFound {
		fmt.Println("category not found")

		return
	}

	task := Task{
		ID:         len(taskStorage) + 1,
		Title:      name,
		Date:       date,
		CategoryId: categoryId,
		DoneTask:   false,
		UserId:     authUser.ID,
	}
	taskStorage = append(taskStorage, task)

}
func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string
	fmt.Println("please enter the category title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the category color")
	scanner.Scan()
	color = scanner.Text()

	fmt.Println("category:", title, color)

	category := Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserId: authUser.ID,
	}
	categoryStorage = append(categoryStorage, category)
}

func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var id, name, email, password string

	fmt.Println("please enter the user name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the user password")
	scanner.Scan()
	password = scanner.Text()

	fmt.Println("user:", name, email, password)
	id = email
	fmt.Println("user Id:", id, "user email:", email, "user password:", password)

	user := User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: password,
	}
	userStorage = append(userStorage, user)

	writeFileUser(user)

}

func loginUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the user password")
	scanner.Scan()
	password = scanner.Text()

	fmt.Println("Login-Form")

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			authUser = &user

			break
		}
	}

	if authUser == nil {
		fmt.Println("email is password not found")

	}

}

func listTask() {
	for _, task := range taskStorage {
		if task.UserId == authUser.ID {
			fmt.Println(task)
		}
	}

}

func loadUserStorage() {
	file, err := os.Open(userStoragePath)
	if err != nil {
		fmt.Println("not open file !", err)
	}
	data, err := os.ReadFile(userStoragePath)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("not read file", err)
	}
	var dataString string = string(data)
	dataString = strings.Trim(dataString, "\n")
	userSlice := strings.Split(dataString, "\n")
	for _, u := range userSlice {
		var user = User{}

		userFild := strings.Split(u, ",")
		for _, fild := range userFild {
			value := strings.Split(fild, ": ")
			if len(value) != 2 {
				fmt.Println("not valid field , skip...", len(value))

				continue
			}
			fildName := strings.ReplaceAll(value[0], " ", "")
			fildValue := value[1]

			switch fildName {
			case "id":
				id, err := strconv.Atoi(fildValue)
				if err != nil {
					fmt.Println("str convert error", err)
				}
				user.ID = id

			case "name":
				user.Name = fildValue

			case "email":
				user.Email = fildValue

			case "password":
				user.Password = fildValue

			}

		}
		fmt.Printf("user: %+v\n", user)
	}

	//fmt.Println(data)
}

func writeFileUser(user User) {
	var file *os.File
	//با مقدادیر زیر ما یک فایل ایجاد،اگر ایجاد شده بود اپند یا به ان چیزی اضافه می کنیم و ان را می خوانیم و اطلاعات یوزر خود را در این فایل سیو می کنیم
	file, err := os.OpenFile(userStoragePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("can,t create or open file", err)

		return
	}
	data := fmt.Sprintf("id: %v, name: %v, email: %v, password: %v\n",
		user.ID, user.Name, user.Email, user.Password)

	var b = []byte(data)
	file.Write(b)

	file.Close()
}
