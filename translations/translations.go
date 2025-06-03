package translations

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	"gitlab.com/conexxxion/conexxxion-backoffice/config"
	"gopkg.in/yaml.v2"
)

var translations map[string]map[string]string

func GetTranslation(language string, sentence string) string {
	if translations == nil {
		return sentence
	}
	trans, exists := translations[language]
	if !exists {
		return sentence
	}
	sentenceTrans, ok := trans[sentence]
	if !ok {
		return sentence
	}
	return sentenceTrans
}

func ExistsLanguage(language string) bool {
	_, exists := translations[language]
	return exists
}

func LoadTranslations() {
	translations = make(map[string]map[string]string)
	r, err := regexp.Compile("^.+\\.[a-z]{2}\\.(yml|yaml)$")
	if err != nil {
		log.Fatal(err)
	}
	transPath := config.GetConfig().TranslationsPath
	if !strings.HasSuffix(transPath, "/") {
		transPath += "/"
	}
	files, err := filepath.Glob(transPath + "*")
	if err != nil {
		log.Fatalf("translations | Reading files from filepath: %s | error: %s\n", transPath, err.Error())
	}
	for _, file := range files {
		if r.Match([]byte(file)) {
			parseTransFile(file)
		}
	}
}

func parseTransFile(name string) {
	n := strings.Split(name, ".")
	if len(n) < 3 {
		return
	}
	lan := n[len(n)-2]
	file, err := os.ReadFile(name)
	if err != nil {
		log.Error("translations | reading yaml file | error: ", err.Error())
		return
	}
	var trans map[string]string
	err = yaml.Unmarshal(file, &trans)
	if err != nil {
		log.Errorf("translations | parsing yaml file <%s> | error: %s\n", name, err.Error())
		return
	}
	translations[lan] = trans
}
