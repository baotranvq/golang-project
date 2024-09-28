package main

import (
    "fmt"
    "my-first-go-app/todolist/todo"  // Đường dẫn đúng đến package todo
)

func main() {
    fmt.Println("Hello, World!")
    todo.PrintTask()  // Gọi hàm từ package todo
}
