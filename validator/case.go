package validator

func GetValidate(key string, vals ...interface{}) []map[string]interface{} {
	result := []map[string]interface{}{}

	switch key {
	case "name":
		result = append(result, name(vals[0].(string)))
	case "startDate":
		result = append(result, startDate(vals[0].(string)))
	case "endDate":
		result = append(result, endDate(vals[0].(string), vals[1].(string)))
	case "color":
		result = append(result, color(vals[0].(string)))
	case "externalId":
		result = append(result, externalId(vals[0].(string)))
	case "status":
		result = append(result, status(vals[0].(string)))
	case "durationUnit":
		result = append(result, durationUnit(vals[0].(string)))
	}

	return result
}
