package api

import "github.com/pkg/errors"

func ConvertPriorityToDomain(priority string) (Priority, error) {
	switch priority {
	case "low":
		return PriorityLow, nil
	case "medium":
		return PriorityMedium, nil
	case "high":
		return PriorityHigh, nil
	default:
		return "", errors.New("invalid priority")
	}
}

func ConvertStatusToDomain(status string) (Status, error) {
	switch status {
	case "ready for work":
		return StatusReadyForWork, nil
	case "in progress":
		return StatusInProgress, nil
	case "done":
		return StatusDone, nil
	default:
		return "", errors.New("invalid status")
	}
}
