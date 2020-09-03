package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockInit struct {
}

func TestGetCustomConf(t *testing.T) {
	t.Run("It should return all product", func(t *testing.T) {

		dbUser := "DB_USER"
		dbUserDefault := "defaultvalue"
		result := GetCustomConf(dbUser, dbUserDefault)
		assert.Equal(t, result, "root") // mengecek len result
	})
}

func TestInitDB(t *testing.T) {
	t.Run("It should return all product", func(t *testing.T) {
		result := InitDB()
		assert.NotNil(t, result) // mengecek len result
	})
}
