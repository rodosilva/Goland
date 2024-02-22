package helpers

import (
	"math/rand"
)

// This is from lesson Packages //
type SomeType struct {
	TypeName   string
	TypeNumber int
}

/////////////////////////////////

func RandomNumber(n int) int {
	value := rand.Intn(n)
	//rand.Seed(time.Now.UnixNano()) //Used to be used on previous versions. Not needed now
	return value

}
