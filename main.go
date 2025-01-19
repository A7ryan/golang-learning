package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

// Global Var
var tasksList = []string{}

// var port = ":8080"

func main() {
	// var welcomeUser = "Welcome Aadi!"
	// var greetNewYear = "Wishing you a very Happy New Year.."

	// tasksList = append(tasksList, welcomeUser, greetNewYear)

	for {
		var menuSelection int
		fmt.Println("Welcome to Todo List Menu")
		fmt.Println("1. List Items")
		fmt.Println("2. Add Items")
		fmt.Println("3. Exit")

		fmt.Scanln(&menuSelection)

		switch menuSelection {
		case 1:
			listTasks()
		case 2:
			addTask()
		case 3:
			fmt.Println("Thanks, have a great day ahead!")
			os.Exit(0)
		}
	}
	// listTasks()
	// var msg = returnPrint()

	// fmt.Printf("http://localhost%s", port)
	// http.HandleFunc("/tasks", dispTaskAPI)

	// http.ListenAndServe(port, nil)
}

func listTasks() {
	for _, tasks := range tasksList {
		fmt.Println(tasks)
	}
}

func addTask() {
	var newTask string
	fmt.Println("Enter New Task: ")
	// fmt.Scanln(&newTask) this will only add data upto 1st white space

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		newTask = scanner.Text()
	}

	tasksList = append(tasksList, newTask)
	fmt.Println("####### Task Added Successfully #######")
}

func returnPrint() string {
	return "Hello, Aadi.."
}

func dispTaskAPI(writer http.ResponseWriter, request *http.Request) {
	for _, task := range tasksList {
		fmt.Fprintf(writer, "%s", task)
	}
}
