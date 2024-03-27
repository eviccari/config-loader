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

type FileProperties struct {
	Name          string   `json:"name"`
	Type          string   `json:"type"`
	PathLocations []string `json:"path_locations"`
	FinalPath     string   `json:"final_path"`
}

func NewFileProperties() *FileProperties {
	return &FileProperties{}
}

func (fp *FileProperties) SetName(name string) *FileProperties {
	fp.Name = name
	return fp
}

func (fp *FileProperties) SetType(fileType string) *FileProperties {
	fp.Type = fileType
	return fp
}

func (fp *FileProperties) SetLocations(locations ...string) *FileProperties {
	fp.PathLocations = append(fp.PathLocations, locations...)
	return fp
}

func (fp *FileProperties) Build() *FileProperties {
	return fp
}

func (fp FileProperties) ToJSONString() (string, error) {
	data, err := json.Marshal(fp)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (fp *FileProperties) Parse() (m map[string]interface{}, err error) {
	file, err := fp.Load()
	if err != nil {
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return
	}
	switch fp.Type {
	case YAML, YML:
		err = yaml.Unmarshal(data, &m)
	case JSON:
		err = json.Unmarshal(data, &m)
	}
	return
}

func (fp *FileProperties) Load() (file *os.File, err error) {
	if len(fp.PathLocations) < 1 {
		err = errors.New("path(s) location(s) cannot be empty")
	}
	for _, pl := range fp.PathLocations {
		if f, openError := os.Open(fmt.Sprintf("%s/%s.%s", pl, fp.Name, fp.Type)); openError == nil {
			file = f
			fp.FinalPath = pl
			return
		}
	}
	jsonData, jsonConvError := fp.ToJSONString()
	if jsonConvError != nil {
		return nil, jsonConvError
	}
	err = fmt.Errorf("error on load config file with given properties: \n%s", jsonData)
	return
}
