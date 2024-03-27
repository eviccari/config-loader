package configloader

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestShouldCreateNewFileProperties(t *testing.T) {
	cl := New().
		Name("app_test").
		FileType(YAML).
		PathLocations(".", "./configs-test").
		Build()

	if cl == nil {
		t.Errorf("FileProperties instance cannot be nil")
	}
}

func TestShouldParseYAMLConfigFile(t *testing.T) {
	cl := New().
		Name("app_test").
		FileType(YAML).
		PathLocations(".", "./configs").
		Build()

	configs, err := cl.Parse()
	if err != nil {
		t.Errorf(err.Error())
	}
	validateFields(configs, t)
}

func TestShouldGenerateJSONString(t *testing.T) {
	cl := New().
		Name("app_test").
		FileType(YAML).
		PathLocations(".", "./configs").
		Build()

	jsonString, err := cl.ToJSONString()
	if err != nil {
		t.Errorf(err.Error())
	}
	newCL := &ConfigLoader{}
	if err = json.Unmarshal([]byte(jsonString), newCL); err != nil {
		t.Errorf(err.Error())
	}
	if el := compare(*cl, *newCL); len(el) > 0 {
		t.Errorf("structures are not equal: \n%v", el)
	}
}

func compare(cl1, cl2 ConfigLoader) (errorsList []error) {
	appendErrors := func(err error) {
		if err != nil {
			errorsList = append(errorsList, err)
		}
	}
	data1, err := json.Marshal(cl1)
	appendErrors(err)
	data2, err := json.Marshal(cl2)
	appendErrors(err)
	if string(data1) != string(data2) {
		appendErrors(fmt.Errorf("structs values are different"))
	}
	return
}

func TestShouldNotLoadFile(t *testing.T) {
	cl := New().
		Name("app_test").
		FileType(YAML).
		PathLocations().
		Build()

	if _, err := cl.Load(); err == nil {
		t.Errorf("expected error: %s, got \"\"", "path(s) location(s) cannot be empty")
	}
}

func validateFields(configs map[string]interface{}, t *testing.T) {
	appName, ok := configs["app_name"]
	if !ok {
		t.Errorf("expect that configs contains AppName")
	}
	if appName != "my-app" {
		t.Errorf("expected AppName value: %s, got %s", "my-app", appName)
	}
	appVersion, ok := configs["app_version"]
	if !ok {
		t.Errorf("expect that configs contains AppVersion")
	}
	if appVersion != "1.0.0" {
		t.Errorf("expected AppVersion: %s, got %s", "1.0.0", appVersion)
	}
	logLevel, ok := configs["log_level"]
	if !ok {
		t.Errorf("expect that configs contains LogLevel")
	}
	if logLevel != -4 {
		t.Errorf("expected LogLevel: %d, got %d", -4, logLevel)
	}
}
