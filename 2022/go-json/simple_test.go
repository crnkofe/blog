package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type simpleTestSuite struct {
	suite.Suite
}

func (s simpleTestSuite) TestSimpleInt() {
	samples :=
		[]sample{
			{ "", "", true},
		    { "5", "5", false},
			{ "-5", "-5", false},
			{ `null`, "0", false}, // ?! null a number
			{ `{}`, "0", true}, // can't unmarshal object to int
			{ `[]`, "0", true}, // can't unmarshal array to int
			{ `"5"`, "5", true}, // can't unmarshal string to int
			{ "9223372036854775807", "9223372036854775807", false}, // math.MaxInt64
			{ "9223372036854775808", "", true}, //overflow
			{ "5e10", "", true}, // floating point, but also a nice int
			{ "0.1", "", true }, // floating point - makes sense
		}

	for i, sm := range samples {
		var si simpleInt
		var err error
		err = json.Unmarshal([]byte(sm.input), &si)
		s.testSampleBody(i, sm, err, si)
	}
}

func (s simpleTestSuite) TestSimpleString() {
	samples :=
		[]sample{
			{ `"works just fine"`, `"works just fine"`, false},
			{ `"čćžš"`, `"čćžš"`, false},
			{ `null`, `""`, false}, // ?! null to empty string
			{ `{}`, "", true}, // can't unmarshal object to string
			{ `[]`, "", true}, // can't unmarshal array to string
			{ `true`, "", true}, // can't unmarshal string to string
			{ "1", "", true}, // int to string is a no-go
			{ "0.1", "", true }, // floating point to str is a no-go
		}

	for i, sm := range samples {
		var sb simpleString
		var err error
		err = json.Unmarshal([]byte(sm.input), &sb)
		s.testSampleBody(i, sm, err, sb)
	}
}

func (s simpleTestSuite) TestSimpleBool() {
	samples :=
		[]sample{
			{ "true", "true", false},
			{ "false", "false", false},
			{ `null`, "false", false}, // ?! null to false
			{ `{}`, "", true}, // can't unmarshal object to bool
			{ `[]`, "", true}, // can't unmarshal array to bool
			{ `"true"`, "", true}, // can't unmarshal string to bool
			{ "1", "", true}, // int to bool not working, but in some languages this will work
			{ "0.1", "", true }, // floating point - makes sense not to unmarshal to bool
		}

	for i, sm := range samples {
		var sb simpleBool
		var err error
		err = json.Unmarshal([]byte(sm.input), &sb)
		s.testSampleBody(i, sm, err, sb)
	}
}

func (s simpleTestSuite) TestSimpleIntArray() {
	samples :=
		[]sample{
			{ "[1, 2, 3, 4, 5]", "[1,2,3,4,5]", false}, // note the whitespace
			{ "[0]", "[0]", false},
			{ `null`, "null", false}, // null array is a null on out?
			{ `{}`, "", true}, // can't unmarshal object to int array
			{ `[]`, "[]", false},
			{ `"true"`, "", true}, // can't unmarshal string to  int array
			{ "1", "", true}, // can't go from int to int array
			{ "0.1", "", true }, // can't go from floating point to int array
		}

	for i, sm := range samples {
		var sb simpleIntArray
		var err error
		err = json.Unmarshal([]byte(sm.input), &sb)
		s.testSampleBody(i, sm, err, sb)
	}
}

func (s simpleTestSuite) TestSimpleStrMap() {
	samples :=
		[]sample{
			{ `{"map":"to", "something":"else"}`, `{"map":"to","something":"else"}`, false}, // note the whitespace
			{ "{}", "{}", false},
			{ `null`, "null", false}, // null map goes to null
			{ `[]`, "", true},
			{ `"true"`, "", true}, // can't unmarshal string to bool
			{ "1", "", true}, // int to bool not working, but in some languages this will work
			{ "0.1", "", true }, // floating point - makes sense not to unmarshal to bool
			{ `{"c":"c", "b":"b", "a":"a"}`, `{"a":"a","b":"b","c":"c"}`, false}, // ?! marshaller is ordering map entries
		}

	for i, sm := range samples {
		var sb simpleStrMap
		var err error
		err = json.Unmarshal([]byte(sm.input), &sb)
		s.testSampleBody(i, sm, err, sb)
	}
}

func (s simpleTestSuite) testSampleBody(i int, sm sample, err error, loaded interface{}) {
	if sm.unmError {
		s.Error(err, fmt.Sprintf("failed case %d", i))
		return
	} else {
		s.NoError(err, fmt.Sprintf("failed case %d", i))
	}
	fmt.Sprintf(" Input: %s", sm.input)
	fmt.Sprintf("Loaded: %v", loaded)

	var data []byte
	data, err = json.Marshal(loaded)
	fmt.Sprintf("Marshalled: %s", data)
	s.Equal(sm.output, string(data), fmt.Sprintf("failed case %d", i))
}

func TestSimpleSuite(t *testing.T) {
	suite.Run(t, new(simpleTestSuite))
}
