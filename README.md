`go get github.com/urban-lib/logging@v1`

### Using

```go
package main

import "github.com/urban-lib/logging"

func main() {
	cfg := logging.Configuration{
		EnableConsole:     true,
		ConsoleJSONFormat: false,
		ConsoleLevel:      "debug",
		EnableFile:        false,
		FileJSONFormat:    false,
		FileLevel:         "error",
		FileLocation:      "/var/log/logging/error.log",
	}
	if err := logging.New(cfg); err != nil {
		panic(err)
	}
	printInfo()
}

func printInfo() {
	logging.Debugf("Debug message")
	logging.Infof("Info message")
	logging.Warnf("Warning message")
	logging.Errorf("Error message")
	logging.Fatalf("Fatal error message")

	logging.WithFields(logging.Fields{"FieldName": "Value"}).Debugf("Request")
	logging.WithFields(logging.Fields{"FieldName": 2}).Infof("Request")
	logging.WithFields(logging.Fields{"FieldName": map[string]interface{}{
		"subField": "subValue",
	}}).Warnf("Request")
	logging.WithFields(logging.Fields{"FieldName": []string{}}).Errorf("Request")
	logging.WithFields(logging.Fields{"FieldName": []int{}}).Fatalf("Request")
}
```