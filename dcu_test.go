package dcu

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v3"
)

// ------------------------------------ DeserializeFromString ------------------------------------

func TestDeserializeFromString_Successful(t *testing.T) {
	// Load input file
	input, err1 := os.Open("./media/compose-file-test.yml")
	bytes, err2 := ioutil.ReadAll(input)
	assert.Nil(t, err1, err2)

	// Execute method
	project, err := DeserializeFromString(string(bytes))

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, 4, len(project.Services))
	for _, service := range project.Services {
		if service.Name == "frontend-wordpress" {
			assert.Equal(t, "test-frontend-wordpress", service.ContainerName)
		} else if service.Name == "db-admin-phpmyadmin" {
			assert.Equal(t, "phpmyadmin/phpmyadmin:latest", service.Image)
		} else if service.Name == "database-mysql" {
			assert.Equal(t, uint32(3306), service.Ports[0].Target)
			assert.Equal(t, uint32(3306), service.Ports[0].Published)
		}
	}
}

func TestDeserializeFromString_Failure(t *testing.T) {
	// Load input file
	input, err1 := os.Open("./media/invalid-compose-file.yml")
	bytes, err2 := ioutil.ReadAll(input)
	assert.Nil(t, err1, err2)

	_, err := DeserializeFromString(string(bytes))
	assert.NotNil(t, err)
}

// ------------------------------------ DeserializeFromFile ------------------------------------

func TestDeserializeFromFile_Successful(t *testing.T) {
	// Execute method
	project, err := DeserializeFromFile("./media/compose-file-test.yml")

	// Assertions
	assert.Nil(t, err)
	assert.NotNil(t, project)

	// Save output files for investigation
	output, err := yaml.Marshal(&project)
	if err != nil {
		return
	}
	os.MkdirAll("./test-output", os.ModePerm)
	ioutil.WriteFile("./test-output/project-output.yml", output, 0777)
}

func TestDeserializeFromFile_Failure1(t *testing.T) {
	_, err := DeserializeFromFile("./media/invalid-compose-file.yml")
	assert.NotNil(t, err)
	assert.Equal(t, "could not parse file", err.Error())
}

func TestDeserializeFromFile_Failure2(t *testing.T) {
	_, err := DeserializeFromFile("./media/not-existing.yml")
	assert.NotNil(t, err)
	assert.Equal(t, "could not open file", err.Error())
}
