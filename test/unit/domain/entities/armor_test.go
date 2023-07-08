package test

import (
	"testing"
	"zombie-golang/internal/domain/entities"

	"github.com/stretchr/testify/assert"
)

func TestArmor(t *testing.T) {
	var armor entities.Armor
	createdArmor := armor.New(entities.ArmorProps{
		Name: "Test",
	})

	assert.Equal(t, createdArmor.Name, "Test")
	assert.NotEmpty(t, createdArmor.Id)
	assert.NotEmpty(t, createdArmor.CreatedAt)
	assert.Nil(t, createdArmor.UpdatedAt)
	assert.Nil(t, createdArmor.DeletedAt)
}
