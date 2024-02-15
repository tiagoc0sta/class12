/* package main

import "fmt"

func recommendTravelLocation(season string) string {

 if season == "spring" {
  return "Japan for cherry blossoms"
 } else if season == "autumn" {
  return "New England for the fall"
 } else if season == "summer" {
  return "Italy beaches are great this time of year"
 } else if season == "winter" {
  return "Switzerland for skiing"
 } else {
  return "Unknown season, unable to recommend a travel location"
 }

}

func main() {

 //Example usage of the recommendTravelLocation function

 fmt.Println(recommendTravelLocation("spring"))
 fmt.Println(recommendTravelLocation("winter"))
 fmt.Println(recommendTravelLocation("summer"))
 fmt.Println(recommendTravelLocation("autumn"))

}*/

/*package main

import "fmt"

// Declaring a struct type
type Person struct {
 Name string
 Age  int
}

func main() {
 // Initializing a struct
 person := Person{
  Age:  30,
  Name: "Alice",
 }

 // Accessing and printing struct fields
 fmt.Println(person.Name) // Output: Alice
 fmt.Println(person.Age)  // Output: 30


 // Modifying struct fields
 person.Age = 31
 fmt.Println(person.Age) // Output: 31
}*/

/*package main

import "fmt"

//Declaring a struct type for Ice Cream

type IceCream struct {
 Flavor string
 Price  float64
}

func main() {

 //Initializing a struct with ice cream details
 iceCream := IceCream{
  Flavor: "Chocolate",
  Price:  2.99,
 }

 //Accessing and printing struct fields

 fmt.Println(iceCream.Flavor)
 fmt.Println(iceCream.Price)

 //Modifying struct fields

 iceCream.Flavor = "French Vanilla"
 iceCream.Price = 5.99

 fmt.Println(iceCream.Flavor)
 fmt.Println(iceCream.Price)

}*/

/*package main

import "fmt"

//Declaring a struct type

type Employee struct {
 Salary   float64
 Vacation int
}


type Movie struct {
	Title string
	Genre string
}

func main() {

 //Initializing a struct
 employee := Employee{
  Salary:   50000.00,
  Vacation: 21,
 }

  //Initializing a struct
	movie := Movie{
		Title: "Inception",
		Genre: "Sci-Fi",
	}


 //Accessing and printing struct fields
 fmt.Println(employee.Salary)
 fmt.Println(employee.Vacation)


 //Accessing and printing struct fields
 fmt.Println(movie.Title)
 fmt.Println(movie.Genre)

 //Modifying struct fields
 employee.Vacation = 25
 fmt.Println(employee.Vacation)

 //Modifying struct fields
 movie.Genre = "Science Fiction"
 fmt.Println(movie.Genre)

}*/

/*package main

import "fmt"

type Person struct {
    Name    string
    Age     int
    Address string
}

func main() {

	 // initialize struct
    people := map[string]Person{
        "Alice": {Name: "Alice Smith", Age: 30, Address: "123 Main st"},
        "Bob":   {Name: "Bob Johnson", Age: 25, Address: "456 Park ave"},
    }

    // Accessing and printing the original values
    fmt.Println("Original values:")
    fmt.Println(people["Alice"])
    fmt.Println(people["Bob"])

    // Modifying a value in the map
    alice := people["Alice"] // Get the struct value from the map
    alice.Age = 31           // Modify the field
    people["Alice"] = alice  // Assign the modified struct value back to the map

    // Accessing and printing the modified value
    fmt.Println("\nModified value of Alice's age:")
    fmt.Println(people["Alice"])
}*/

/*
//Pointers///////////////
package main

import "fmt"

func modifyValue(val *int) {
 *val = 50
}

func main() {
 x := 42
 fmt.Println("Before:", x) // Output: Before: 42

 modifyValue(&x)
 fmt.Println("After:", x) // Output: After: 50
}
*/

/*package main

import "fmt"

//Modifying struct fields with pointers

type Person struct {
 Name string
 Age  int
}

func updateAge(p *Person, newAge int) {
 p.Age = newAge
}

func main() {

 person := Person{Name: "Alice", Age: 30}

 fmt.Println("Before", person)

 updateAge(&person, 31)

 fmt.Println("After: ", person)

}*/

/*
package main

import "fmt"

//Returning a pointer from a function

func createIntPointer() *int {
 x := 100

 return &x
}

func main() {
 ptr := createIntPointer()
 fmt.Println(*ptr)

}
*/

/*package main

import "fmt"

type Point struct {
 X, Y int
}

func main() {

 p := Point{X: 10, Y: 20}

 ptr := &p

 fmt.Println("X: ", (*ptr).X)
 fmt.Println("Y: ", ptr.Y)

}*/

/*package main

import "fmt"

func modifySlice(slice *[]int) {
 (*slice)[0] = 100

}

func main() {

 nums := []int{1, 2, 3, 4}
 fmt.Println("Before:", nums)

 modifySlice(&nums)
 fmt.Println("After: ", nums)

}*/

package main

import "fmt"

//Pass-by-Value vs Pass-by-Reference

func modifyValue(val int) {

 val = 50

}

func modifyPointer(val *int) {
 *val = 50
}


func main() {

 x := 42

 fmt.Println("Before:", x)

 modifyValue(x)

 fmt.Println("After (value):", x)

 modifyPointer(&x)

 fmt.Println("After (pointer):", x)

}

