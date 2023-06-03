package main

import (
	"Files/data"
	"fmt"
	"time"
)

func main() { //main go routine

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

	//Call Go Routines
	go printMessage("Go Routines")
	go printMessage("Go Routines2")
	printMessage("Main Routine")

	//call Channels

	var channel chan string
	go testChannel("Channel 1", channel)

	response := <-channel
	fmt.Println(response)

}

////////////////////Go Routines ////////////////////////

func printMessage(text string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text+" ", i)
		time.Sleep(1000 * time.Millisecond)
	}
}

func testChannel(text string, channel chan string) {

	for i := 0; i < 10; i++ {
		fmt.Printf("Channel %v", text+" ", i)
		time.Sleep(800 * time.Millisecond)
	}

	response := <-channel
	fmt.Printf("Channel %v", response)
}
