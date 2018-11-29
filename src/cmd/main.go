package main

import (
  "fmt"
  "os"
  flag "github.com/ogier/pflag"
	"encoding/xml"
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

func json_output(args...interface{}){
	encoder("json",args)
}

func xml_output(args...interface{}){
	encoder("xml",args)
}

func encoder (format string , args...interface{})(resp []byte , err error){
	switch format{
		case "xml":
			resp, err = xml.Marshal(args[0]) 
			resp = []byte (xml.Header+ string(resp))
			var x CompleteMultipartUpload
			xml.Unmarshal(resp, &x)
			fmt.Print("calling json encoder from xml\n")
			encoder("json", x)
		case "json":
			resp, err = json.Marshal(args[0])
	}
	if err==nil{
		fmt.Printf("format : %s, digest: %s\n\n", format, resp)
	}else{ fmt.Print(err)}
	
	return
}

func init()  {
	ty:= Try{PartNumber:0, ETag:"000"}
	ty1:= Try{PartNumber:1, ETag:"001"}
	cmp:= CompleteMultipartUpload{}
	cmp.Parts=append(cmp.Parts, ty)
	cmp.Parts=append(cmp.Parts,ty1)
	json_output(cmp)
	xml_output(cmp)
	
	flag.StringVarP(&input.Name, "name", "n", "", "Enter your name")
}
