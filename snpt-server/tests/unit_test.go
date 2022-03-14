package tests

import (
	util "snpt/lib"
	"testing"
)

func TestEnvVarValue(t *testing.T) {
	key := "ENV_VAR_TEST"
	value := util.GoDotEnvVariable(key)
	expected := "petermeier"
	if value != expected {
		t.Errorf("Value was incorrect, got %s, want %s", value, expected)
	}
}

func TestEnvVarNonExistentVeliu(t *testing.T) {
	key := "ENV_VAR_TEST_NO"
	value := util.GoDotEnvVariable(key)
	expected := "NOT FOUND"
	if value != expected {
		t.Errorf("Value was incorrect, got %s, want %s", value, expected)
	}
}

func TestRandomString(t *testing.T) {
	expected := 5
	value := util.GenerateRandomString(5, false)
	if len(value) != expected {
		t.Errorf("Value was incorrect, got %s, want %s", value, expected)
	}
}
