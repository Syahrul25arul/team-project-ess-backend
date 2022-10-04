package database

import (
	"employeeSelfService/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClientDb(t *testing.T) {
	// prepare database and repository
	config.SetupEnv("../.env")
	config.SanityCheck()

	db := GetClientDb()
	assert.NotNil(t, db)
}
