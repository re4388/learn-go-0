package learn

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) Eat() {
	fmt.Println("Animal is eating")
}

/*
In Go, code reuse through composition is achieved by embedding one struct into another.
This approach is known as composition-based reuse
which allows the embedded struct (Animal) to provide its functionality to the containing struct (Dog).

unlike TS, Go 這裡是用 composition, 不是用繼承
*/
type Dog struct {
	Animal
	breed string
}

func (d *Dog) Bark() {
	fmt.Println("Dog is barking")
}

type Cat struct {
	Animal
	color string
}

func (c *Cat) Meow() {
	fmt.Println("Cat is meowing")
}

func RUN_struct1() {
	dog := Dog{Animal: Animal{name: "Doggo"}, breed: "Labrador"}
	cat := Cat{Animal: Animal{name: "Kitty"}, color: "White"}

	dog.Eat()  // Output: Animal is eating
	dog.Bark() // Output: Dog is barking

	cat.Eat()  // Output: Animal is eating
	cat.Meow() // Output: Cat is meowing
}
