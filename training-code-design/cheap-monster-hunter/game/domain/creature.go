package domain

// Creature roles as huger or monster in the game world.
type Creature struct {
	Life         int `json:"life"`
	Damage       int `json:"damage"`
	AttackPower  int `json:"attack_power"`
	DefencePower int `json:"defence_power"`
}

// IsLife means the creature is life or not.
func (c *Creature) IsLife() bool {
	return c.Life > c.Damage
}

// DefendFrom means the creature defend an attack from other creature.
func (c *Creature) DefendFrom(other *Creature) (rest int, isLife bool) {
	c.Damage += other.AttackPower - c.DefencePower
	rest = max(c.Life-c.Damage, 0)
	return rest, rest > 0
}

// AttackTo means the creature attacks to other creature.
func (c *Creature) AttackTo(other *Creature) (rest int, killed bool) {
	rest, isLife := other.DefendFrom(c)
	return rest, !isLife
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// CreatureRace is used generating a creature with template.
type CreatureRace struct {
	BaseCreature Creature
	GrowthRates  Creature
}

// NewCreature generate a creature with level.
func (mr *CreatureRace) NewCreature(level int) *Creature {
	child := &Creature{}
	*child = mr.BaseCreature
	child.Life += mr.GrowthRates.Life * level
	child.AttackPower += mr.GrowthRates.AttackPower * level
	child.DefencePower += mr.GrowthRates.DefencePower * level
	return child
}
