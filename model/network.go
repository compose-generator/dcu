package model

// Network represents the YAML structure of a network configuration in a docker compose file
type Network struct {
	External ExtneralNetwork `yaml:"external,omitempty"`
	Ipam     IPAMNetwork     `yaml:"ipam,omitempty"`
}