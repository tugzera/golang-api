package entities

type ArmorProps struct {
	Name string
}

type Armor struct {
	BaseEntity
	Name string
}

func (armor *Armor) New(props ArmorProps) *Armor {
	return &Armor{
		BaseEntity: *armor.BaseEntity.New(),
		Name:       props.Name,
	}
}
