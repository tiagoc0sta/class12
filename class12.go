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


package main

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
}
