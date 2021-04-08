package dcu

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/compose-generator/dcu/model"

	yaml "gopkg.in/yaml.v3"
)

// DeserializeFromString takes a yaml string and converts it to a ComposeFile object
func DeserializeFromString(yamlString string) (composeFile model.ComposeFile, err error) {
	err = yaml.Unmarshal([]byte(yamlString), &composeFile)
	return
}

// DeserializeFromFile reads from a yaml file and converts it to a ComposeFile object
func DeserializeFromFile(path string) (composeFile model.ComposeFile, err error) {
	if !strings.HasSuffix(path, ".yml") && !strings.HasSuffix(path, ".yaml") {
		return model.ComposeFile{}, errors.New("the file must be of file type yml or yaml")
	}
	yamlFile, err := os.Open(path)
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(yamlFile)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(bytes, &composeFile)
	return
}

// SerializeToString converts a ComposeFile object to a yaml string
func SerializeToString(composeFile model.ComposeFile) (yamlString string, err error) {
	bytes, err := yaml.Marshal(&composeFile)
	return string(bytes), err
}

// SerializeToFile wriet a ComposeFile object to a yaml file
func SerializeToFile(composeFile model.ComposeFile, path string) (err error) {
	if !strings.HasSuffix(path, ".yml") && !strings.HasSuffix(path, ".yaml") {
		return errors.New("the file must be of file type yml or yaml")
	}
	output, err := yaml.Marshal(&composeFile)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(path, output, 0777)
	return
}
