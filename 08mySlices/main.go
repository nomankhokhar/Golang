package main

import (
	"fmt"
)

type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("Driving...")
}

func (c Car) GetName() string {
	return c.Name
}

func todo() {
	arr := []int{1, 2, 3, 4}
	arr2 := []string{"fg", "dfg", "dfg"}
	arr2 = append(arr2, "is", "Noman")
	fmt.Println(arr, arr2)
}

func main() {
	// fmt.Println("Welcome to my Slice")

	// var fruits = []string{"Apple", "Tomato", "Kela"}

	// fmt.Println(fruits)

	// fruits = append(fruits, "Nomi", "Ali")

	// fmt.Println(fruits)

	// fruits = append(fruits[1:])

	// fmt.Println(fruits)

	// highScore := make([]int, 4)

	// highScore[0] = 10
	// highScore[1] = 11
	// highScore[2] = 12
	// highScore[3] = 13

	// highScore = append(highScore, 14, 15, 17)

	// sort.Ints(highScore)

	// fmt.Println(highScore)

	// courses := []string{"Course1", "Course2", "Course3", "Course4", "Course1", "Course2", "Course3", "Course4"}

	// fmt.Println("Enter the course index you want to remove: ")

	// reader := bufio.NewReader(os.Stdin)

	// input, _ := reader.ReadString('\n')
	// input = strings.TrimSpace(input)

	// numRating, err := strconv.ParseInt(input, 10, 64)
	// if err != nil {
	// 	fmt.Println("Invalid input. Please enter a valid course index.")
	// 	return
	// }

	// if numRating < 0 || int(numRating) >= len(courses) {
	// 	fmt.Println("Invalid course index. Please enter a valid course index.")
	// 	return
	// }

	// courses = append(courses[:numRating], courses[numRating+1:]...)

	// fmt.Println(courses)

	todo()

	c := Car{
		Name:    "GTR",
		Age:     1,
		ModelNo: 2,
	}

	c.Print()
	c.Drive()

	fmt.Println(c.GetName())

}
