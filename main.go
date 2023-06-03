package main

import (
	"Files/data"
	"fmt"
)

func main() {

	/*rootPath, _ := os.Getwd()
	c, err := utils.ReadFileFromText((rootPath + "/data/text.txt"))
	if err == nil {
		fmt.Print(c)
		newContent := fmt.Sprintf("Original :%v\n   Double Original:%v%v", c, c, c)
		utils.WriteToFile((rootPath + "/data/text2.txt"), newContent)

	} else {
		fmt.Print("Error reading file", err)

	}*/

	////////////////////Call Instructor Struct ////////////////////////

	max := data.Instructor{Id: 45, FirstName: "Max", LastName: "Max", Score: 100}
	print(max.Print())

	kyle := data.NewInstructor("Kyle", 100)
	print(kyle.Print())

	goCourse := data.Course{Id: 1, Name: "Go Fundamentals", Slug: "go", Legacy: true, Duration: 3.2, Instructor: max}

	fmt.Printf("%v", goCourse)
	swiftWS := data.NewWorkshop("Sift", max)

	fmt.Printf("%v", swiftWS)
	print(kyle.Print())

	var courses [2]data.Course
	courses[0] = goCourse
	//courses[1] = swiftWS  // this will not work because swiftWS is a workshop not a course  (Not Inheritance)

	for _, course := range courses {
		fmt.Printf("%v", course)
	}

	//	Implementing Interfaces
	var courses2 [2]data.Singable
	courses2[0] = goCourse
	courses2[1] = swiftWS

	for _, course := range courses2 {
		fmt.Printf("%v", course)
	}

	var things []any //Global Interface

	things = append(things, "34")
	things = append(things, "34")
	things = append(things, swiftWS)

	things[2].(data.Workshop).SignUp()

}
