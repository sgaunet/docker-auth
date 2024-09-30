package dockerauth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
)

// DefaultConfigFile is the default path to the docker config file
var DefaultConfigFile = "~/.docker/config.json"

// EncodeLoginPassword encodes the login and password into a base64 string
func EncodeLoginPassword(login, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(login + ":" + password))
}

// LoadDockerConfig loads the docker config file and returns the JSON object
// The JSON object is a map of string to interface
// If configFile is empty, it defaults to ~/.docker/config.json
func LoadDockerConfig(configFile string) (map[string]interface{}, error) {
	generic := make(map[string]interface{})

	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

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
	if payload == nil {
		return errors.New("payload is nil")
	}
	if payload["auths"] == nil {
		payload["auths"] = make(map[string]interface{})
	}
	payload["auths"].(map[string]interface{})[registry] = map[string]interface{}{
		"auth": EncodeLoginPassword(login, password),
	}
	return nil
}
