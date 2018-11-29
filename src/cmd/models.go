package main 

type Input struct {
	Name string
}

type Try struct {
	PartNumber int `xml:"PartNumber" json:"part_number"`
	ETag string `xml:"ETag" json:"etag"` 
}

type CompleteMultipartUpload struct{
	Parts []Try`xml:"Part" json:"parts"`
}


