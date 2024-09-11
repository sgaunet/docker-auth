package dockerauth

import (
	"encoding/base64"
	"encoding/json"
	"os"
)

// EncodeLoginPassword encodes the login and password into a base64 string
func EncodeLoginPassword(login, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(login + ":" + password))
}

// LoadDockerConfig loads the docker config file and returns the JSON object
// The JSON object is a map of string to interface
// If configFile is empty, it defaults to ~/.docker/config.json
func LoadDockerConfig(configFile string) (map[string]interface{}, error) {
	var generic map[string]interface{}

	if configFile == "" {
		configFile = "~/.docker/config.json"
	}
	// Read the file
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the JSON
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&generic)
	if err != nil {
		return nil, err
	}

	return generic, nil
}

// SaveDockerConfig saves the docker config file with the given JSON object
// The JSON object is a map of string to interface
// If configFile is empty, it defaults to ~/.docker/config.json
func SaveDockerConfig(configFile string, generic map[string]interface{}) error {
	if configFile == "" {
		configFile = "~/.docker/config.json"
	}
	// Open the file
	file, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the JSON
	encoder := json.NewEncoder(file)
	err = encoder.Encode(generic)
	if err != nil {
		return err
	}

	return nil
}

// AddAuthToDockerConfig adds the auth to the docker config file
// The payload is the JSON object returned by LoadDockerConfig
// The registry is the registry to add the auth to
// The login and password are the credentials to add
func AddAuthToDockerConfig(payload map[string]interface{}, registry, login, password string) error {
	if payload["auths"] == nil {
		payload["auths"] = make(map[string]interface{})
	}
	payload["auths"].(map[string]interface{})[registry] = EncodeLoginPassword(login, password)
	return nil
}
