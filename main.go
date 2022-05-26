package main

import (
    "fmt"
    "strconv"
)

var toplevel int = 200 // first letter is lower case: package private
var TOPLEVEL int = 300 // first letter is upper case: public

var (
    name string = "Jelte"
    age int = 26
    major string = "CS"
)

func main() { 
    // block scoped variables inside func
    var thisisfloat32 float32 = 10
    var thisisfloat64 = 10.
    var thisisint = 10
	fmt.Printf("%v %T \n", thisisint, thisisint)
	fmt.Printf("%v %T \n", thisisfloat32, thisisfloat32)
	fmt.Printf("%v %T \n", thisisfloat64, thisisfloat64)
    fmt.Printf("%v %T \n", toplevel, toplevel)

    var name string = "this is okay" 
    fmt.Println(name)
    // this is not: name := "this is not okay"

    var start_as_int int = 4523
    var should_be_float float32
    should_be_float = float32(start_as_int)
    fmt.Printf("%v %T \n", should_be_float, should_be_float)

    var smiley int = 128512 // int -> string conversion takes unicode character
    fmt.Println(string(smiley))

    var ascii_string int = 51 // conversion to ascii string
    fmt.Println(strconv.Itoa(ascii_string))

    primitive_types()
}

func primitive_types() {
    var b bool = false
    fmt.Printf("%v %T\n", b, b)

    var prim_bool bool // uninitialized variable has 0 value
    fmt.Printf("%v %T\n", prim_bool, prim_bool)

    var prim_int int // value 0
    fmt.Printf("%v %T\n", prim_int, prim_int)

    var prim_uint uint16 // value 0. this int is unsigned
    fmt.Printf("%v %T\n", prim_uint, prim_uint)

    arithmatics()
}

func arithmatics() {
    x := 14
    y := 3
    fmt.Printf("x = %v, y = %v\n", x, y)
    fmt.Print("x - y: ")
    fmt.Println(x - y)
    fmt.Print("x / y: ")
    fmt.Println(x / y)
    fmt.Print("x + y: ")
    fmt.Println(x + y)
    fmt.Print("x * y: ")
    fmt.Println(x * y)
    fmt.Print("x % y: ")
    fmt.Println(x % y)
    fmt.Print("x & y: ")
    fmt.Println(x & y)
    fmt.Print("x | y: ")
    fmt.Println(x | y)

    // operating on 2 different types of int is not allowed!

    /*
    var x int8 = 10
    var y int16 = 40
    fmt.Print(x + y) <-- no no!
    */
}
