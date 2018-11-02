package main

import (
  "fmt"
  "os"
  flag "github.com/ogier/pflag"
)
var (
	input Input
)
func main()  {
  flag.Parse()
  if flag.NFlag() == 0 {
     fmt.Printf("Usage: %s [options]\n", os.Args[0])
     fmt.Println("Options:")
     flag.PrintDefaults()
     os.Exit(1)
  }
  if input.Name!=""{
  	fmt.Printf("Hello, %s!\n", input.Name)
  }
}

func init()  {
	flag.StringVarP(&input.Name, "name", "n", "", "Enter your name")
}
