package main

import "fmt"

// ex3 as below

type Attacker interface {
	Name() string
	Attack(tacker DamageTacker)
}

type Player struct {
	name string
	deal int
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Attack(m DamageTacker) {
	m.DealDamage(p, p.deal)
}

type DamageTacker interface {
	DealDamage(Attacker, int)
}

type Monster struct {
	hp int
}

func (m *Monster) DealDamage(attacker Attacker, damage int) {
	m.hp -= damage
	if m.hp <= 0 {
		fmt.Println(attacker.Name(), "가 나를 죽였다")
	}
}

func main() {
	mango := &Player{name: "Mango", deal: 100}
	unit := &Monster{hp: 500}

	for i := 0; i < 5; i++ {
		mango.Attack(unit)
	}
}
