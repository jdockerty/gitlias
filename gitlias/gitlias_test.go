package gitlias

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const actionsPattern = "actions"

type Actions struct {
	suite.Suite
	configFile *os.File
}

func (a *Actions) SetupTest() {
	assert := assert.New(a.T())

	f, err := ioutil.TempFile(a.T().TempDir(), actionsPattern)
	assert.Nil(err)

	testFileContents := `
	[alias]
		[alias.test1]
		user = "test-user"
		email = "test@example.com"

		[alias.test2]
		user = "test-user-2"
		email = "test-user@example.com"
	`

	_, err = f.WriteString(testFileContents)
	assert.Nil(err)

	a.configFile = f // This file is deleted and closed by the test suite automatically at the end of each test.
}

// TestList will ensure that our configured aliases are returned as expected.
func (a *Actions) TestList() {
	aliases, _ := List(a.configFile.Name())

	assert.ElementsMatch(a.T(), aliases, []string{"test1", "test2"})
}

func (a *Actions) TestAdd() {
	assert := assert.New(a.T())

	aliasName := "testAddAlias"
	user := "testAddUser"
	email := "testAddEmail"
	withNewAlias, _ := Add(a.configFile.Name(), aliasName, user, email)

	_, ok := withNewAlias.Alias[aliasName]
	assert.Truef(ok, "%s should exist in the Alias map", aliasName)
	assert.Equal(withNewAlias.Alias[aliasName].User, user)
	assert.Equal(withNewAlias.Alias[aliasName].Email, email)
}

func TestActionsSuite(t *testing.T) {
	suite.Run(t, new(Actions))
}

func TestInvalidPathErrors(t *testing.T) {
	_, err := Get("/does/not/exist")
	assert.Error(t, err, "expected an error with a fake path, got %s", err)
}
