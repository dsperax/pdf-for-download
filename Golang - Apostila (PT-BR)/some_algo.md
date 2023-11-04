## Remove null fields from JSON:

```
func removeNullFields(data []byte) ([]byte, error) {
	var obj interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, err
	}

	cleanedData := removeNull(obj)

	result, err := json.Marshal(cleanedData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func removeNull(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		cleaned := make(map[string]interface{})
		for key, value := range v {
			if value != nil {
				cleaned[key] = removeNull(value)
			}
		}
		return cleaned
	case []interface{}:
		cleaned := make([]interface{}, 0)
		for _, value := range v {
			if value != nil {
				cleaned = append(cleaned, removeNull(value))
			}
		}
		return cleaned
	default:
		return v
	}
}
```

## Remove empity fields from JSON

```
func removeEmptyFields(data []byte) ([]byte, error) {
	var obj interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, err
	}

	cleanedData := removeEmpty(obj)

	result, err := json.Marshal(cleanedData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func removeEmpty(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		cleaned := make(map[string]interface{})
		for key, value := range v {
			if value != nil {
				switch value := value.(type) {
				case string:
					if value != "" {
						cleaned[key] = removeEmpty(value)
					}
				case map[string]interface{}:
					cleanedValue := removeEmpty(value)
					if len(cleanedValue.(map[string]interface{})) > 0 {
						cleaned[key] = cleanedValue
					}
				default:
					cleaned[key] = removeEmpty(value)
				}
			}
		}
		return cleaned
	case []interface{}:
		cleaned := make([]interface{}, 0)
		for _, value := range v {
			if value != nil {
				switch value := value.(type) {
				case string:
					if value != "" {
						cleaned = append(cleaned, removeEmpty(value))
					}
				case map[string]interface{}:
					cleanedValue := removeEmpty(value)
					if len(cleanedValue.(map[string]interface{})) > 0 {
						cleaned = append(cleaned, cleanedValue)
					}
				default:
					cleaned = append(cleaned, removeEmpty(value))
				}
			}
		}
		return cleaned
	default:
		return v
	}
}
```
