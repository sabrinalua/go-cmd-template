package main 

func main() {
	a:=App{}
	// v:= gos3.InitiateMultipartUploadResult{Namespace:"http://myexample.com"}
	// c:= gos3.ListPartsResult{}
	// c.Part = make([]gos3.PartInfoWithModded,0)
	// r, _:= gos3.EncodeXml(false, c)
	// fmt.Printf("%s\n",r)
	a.Init()
  	a.Run()
}