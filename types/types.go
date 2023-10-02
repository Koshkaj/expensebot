package types

type Document struct {
	Id           string `bson:"_id" json:"id"`
	Filename     string `bson:"fileName" json:"fileName"`
	MimeType     string `bson:"mimeType" json:"mimeType"`
	FileDataJSON string `bson:"fileDataJSON" json:"fileDataJSON"`
}

type File struct {
	Id        string
	Name      string
	Extension string
	MimeType  string
	Data      []byte
}
