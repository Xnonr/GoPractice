// Go Tool Command Line Commands

// Build the executable
// go build Go_Essential_Training.go

// Runs the executable
// ./Go_Essential_Training

// Compiles and executes the code in one go
// go run Go_Essential_Training.go

// Provides a list of Go Tool commands available via the command line terminal
// go help

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {

	// Hello World
	fmt.Println("Hello World!")

	// Numbers

	// Several integer types are available in Go
	// Signed Integers: int, int8, int16, int32, int64
	// Unsigned Integers: uint, uint8, uint16, uint32, uint64

	// Long Variable Declaration

	// Defines a variable
	var x int
	var y int

	// Assigns value to a variable
	x = 1
	y = 2

	// Prints of the types of variables 'x' and 'y'
	fmt.Printf("x = %v, type of %T\n", x, x)
	fmt.Printf("y = %v, type of %T\n", y, y)

	mean1 := (x + y) / 2.0
	fmt.Printf("Result: %v - Type of %T\n", mean1, mean1)

	// Only 2 types of float types in Go
	// float32, float64
	var a float64
	var b float64

	a = 3
	b = 4

	fmt.Printf("a = %v, type of %T\n", a, a)
	fmt.Printf("b = %v, type of %T\n", b, b)

	mean2 := (a + b) / 2.0
	fmt.Printf("Result: %v - Type of %T\n", mean2, mean2)

	// Short Variable Declaration
	// Defines and assigns value to a variable at the same time
	// Go uses type inference and determines on its own what type the variable is given its assigned value
	c := 5
	d := 6.0

	fmt.Printf("a = %v, type of %T\n", c, c)
	fmt.Printf("b = %v, type of %T\n", d, d)

	// Assign multiple values in one line
	e, f := 7, "eight"

	fmt.Printf("a = %v, type of %T\n", e, e)
	fmt.Printf("b = %v, type of %T\n", f, f)

	// You cannot declare a variable and not use it in go, unused variables are not allowed, otherwise a compilation error will occur
	//g := 9

	// Go does not allow mathematical operations between variables of different types, it is a very strict language
	//fmt.Printf("5 + 6.0 = %v", (c + d))

	// Conditionals

	h := 10

	// If Statement
	if h > 5 {
		fmt.Println("The 'h' variable has a value greater than 5.")
	}

	// If - Else Statement
	if h > 10 {
		fmt.Println("The 'h' variable has a value greater than 10.")
	} else {
		fmt.Println("The 'h' variable has a value less than 10.")
	}

	// Logical Operators

	// And
	if h > 5 && h < 15 {
		fmt.Println("The 'h' variable has a value less than 15 but greater than 5.")
	}

	// Or
	if h < 15 || h > 5 {
		fmt.Println("The 'h' variable has a value less than 15 but greater than 5.")
	}

	// Initialization statements are allowed within if statements
	if fraction := a / b; fraction > 0.5 {
		fmt.Println("The 'a' variable is worth more than at least half of that of variable 'b'.")
	}

	// Switch Statement
	i := 11

	switch i {
	case 10:
		fmt.Println("The 'i' variable has a value of 10.")
	case 11:
		fmt.Println("The 'i' variable has a value of 11.")
	case 12:
		fmt.Println("The 'i' variable has a value of 12.")
	default:
		fmt.Println("The 'i' variable has an unknown value.")
	}

	// Looping

	// For Loop
	// The only kind of loop available in Go
	for count1 := 0; count1 < 3; count1++ {
		fmt.Println(count1)
	}

	// Breaking Loops
	// Terminates and stops the loop upon reaching a 'break' statement
	for count2 := 0; count2 < 3; count2++ {

		if count2 > 1 {
			break
		}

		fmt.Println(count2)
	}

	// Continuing Loops
	// Proceeds to the next loop without executing any code which follows upon reaching a 'continue' statement
	for count3 := 0; count3 < 3; count3++ {

		if count3 < 1 {
			continue
		}

		fmt.Println(count3)
	}

	// Another method for writing out a For Loop
	// Similar as to how a While Loop type would be treated in other programming languages
	count4 := 0
	for count4 < 3 {
		fmt.Println(count4)
		count4++
	}

	// Last method for writing out a For Loop
	// Similar as to how a While True or otherwise infinite loop would be treated in other programming languages
	count5 := 0
	for {

		if count5 > 2 {
			break
		}

		fmt.Println(count5)
		count5++
	}

	// Strings
	// Are unicode enabled in Go, can use special characters

	book1 := "Ender's Game"

	// Prints out the variable's String value
	fmt.Println(book1)

	// Prints out the length of the variable's String value
	fmt.Println(len(book1))

	// Strings are immutable in Go

	// String Slicing

	// Slice string[start:end], index starts at 0, 1/2 empty range (last index not included)
	fmt.Println(book1[0:8])

	// Slice[start:], all of the string for the starting index up to the end
	fmt.Println(book1[8:])

	// Slice[:end], all of the string up to the endind index from the start
	fmt.Println(book1[:8])

	// String Concatenation
	fmt.Println(book1[0:] + " by Orson Scott Card")

	// Raw Strings
	// Ignore special characters such as '\n' or '\t' and can be written across multiple lines
	multiLineString1 := `
		Hello
		World
		`
	fmt.Println(multiLineString1)

	// Slices
	// Essentially similar in function to Arrays and Lists in other programming languages

	starWarsCharacters := []string{"Han", "Leia", "Luke"}
	fmt.Printf("Star Wars Characters: %v - Type %T \n", starWarsCharacters, starWarsCharacters)

	fmt.Println(len(starWarsCharacters))

	fmt.Println(starWarsCharacters[1])

	for swcIndex1 := 0; swcIndex1 < len(starWarsCharacters); swcIndex1++ {
		fmt.Println(starWarsCharacters[swcIndex1])
	}

	// Single value range
	for swcIndex2 := range starWarsCharacters {
		fmt.Println(swcIndex2)
	}

	// Double value range
	for swcIndex3, swcName1 := range starWarsCharacters {
		fmt.Printf("%s has the index of %d \n", swcName1, swcIndex3)
	}

	// Double value range whilst ignoring the index via the usage of '_'
	// '_' is a special variable in Go which can be left as undefined
	for _, swcName2 := range starWarsCharacters {
		fmt.Println(swcName2)
	}

	// Appending to Slices
	starWarsCharacters = append(starWarsCharacters, "Sheev")
	fmt.Println(starWarsCharacters)

	// Maps
	// Data structure where keys point to values

	starWarsStarshipSpeeds := map[string]float32{
		"MilleniumFalcon": 2.0,
		"X-Wing":          1.0,
		"TieFighter":      0.0, // Trailing comma required when using multiple lines
	}

	fmt.Println(len(starWarsStarshipSpeeds))

	// Retrieving a value
	fmt.Println(starWarsStarshipSpeeds["MilleniumFalcon"])

	// A zero value is returned when a key is not found within the map
	fmt.Println(starWarsStarshipSpeeds["StarDestroyer"])

	// Determining whether a value exists or is not present within a map
	value1, presentInMap1 := starWarsStarshipSpeeds["StarDestroyer"]
	if !presentInMap1 {
		fmt.Println("StarDestroyer not found.")
	} else {
		fmt.Println(value1)
	}

	// Setting a value
	starWarsStarshipSpeeds["MilleniumFalcon"] = 3
	fmt.Println(starWarsStarshipSpeeds)

	// Deleting a value
	delete(starWarsStarshipSpeeds, "TieFighter")
	fmt.Println(starWarsStarshipSpeeds)

	// Iterating over a Map using a For Loop via a signle value using the map keys
	for mapKey1 := range starWarsStarshipSpeeds {
		fmt.Println(mapKey1)
	}

	for mapKey2, mapValue1 := range starWarsStarshipSpeeds {
		fmt.Printf("%s ->  %.2f \n", mapKey2, mapValue1)
	}

	// Functions

	additionValue := addition(1, 2)
	fmt.Println(additionValue)

	div1, mod1 := modularDivision(7, 2)
	fmt.Printf("div = %d, mod = %d \n", div1, mod1)

	// Behaves like a point, where a copy of the slice is created but it points to the same location in memory
	doubleAtValues := []int{1, 2, 3, 4}
	doubleAt(doubleAtValues, 2)
	fmt.Println(doubleAtValues)

	// Go passes parameters by value, creating a copy of the value and passing it to the function
	doubledValue := 10
	double(doubledValue)
	fmt.Println(doubledValue)

	// Passing in the pointer with the '&' symbol
	doubleWithPointer(&doubledValue)
	fmt.Println(doubledValue)

	// Error Returns

	squareRoot1, error1 := squareRoot(2.0)
	if error1 != nil {
		fmt.Printf("ERROR: %s \n", error1)
	} else {
		fmt.Println(squareRoot1)
	}

	squareRoot2, error2 := squareRoot(-2.0)
	if error2 != nil {
		fmt.Printf("ERROR: %s \n", error2)
	} else {
		fmt.Println(squareRoot2)
	}

	// Defer

	// Go has a garbage collector, so memory management is not an issue you typically have to deal with as a programmer in Go
	// The garbage collector will get rid of allocated objects that are no long being used or have a use

	// Defers are called in the reverse order

	// The process in Go of running a function, checking for errors and the defering is very common in Go

	// Structs

	starfighter1 := Starfighter{1, "T-65B X-Wing Starfighter", "Alliance to Restore the Republic"}

	fmt.Println(starfighter1)
	fmt.Printf("%v \n", starfighter1)
	fmt.Println(starfighter1.Name)

	// Can assign structures' values by their variables and not just their positions
	starfighter2 := Starfighter{
		Allegiance: "First Galactic Empire",
		Name:       "TIE/LN Space Superiority Starfighter",
		ID:         2,
	}

	fmt.Println(starfighter2)

	// Can omit variables, which will default said omitted variables' values to their respective '0' values
	starfighter3 := Starfighter{
		ID:   3,
		Name: "Imperial-Class Star Destroyer",
	}

	fmt.Println(starfighter3)

	// Methods
	// Methods can be defined on structures

	fmt.Println(starfighter1.getStarfighterAllegiance())
	starfighter1.setStarfighterAllegiance("New Republic")

	fmt.Println(starfighter1.getStarfighterAllegiance())

	// Creating Structures with New Functions
	// Allows for the carrying out of data validation

	starfighter4, error3 := NewStarfighter(4, "Variable Geometry Self-Propelled Battle Droid, Mark I", "Confederacy of Independent Systems")

	if error3 != nil {
		fmt.Println("ERROR: ", error3)
	} else {
		fmt.Printf("%v \n", starfighter4)
	}

	// Interface
	// Provides functions whos actions can apply to multiple types of objects

	// Generics
	// Iterfaces can now become sets of types
	fmt.Println(min([]float64{2, 1, 3}))
	fmt.Println(min([]string{"B", "A", "C"}))

	// Panic
	// Generally in Go, errors are used to signal a problem, but not all issue can be handled via errors

	values1 := []int{1, 2, 3}
	nonExistingValue := values1[10] // This will cause a Panic due to an index out of rance error
	fmt.Println(nonExistingValue)

	// Panics can be created on demand
	// Not considered idiomatic Go, so don't Panic whenever possible
	//panic("Luke is in a panic.")

	value2, error4 := safeValue(values1, 10)
	if error4 != nil {
		fmt.Println("ERROR: ", error4)
	}
	fmt.Println("Value 2: ", value2)

	// GoRoutines
	// Go's concurrency routines
	// Very light, easily able to spin up 1000s of them on a single machine as opposed to processes or threads

	urls1 := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}

	// Serial Version
	start1 := time.Now()
	sitesSerial(urls1)
	fmt.Println(time.Since(start1))

	// Concurrent Version
	start2 := time.Now()
	sitesConcurrent(urls1)
	fmt.Println(time.Since(start2))

	// Channels
	// Prefered method to communicate between Go Routines
	// Think of it as a 1 way directional pipe with a single allowed data type
	// receive <- [channel] <- send
	// There must be both a send and receiver, a lack of either one will result in the channel being blocked

	// Creates the Channel
	channel1 := make(chan int)

	// Sends information via the Channel
	// Blocks everything until someone reads from the Channel, a Panic would appear
	//channel1 <- 353

	// Use a Go Routine to prevent a deadlock and therefore a fatal error
	// Will run concurrently to the Main Go Routine
	go func() {
		channel1 <- 353
	}()

	// Receives information via the Channel
	value3 := <-channel1
	fmt.Printf("Got %d \n", value3)

	const count6 = 3
	// Sending values repeatedly via a Go Routine into the Channel
	go func() {
		for i := 0; i < count6; i++ {
			fmt.Printf("Sending %d \n", i)
			channel1 <- i
			time.Sleep(time.Second)
		}
	}()

	// Receiving the values from the Channel and printing them out
	for i := 0; i < count6; i++ {
		value := <-channel1
		fmt.Printf("Received %d \n", value)
	}

	// Close the signal from the Channel when done
	go func() {
		for i := 0; i < count6; i++ {
			fmt.Printf("Sending %d \n", i)
			channel1 <- i
			time.Sleep(time.Second)
		}
		close(channel1)
	}()

	// Allows for the use of range at the receiving end of the Channel
	for i := range channel1 {
		fmt.Printf("Received %d \n", i)
	}

	// Buffered Channels
	// 1 value can be sent in without blocking, but the 2nd will be
	channel2 := make(chan int, 1)
	channel2 <- 19
	value4 := <-channel2
	fmt.Println(value4)

	// Select
	// Allows for the ability to work with several channels at the same time
	channel3, channel4 := make(chan int), make(chan int)

	go func() {
		channel3 <- 42
	}()

	// Once one of the channels becomes available, both from a sending and receiving standpoint, it will enter a specific case
	select {
	case value4 := <-channel3:
		fmt.Printf("Got %d from Channel 3 \n", value4)
	case value4 := <-channel4:
		fmt.Printf("Got %d from Channel 4 \n", value4)
	}

	// Select is especially useful for Timeouts

	channel5 := make(chan float64)

	go func() {
		time.Sleep(100 * time.Millisecond)
		channel5 <- 3.14
	}()

	select {
	case value4 := <-channel5:
		fmt.Printf("Got %f from Channel 5 \n", value4)

	// If after 20 milliseconds nothing is received from Channel 5, a Timeout is printed
	case <-time.After(20 * time.Millisecond):
		fmt.Println("Timeout")
	}

	// JSON
	var jsonStruct1 = `{
		"user": "Han Solo"
		"type": "deposit"
		"amount": 150.50
	}`

	// Simulates an I/O Reader and file / socket
	stringReader1 := strings.NewReader(jsonStruct1)

	// JSON Decoder created on top of the reader
	decoder1 := json.NewDecoder(stringReader1)

	// The final struct of this process
	var request1 Request

	// Calling the decoder on the pointer to the struct
	if error5 := decoder1.Decode(&request1); error5 != nil {
		log.Fatalf("ERROR: Cannot decode - %s", error5)
	}

	fmt.Printf("Got: %v \n", request1)

	// Creating a response
	previousBalance1 := 1_000_000.0

	response1 := map[string]interface{}{
		"ok":      true,
		"balance": previousBalance1 + request1.Amount,
	}

	// Encoding the response in JSON
	encoder1 := json.NewEncoder(os.Stdout)

	if error6 := encoder1.Encode(response1); error6 != nil {
		log.Fatalf("ERROR: Cannot encode - %s", error6)
	}

	// HTTP Calls

	// GET Request
	response2, error7 := http.Get("https://httpbin.org/get")

	if error7 != nil {
		log.Fatalf("ERROR: Cannot call httpbin.org")
	}

	defer response2.Body.Close()

	io.Copy(os.Stdout, response2.Body)

	// POST Request
	job1 := &Job{
		User:   "Saitama",
		Action: "punch",
		Count:  1,
	}

	// In memory reader and writter
	var buffer1 bytes.Buffer

	// Encodes the data into the buffer
	encoder2 := json.NewEncoder(&buffer1)
	if error8 := encoder2.Encode(job1); error8 != nil {
		log.Fatalf("ERROR: Cannot encode job - %s", error8)
	}

	// 3 Parameters for Post
	// The URL, the content type which is JSON, a body which is an eye reader, in this case the in memory buffer
	response3, error9 := http.Post("https://httpbin.org/post", "application/json", &buffer1)

	if error9 != nil {
		log.Fatalf("Error: Cannot call httpbin.org")
	}

	defer response3.Body.Close()

	io.Copy(os.Stdout, response3.Body)

	// Timeouts & Size Limits

	// Creates a context with a time limit of 3 seconds
	context1, cancel1 := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel1()

	// Giving the previous context, use this method, the URL and the error
	request4, error10 := http.NewRequestWithContext(context1, http.MethodGet, "https://httpbin.org/ip", nil)

	if error10 != nil {
		log.Fatal(error10)
	}

	response4, error11 := http.DefaultClient.Do(request4)

	if error11 != nil {
		log.Fatal(error11)
	}

	defer response4.Body.Close()

	// Define how much data is allowed to be read, in this case 1 megabyte
	const megabyte = 1 << 20

	// Wraps the response body with an I/O Limit Reader, limiting how much data is read
	limitReader1 := io.LimitReader(response4.Body, megabyte)
	io.Copy(os.Stdout, limitReader1)

	// HTTP Server

}

// Functions

// Adds 2 integers together and returns the sum
func addition(num1 int, num2 int) int {
	return num1 + num2
}

// Divides 2 integers by each other, returning the quotient and remainder
func modularDivision(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}

func doubleAt(values []int, index int) {
	values[index] *= 2
}

func double(value int) {
	value *= 2
}

func doubleWithPointer(value *int) {
	*value *= 2
}

// Go functions can return multiple objects
func squareRoot(num float64) (float64, error) {

	if num < 0 {
		// Customary to return a 0 value when an error occurs
		return 0.0, fmt.Errorf("Attempting to find the square root of a negative value (%f)", num)
	}

	// 'nil' returned signifying no error
	return math.Sqrt(num), nil
}

func acquire(name string) (string, error) {
	return name, nil
}

func release(name string) {
	fmt.Printf("Cleaning up %s \n", name)
}

func worker() {
	r1, error := acquire("A")

	if error != nil {
		fmt.Println("ERROR: ", error)
		return
	}

	// This 'defer' is a defered function call, meaning the release will get called only when the worker function exits
	defer release(r1)

	fmt.Println("worker")
}

func NewStarfighter(ID int, Name string, Allegiance string) (*Starfighter, error) {
	if ID == 0 {
		return nil, fmt.Errorf("No ID number assigned.")
	}

	if Name == "" {
		return nil, fmt.Errorf("No Name assigned.")
	}

	newStarfighter := Starfighter{
		ID:         ID,
		Name:       Name,
		Allegiance: Allegiance,
	}

	return &newStarfighter, nil
}

func BlasterFire() string {
	return "PewPewPew!"
}

func StarfighterSoundEffects(sf Starfighter) string {
	return (sf.Name + " of the " + sf.Allegiance + " opens fire. " + BlasterFire())
}

func min[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("Minumum of an empty slice.")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v < m {
			m = v
		}
	}

	return m, nil
}

func safeValue(values []int, index int) (n int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	return values[index], nil
}

func returnType(url string) {
	response, error := http.Get(url)
	if error != nil {
		fmt.Printf("ERROR: %s \n", error)
		return
	}

	defer response.Body.Close()
	contentType := response.Header.Get("content-type")
	fmt.Printf("%s -> %s \n", url, contentType)
}

// Serial version
func sitesSerial(urls []string) {
	for _, url := range urls {
		returnType(url)
	}
}

// Concurrent version
// Starts a Go Routine for every single URL
func sitesConcurrent(urls []string) {
	var waitGroup sync.WaitGroup
	for _, url := range urls {

		// Adds 1 to the wait group for every single Go Routine
		waitGroup.Add(1)
		go func(url string) {
			returnType(url)

			// Go Routine signals it is done
			waitGroup.Done()
		}(url)
	}

	// Waits for all of the Go Routines to complete
	waitGroup.Wait()
}

// Methods

func (sf Starfighter) getStarfighterAllegiance() string {
	return sf.Allegiance
}

// Pointer Receiver Method which will actual change the object it is acting upon
func (sf *Starfighter) setStarfighterAllegiance(newAllegiance string) {
	sf.Allegiance = newAllegiance
}

// Structures

type Starfighter struct {
	ID         int
	Name       string
	Allegiance string
}

// Interfaces

type WeaponsFire interface {
	BlasterFire() string
}

// Generics
type Ordered interface {
	int | float64 | string
}

// JSON

type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

// Job Description
type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}
