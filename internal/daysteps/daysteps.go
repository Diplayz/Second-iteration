package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	var stepsStr, durationStr string

	if strings.Contains(datastring, ",") {
		splitData := strings.Split(datastring, ",")
		if len(splitData) != 2 {
			return fmt.Errorf("invalid format: expected 2 fields, got %d", len(splitData))
		}
		stepsStr, durationStr = splitData[0], splitData[1]

	} else if strings.Contains(datastring, " ") {
		splitData := strings.Split(datastring, " ")
		if len(splitData) != 2 {
			return fmt.Errorf("invalid format: expected 2 fields, got %d", len(splitData))
		}
		stepsStr, durationStr = splitData[0], splitData[1]
	} else {
		return fmt.Errorf("incorrect data")
	}
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("incorrect transformation of steps: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("the number of steps must be positive")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("incorrect transformation of duration: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("duration must be positive")
	}
	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 || ds.Weight <= 0 || ds.Height <= 0 || ds.Duration <= 0 {
		return "", fmt.Errorf("incorrect data: steps=%d, weight=%.1f, height=%.1f, duration=%v",
			ds.Steps, ds.Weight, ds.Height, ds.Duration)
	}

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	spentRunCalories, err := spentenergy.RunningSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("error calculating calories: %w", err)
	}

	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, spentRunCalories)

	return info, nil
}
