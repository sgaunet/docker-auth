package dockerauth_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sgaunet/docker-auth/pkg/dockerauth"
	"github.com/stretchr/testify/assert"
)

// func TestChangeJSON(t *testing.T) {
// 	var generic map[string]interface{}
// 	// generic = make(map[string]interface{})
// 	// Test for changing the JSON data
// 	// Create a new JSON object
// 	jsonData := []byte(`{"name": "John Doe", "age": 25}`)
// 	err := json.Unmarshal(jsonData, &generic)
// 	assert.Nil(t, err)

// 	generic["auth"] = "true"
// 	// Check if the name has been changed
// 	if generic["name"] != `John Doe` {
// 		t.Error("Name not changed")
// 	}

// 	// Marshall the JSON object
// 	newJSON, err := json.Marshal(generic)
// 	assert.Nil(t, err)
// 	assert.Equal(t, `{"age":25,"auth":"true","name":"John Doe"}`, string(newJSON))
// }

func TestAddAuthToDockerConfig(t *testing.T) {
	t.Run("add one auth", func(t *testing.T) {
		// Test for adding the auth to the docker config
		generic := make(map[string]interface{})

		// Add the auth to the JSON object
		err := dockerauth.AddAuthToDockerConfig(generic, "https://index.docker.io/v1/", "login", "password")
		assert.Nil(t, err)

		expected := map[string]interface{}{"auths": map[string]interface{}{"https://index.docker.io/v1/": "bG9naW46cGFzc3dvcmQ="}}
		var expectedJSON []byte
		expectedJSON, err = json.Marshal(expected)
		assert.Nil(t, err)

		resultJSON, err := json.Marshal(generic)
		assert.Nil(t, err)

		// Check if the auth has been added
		if !cmp.Equal(expectedJSON, resultJSON) {
			fmt.Println(cmp.Diff(expectedJSON, resultJSON))
			t.Error("resultJSON not equal to expectedJSON")
		}
	})

	t.Run("add multiple auth", func(t *testing.T) {
		generic := make(map[string]interface{})

		// Add the auth to the JSON object
		err := dockerauth.AddAuthToDockerConfig(generic, "https://index.docker.io/v1/", "login", "password")
		assert.Nil(t, err)
		err = dockerauth.AddAuthToDockerConfig(generic, "registry.gitlab.com", "login", "password")
		assert.Nil(t, err)

		expected := map[string]interface{}{"auths": map[string]interface{}{"https://index.docker.io/v1/": "bG9naW46cGFzc3dvcmQ=", "registry.gitlab.com": "bG9naW46cGFzc3dvcmQ="}}
		var expectedJSON []byte
		expectedJSON, err = json.Marshal(expected)
		assert.Nil(t, err)

		resultJSON, err := json.Marshal(generic)
		assert.Nil(t, err)

		// Check if the auth has been added
		if !cmp.Equal(expectedJSON, resultJSON) {
			fmt.Println(cmp.Diff(expectedJSON, resultJSON))
			t.Error("resultJSON not equal to expectedJSON")
		}
	})

	t.Run("add multiple auth and update one", func(t *testing.T) {
		generic := make(map[string]interface{})

		// Add the auth to the JSON object
		err := dockerauth.AddAuthToDockerConfig(generic, "https://index.docker.io/v1/", "login", "password")
		assert.Nil(t, err)
		err = dockerauth.AddAuthToDockerConfig(generic, "registry.gitlab.com", "login", "password")
		assert.Nil(t, err)
		err = dockerauth.AddAuthToDockerConfig(generic, "registry.gitlab.com", "login2", "password2")
		assert.Nil(t, err)

		expected := map[string]interface{}{"auths": map[string]interface{}{"https://index.docker.io/v1/": "bG9naW46cGFzc3dvcmQ=", "registry.gitlab.com": "bG9naW4yOnBhc3N3b3JkMg=="}}
		var expectedJSON []byte
		expectedJSON, err = json.Marshal(expected)
		assert.Nil(t, err)

		resultJSON, err := json.Marshal(generic)
		assert.Nil(t, err)

		// Check if the auth has been added
		if !cmp.Equal(expectedJSON, resultJSON) {
			fmt.Println(cmp.Diff(expectedJSON, resultJSON))
			t.Error("resultJSON not equal to expectedJSON")
		}
	})
}
