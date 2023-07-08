package test

import (
	"testing"
	"zombie-golang/internal/domain/entities"

	"github.com/stretchr/testify/assert"
)

func TestWeapon(t *testing.T) {
	var weapon entities.Weapon
	createdWeapon := weapon.New(entities.WeaponProps{
		Name:  "Porrete",
		Power: 12,
		Speed: 4,
	})
	assert.Equal(t, createdWeapon.Name, "Porrete")
	assert.Equal(t, createdWeapon.Power, 12)
	assert.Equal(t, createdWeapon.Speed, 4)

}
