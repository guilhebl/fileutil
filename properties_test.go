package props

import (
	"testing"
)

func TestReadPropertiesFile(t *testing.T) {
	props, err := ReadPropertiesFile("sample_test.properties")
	if err != nil {
		t.Error("Error while reading properties file")
	}

	if props["host"] != "localhost" || props["proxyHost"] != "test" || props["protocol"] != "https://" || props["chunk"] != "" {
		t.Error("Error properties not loaded correctly")
	}

	if props["host"] != "localhost" || props["proxyHost"] != "test" || props["protocol"] != "https://" || props["chunk"] != "" {
		t.Error("Error properties not loaded correctly")
	}

	// should ignore whitespace and comments
	if props["z"] != "" || props["xyz"] != "hello there" {
		t.Error("Error comments and spaced values not loaded correctly")
	}

}