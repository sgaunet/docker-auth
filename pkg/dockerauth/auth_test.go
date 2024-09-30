package dockerauth_test

import (
	"encoding/json"
	"fmt"
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
