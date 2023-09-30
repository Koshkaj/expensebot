package types

type Document struct {
	Id       string `bson:"_id" json:"id"`
	Filename string `bson:"fileName" json:"fileName"`
	MimeType string `bson:"mimeType" json:"mimeType"`
}

type File struct {
	Name      string
	Extension string
	MimeType  string
	Data      []byte
}
