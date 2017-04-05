package cat

//Cat a struct with an Id and a Cat picture
type Cat struct {
	ID  string `json:"id"`
	Pic string `json:"pic"`
	Err string `json:"error"`
}

//Cats a slice of beautiful cat
type Cats []Cat
