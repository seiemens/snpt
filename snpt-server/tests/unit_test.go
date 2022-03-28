/*
Created by: Ramon
Date: 14.3.22
*/
package tests

import (
	"context"
	util "snpt/lib"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

//test to see if you can read snpts from database
func TestGetMongoSnippetByKey(t *testing.T) {
	expected := "hello2"
	Uri := util.GoDotEnvVariable("MONGO_URL")
	util.Client, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI(Uri))
	value := util.GetMongoSnippetByKey("id", "D9Cq8V")
	if value[0].Content != expected {
		t.Errorf("Value was incorrect, got %s, want %s", value[0].Content, expected)
	}
}
