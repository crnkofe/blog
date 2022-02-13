package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type sampleCustomer struct {
	ID         *int               `json:"id,omitempty"`
	Email      *string            `json:"email,omitempty"`
	Archived   *bool              `json:"archived,omitempty"`
	Traits     *[]int             `json:"traits,omitempty"`
	Properties *map[string]string `json:"properties,omitempty"`
}

func showcaseNulls() {
	fmt.Println("* demo unmarshal `null` to basic types/list/map then marshal *")
	var si int
	err := json.Unmarshal([]byte(`null`), &si)
	if err != nil {
		log.Fatal(err)
	}

	var sb bool
	err = json.Unmarshal([]byte(`null`), &sb)
	if err != nil {
		log.Fatal(err)
	}

	var sst string
	err = json.Unmarshal([]byte(`null`), &sst)
	if err != nil {
		log.Fatal(err)
	}

	sm := map[string]string{}
	err = json.Unmarshal([]byte(`null`), &sm)
	if err != nil {
		log.Fatal(err)
	}

	siarr := make([]int, 0)
	err = json.Unmarshal([]byte(`null`), &siarr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("int:               %5v | %5s\n", si, marshal(si))
	fmt.Printf("bool:              %5v | %5s\n", sb, marshal(sb))
	fmt.Printf("string:            %5v | %5s\n", sst, marshal(sst))
	fmt.Printf("map[string]string: %5v | %5s\n", sm, marshal(sm))
	fmt.Printf("[]int:             %5v    | %5s\n", siarr, marshal(siarr))
}

func showcaseStructs() {
	fmt.Println("* demo unmarshal then marshal sampleCustomer *")

	var c0 sampleCustomer
	err := json.Unmarshal([]byte(`{"id":12345,"email":"john@kofe.com","archived":true,"traits":[5,4,3,2,1],"properties":{"a":"b"}}`), &c0)
	if err != nil {
		log.Fatal(err)
	}

	var c1 sampleCustomer
	err = json.Unmarshal([]byte(`{}`), &c1)
	if err != nil {
		log.Fatal(err)
	}

	var c2 sampleCustomer
	err = json.Unmarshal([]byte(`{"traits":[], "properties":{}}`), &c2)
	if err != nil {
		log.Fatal(err)
	}

	var c3 sampleCustomer
	c3.Traits = &[]int{}
	err = json.Unmarshal([]byte(`{"properties":null}`), &c3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(`({"id":12345,"email":"john@kofe.com","archived":true,traits":[5,4,3,2,1],"properties":{"a":"b"}} -> ` + "\n")
	fmt.Println(" " + string(marshal(c0)))
	fmt.Printf(`({}                                 -> %s`+"\n", marshal(c1))
	fmt.Printf(`({"traits":[], "properties":{}}     -> %s`+"\n", marshal(c2))
	fmt.Printf(`({"traits":null, "properties":null} -> %s`+"\n", marshal(c3))
}

func showcaseOrder() {
	fmt.Println("* demo unmarshal ordered sampleCustomer *")

	c1str := `{"id":12345,"email":"john@mail.com","archived":true,"traits":[5,4,3,2,1],"properties":{"a":"b"}}`

	var c1 sampleCustomer
	err := json.Unmarshal([]byte(c1str), &c1)
	if err != nil {
		log.Fatal(err)
	}

	c2str := `{"properties":{"a":"b"},"traits":[5,4,3,2,1],"id":12345,"email":"john@mail.com","archived":true}`

	var c2 sampleCustomer
	err = json.Unmarshal([]byte(c2str), &c2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Unmarshal %s ->\n", c1str)
	fmt.Printf("\t  %s\n", marshal(c1))

	fmt.Printf("Unmarshal %s ->\n", c2str)
	fmt.Printf("\t  %s\n", marshal(c2))
}

func showcaseDuplicates() {
	fmt.Println("* demo duplicate fields *")

	c1str := `{"id":12345,"email":"john@mail.com","id":67891,"email":"squirrel@mail.com"}`
	var c1 sampleCustomer
	err := json.Unmarshal([]byte(c1str), &c1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Unmarshal %s ->\n", c1str)
	fmt.Printf("\t  %s\n", marshal(c1))
}

func showcaseArrayTypes() {
	fmt.Println("* demo multi-type array *")

	strNums := `[true,1,"test",3000.0,{"a":"b"},[2,3]]`
	var nums []interface{}
	err := json.Unmarshal([]byte(strNums), &nums)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Unmarshal %s ->\n", strNums)
	fmt.Printf("\t  %s\n", marshal(nums))

	for _, n := range nums {
		fmt.Printf("%5v: %s\n", n, reflect.TypeOf(n))
	}
}

func marshal(si interface{}) []byte {
	result, err := json.Marshal(si)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func main() {
	showcaseNulls()
	fmt.Println()
	showcaseStructs()
	fmt.Println()
	showcaseOrder()
	fmt.Println()
	showcaseDuplicates()
	fmt.Println()
	showcaseArrayTypes()
}
