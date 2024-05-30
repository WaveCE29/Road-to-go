package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Python", "Java"}
	fmt.Println("Course Name:", courseName)
	courseName = append(courseName, "Ruby", "C#", "Go", "C++", "JavaScript")
	fmt.Println("Course Name:", courseName)
	// [4:7] means from index 4 to 6
	courseBack := courseName[4:7]
	fmt.Println("Course Back:", courseBack)
	// [:7] means from index 0 to 6
	courseBack = courseName[:6]
	fmt.Println("Course Back:", courseBack)

}
