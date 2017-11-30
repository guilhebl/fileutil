package props

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

type Properties map[string]string

func (p Properties) GetProperty(k string) string {
	return p[k]
}

func (p Properties) GetIntProperty(k string) int64 {
	prop, _ := strconv.ParseInt(p[k], 10, 0)
	return prop
}

// Reads a file in the properties format
func ReadPropertiesFile(filename string) (Properties, error) {
	config := Properties{}

	if len(filename) == 0 {
		return config, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal, comment := strings.Index(line, "="), strings.Index(line, "#"); equal >= 0 && comment == -1 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}
