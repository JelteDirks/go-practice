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
    // primitive_types()
    // constants()
    // numbers_etc()
    arrays_and_slices()
}

func arrays_and_slices() {
    var my_grades = [2]float32 {9.5, 7.8} // contiguous in memory!
    fmt.Printf("My grades: %v\n", my_grades)
    
    var grades_no_size = [...]float32 {9.5, 7.8} // or without size
    grades_no_size[1] = 9.1
    fmt.Printf("Grades no size: %v\n", grades_no_size)
}

func numbers_etc() {
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
func constants() {
    const dont_change int = 62
    fmt.Printf("Don't change this one! %v\n", dont_change)
    
    // const no_no = math.Sin(1.53) <-- needs evaluation, can not be constant!
    // use primitives boi!
    // implicit conversion when const a = 52 <-- replace all occurances of a

    const (
        a = iota // starts at 0
        b = iota
    )

    fmt.Printf("a = %v, b = %v\n", a, b)
    
    // but...

    const (
        _ = iota // this is not stored
        c = iota // starts at 1 now
        d = iota
    )

    fmt.Printf("c = %v, d = %v\n", c, d)

    // iota is a scoped incrementer ... ?

    const (
        lsb = 1 << iota
        second_bit
        third_bit
        msb
    )
    var bit_pattern byte = lsb | second_bit | third_bit | msb
    fmt.Printf("%b = %v\n", bit_pattern, bit_pattern)
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

    var prim_str string = "Jeake!"
    fmt.Printf("My nickname is: %v\n", prim_str)
    fmt.Printf("What type is it actually? prim_str[2]=%v and type %T\n", prim_str[2], prim_str[2])

    // strings are uint8/byte (UTF-8) arrays
    // prim_str[2] = "d" <-- can not assign string to byte!
    // prim_str[5] = "c" <-- but also: can not manipulate string value!

    var prim_str2 string = " or just... Jelte"
    fmt.Println(prim_str + prim_str2) // but we can... concatenate!
    
    var prim_rune rune = 'd' // rune is a type alias for uint32
    fmt.Printf("Some rune shit? %v %T\n", prim_rune, prim_rune)

    var name string = "Jelte"
    var name_in_bytes = []byte(name)
    fmt.Printf("My name in bytes: %v\n", name_in_bytes)
    some_operators()
    some_floatybois()
}
func some_floatybois() {
    var x float32 = 3.14
    var y float32 = 6.022e-23
    var z float32 = 2.9E9 

    fmt.Printf("3.14 = %v\n", x)
    fmt.Printf("6.022e-23 = %v\n", y)
    fmt.Printf("2.9E9 = %v\n", z)
}
func some_operators() {
    var x int8 = 10 // 1010
    var y int8 = 3 // 0011
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

    fmt.Print("x & y: ") // AND: 0010
    fmt.Println(x & y)
    fmt.Print("x | y: ") // OR: 1011
    fmt.Println(x | y)
    fmt.Print("x ^ y: ") // XOR: 1001
    fmt.Println(x ^ y)
    fmt.Print("x &^ y: ") // AND NOT: 1000
    fmt.Println(x &^ y)

    fmt.Print("y << 2: ") // Left SHIFT: 1100
    fmt.Println(y << 2) 
    
    fmt.Print("y >> 2: ") // RIGHT SHIFT: 0000
    fmt.Println(y >> 2) 

    // operating on 2 different types of int is not allowed!
    var a int8 = 10
    var b int16 = 40
    // fmt.Print(a + b) <-- no no!
    fmt.Println(a + int8(b)) // <-- yes yes!
}
