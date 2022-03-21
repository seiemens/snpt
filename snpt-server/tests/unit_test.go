/*
Created by: Ramon
Date: 14.3.22
*/
package tests

import (
	util "snpt/lib"
	"testing"
)

//Tests the EnvVariables for its value
func TestEnvVarValue(t *testing.T) {
	key := "ENV_VAR_TEST"
	value := util.GoDotEnvVariable(key)
	expected := "petermeier"
	if value != expected {
		t.Errorf("Value was incorrect, got %s, want %s", value, expected)
	}
}

//tests if the env var vas nonexistent
func TestEnvVarNonExistentVeliu(t *testing.T) {
	key := "ENV_VAR_TEST_NO"
	value := util.GoDotEnvVariable(key)
	expected := "NOT FOUND"
	if value != expected {
		t.Errorf("Value was incorrect, got %s, want %s", value, expected)
	}
}

//tests the RandomString func, which returns a randongenerated string
func TestRandomString(t *testing.T) {
	expected := 5
	value := len(util.GenerateRandomString(5, false))
	if value != expected {
		t.Errorf("Value was incorrect, got %d, want %d", value, expected)
	}
}
