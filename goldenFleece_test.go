package GoldenFleece

import "testing"

const json = `{
  "number": 1,
  "int": 2,
  "float": 3.3,
  "string": "Hello World",
  "bool": true,
  "bool2": false,
  "simple-array": ["foo", "bar", "4", 5, 6.6],
  "map": {
    "foo": "bar",
    "num2": 7,
    "array2": ["kung", "foo"],
    "map2": {
      "num3": 99
    }
  },
  "complex-array": [
    {
      "dub": "bub"
    },
    {
      "answer": 42,
      "the-question": "what is the meaning to:",
      "the-question-part-2": ["life", "the universe", "everything"]
    }
  ]
}`

func TestGetInt(t *testing.T) {
	//Arrange
	data, err := Loads(json)
	if err != nil {
		t.Error("Could not load json files")
	}

	// Check simple path
	result, err := data.GetInt("int")
	if result != 2 {
		t.Error("Could not get correct value")

	}

	// Check complex path
	result, err = data.GetInt("map", "map2", "num3")
	if result != 99 {
		t.Error("Could not get correct value")
	}

	// Check no key error
	result, err = data.GetInt("blah")
	if result != 0 || err == nil {
		t.Error("Could not get correct value", result, err)
	}

	// Check bad data type error
	result, err = data.GetInt("string")
	if result != 0 || err == nil {
		t.Error("Could not get correct value")
	}
}

func TestLoads(t *testing.T) {
	_, err := Loads("test")
	if err == nil {
		t.Error("Could not load bad string")
	}
}

func TestDump(t *testing.T) {
	//Arrange
	data, err := Loads(json)
	if err != nil {
		t.Error("Could not load json files")
	}

	result := data.Dump()
	if len(result) == 0 {
		t.Error("Could not dump json")
	}

	resultString := data.Dumps(4)
	if len(resultString) == 0 {
		t.Error("Could not dumps json")
	}
}

func TestGetString(t *testing.T) {
	//Arrange
	data, err := Loads(json)
	if err != nil {
		t.Error("Could not load json files")
	}

	//Assert
	result, err := data.GetString("string")
	if result != "Hello World" {
		t.Error("Could not get correct value")
	}

	result, err = data.GetString("map", "foo")
	if result != "bar" {
		t.Error("Could not get correct value")
	}

	// Check no key error
	result, err = data.GetString("blah")
	if result != "" || err == nil {
		t.Error("Could not get correct value", result, err)
	}

	// Check bad data type error
	result, err = data.GetString("int")
	if result != "" || err == nil {
		t.Error("Could not get correct value")
	}
}

func TestGetBool(t *testing.T) {
	//Arrange
	data, err := Loads(json)
	if err != nil {
		t.Error("Could not load json files")
	}

	//Assert
	result, err := data.GetBool("bool")
	if result == false {
		t.Error("Could not get correct value")
	}

	result, err = data.GetBool("bool2")
	if result == true {
		t.Error("Could not get correct value")
	}

	// Check no key error
	result, err = data.GetBool("blah")
	if result != false || err == nil {
		t.Error("Could not get correct value", result, err)
	}

	// Check bad data type error
	result, err = data.GetBool("int")
	if result != false || err == nil {
		t.Error("Could not get correct value")
	}
}

func TestGetFloat(t *testing.T) {
	//Arrange
	data, err := Loads(json)
	if err != nil {
		t.Error("Could not load json files")
	}

	//Assert
	result, err := data.GetFloat("float")
	if result != 3.3 {
		t.Error("Could not get correct value")
	}

	// Check no key error
	result, err = data.GetFloat("blah")
	if result != 0.0 || err == nil {
		t.Error("Could not get correct value", result, err)
	}

	// Check bad data type error
	result, err = data.GetFloat("string")
	if result != 0.0 || err == nil {
		t.Error("Could not get correct value")
	}
}

func TestGetArray(t *testing.T) {
	//Arrange
	data, err := Loads(json)
	if err != nil {
		t.Error("Could not load json files")
	}

	//Act

	simpleExpectedValues := []interface{}{
		"foo", "bar", "4", 5.0, 6.6,
	}
	result, err := data.GetArray("simple-array")
	for i, v := range result {
		if v != simpleExpectedValues[i] {
			t.Error("Could not get correct value", v)
		}
	}

	complexExpectedValues := []interface{}{
		map[string]interface{}{
			"dub": "bub",
		},
		map[string]interface{}{
			"answer":              42.0,
			"the-question":        "what is the meaning to:",
			"the-question-part-2": []interface{}{"life", "the universe", "everything"},
		},
	}
	result, err = data.GetArray("complex-array")
	for i, v := range result {
		for k, vv := range v.(map[string]interface{}) {
			var ev interface{}
			if k == "the-question-part-2" {
				ev = complexExpectedValues[i].(map[string]interface{})[k]
				for x, vvv := range vv.([]interface{}) {
					if vvv != ev.([]interface{})[x] {
						t.Error("Could not get correct value", vvv)
					}
				}
			} else {
				ev = complexExpectedValues[i].(map[string]interface{})[k]
				if vv != ev {
					t.Error("Counld not get correct value", vv, ev)
				}
			}
		}
	}

	// Check no key error
	result, err = data.GetArray("blah")
	if err == nil {
		t.Error("Could not get correct value", result, err)
	}

	// Check bad data type error
	result, err = data.GetArray("string")
	if err == nil {
		t.Error("Could not get correct value")
	}

	// Check bad data type error
	result, err = data.GetArray("simple-array", "test")
	if err == nil {
		t.Error("Could not get correct value")
	}
}

func TestGetMap(t *testing.T) {
	//Arrange
	data, err := Loads(json)
	if err != nil {
		t.Error("Could not load json files")
	}

	//Assert
	result, err := data.GetMap("map")
	if result["foo"] != "bar" {
		t.Error("Could not get correct value")
	}

	// Check no key error
	result, err = data.GetMap("blah")
	if err == nil {
		t.Error("Could not get correct value", result, err)
	}

	// Check bad data type error
	result, err = data.GetMap("string")
	if err == nil {
		t.Error("Could not get correct value")
	}
}

func TestGetArrayMap(t *testing.T) {
	//Arrange
	data, err := Loads(json)
	if err != nil {
		t.Error("Could not load json files")
	}

	//Assert
	result, err := data.GetArrayMap(1, "complex-array")
	if result["answer"] != 42.0 {
		t.Error("Could not get correct value")
	}

	// Check no key error
	result, err = data.GetArrayMap(1, "blah")
	if err == nil {
		t.Error("Could not get correct value", result, err)
	}

	// Check bad index value
	result, err = data.GetArrayMap(5, "complex-array")
	if err == nil {
		t.Error("Could not get correct value")
	}

	// Check bad data type error
	result, err = data.GetArrayMap(1, "string")
	if err == nil {
		t.Error("Could not get correct value")
	}

	// Check bad data type error
	result, err = data.GetArrayMap(1, "simple-array")
	if err == nil {
		t.Error("Could not get correct value")
	}
}
