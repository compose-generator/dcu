package dcu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------------------------------------ DeserializeFromString ------------------------------------

func TestDeserializeFromString_Successful(t *testing.T) {
	composeFile, err := DeserializeFromFile("./media/compose-file-test.yml")
	assert.Nil(t, err)
	assert.Equal(t, "test-frontend-wordpress", composeFile.Services["frontend-wordpress"].ContainerName)
	assert.Equal(t, "phpmyadmin/phpmyadmin:latest", composeFile.Services["db-admin-phpmyadmin"].Image)
	assert.Equal(t, "3306:3306", composeFile.Services["database-mysql"].Ports[0])
}

func TestDeserializeFromString_Failure(t *testing.T) {
	_, err := DeserializeFromFile("./media/not-existing.yml")
	assert.NotNil(t, err)
}

// ------------------------------------ DeserializeFromFile ------------------------------------

