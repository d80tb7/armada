package exec

import (
	"runtime"
	"testing"
)
import "github.com/stretchr/testify/assert"

func testScript() string {
	if runtime.GOOS == "windows" {
		return "./testdata/test-exec.bat"
	} else {
		return "./testdata/test-exec.sh"
	}
}

func TestAuthenticatorHappyPath(t *testing.T) {

	cd := CommandDetails{
		Cmd: testScript(),
		Env: []EnvVar{
			{Name: "EXEC_TEST_OUTPUT", Value: "aToken"},
			{Name: "EXEC_TEST_EXIT_CODE", Value: "0"},
		},
		Interactive: false,
	}
	a := NewAuthenticator(cd)

	md, err := a.GetRequestMetadata(nil, "")

	assert.Empty(t, err)
	assert.Equal(t, map[string]string{"authorization": "Bearer aToken"}, md)
}

func TestAuthenticatorCmdFails(t *testing.T) {

	cd := CommandDetails{
		Cmd: testScript(),
		Env: []EnvVar{
			{Name: "EXEC_TEST_OUTPUT", Value: "aToken"},
			{Name: "EXEC_TEST_EXIT_CODE", Value: "1"},
		},
		Interactive: false,
	}
	a := NewAuthenticator(cd)

	md, err := a.GetRequestMetadata(nil, "")

	assert.NotEmpty(t, err)
	assert.Empty(t, md)
}

func TestAuthenticatorMissingCmd(t *testing.T) {

	cd := CommandDetails{
		Cmd: "not_a_valid_command.sh",
		Env: []EnvVar{
			{Name: "EXEC_TEST_OUTPUT", Value: "aToken"},
			{Name: "EXEC_TEST_EXIT_CODE", Value: "0"},
		},
		Interactive: false,
	}
	a := NewAuthenticator(cd)

	md, err := a.GetRequestMetadata(nil, "")

	assert.NotEmpty(t, err)
	assert.Empty(t, md)
}

func TestAuthenticatorNoToken(t *testing.T) {

	cd := CommandDetails{
		Cmd: testScript(),
		Env: []EnvVar{
			{Name: "EXEC_TEST_OUTPUT", Value: ""},
			{Name: "EXEC_TEST_EXIT_CODE", Value: "0"},
		},
		Interactive: false,
	}
	a := NewAuthenticator(cd)

	md, err := a.GetRequestMetadata(nil, "")

	assert.NotEmpty(t, err)
	assert.Empty(t, md)
}
