package validator

import (
	"math"
	"regexp"
	"strings"
	"time"
)

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func color(color string) map[string]interface{} {
	match, _ := regexp.MatchString(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`, color)

	if !match {
		return map[string]interface{}{
			"error":   true,
			"message": "Color must be a valid hex color code. Example: #FF0000",
		}
	}

	return map[string]interface{}{
		"error":   false,
		"message": "Status is valid",
	}
}

func name(name string) map[string]interface{} {
	if len(name) > 255 {
		return map[string]interface{}{
			"error":   true,
			"message": "Name must be less than 255 characters",
		}
	}

	return map[string]interface{}{
		"error":   false,
		"message": "Status is valid",
	}
}

func startDate(startDate string) map[string]interface{} {
	_, err := time.Parse("2006-01-02T15:04:05Z", startDate)
	if err != nil {
		return map[string]interface{}{
			"error":   true,
			"message": "Start date must be a valid ISO8601 date. Example: 2022-12-31T14:59:00Z",
		}
	}

	return map[string]interface{}{
		"error":   false,
		"message": "Status is valid",
	}
}

func endDate(endDate string, startDate string) map[string]interface{} {
	endDateTime, err := time.Parse("2006-01-02T15:04:05Z", endDate)
	if err != nil {
		return map[string]interface{}{
			"error":   true,
			"message": "End date must be a valid ISO8601 date. Example: 2022-12-31T14:59:00Z",
		}
	}

	startDateTime, _ := time.Parse("2006-01-02T15:04:05Z", startDate)

	if endDateTime.Before(startDateTime) {
		return map[string]interface{}{
			"error":   true,
			"message": "End date must be greater than start date",
		}
	}

	return map[string]interface{}{
		"error":   false,
		"message": "Status is valid",
	}
}
func externalId(externalID string) map[string]interface{} {
	if len(externalID) > 255 {
		return map[string]interface{}{
			"error":   true,
			"message": "External ID must be less than 255 characters",
		}
	}

	return map[string]interface{}{
		"error":   false,
		"message": "Status is valid",
	}
}

func status(status string) map[string]interface{} {
	validStatuses := []string{"NEW", "PLANNED", "DELETED"}
	if !contains(validStatuses, status) {
		return map[string]interface{}{
			"error":   true,
			"message": "Status must be one of the following: " + strings.Join(validStatuses, ", "),
		}
	} else if status == "DELETED" {
		return map[string]interface{}{
			"error":   true,
			"message": "Status cannot be DELETED",
		}
	}

	return map[string]interface{}{
		"error":   false,
		"message": "Status is valid",
	}
}

func durationUnit(durationUnit string) map[string]interface{} {
	validDurationUnits := []string{"HOURS", "DAYS", "WEEKS"}

	if durationUnit != "" && !contains(validDurationUnits, durationUnit) {
		return map[string]interface{}{
			"error":   true,
			"message": "Duration unit must be one of the following: " + strings.Join(validDurationUnits, ", "),
		}
	}

	return map[string]interface{}{
		"error":   false,
		"message": "Status is valid",
	}
}

func calculateDuration(startDate, endDate string, durationUnit string) float64 {
	layout := "2006-01-02T15:04:05Z"

	startDateTime, _ := time.Parse(layout, startDate)
	endDateTime, _ := time.Parse(layout, endDate)

	durationInSeconds := endDateTime.Sub(startDateTime).Seconds()

	switch durationUnit {
	case "HOURS":
		return round(durationInSeconds/3600, 2)
	case "WEEKS":
		return round(durationInSeconds/(7*24*3600), 2)
	default: // DAYS
		return round(durationInSeconds/(24*3600), 2)
	}
}

func round(value float64, decimalPlaces int) float64 {
	rounding := math.Pow(10, float64(decimalPlaces))
	return math.Round(value*rounding) / rounding
}
