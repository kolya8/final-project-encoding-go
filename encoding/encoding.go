package encoding

import (
	"encoding/json"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	jsonFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonFile, &j.DockerCompose)
	if err != nil {
		return err
	}

	yamlData, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		return err
	}

	yamlFile, err := os.Create(j.FileOutput)
	if err != nil {
		return err
	}

	defer yamlFile.Close()

	_, err = yamlFile.Write(yamlData)
	if err != nil {
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	yamlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &y.DockerCompose)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		return err
	}

	jsonFile, err := os.Create(y.FileOutput)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
