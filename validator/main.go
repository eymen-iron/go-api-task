package validator

import (
	"math"
	"regexp"
	"strings"
	"time"
)

type ValidatorMessage struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func Color(color string) ValidatorMessage {
	match, _ := regexp.MatchString(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`, color)

	if !match {
		return ValidatorMessage{
			Error:   true,
			Message: "Color must be a valid hex color code. Example: #FF0000",
		}
	}

	return ValidatorMessage{
		Error:   false,
		Message: "Status is valid",
	}
}

func Name(name string) ValidatorMessage {
	if len(name) > 255 {
		return ValidatorMessage{
			Error:   true,
			Message: "Name must be less than 255 characters",
		}
	}

	return ValidatorMessage{
		Error:   false,
		Message: "Status is valid",
	}
}

func StartDate(startDate string) ValidatorMessage {
	_, err := time.Parse("2006-01-02T15:04:05Z", startDate)
	if err != nil {
		return ValidatorMessage{
			Error:   true,
			Message: "Start date must be a valid ISO8601 date. Example: 2022-12-31T14:59:00Z",
		}
	}

	return ValidatorMessage{
		Error:   false,
		Message: "Status is valid",
	}
}

func EndDate(endDate string, startDate string) ValidatorMessage {
	endDateTime, err := time.Parse("2006-01-02T15:04:05Z", endDate)
	if err != nil {
		return ValidatorMessage{
			Error:   true,
			Message: "End date must be a valid ISO8601 date. Example: 2022-12-31T14:59:00Z",
		}
	}

	startDateTime, _ := time.Parse("2006-01-02T15:04:05Z", startDate)

	if endDateTime.Before(startDateTime) {
		return ValidatorMessage{
			Error:   true,
			Message: "End date must be greater than start date",
		}
	}

	return ValidatorMessage{
		Error:   false,
		Message: "Status is valid",
	}
}
func ExternalId(externalID string) ValidatorMessage {
	if len(externalID) > 255 {
		return ValidatorMessage{
			Error:   true,
			Message: "External ID must be less than 255 characters",
		}
	}

	return ValidatorMessage{
		Error:   false,
		Message: "Status is valid",
	}
}

func Status(status string) ValidatorMessage {
	validStatuses := []string{"NEW", "PLANNED", "DELETED"}
	if !contains(validStatuses, status) {
		return ValidatorMessage{
			Error:   true,
			Message: "Status must be one of the following: " + strings.Join(validStatuses, ", "),
		}
	} else if status == "DELETED" {
		return ValidatorMessage{
			Error:   true,
			Message: "Status cannot be DELETED",
		}
	}

	return ValidatorMessage{
		Error:   false,
		Message: "Status is valid",
	}
}

func DurationUnit(durationUnit string) ValidatorMessage {
	validDurationUnits := []string{"HOURS", "DAYS", "WEEKS"}

	if durationUnit != "" && !contains(validDurationUnits, durationUnit) {
		return ValidatorMessage{
			Error:   true,
			Message: "Duration unit must be one of the following: " + strings.Join(validDurationUnits, ", "),
		}
	}

	return ValidatorMessage{
		Error:   false,
		Message: "Status is valid",
	}
}

func CalculateDuration(startDate, endDate string, durationUnit string) int {
	layout := "2006-01-02T15:04:05Z"

	startDateTime, _ := time.Parse(layout, startDate)
	endDateTime, _ := time.Parse(layout, endDate)

	durationInSeconds := endDateTime.Sub(startDateTime).Seconds()

	switch durationUnit {
	case "HOURS":
		return int(round(durationInSeconds/3600, 2))
	case "WEEKS":
		return int(round(durationInSeconds/(7*24*3600), 2))
	default: // DAYS
		return int(round(durationInSeconds/(24*3600), 2))
	}
}

func round(value float64, decimalPlaces int) float64 {
	rounding := math.Pow(10, float64(decimalPlaces))
	return math.Round(value*rounding) / rounding
}
