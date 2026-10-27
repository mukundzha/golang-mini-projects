
package main

import "fmt"

func main() {
    tasks := []string{}

    for {
        fmt.Println("\n1. Add task")
        fmt.Println("2. List tasks")
        fmt.Println("3. Exit")
		fmt.Print("\nChoose: ")

        var choice int
        fmt.Scan(&choice)

        if choice == 1 {
            var task string
            fmt.Print("\nEnter task: ")
            fmt.Scan(&task)

            tasks = append(tasks, task)
            fmt.Println("Task added!")

        } else if choice == 2 {
			for i := 0; i < len(tasks); i++ {
				fmt.Print("")
                fmt.Println(i+1, tasks[i])
            }

        } else if choice == 3 {
            break

        } else {
            fmt.Println("\nInvalid choice")
        }
    }
}

