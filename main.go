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
}
