package domain

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestNewCreatureFromRace(t *testing.T) {
	race := CreatureRace{
		BaseCreature: Creature{
			Life:         10,
			AttackPower:  10,
			DefencePower: 10,
		},
		GrowthRates: Creature{
			Life:         3,
			AttackPower:  3,
			DefencePower: 3,
		},
	}
	child := race.NewCreature(7)
	if child.Life != 31 {
		t.Errorf("life not match %d != %d", child.Life, 31)
	}
	if child.AttackPower != 31 {
		t.Errorf("AttackPower not match %d != %d", child.AttackPower, 31)
	}
	if child.DefencePower != 31 {
		t.Errorf("DefencePower not match %d != %d", child.DefencePower, 31)
	}
}

func TestIsLife(t *testing.T) {
	c := &Creature{}
	if c.IsLife() {
		t.Errorf("zero value of creature is dead but live!! %v", c)
	}
	c.Life = 1
	if !c.IsLife() {
		t.Errorf("creature has life > 0 live but dead %v", c)
	}
}
