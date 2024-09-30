package dockerauth_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sgaunet/docker-auth/pkg/dockerauth"
	"github.com/stretchr/testify/assert"
)

func TestAddAuthToDockerConfig(t *testing.T) {
	t.Run("add one auth", func(t *testing.T) {
		var expectedJSON []byte
		generic := make(map[string]interface{})
		// Add the auth to the JSON object
		err := dockerauth.AddAuthToDockerConfig(generic, "https://index.docker.io/v1/", "login", "password")
		assert.Nil(t, err)

		expected := map[string]interface{}{"auths": map[string]interface{}{
			"https://index.docker.io/v1/": map[string]interface{}{
				"auth": "bG9naW46cGFzc3dvcmQ=",
			},
		},
		}
		expectedJSON, err = json.Marshal(expected)
		assert.Nil(t, err)

		resultJSON, err := json.Marshal(generic)
		assert.Nil(t, err)

		// Check if the auth has been added
		if !cmp.Equal(expectedJSON, resultJSON) {
			t.Error("resultJSON not equal to expectedJSON")
			fmt.Println(cmp.Diff(expectedJSON, resultJSON))
		}
	})

	t.Run("add multiple auth", func(t *testing.T) {
		generic := make(map[string]interface{})
		err := dockerauth.AddAuthToDockerConfig(generic, "https://index.docker.io/v1/", "login", "password")
		assert.Nil(t, err)
		err = dockerauth.AddAuthToDockerConfig(generic, "registry.gitlab.com", "login", "password")
		assert.Nil(t, err)

		expected := map[string]interface{}{"auths": map[string]interface{}{
			"https://index.docker.io/v1/": map[string]interface{}{
				"auth": "bG9naW46cGFzc3dvcmQ=",
			},
			"registry.gitlab.com": map[string]interface{}{
				"auth": "bG9naW46cGFzc3dvcmQ=",
			},
		},
		}
		var expectedJSON []byte
		expectedJSON, err = json.Marshal(expected)
		assert.Nil(t, err)

		resultJSON, err := json.Marshal(generic)
		assert.Nil(t, err)

		if !cmp.Equal(expectedJSON, resultJSON) {
			t.Error("resultJSON not equal to expectedJSON")
			fmt.Println(cmp.Diff(expectedJSON, resultJSON))
		}
	})

	t.Run("add multiple auth and update one", func(t *testing.T) {
		generic := make(map[string]interface{})

		err := dockerauth.AddAuthToDockerConfig(generic, "https://index.docker.io/v1/", "login", "password")
		assert.Nil(t, err)
		err = dockerauth.AddAuthToDockerConfig(generic, "registry.gitlab.com", "login", "password")
		assert.Nil(t, err)
		err = dockerauth.AddAuthToDockerConfig(generic, "registry.gitlab.com", "login2", "password2")
		assert.Nil(t, err)

		// expected := map[string]interface{}{"auths": map[string]interface{}{"https://index.docker.io/v1/": "bG9naW46cGFzc3dvcmQ=", "registry.gitlab.com": "=="}}
		expected := map[string]interface{}{"auths": map[string]interface{}{
			"https://index.docker.io/v1/": map[string]interface{}{
				"auth": "bG9naW46cGFzc3dvcmQ=",
			},
			"registry.gitlab.com": map[string]interface{}{
				"auth": "bG9naW4yOnBhc3N3b3JkMg==",
			},
		},
		}
		var expectedJSON []byte
		expectedJSON, err = json.Marshal(expected)
		assert.Nil(t, err)

		resultJSON, err := json.Marshal(generic)
		assert.Nil(t, err)

		if !cmp.Equal(expectedJSON, resultJSON) {
			t.Error("resultJSON not equal to expectedJSON")
			fmt.Println(cmp.Diff(expectedJSON, resultJSON))
		}
	})
}

func TestLoadDockerConfigFile(t *testing.T) {
	t.Run("load empty file", func(t *testing.T) {
		generic, err := dockerauth.LoadDockerConfig("testdata/emtpy.json")
		assert.NotNil(t, err)
		assert.Nil(t, generic)
	})

	t.Run("load non empty file", func(t *testing.T) {
		generic, err := dockerauth.LoadDockerConfig("testdata/emptyjson.json")
		assert.Nil(t, err)
		assert.NotNil(t, generic)
	})

	t.Run("load invalid json file", func(t *testing.T) {
		generic, err := dockerauth.LoadDockerConfig("testdata/invalidjson.json")
		assert.NotNil(t, err)
		assert.Nil(t, generic)
	})
}

func TestSaveDockerConfigFile(t *testing.T) {
	t.Run("save empty json", func(t *testing.T) {
		testFile := "/tmp/emptyjson.json"
		err := dockerauth.SaveDockerConfig(testFile, make(map[string]interface{}))
		assert.Nil(t, err)
		payload, err := dockerauth.LoadDockerConfig(testFile)
		assert.Nil(t, err)
		assert.NotNil(t, payload)
		assert.Equal(t, 0, len(payload))
		// cleanup
		_ = os.Remove(testFile)
	})

	t.Run("save valid auth", func(t *testing.T) {
		testFile := "/tmp/validauth.json"
		payload := make(map[string]interface{})
		err := dockerauth.AddAuthToDockerConfig(payload, "https://index.docker.io/v1/", "login", "password")
		assert.Nil(t, err)
		err = dockerauth.SaveDockerConfig(testFile, payload)
		assert.Nil(t, err)
		payload, err = dockerauth.LoadDockerConfig(testFile)
		assert.Nil(t, err)
		assert.NotNil(t, payload)
		assert.Equal(t, 1, len(payload["auths"].(map[string]interface{})))
		// cleanup
		_ = os.Remove(testFile)
	})
}
