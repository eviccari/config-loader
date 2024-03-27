package filehandlers

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestShouldCreateNewFileProperties(t *testing.T) {
	fp := NewFileProperties().
		SetName("app_test").
		SetType(YAML).
		SetLocations(".", "./configs-test").
		Build()

	if fp == nil {
		t.Errorf("FileProperties instance cannot be nil")
	}
}

func TestShouldParseYAMLConfigFile(t *testing.T) {
	fp := NewFileProperties().
		SetName("app_test").
		SetType(YAML).
		SetLocations(".", "./configs").
		Build()

	configs, err := fp.Parse()
	if err != nil {
		t.Errorf(err.Error())
	}
	validateFields(configs, t)
}

func TestShouldGenerateJSONString(t *testing.T) {
	fp := NewFileProperties().
		SetName("app_test").
		SetType(YAML).
		SetLocations(".", "./configs").
		Build()

	jsonString, err := fp.ToJSONString()
	if err != nil {
		t.Errorf(err.Error())
	}
	newFP := &FileProperties{}
	if err = json.Unmarshal([]byte(jsonString), newFP); err != nil {
		t.Errorf(err.Error())
	}
	if el := compare(*fp, *newFP); len(el) > 0 {
		t.Errorf("structures are not equal: \n%v", el)
	}
}

func compare(fp1, fp2 FileProperties) (errorsList []error) {
	appendErrors := func(err error) {
		if err != nil {
			errorsList = append(errorsList, err)
		}
	}
	data1, err := json.Marshal(fp1)
	appendErrors(err)
	data2, err := json.Marshal(fp2)
	appendErrors(err)
	if string(data1) != string(data2) {
		appendErrors(fmt.Errorf("structs values are different"))
	}
	return
}

func TestShouldNotLoadFile(t *testing.T) {
	fp := NewFileProperties().
		SetName("app_test").
		SetType(YAML).
		SetLocations().
		Build()

	if _, err := fp.Load(); err == nil {
		t.Errorf("expected error: %s, got \"\"", "path(s) location(s) cannot be empty")
	}
}

func validateFields(configs map[string]interface{}, t *testing.T) {
	appName, ok := configs["AppName"]
	if !ok {
		t.Errorf("expect that configs contains AppName")
	}
	if appName != "my-app" {
		t.Errorf("expected AppName value: %s, got %s", "my-app", appName)
	}
	appVersion, ok := configs["AppVersion"]
	if !ok {
		t.Errorf("expect that configs contains AppVersion")
	}
	if appVersion != "1.0.0" {
		t.Errorf("expected AppVersion: %s, got %s", "1.0.0", appVersion)
	}
	logLevel, ok := configs["LogLevel"]
	if !ok {
		t.Errorf("expect that configs contains LogLevel")
	}
	if logLevel != -4 {
		t.Errorf("expected LogLevel: %d, got %d", -4, logLevel)
	}
}
