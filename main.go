package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	ID       int
	Title    string
	Date     string
	Category string
	DoneTask bool
	UserId   int
}

var userStorage []User
var authUser *User

var taskStorage []Task

func main() {
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

	fmt.Println("please enter the task category")
	scanner.Scan()
	category = scanner.Text()

	task := Task{
		ID:       len(taskStorage) + 1,
		Title:    name,
		Date:     date,
		Category: category,
		DoneTask: false,
		UserId:   authUser.ID,
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
