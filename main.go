package main

import "fmt"

func main() {
    // Student information
    name := "Indresh"
    age := 22

    // Marks in three subjects
    subject1 := 80
    subject2 := 75
    subject3 := 90

    // Calculate total marks
    total := subject1 + subject2 + subject3

    // Calculate average
    average := float64(total) / 3.0

    // Check whether the student passed
    passed := average >= 40

    // Display the student information
    fmt.Println("----- Student Report -----")
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Subject 1:", subject1)
    fmt.Println("Subject 2:", subject2)
    fmt.Println("Subject 3:", subject3)
    fmt.Println("Total Marks:", total)
    fmt.Printf("Average Marks: %.2f\n", average)
    fmt.Println("Passed:", passed)
}