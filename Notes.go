1. if else mellodrama:
https://go.dev/play/p/MaYePDBFREk

2. marshalling -> converting go objects into json is called marshalling , and the opposite of it is called unmarshalling
https://go.dev/play/p/Z9mcyYrRTOP

3. marshalling and unmarshalling of maps
https://go.dev/play/p/ZmdS0Hffwk8

package main

import (
 "fmt"
)

func main() {
 byteArray := []byte{97, 98, 99, 100, 101, 102}
 fmt.Println(byteArray)
}
Output

[97 98 99 100 101 102]

5. Redeclartion of variables

    // this works 
    a,b := 4,5
    fmt.Println(a,b)
    b,c := 6,7
    fmt.Println(b,c)

    // this does not work
    a := 5
    a := 6

    because 
    You can't redeclare a variable which has already been declared in the same block. However, variables can be redeclared in short multi-variable declarations where at least one new variable is introduced.

  

