package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"utils"
)

// adapted from https://stackoverflow.com/questions/29366038/looping-iterate-over-the-second-level-nested-json-in-go-lang
func countFromMap(m map[string]interface{}, excludeRed bool) int {
	localCount := 0
	for _, val := range m {
		switch concreteVal := val.(type) {
		case string:
			if excludeRed && concreteVal == "red" {
				return 0
			}
		case map[string]interface{}:
			localCount += countFromMap(val.(map[string]interface{}), excludeRed)
		case []interface{}:
			localCount += countFromArray(val.([]interface{}), excludeRed)
		case int:
			localCount += concreteVal
		case float64:
			localCount += int(concreteVal)
		default:
			panic(val)
		}
	}

	return localCount
}

func countFromArray(array []interface{}, excludeRed bool) int {
	localCount := 0
	for _, val := range array {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			localCount += countFromMap(val.(map[string]interface{}), excludeRed)
		case []interface{}:
			localCount += countFromArray(val.([]interface{}), excludeRed)
		case int:
			localCount += concreteVal
		case float64:
			localCount += int(concreteVal)
		case string:
			// todo
		default:
			panic(val)
		}
	}

	return localCount
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(countFromMap(result, false))
	utils.WriteIntegerOutput(countFromMap(result, false), "1")

	fmt.Println(countFromMap(result, true))
	utils.WriteIntegerOutput(countFromMap(result, true), "2")
}
