package main

type simpleInt int
type simpleString string
type simpleIntArray []int
type simpleStrMap map[string]string
type simpleBool bool

type customer struct {
	ID         simpleInt      `json:"id"`
	Email      simpleString   `json:"email"`
	Archived   simpleBool     `json:"archived"`
	Traits     simpleIntArray `json:"traits"`
	Properties simpleStrMap   `json:"properties"`
}

type customerOmit struct {
	ID         simpleInt      `json:"id,omitempty"`
	Email      simpleString   `json:"email,omitempty"`
	Archived   simpleBool     `json:"archived,omitempty"`
	Traits     simpleIntArray `json:"traits,omitempty"`
	Properties simpleStrMap   `json:"properties,omitempty"`
}

type customerPtr struct {
	ID         *simpleInt      `json:"id"`
	Email      *simpleString   `json:"email"`
	Archived   *simpleBool     `json:"archived"`
	Traits     *simpleIntArray `json:"traits"`
	Properties *simpleStrMap   `json:"properties"`
}

type customerPtrOmit struct {
	ID         *simpleInt      `json:"id,omitempty"`
	Email      *simpleString   `json:"email,omitempty"`
	Archived   *simpleBool     `json:"archived,omitempty"`
	Traits     *simpleIntArray `json:"traits,omitempty"`
	Properties *simpleStrMap   `json:"properties,omitempty"`
}

type sample struct {
	input    string
	output   string
	unmError bool
}
