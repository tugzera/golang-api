package entities

import "errors"

type Zombie struct {
	BaseEntity
	Name   string
	Level  int
	Armor  *Armor
	Weapon *Weapon
}

type ZombieProps struct {
	Name  string
	Level int
}

func (zombie *Zombie) New(props ZombieProps) *Zombie {
	return &Zombie{
		BaseEntity: *zombie.BaseEntity.New(),
		Name:       props.Name,
		Level:      props.Level,
	}
}

func (zombie *Zombie) AddWeapon(weapon Weapon) {
	zombie.Weapon = &weapon
}

func (zombie *Zombie) AddArmor(armor Armor) {
	zombie.Armor = &armor
}

func (zombie *Zombie) CalculateDamage() (int, error) {
	if zombie.Weapon == nil {
		return 0, errors.New("Weapon is not defined")
	}
	return zombie.Weapon.Power * zombie.Weapon.Speed * zombie.Level, nil
}
