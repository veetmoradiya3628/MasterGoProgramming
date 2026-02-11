package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckUsernames(t *testing.T) {
	got := CheckUsername("test123")
	if got != true {
		assert.Equal(t, true, got, "got %v\n want %v", got, true)
	}
}

// Table driven testing
func TestCheckUsernameTable(t *testing.T) {

	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{"too short", "test1", false},
		{"empty", "", false},
		{"contains admin", "", false},
		{"valid", "getusername", true},
	}

	for _, tc := range testCases {
		got := CheckUsername(tc.input)
		if got != tc.want {
			assert.Equal(t, tc.want, got, "got %v\n want %v", got, tc.want)
		}
	}
}

func TestCheckUsernameTableWithSubTest(t *testing.T) {

	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{"too short", "test1", false},
		{"empty", "", false},
		{"contains admin", "", false},
		{"valid", "getusername", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CheckUsername(tc.input)
			assert.Equal(t, tc.want, got, "got %v\n want %v", got, tc.want)
		})
	}
}

func TestLogin(t *testing.T) {
	err, ok := Login("testusername")
	assert.NoError(t, err)
	assert.True(t, ok)
}
