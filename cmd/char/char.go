package main 

import "fmt"

func main(){
  var myByte byte = 'a'
  var myRune rune = 'b'

  fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)
}
