package cat

//Cat a struct with an Id and a Cat picture
type Cat struct {
	ID  string `json:"id"`
	Pic []byte `json:"pic"`
}

//Cats a slice of beautiful cat
type Cats []Cat
