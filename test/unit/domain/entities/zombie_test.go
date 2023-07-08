package test

import (
	"testing"
	"zombie-golang/internal/domain/entities"

	"github.com/stretchr/testify/assert"
)

func TestZombie(t *testing.T) {
	var zombie entities.Zombie
	createdZombie := zombie.New(entities.ZombieProps{
		Name:  "Carlos",
		Level: 10,
	})
	assert.Equal(t, createdZombie.Name, "Carlos")
	assert.Equal(t, createdZombie.Level, 10)

	var weapon entities.Weapon

	createdWeapon := weapon.New(entities.WeaponProps{
		Name:  "Porrete",
		Power: 100,
		Speed: 15,
	})
	createdZombie.AddWeapon(*createdWeapon)

	damage, error := createdZombie.CalculateDamage()
	assert.Equal(t, damage, 15000)
	assert.Equal(t, error, nil)
}
