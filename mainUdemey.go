package main

import (
	"fmt"
)
func main() {
	var x int = 1
	var y int
	var  ip *int = nil

	ip = &x
	y = *ip
	fmt.Println(x, y, *ip, &ip)

  // new() is writtern in golang that allocate memory and assign to the ptr

	ptr := new(int)
	*ptr = 3

	fmt.Println(*ptr)

	fmt.Println("Hello Noman Ali you need to work hard intead of listiong the Motivational Video!")

	var z *int = nil
	z = foo()
	fmt.Println(*z)

	var a int32 = 1;
	var b int16 = 2;
	a = int32(b)
	fmt.Println(a)

	// Floating Point

	var xy float32 = 123.33;
	var zy float64 = 1.24324e1

	fmt.Println(xy, zy);

	// String
	str := "Hello"

	fmt.Println(str)

	// Unicode Package All these Return Boolean values
	fmt.Println(unicode.IsLetter(str[0]))
  	fmt.Println(unicode.IsSpace(str[0]))
	fmt.Println(unicode.IsNumber(str[0]))
	fmt.Println(unicode.IsLower(str[0]))
	fmt.Println(unicode.IsPunct(str[0]))

	// String Package to Manipulate String

	// compare(), Contains(), HasPrefix(), Index()
	Replace(s, old, new, n)
	Split(s, sep, n)
	SplitAfter(s, sep, n)
	SplitBefore(s, sep, n)
	ToLower(s)
	ToUpper(s)
	Trim(s)
	TrimLeft(s)
	TrimRight(s)
	TrimSpace(s)

	// Strconv Package
	Atoi(s) //-> convert string into int
	Atof(s) //-> convert string into float
    Atoi64(s) //-> convert string into int64
    Atof64(s) //-> convert string into float64
    Atob(s) //-> convert string into bool
    Atoui(s) //-> convert string into uint
    Atoui64(s) //-> convert string into uint64
    Atoi32(s) //-> convert string into int32
    Atof32(s) //-> convert string into float32
    Atoi16(s) //-> convert string into int16
    Atof16(s) //-> convert string into float16
    Atoi8(s) //-> convert string into int8
    Atof8(s) //-> convert string into float8

	FormateFloat(s string)
	ParseFloat(s string)

	const lol = 1.3
	const (four = 4
			hi = "hi")
	fmt.Println(four , hi)

	// iota is use to Generate to Constant

	type Grade int;
	const (
        A Grade = iota
        B
        C
        D
		F
	)

	fmt.Println(A, B, C, D, F)

	if x < 100{
		fmt.Println(x)
	}

	for i:=0 ; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("LOL",i);
	}

	i := 0
	for i < 10 {
        fmt.Println(i)
        i++
    }

	for{
	 	fmt.Println("hi")
	 }

	x:= 1

	switch x{
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		default:
			fmt.Println("No You")
	}

	var appleNum int

	fmt.Println("Number of elements : ")

 	fmt.Scan(&appleNum)

	fmt.Println(appleNum)

}

func foo() *int {
	x:= 1
	return &x
}
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var in float64
	fmt.Printf("please enter a float: ")
	fmt.Scan(&in)
	x := int(in)
	fmt.Printf("your input truncated: ")
	fmt.Printf(strconv.Itoa(x) + "\n")
	return
}

package main

import (
	"fmt"
)

func main() {
	var Word string
	fmt.Println("Enter String No. ")
	fmt.Scanln(&Word)

	if strings.Contains(Word, "a") && strings.HasPrefix(Word,"i") && strings.HasSuffix(Word, "n") {
		fmt.Println("Found!")
	}else
	{
		fmt.Println("Not Found!")
	}

}

func main() {
  i, _ := strconv.Atoi("10")
  y := i * 2
  fmt.Println(y)
}

func main() {
  s := strings.Replace("ianianian", "ni", "in", 2)
  fmt.Println(s)
}

func main() {
  var xtemp int
  x1 := 0
  x2 := 1
  for x:=0; x<5; x++ {
    xtemp = x2
    x2 = x2 + x1
    x1 = xtemp
  }
  fmt.Println(x2)
}

package main

import "fmt"

func main() {
var word string
fmt.Scanln(&word)
word = strings.ToLower(word)
word = strings.TrimSpace(word)
fmt.Printf(word)
if strings.HasPrefix(word, "i") && strings.Contains(word, "a") && strings.HasSuffix(word, "n") {
	fmt.Printf("Found")
} else {
	fmt.Printf("Not Found")
}

// Array and Slice

var x [5]int = [5]int{1, 2, 3, 4, 5};

fmt.Println((x[4]))

y := [3]int {1, 2, 3}

for i, v := range y {
	fmt.Printf("index %d, value %d\n", i,v)
}

arr := []string{"a", "b", "c", "d","e","f","g","h","i","j","k","l","m"}

s1 := arr[1:5]
s2 := arr[3:10]

fmt.Println(s1)
fmt.Println(s2)

fmt.Println(len(s1) , cap(s1))

// Make Slice Function

sli := make([]int, 5)
fmt.Println(sli) // Output: [0 0 0 0 0]

sli = make([]int, 5, 10)
fmt.Println(sli) // Output: [0 0 0 0 0]

sli = append(sli, 100)
fmt.Println(sli) // Output: [0 0 0 0 0 100]

var idMap map[string]int
idMap  = make(map[string]int)

fmt.Println(idMap)

// using Map Literal Function

idMap := map[string]int{
	"Joe" : 123,
	"Noman" : 124,
}

fmt.Println(idMap["Joe"]) // Output: 123

// Map Function

// two value assignment tests for existance of the key

value, ok := idMap["Joe"]

fmt.Println(value, ok) // Output: 123 true

// Iterating Function

for key, val := range idMap {
	fmt.Println(key, val)
}

// Structs Function

// this should define at the beginning AND OUTsITE THE MAIN

type Person struct {
	name string
	addr string
	phone string
}

p1 := Person{
	name : "Noman Ali",
	addr : "1234567890",
  phone : "1234567890",
}

	fmt.Println(p1, p1.name)

	p2 := new(Person)

	p2.name = "Noman Ali"

	fmt.Println(p2)

}

package main

import (
	"fmt"
	"os"
)

func main() {
	slice := make([]int, 0, 3) // Create an empty slice with initial capacity 3

	for {
		var input string
		fmt.Print("Enter an integer (or 'X' to quit): ")
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		if input == "X" || input == "x"{
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer or 'X' to quit.")
			continue
		}

		slice = append(slice, num)
		sort.Ints(slice) // Sort the slice in ascending order

		fmt.Println("Sorted slice:", slice)
	}
}

func main() {
  x := []int {4, 8, 5}
  y := -1
  for _, elt := range x {
    if elt > y {
      y = elt
    }
	fmt.Println(elt)
  }
  fmt.Print(y)
}

func main() {
  x := [...]int {4, 8, 5}
  y := x[0:2] // 4 , 8
  z := x[1:3] // 8 , 5
  y[0] = 1
  z[1] = 3
  z[0] = 4
  fmt.Print(x)
}

func main() {
  x := [...]int {1, 2, 3, 4, 5}
  y := x[0:2]
  z := x[1:4]
  fmt.Print(len(y), cap(y), len(z), cap(z))
}

func main() {
  x := map[string]int {
    "ian": 1, "harris": 2}
  for i, j := range x {
    if i == "harris" {
      fmt.Print(i, j)
    }
  }
}

type P struct {
    x string
    y int
}
func main() {
  b := P{"x", -1}
  a := [...]P
      {
        P{"a", 10},
        P{"b", 2},
        P{"c", 3}
      }

  for _, z := range a {
    if z.y > b.y {
      b = z
    }
  }

  fmt.Println(b.x)
}

func main() {
  s := make([]int, 0, 3)
  s = append(s, 100)
  fmt.Println(len(s), cap(s))
}

// RFC Request for Comments

type Person struct {
	name string
	addr string
	phone string
}

func main(){
  P1 := Person{ name : "Noman Ali", addr : "a at", phone : "1234567890"}
  
  // JSON All Unicode characters

  barr, err := json.Marshal(P1)

  var p2 Person

  errUn := json.Unmarshal(barr, &p2)

  fmt.Println(barr , err , errUn);

  dat , e := ioutil.ReadFile("test.txt")

  data = "Hello Noman Ali Write File"

  err = ioutil.WriteFile("test.txt", data, 0777)

  fmt.Println(dat , e)

  f, err := os.Open("test.txt")
  barr := make([]byte, 10)
  nb, err := f.Read(barr)
  defer f.Close()

  fmt.Println(f,err,nb)

}
