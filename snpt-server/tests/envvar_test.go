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
	} else {
		t.Log("env var test successful")
	}
}

func TestEnvVarNonExistentVeliu(t *testing.T) {
	key := "ENV_VAR_TEST_NO"
	value := util.GoDotEnvVariable(key)
	expected := "NOT FOUND"
	if value != expected {
		t.Errorf("Value was incorrect, got %s, want %s", value, expected)
	} else {
		t.Log("Non existent env var test successful")
	}
}
