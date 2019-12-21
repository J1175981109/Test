package main

type hero struct {
	name string
	MaxHp int
	Hp int
	MaxMp int
	Mp int
	Skill []skill
}

type A interface {
	hero
	Do(target B)
	Print()
	AttackNum() int
}
