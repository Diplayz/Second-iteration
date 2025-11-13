package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	splitData := strings.Split(datastring, ",")
	if len(splitData) != 3 {
		return fmt.Errorf("invalid format: expected 3 fields, got %d", len(splitData))
	}
	steps, err := strconv.Atoi(splitData[0])
	if err != nil {
		return fmt.Errorf("incorrect transformation of steps: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("the number of steps must be positive")
	}
	t.Steps = steps

	t.TrainingType = splitData[1]

	duration, err := time.ParseDuration(splitData[2])
	if err != nil {
		return fmt.Errorf("incorrect transformation of duration: %w", err)
	}

	if duration <= 0 {
		return fmt.Errorf("duration must be positive")
	}
	t.Duration = duration
	return nil
}
func (t Training) ActionInfo() (string, error) {
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	distance := spentenergy.Distance(t.Steps, t.Height)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)

	return info, nil
}
