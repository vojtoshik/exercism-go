package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

const lettersInAlphabet = 26

var randomSource = rand.NewSource(time.Now().UnixNano())
var randomGenerator = rand.New(randomSource)

// Robot is a struct representing a robot
type Robot struct {
	name string
}

// Name returns name of the robot
// it also generates one if the robot doesn't have one at the moment of invocation
func (r *Robot) Name() (string, error) {

	if r.name == "" {
		r.name = randLetter() + randLetter() + randNum() + randNum() + randNum()
	}

	return r.name, nil
}

// Reset resets robot to the "factory" state, makes it nameless
func (r *Robot) Reset() {
	r.name = ""
}

func randLetter() string {
	return string(rune('A' + randomGenerator.Intn(lettersInAlphabet)))
}

func randNum() string {
	return strconv.Itoa(randomGenerator.Intn((10)))
}
