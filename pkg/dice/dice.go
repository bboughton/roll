package dice

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Dice struct {
	Sides int
	Count int
	Mod   int
	Seed  int
}

func (d Dice) Roll() int {
	result := 0
	for i := 0; i < d.Count; i++ {
		result += d.doRoll()
	}
	return result
}

func (d Dice) doRoll() int {
	var seed int64
	if d.Seed > 0 {
		seed = int64(d.Seed)
	} else {
		// Artificially add some delays between invocations of this method
		time.Sleep(3 * time.Nanosecond)
		seed = time.Now().UnixNano()
	}
	r := rand.New(rand.NewSource(seed))
	return r.Intn(d.Sides) + 1 + d.Mod
}

func Parse(v string) (Dice, error) {
	var err error
	mod := 0
	i := strings.Index(v, "+")
	if i > 0 {
		mod, err = strconv.Atoi(v[i+1:])
		if err != nil {
			return Dice{}, errors.New("parse: failed to parse modifier")
		}
		v = v[:i]
	}

	i = strings.Index(v, "d")
	sides, err := strconv.Atoi(v[i+1:])
	if err != nil {
		return Dice{}, errors.New("parse: failed to parse sides")
	}

	count := 1
	if len(v[:i]) > 0 {
		count, err = strconv.Atoi(v[:i])
		if err != nil {
			return Dice{}, errors.New("parse: failed to parse count")
		}
	}

	return Dice{
		Sides: sides,
		Count: count,
		Mod:   mod,
	}, nil
}
