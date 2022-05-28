package main

import (
    "fmt"
    "strconv"
    "io/ioutil"
    "log"
    "net/http"
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
    // arrays_and_slices()
    // structs_and_maps()
    // if_and_switch()
    // loops_and_stuff()
    i_panic()
    defer_and_panic()
    is_this_server()
}

func is_this_server() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
       w.Write([]byte("Awesome!")) 
    })
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        panic(err.Error())
    }
}

func i_panic() { // recovered from the panic!
    fmt.Println("uh oh")
    defer func() {
        if e := recover() ; e != nil {
            fmt.Println(e)
        }
    }()
    panic("deez nuts")
}

func defer_and_panic() {
    fmt.Println("start")
    defer fmt.Println("deferred baby!")
    defer fmt.Println("But wait?") // LIFO is weird, no?
    fmt.Println("end")
    example1()
}

func example1() {
    response, fetchErr := http.Get("http://www.google.com/robots.txt")
    defer response.Body.Close()
    if fetchErr != nil {
        log.Fatal(fetchErr)
    }
    content, readErr := ioutil.ReadAll(response.Body)
    if readErr != nil {
        log.Fatal(readErr)
    }
    fmt.Printf("response=%s", content)
}

func loops_and_stuff() {
    for i := 0 ; i < 2; i++ {
        fmt.Println("simple loop")
    }
    
    for i, j := 0, 0 ; i < 2; i, j = i+1, j+1 {
        fmt.Println("two variable loop!")
    }

    for i := 0 ; i < 5 ; i++ {
        i = 5
        fmt.Println("i is immediately 5")
    }

    i := 1
    for ; i < 2; i++ {
        fmt.Println("no init?")
    }

    for ; i < 5 ; {
        fmt.Println("this takes a while...")
        i++
    }
    
    for i < 6 {
        fmt.Println("this takes a clean while...")
        i++
    }

    for {
        fmt.Println("this will go on forever")
        break // or will it?
    }

    a := [3]int {42, 503, 33}
    for key, value := range a { // iterator much?
        fmt.Printf("key=%v, value=%v\n", key, value) 
    }
}

func if_and_switch() {
    if true {
        fmt.Println("I mean come on...")
    }

    var jeake = map[string]string {
        "name": "Jelte",
        "size": "Medium",
    }

    if name, has_name := jeake["name"] ; has_name {
        fmt.Printf("name is present: %v\n", name) // name is block scoped!
    }
    
    // nothing more special with the ifs

    switch 3 {
    case 2, 10, 5:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    default:
        fmt.Println("default")
    }
    
    i := 40
    switch {
    case i <= 100:
        fmt.Println("<=100")
        fallthrough
    case i <= 1000:
        fmt.Println("<=1000")
    }
}

func structs_and_maps() {
    var person1 map[string]string = map[string]string {
        "name": "Jelte",
        "age": "26",
    }
    fmt.Printf("person1=%v\n", person1)
    
    var person2 = make(map[string]string)
    person2 = map[string]string{"name": "Frenk",}
    fmt.Printf("person2=%v\n", person2);
    fmt.Printf("person1.name=%v\n", person1["name"])
    fmt.Printf("person2.name=%v\n", person2["name"])

    _ , has_age := person2["age"]
    fmt.Printf("Does person2 have age? %v\n", has_age)
    name, has_name := person2["name"]
    fmt.Printf("Does person2 have name? %v %v\n", has_name, name)

    // Structs

    type Person struct {
        age int
        name string
        size string
    }

    var jelte = Person {
        age: 26,
        name: "Jelte",
        size: "Medium",
    }
    fmt.Printf("jelte=%v\n", jelte)

    type Human struct {
        rights bool
        eyes string
    }

    type Student struct {
        Person
        Human
        uni string
        country string
    }

    var jelte2 = Student {
        uni: "TU/e",
        country: "NL",
        Person: Person { age:26, name: "jelte", },
        Human: Human { rights: true, eyes: "blue", },
    }
    fmt.Printf("jelte2=%v\n", jelte2)
}

func arrays_and_slices() {
    /*
    Arrays are handled as values, not as references!
    Copying an array copies all values, not just a reference!
    */
    var my_grades = [2]float32 {9.5, 7.8} // contiguous in memory!
    fmt.Printf("My grades: %v\n", my_grades)
    
    var grades_no_size = [...]float32 {9.5, 7.8} // or without size
    grades_no_size[1] = 9.1
    fmt.Printf("Grades no size: %v\n", grades_no_size)

    var names [3]string 
    names[0] = "Jelte"
    names[1] = "Deez"
    names[2] = "Nuts"

    fmt.Printf("%v has length %v\n", names, len(names))

    new_names := names // so this is a copy?
    new_names[0] = "69"

    fmt.Printf("%v? That is unexpected!\n", new_names)

    pnt_names := &names // so this is a pointer?
    pnt_names[1] = "Pointed!"
    
    fmt.Printf("pnt_names = %v, names = %v\n", pnt_names, names)

    // Slices!
    
    var slicey_boi = []int {2, 4, 8}
    fmt.Printf("slicey_boi=%v, len=%v, cap=%v\n", slicey_boi, 
        len(slicey_boi), cap(slicey_boi))

    var first_two = slicey_boi[:2]
    fmt.Printf("first_two=%v, len=%v, cap=%v\n", first_two, 
        len(first_two), cap(first_two))

    var last_one = slicey_boi[2:]
    fmt.Printf("last_one=%v, len=%v, cap=%v\n", last_one, 
        len(last_one), cap(last_one))

    var made_slice = make([]int, 0, 20)
    fmt.Printf("made_slice=%v, len=%v, cap=%v\n", made_slice, 
        len(made_slice), cap(made_slice))
    
    made_slice = append(made_slice, 4)
    fmt.Printf("made_slice=%v, len=%v, cap=%v\n", made_slice, 
        len(made_slice), cap(made_slice))
    
    made_slice = append(made_slice, 5, 2, 523, 66, 23, 68)
    fmt.Printf("made_slice=%v, len=%v, cap=%v\n", made_slice, 
        len(made_slice), cap(made_slice))

    var spread_me = []int { 4, 523}
    made_slice = append(made_slice, spread_me...) // spread operator baby!
    fmt.Printf("made_slice=%v, len=%v, cap=%v\n", made_slice, 
        len(made_slice), cap(made_slice))
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
