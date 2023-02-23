package main

import "fmt"

func main() {
  var a = 10
  var ptr = &a

  fmt.Println("Initial values :")
  fmt.Printf("%v\t\t%T\n", a, a)
  fmt.Printf("%v\t\t%T\n", ptr, ptr)
  fmt.Printf("%v\t\t%T\n", *ptr, *ptr)

  a *= 5
  
  fmt.Println("Initial values :")
  fmt.Printf("%v\t\t%T\n", a, a)
  fmt.Printf("%v\t\t%T\n", ptr, ptr)
  fmt.Printf("%v\t\t%T\n", *ptr, *ptr)

  *ptr *= 5

  fmt.Println("Initial values :")
  fmt.Printf("%v\t\t%T\n", a, a)
  fmt.Printf("%v\t\t%T\n", ptr, ptr)
  fmt.Printf("%v\t\t%T\n", *ptr, *ptr)

  secondPoiter := &ptr

  fmt.Println("Initial values :")
  fmt.Printf("%v\t\t%T\n", a, a)
  fmt.Printf("%v\t\t%T\n", ptr, ptr)
  fmt.Printf("%v\t\t%T\n", *ptr, *ptr)
  fmt.Printf("%v\t\t%T\n", secondPoiter, secondPoiter)
  fmt.Printf("%v\t\t%T\n", *secondPoiter, *secondPoiter)

  thirdPtr := &secondPoiter

  fmt.Println("Initial values :")
  fmt.Printf("%v\t\t%T\n", a, a)
  fmt.Printf("%v\t\t%T\n", ptr, ptr)
  fmt.Printf("%v\t\t%T\n", *ptr, *ptr)
  fmt.Printf("%v\t\t%T\n", secondPoiter, secondPoiter)
  fmt.Printf("%v\t\t%T\n", *secondPoiter, *secondPoiter)
  fmt.Printf("%v\t\t%T\n", thirdPtr, thirdPtr)
  fmt.Printf("%v\t\t%T\n", **thirdPtr, **thirdPtr)



}