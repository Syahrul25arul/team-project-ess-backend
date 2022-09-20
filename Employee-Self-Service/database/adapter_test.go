package database

import (
	"employeeSelfService/config"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClientDb(t *testing.T) {
	// prepare database and repository
	config.SetupEnv("../.env")
	config.SanityCheck()

	db := GetClientDb()

	fmt.Println(db)

	assert.NotNil(t, db)
}
