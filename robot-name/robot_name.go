package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const lettersInAlphabet = 26

// Robot is a struct representing a robot
type Robot struct {
	name string
}

// NameProvider generates, contains and provides a sequence of unique names
type NameProvider struct {
	codes  []int
	cursor int
}

var randomSource = rand.NewSource(time.Now().UnixNano())
var randomGenerator = rand.New(randomSource)

var uniqNamesProvider = NewNameProvider(randomGenerator)

// Name returns name of the robot
// it also generates one if the robot doesn't have one at the moment of invocation
func (r *Robot) Name() (string, error) {

	if r.name == "" {
		name, err := uniqNamesProvider.Get()

		if err != nil {
			return "", err
		}

		r.name = name
	}

	return r.name, nil
}

// Reset resets robot to the "factory" state, makes it nameless
func (r *Robot) Reset() {
	r.name = ""
}

// NewNameProvider creates a new instance of NewProvider and populates it with uniq names
func NewNameProvider(rg *rand.Rand) NameProvider {
	p := NameProvider{}

	p.seed(rg, lettersInAlphabet*lettersInAlphabet*1000)

	return p
}

// IsExhausted checks if NameProvider still has unused uniq names
func (p NameProvider) IsExhausted() bool {
	return len(p.codes) <= p.cursor
}

// Get returns a uniq name
func (p *NameProvider) Get() (string, error) {
	if p.IsExhausted() {
		return "", errors.New("the pool of names is exhausted")
	}

	result := p.codes[p.cursor]
	p.cursor++

	return numberToName(result), nil
}

func (p *NameProvider) seed(rg *rand.Rand, maxNames int) {
	p.codes = make([]int, maxNames)

	for i := 0; i < maxNames; i++ {
		p.codes[i] = i
	}

	rg.Shuffle(maxNames, func(i, j int) {
		p.codes[i], p.codes[j] = p.codes[j], p.codes[i]
	})
}

func numberToName(number int) string {
	l1 := number / 26000
	l2 := number % 26000 / 1000
	code := number % 1000

	return string('A'+l1) + string('A'+l2) + fmt.Sprintf("%03d", code)
}
