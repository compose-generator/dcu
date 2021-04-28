package dcu

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	loader "github.com/compose-spec/compose-go/loader"
	spec "github.com/compose-spec/compose-go/types"

	yaml "gopkg.in/yaml.v3"
)

// ---------------------------------------------------------- Serializing / Deserializing ----------------------------------------------------------

// DeserializeFromString takes a yaml string and converts it to a Project object
func DeserializeFromString(yamlString string) (spec.Project, error) {
	return DeserializeFromBytes([]byte(yamlString))
}

// DeserializeFromBytes takes a byte array and converts it to a Project object
func DeserializeFromBytes(bytes []byte) (composeFile spec.Project, err error) {
	config, err := loader.ParseYAML(bytes)
	if err != nil {
		return spec.Project{}, errors.New("could not parse file")
	}
	configDetails := spec.ConfigDetails{
		ConfigFiles: []spec.ConfigFile{
			{Config: config},
		},
	}
	projectRef, err := loader.Load(configDetails)
	if err != nil {
		return spec.Project{}, errors.New("something went wrong while parsing file")
	}
	return *projectRef, nil
}

// DeserializeFromFile reads from a yaml file and converts it to a Project object
func DeserializeFromFile(path string) (spec.Project, error) {
	if !strings.HasSuffix(path, ".yml") && !strings.HasSuffix(path, ".yaml") {
		return spec.Project{}, errors.New("the file must be of file type yml or yaml")
	}
	yamlFile, err := os.Open(path)
	if err != nil {
		return spec.Project{}, errors.New("could not open file")
	}
	bytes, err := ioutil.ReadAll(yamlFile)
	if err != nil {
		return spec.Project{}, errors.New("could not read file")
	}
	return DeserializeFromBytes(bytes)
}

// SerializeToString converts a ComposeFile object to a yaml string
func SerializeToString(project spec.Project) (yamlString string, err error) {
	bytes, err := yaml.Marshal(&project)
	return string(bytes), err
}

// SerializeToFile writes a ComposeFile object to a yaml file
func SerializeToFile(project spec.Project, path string) (err error) {
	project.Services.MarshalYAML()
	if !strings.HasSuffix(path, ".yml") && !strings.HasSuffix(path, ".yaml") {
		return errors.New("the file must be of file type yml or yaml")
	}
	output, err := yaml.Marshal(&project)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(path, output, 0777)
	return
}

// ---------------------------------------------------------------- Helper functions ---------------------------------------------------------------

// GetVolumePathsFromComposeFilePath deserializes a compose file and returns all paths of volumes
func GetVolumePathsFromComposeFilePath(composeFilePath string) []string {
	composeFile, err := DeserializeFromFile(composeFilePath)
	if err != nil {
		panic(err)
	}
	return GetVolumePathsFromComposeFile(composeFile)
}

// GetVolumePathsFromComposeFile returns all paths of volumes
func GetVolumePathsFromComposeFile(project spec.Project) (filePaths []string) {
	for _, service := range project.Services {
		for _, volume := range service.Volumes {
			source := volume.Source
			if strings.HasPrefix(source, "./") || strings.HasPrefix(source, "/") {
				filePaths = append(filePaths, source)
			}
		}
	}
	return
}

// GetEnvFilePathsFromComposeFilePath deserializes a compose file and returns all paths of env files
func GetEnvFilePathsFromComposeFilePath(composeFilePath string) []string {
	composeFile, err := DeserializeFromFile(composeFilePath)
	if err != nil {
		panic(err)
	}
	return GetEnvFilePathsFromComposeFile(composeFile)
}

// GetEnvFilePathsFromComposeFile returns all paths of env files
func GetEnvFilePathsFromComposeFile(project spec.Project) (filePaths []string) {
	for _, service := range project.Services {
		for _, envFilePath := range service.EnvFile {
			filePaths = appendStringToSliceIfMissing(filePaths, envFilePath)
		}
	}
	return
}
