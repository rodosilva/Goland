package main

import "testing"

//////////////////////// This is a simpler way to do Tests ////////////////////////

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////

//////////////////////// This is the more complex way to do Tests ////////////////////////

// The func has to start with Test___
func TestDivide(t *testing.T) { //*testing.T build in package
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have") //Throwing a testing error
	}
}

func TestBadDivide(t *testing.T) { //*testing.T build in package
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not get an error when we should have") //Throwing a testing error
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////

// Comands //
// go test -cover // To see coverage
// go test -coverprofile=coverage.out && go tool cover -html=coverage.out // To pop up a html coverage
