package entities

type Weapon struct {
	BaseEntity
	Name  string
	Power int
	Speed int
}

type WeaponProps struct {
	Name  string
	Power int
	Speed int
}

func (weapon *Weapon) New(props WeaponProps) *Weapon {
	return &Weapon{
		BaseEntity: *weapon.BaseEntity.New(),
		Name:       props.Name,
		Speed:      props.Speed,
		Power:      props.Power,
	}
}
