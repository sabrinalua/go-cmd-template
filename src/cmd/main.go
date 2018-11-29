package main

import (
  "fmt"
  "os"
  flag "github.com/ogier/pflag"
	//"encoding/xml"
	"encoding/json"
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
	ty:= Try{PartNumber:0, ETag:"000"}
	ty1:= Try{PartNumber:1, ETag:"001"}
	cmp:= CompleteMultipartUpload{}
	cmp.Parts=append(cmp.Parts, ty)
	cmp.Parts=append(cmp.Parts,ty1)
	x, err:= json.Marshal(cmp)	
	if err==nil{
		fmt.Printf("xml %s\n",x)
	}else{
		fmt.Print(err)
	}
	
	flag.StringVarP(&input.Name, "name", "n", "", "Enter your name")
}
