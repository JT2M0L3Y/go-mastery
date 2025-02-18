// Instructions:
// Complete each task below without looking at references.
// Try to recall the syntax from memory and write valid Go code.
package main

import (
	"fmt"
)

func multiReturn() (int, string) {
	return 1, "Hey"
}

func shadowEx() {
	x := 20
	fmt.Println("Inner: ", x)
}

func closureEx() func() string {
	name := "John"

	return func() string {
		name = "Hi " + name
		return name
	}
}

func main() {
	// 1. Declare a package-level variable.
	var greeting string = "Hello, World!"
	fmt.Println(greeting)

	// 2. Declare a variable `name` of type string and assign it your name.
	var name string = "JT2M0L3Y"
	fmt.Println(name)

	// 3. Declare an integer variable `age` and assign it a value.
	var age int = 22
	fmt.Println(age)

	// 4. Declare a float variable `height` and assign it a value.
	var height float64 = 5.68
	fmt.Println(height)

	// 5. Declare a boolean variable `isCodingFun` and assign it `true`.
	var isCodingFun bool = true
	fmt.Println(isCodingFun)

	// 6. Declare multiple variables in a single line `city`, `country`, `population`.
	var city, country, population = "Silverdale", "USA", 25000
	fmt.Println(city)
	fmt.Println(country)
	fmt.Println(population)

	// 7. Use the shorthand syntax to declare and initialize a variable for your favorite language.
	favoriteLang := "Go"

	// 8. Declare a constant `pi` with a value of 3.14159.
	const pi = 3.14159
	fmt.Println(pi)

	// 9. Swap two variables without using a temporary variable.
	var one, two = 1, 2
	one, two = two, one
	fmt.Println(one, two)

	// 10. Print the **type** of a variable using the golang formatting lib
	fmt.Printf("%T", favoriteLang)
	fmt.Println()

	// 11. Declare a zero-valued variable without assigning anything. Print its default value.
	var zero float64 = 0.0
	fmt.Println(zero)

	// 12. Declare a pointer variable that points to an integer and modify the value using the pointer.
	var pt1 *int = &one
	*pt1 = 4
	fmt.Println(one)

	// 13. Create a function that returns multiple values and store them in variables.
	newint, newStr := multiReturn()
	fmt.Println(newint, newStr)

	// 14. Convert an integer to a float and vice versa.
	floatVal := float64(age)
	fmt.Printf("%T", floatVal)
	fmt.Println()

	// 15. Use `const` with an untyped constant and assign it to different typed variables.
	const constVal = 42
	var floatVar float64 = constVal
	var intVar int = constVal
	fmt.Println(floatVar, intVar)

	// 16. Declare and use a shadowed variable inside a function.
	x := 10
	shadowEx()
	fmt.Println("Outer: ", x)

	// 17. Use the new keyword to allocate memory for an integer and modify its value.
	var allocInt = new(int)
	*allocInt = 5
	fmt.Println(allocInt, *allocInt)

	// 18. Use shorthand assignment inside an `if` statement.
	if (x > 5) {
		x := 4
		fmt.Println(x)
	}
	fmt.Println(x)

	// 19. Use a function closure to modify an outer variable.
	msg := closureEx()
	fmt.Println(msg())
}
