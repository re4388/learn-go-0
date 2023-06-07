package learn

import "fmt"

func RUN_fmt() {
	var name string
	var age int

	fmt.Println("Enter your name and age:")
	_, err := fmt.Scanf("%s %d", &name, &age)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Name: %s\nAge: %d\n", name, age)
}
