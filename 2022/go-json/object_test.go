package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type objectTestSuite struct {
	suite.Suite
}

func (s objectTestSuite) TestCustomerMarshal() {
	samples :=
		[]sample{
			{
				`{}`,
				`{"id":0,"email":"","archived":false,"traits":null,"properties":null}`,
				false,
			},
			{
				`{"traits":[], "properties": {}}`,
				`{"id":0,"email":"","archived":false,"traits":[],"properties":{}}`,
				false,
			},
			{
				`{"traits":null, "properties":null}`,
				`{"id":0,"email":"","archived":false,"traits":null,"properties":null}`,
				false,
			},
			{
				`{"id":12345,"email":"john@mail.com","archived":true,"traits":[5,4,3,2,1],"properties":{"a":"b"}}`,
				`{"id":12345,"email":"john@mail.com","archived":true,"traits":[5,4,3,2,1],"properties":{"a":"b"}}`,
				false,
			},
			{
				`{"properties":{"a":"b"},"traits":[5,4,3,2,1],"id":12345,"email":"john@mail.com","archived":true}`,
				`{"id":12345,"email":"john@mail.com","archived":true,"traits":[5,4,3,2,1],"properties":{"a":"b"}}`,
				false,
			}, // note that serialization order matches struct order as far as fields are concerned
		}

	for i, sm := range samples {
		var c customer
		var err error
		err = json.Unmarshal([]byte(sm.input), &c)
		s.testSampleBody(i, sm, err, c)
	}
}

func (s objectTestSuite) TestCustomerOmitMarshal() {
	samples :=
		[]sample{
			{
				`{}`,
				`{}`,
				false,
			},
			{
				`{"traits":null, "properties":null}`,
				`{}`,
				false,
			},  // null array and map are omitted
			{
				`{"traits":[], "properties":{}}`,
				`{}`,
				false,
			}, // empty array and map are omitted as well
			{
				`{"id":12345,"email":"john@mail.com","archived":true,"traits":[5,4,3,2,1],"properties":{"a":"b"}}`,
				`{"id":12345,"email":"john@mail.com","archived":true,"traits":[5,4,3,2,1],"properties":{"a":"b"}}`,
				false,
			}, // no changes here
			{
				`{"properties":{"a":"b"},"traits":[5,4,3,2,1],"id":12345,"email":"john@mail.com","archived":true}`,
				`{"id":12345,"email":"john@mail.com","archived":true,"traits":[5,4,3,2,1],"properties":{"a":"b"}}`,
				false,
			}, // no changes here
		}

	for i, sm := range samples {
		var c customerOmit
		var err error
		err = json.Unmarshal([]byte(sm.input), &c)
		s.testSampleBody(i, sm, err, c)
	}
}

func (s objectTestSuite) TestCustomerPtrMarshal() {
	samples :=
		[]sample{
			{
				`{}`,
				`{"id":null,"email":null,"archived":null,"traits":null,"properties":null}`,
				false,
			}, // with pointers everything becomes null now
			{
				`{"traits":null, "properties":null}`,
				`{"id":null,"email":null,"archived":null,"traits":null,"properties":null}`,
				false,
			},  // no difference between this case and first case (can't separate between being omitted and unknown)
			{
				`{"traits":[], "properties":{}}`,
				`{"id":null,"email":null,"archived":null,"traits":[],"properties":{}}`,
				false,
			}, // kind of obvious with using pointers without omitempty but is it desirable?
			{
				`{"id":12345,"email":"john@mail.com","archived":true,"traits":[5,4,3,2,1],"properties":{"a":"b"}}`,
				`{"id":12345,"email":"john@mail.com","archived":true,"traits":[5,4,3,2,1],"properties":{"a":"b"}}`,
				false,
			}, // no changes here
			{
				`{"properties":{"a":"b"},"traits":[5,4,3,2,1],"id":12345,"email":"john@mail.com","archived":true}`,
				`{"id":12345,"email":"john@mail.com","archived":true,"traits":[5,4,3,2,1],"properties":{"a":"b"}}`,
				false,
			}, // no changes here
		}

	for i, sm := range samples {
		var c customerPtr
		var err error
		err = json.Unmarshal([]byte(sm.input), &c)
		s.testSampleBody(i, sm, err, c)
	}
}

func (s objectTestSuite) TestCustomerPtrOmitMarshal() {
	samples :=
		[]sample{
			{
				`{}`,
				`{}`,
				false,
			}, // with pointers and omit nulls disappear
			{
				`{"traits":null, "properties":null}`,
				`{}`,
				false,
			},  // nulls still disappear even if set
			{
				`{"traits":[], "properties":{}}`,
				`{"traits":[],"properties":{}}`,
				false,
			}, // we get empty lists and maps which is rather nice
			{
				`{"id":12345,"email":"john@mail.com","archived":true,"traits":[5,4,3,2,1],"properties":{"a":"b"}}`,
				`{"id":12345,"email":"john@mail.com","archived":true,"traits":[5,4,3,2,1],"properties":{"a":"b"}}`,
				false,
			}, // no changes here
			{
				`{"properties":{"a":"b"},"traits":[5,4,3,2,1],"id":12345,"email":"john@mail.com","archived":true}`,
				`{"id":12345,"email":"john@mail.com","archived":true,"traits":[5,4,3,2,1],"properties":{"a":"b"}}`,
				false,
			}, // no changes here
		}

	for i, sm := range samples {
		var c customerPtrOmit
		var err error
		err = json.Unmarshal([]byte(sm.input), &c)
		s.testSampleBody(i, sm, err, c)
	}
}

func (s objectTestSuite) TestEmptiness() {
	var id simpleInt
	var email simpleString
	var archived simpleBool
	var traits simpleIntArray
	var props simpleStrMap

	customer1 := customerPtrOmit{
		ID:         &id,
		Email:      &email,
		Archived:   &archived,
		Traits:     &traits,
		Properties: &props,
	}

	var id2 simpleInt = 0
	var email2 simpleString = ""
	var archived2 simpleBool = false

	customer2 := customerPtrOmit{
		ID:         &id2,
		Email:      &email2,
		Archived:   &archived2,
	}

	var id3 simpleInt
	var email3 simpleString
	var archived3 simpleBool
	var traits3 = simpleIntArray{}
	var props3 = simpleStrMap{}
	customer3 := customerPtrOmit{
		ID:         &id3,
		Email:      &email3,
		Archived:   &archived3,
		Traits:     &traits3,
		Properties: &props3,
	}

	var id4 simpleInt
	var email4 simpleString
	var archived4 simpleBool
	var uninitIntArray simpleIntArray
	var traits4 = &uninitIntArray
	var uninitStrMap simpleStrMap
	var props4 = &uninitStrMap
	customer4 := customerPtrOmit{
		ID:         &id4,
		Email:      &email4,
		Archived:   &archived4,
		Traits:     traits4,
		Properties: props4,
	}

	data, err := json.Marshal(customer1)
	s.NoError(err)

	data2, err := json.Marshal(customer2)
	s.NoError(err)

	data3, err := json.Marshal(customer3)
	s.NoError(err)

	data4, err := json.Marshal(customer4)
	s.NoError(err)

	s.Equal(`{"id":0,"email":"","archived":false,"traits":null,"properties":null}`, string(data))
	s.Equal(`{"id":0,"email":"","archived":false}`, string(data2))
	s.Equal(`{"id":0,"email":"","archived":false,"traits":[],"properties":{}}`, string(data3))
	s.Equal(`{"id":0,"email":"","archived":false,"traits":null,"properties":null}`, string(data4))
}

func (s objectTestSuite) testSampleBody(i int, sm sample, err error, customer interface{}) {
	if sm.unmError {
		s.Error(err, fmt.Sprintf("failed case %d", i))
		return
	} else {
		s.NoError(err, fmt.Sprintf("failed case %d", i))
	}
	fmt.Sprintf(" Input: %s", sm.input)
	fmt.Sprintf("Loaded: %v", customer)

	var data []byte
	data, err = json.Marshal(customer)
	fmt.Sprintf("Marshalled: %s", data)
	s.Equal(sm.output, string(data), fmt.Sprintf("failed case %d", i))
}


func TestObjectSuite(t *testing.T) {
	suite.Run(t, new(objectTestSuite))
}

