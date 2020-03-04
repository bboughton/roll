package dice_test

import (
	"testing"

	"github.com/bboughton/roll/pkg/dice"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name string
		arg  string
		dice dice.Dice
	}{
		{
			name: "one d20",
			arg:  "d20",
			dice: dice.Dice{Sides: 20, Count: 1},
		},
		{
			name: "two d20",
			arg:  "2d20",
			dice: dice.Dice{Sides: 20, Count: 2},
		},
		{
			name: "two d6",
			arg:  "2d6",
			dice: dice.Dice{Sides: 6, Count: 2},
		},
		{
			name: "one d4",
			arg:  "1d4",
			dice: dice.Dice{Sides: 4, Count: 1},
		},
		{
			name: "three d6",
			arg:  "3d6",
			dice: dice.Dice{Sides: 6, Count: 3},
		},
		{
			name: "d6 with modifier",
			arg:  "1d6+1",
			dice: dice.Dice{Sides: 6, Count: 1, Mod: 1},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// given
			arg := tc.arg

			// when
			d, err := dice.Parse(arg)

			// then
			assert.Equal(t, tc.dice, d)
			assert.NoError(t, err)
		})
	}
}

func TestParse_Errors(t *testing.T) {
	cases := []struct {
		name string
		arg  string
	}{
		{
			name: "invalid modifier",
			arg:  "d20+a",
		},
		{
			name: "invalid sides",
			arg:  "da",
		},
		{
			name: "invalid count",
			arg:  "ad20",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// given
			arg := tc.arg

			// when
			_, err := dice.Parse(arg)

			// then
			assert.Error(t, err)
		})
	}
}

func TestDice_Roll(t *testing.T) {
	cases := []struct {
		name   string
		dice   dice.Dice
		result int
	}{
		{
			name:   "one d 6",
			dice:   dice.Dice{Sides: 6, Count: 1, Mod: 0, Seed: 1},
			result: 6,
		},
		{
			name:   "two d 6",
			dice:   dice.Dice{Sides: 6, Count: 2, Mod: 0, Seed: 1},
			result: 12,
		},
		{
			name:   "two d 6 with plus 1 modifier",
			dice:   dice.Dice{Sides: 6, Count: 2, Mod: 1, Seed: 1},
			result: 14,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// given
			dice := tc.dice

			// when
			got := dice.Roll()

			// then
			assert.Equal(t, tc.result, got)
		})
	}
}
