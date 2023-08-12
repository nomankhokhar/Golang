package main

import "fmt"

const (
	lol           string = "Noman"
	MaNhibatayoGa        = "Na Bta bhai"
)

const (
	i = iota
	j
	k
)

type Student struct {
	name     string
	RollNo   int
	subjects []string
}

func main() {
	// var i int = 12
	// i = 13
	// No := "lol"
	// flo := 12
	// flu := string(flo)
	// fmt.Println("Hello World", i, No, flo, flu)
	// fmt.Printf("%v , %T", i, i)
	// fmt.Printf("%v , %T", flo, flo)
	// var boo bool = true
	// fmt.Println("%v %T", boo, boo)

	// // complex is Builtin Function that is use for Complex Values
	// var comPlex complex128 = complex(1, 3)

	// fmt.Println(real(comPlex), imag(comPlex))
	// fmt.Println("%v %T", comPlex, comPlex)

	// s := "This is a String"
	// s1 := []byte(s)
	// fmt.Printf("%v , %T", s1, s1)

	// // const work with all data types flloat64 32
	// const CNST int = 12

	// fmt.Println("\n", CNST, lol, MaNhibatayoGa)

	// fmt.Println(i, j, k)

	// // Arrays in Golang
	// 	var amount [3]int = [3]int{10, 20, 30}
	// 	amout := [...]int{10, 20, 30, 45, 34, 53, 4, 534, 5, 34, 53, 45, 34, 53, 4}
	// 	b := amout[2:5]
	// 	fmt.Printf("Amuont: %v\n", amount)
	// 	fmt.Printf("Amuont: %v\n", amout)
	// 	fmt.Printf("Amuont: %v\n", b)

	// 	var TwoDArray [3][3]int = [3][3]int{
	// 		[3]int{1, 0, 0},
	// 		[3]int{1, 1, 0},
	// 		[3]int{1, 0, 1},
	// 	}

	// 	var Slice []float64 = []float64{1.2, 3, 4, 4}
	// 	fmt.Println(TwoDArray)
	// 	fmt.Println(Slice)

	// var a []int = make([]int, 3)
	// a[0] = 4
	// a[1] = 14
	// a [2] = 44

	// var b []int = append(a, 5)
	// fmt.Println(len(a))
	// fmt.Println(b)

	shopinCart := map[string]int{
		"KeyBoard": 100,
		"Mouse":    50,
		"Laptop":   1000,
	}

	for k, v := range shopinCart {
		fmt.Println(k, v)
	}

	// _, ok := shopinCart["KeyBoard"]
	// if ok {
	// 	fmt.Println("Avail")
	// }

	// if _, ok := shopinCart["KeyBoard"]; ok {
	// 	fmt.Println("Avail")
	// }

	// shopinCart["Book"] = 30
	// has, ok := shopinCart["Book"]

	// // Books will get ByRefrence address of Shoping
	// books := shopinCart

	// fmt.Println(has, ok)
	// fmt.Println(shopinCart["Mouse"])
	// delete(books, "Book")

	// fmt.Println(shopinCart)

	// Structure of Struct

	// student := Student{
	// 	"Noman Ali",
	// 	1242,
	// 	[]string{
	// 		"Maths",
	// 		"Physics",
	// 		"Chemistry",
	// 	},
	// }

	// NomanAli := Student{
	// 	name:   "Noman Ali",
	// 	RollNo: 1242,
	// 	subjects: []string{
	// 		"Maths",
	// 		"Physics",
	// 		"Chemistry",
	// 	},
	// }

	// NomanAli.name = "Noman Khan"
	// fmt.Println(NomanAli)
	// fmt.Println(student)
	// fmt.Println(student.name)
	// fmt.Println(student.RollNo)
	// fmt.Println(student.subjects)

	// Pc := Computer{}

	// Pc.cores = 12
	// Pc.memoryCapacity = 1000
	// Pc.memoryType = "SSD"
	// Pc.price = 1212
	// Pc.processorName = "NomanAli"

	// fmt.Println(Pc)
	// fmt.Println(Pc.Memory)
	// fmt.Println(Pc.Processor)

	// com := Computer{
	// 	Processor: Processor{
	// 		processorName: "NomiBhai",
	// 		cores:         12,
	// 	},
	// 	Memory: Memory{
	// 		memoryCapacity: 12121,
	// 		memoryType:     "SSD",
	// 	},
	// 	price: 124234,
	// }

	// fmt.Println(com)

	// bla := []int{1, 2, 3, 4, 45, 6, 54, 7, 7}

	// for _, v := range bla {
	// 	fmt.Println(v)
	// }

	// for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
	// 	fmt.Println(i, j)
	// }
}

type Processor struct {
	processorName string
	cores         int
}

type Memory struct {
	memoryCapacity int
	memoryType     string
}

type Computer struct {
	Processor
	Memory
	price int
}
