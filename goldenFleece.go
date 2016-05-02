// Package GoldenFleece provides a simple api for handeling JSON structures
// in which the programmer does not exactly know the data model or the data
// model is dynamic and may or may not have extra properties present.
package GoldenFleece

import (
	goJson "encoding/json"
	"errors"
)

// JSON object
type JSON struct {
	data interface{}
	raw  []byte
}

// Load parses a byte array and returns the JSON struct ready to use.
func Load(rawData []byte) (JSON, error) {
	newJSON := new(JSON)
	newJSON.raw = rawData

	// Unmarshal bytes into json
	err := goJson.Unmarshal(rawData, &newJSON.data)
	if err != nil {
		return *newJSON, errors.New("Could not load json")
	}

	return *newJSON, nil
}

// Loads string and get's it ready for parsing
func Loads(rawData string) (JSON, error) {
	return Load([]byte(rawData))
}

// Dump a byte array of the json data(Basically it just gives you the raw data)
// that was used to create the json object
func (j *JSON) Dump() []byte {
	return j.raw
}

// Dumps a string of the json data with a indent value specified
func (j *JSON) Dumps(indent int) string {
	// Generate indent string
	var indentString string
	for i := 0; i < indent; i++ {
		indentString += " "
	}

	// Return indented string value
	data, _ := goJson.MarshalIndent(j.data, "", indentString)
	return string(data)
}

// GetInt will return Int from json
func (j *JSON) GetInt(keys ...string) (int, error) {
	// Get data from path
	value, err := j.followPath(keys...)
	if err != nil {
		return 0, err
	}

	// Type assert to float64
	floatValue, ok := value.(float64)
	if !ok {
		return 0, errors.New("Could not type assert to float64")
	}

	// Cast to int and return
	return int(floatValue), nil
}

// GetString will return String from json
func (j *JSON) GetString(keys ...string) (string, error) {
	// Get data from path
	value, err := j.followPath(keys...)
	if err != nil {
		return "", err
	}

	// Type assert to string
	stringValue, ok := value.(string)
	if !ok {
		return "", errors.New("Could not type assert to string")
	}

	return stringValue, nil
}

// GetBool will return a boolean value from json
func (j *JSON) GetBool(keys ...string) (bool, error) {
	// Get data from path
	value, err := j.followPath(keys...)
	if err != nil {
		return false, err
	}

	// Type assert to bool
	boolValue, ok := value.(bool)
	if !ok {
		return false, errors.New("Could not type assert to boolean")
	}

	return boolValue, nil
}

// GetFloat will return a float value from json
func (j *JSON) GetFloat(keys ...string) (float64, error) {
	// Get data from path
	value, err := j.followPath(keys...)
	if err != nil {
		return 0.0, err
	}

	// Type cast to float64
	floatValue, ok := value.(float64)
	if !ok {
		return 0.0, errors.New("Could not type assert to Float")
	}

	return floatValue, nil
}

// GetArray will return an array from json
func (j *JSON) GetArray(keys ...string) ([]interface{}, error) {
	// Get data from path
	value, err := j.followPath(keys...)
	if err != nil {
		return []interface{}{}, err
	}

	// Type cast to array of interfaces
	interfaceValue, ok := value.([]interface{})
	if !ok {
		return []interface{}{}, errors.New("Could not type assert to array")
	}

	return interfaceValue, nil
}

// GetMap will return a map/dict from json
func (j *JSON) GetMap(keys ...string) (map[string]interface{}, error) {
	// Get data from path
	value, err := j.followPath(keys...)
	if err != nil {
		return map[string]interface{}{}, err
	}

	// Type cast to map
	mapValue, ok := value.(map[string]interface{})
	if !ok {
		return map[string]interface{}{}, errors.New("Could not type assert to map")
	}

	return mapValue, nil
}

// GetArrayMap returns map at index N in array
func (j *JSON) GetArrayMap(index int, keys ...string) (map[string]interface{}, error) {
	// Get data from path
	value, err := j.followPath(keys...)
	if err != nil {
		return map[string]interface{}{}, err
	}

	// Type cast to array
	arrayValue, ok := value.([]interface{})
	if !ok {
		return map[string]interface{}{}, errors.New("Could not type assert to array")
	}

	// Get index value
	if len(arrayValue) < (index + 1) {
		return map[string]interface{}{}, errors.New("Could not get index value")
	}

	// Type cast to map
	mapValue, ok := arrayValue[index].(map[string]interface{})
	if !ok {
		return map[string]interface{}{}, errors.New("Could not type assert to map")
	}

	return mapValue, nil
}

// Private helper functions

// followPath down the json rabit hole
func (j *JSON) followPath(keys ...string) (interface{}, error) {
	data := j.data
	for _, k := range keys {
		tempData, err := getData(k, data)
		if err != nil {
			return nil, err
		}
		data = tempData

	}
	return data, nil
}

// getData from json
func getData(key string, data interface{}) (interface{}, error) {
	isValidMap := false

	//Convert data to map
	switch data.(type) {
	case map[string]interface{}:
		isValidMap = true
	}

	// Check if this is a valid map that we want to follow down the rabbit hole
	if !isValidMap {
		return data, errors.New("Could not follow path")
	}

	tempData := data.(map[string]interface{})

	// Check that the key exists
	value, ok := tempData[key]
	if !ok {
		return nil, errors.New("Key does not exist")
	}

	return value, nil
}
