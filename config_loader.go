package filehandlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	YAML = "yaml"
	YML  = "yml"
	JSON = "json"
)

type ConfigLoader struct {
	Name_          string   `json:"name"`
	FileType_      string   `json:"file_type"`
	PathLocations_ []string `json:"path_locations"`
	FinalPath      string   `json:"final_path"`
}

func New() *ConfigLoader {
	return &ConfigLoader{}
}

func (cl *ConfigLoader) Name(name string) *ConfigLoader {
	cl.Name_ = name
	return cl
}

func (cl *ConfigLoader) FileType(fileType string) *ConfigLoader {
	cl.FileType_ = fileType
	return cl
}

func (cp *ConfigLoader) PathLocations(locations ...string) *ConfigLoader {
	cp.PathLocations_ = append(cp.PathLocations_, locations...)
	return cp
}

func (cl *ConfigLoader) Build() *ConfigLoader {
	return cl
}

func (cl ConfigLoader) ToJSONString() (string, error) {
	data, err := json.Marshal(cl)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (cl *ConfigLoader) Parse() (m map[string]interface{}, err error) {
	file, err := cl.Load()
	if err != nil {
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return
	}
	switch cl.FileType_ {
	case YAML, YML:
		err = yaml.Unmarshal(data, &m)
	case JSON:
		err = json.Unmarshal(data, &m)
	}
	return
}

func (cl *ConfigLoader) Load() (file *os.File, err error) {
	if len(cl.PathLocations_) < 1 {
		err = errors.New("path(s) location(s) cannot be empty")
	}
	for _, pl := range cl.PathLocations_ {
		if f, openError := os.Open(fmt.Sprintf("%s/%s.%s", pl, cl.Name_, cl.FileType_)); openError == nil {
			file = f
			cl.FinalPath = pl
			return
		}
	}
	jsonData, jsonConvError := cl.ToJSONString()
	if jsonConvError != nil {
		return nil, jsonConvError
	}
	err = fmt.Errorf("error on load config file with given properties: \n%s", jsonData)
	return
}
