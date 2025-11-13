package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect data")
	}
	spentCalories := (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH
	return spentCalories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect data")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	spentCalories := (weight * meanSpeed * duration.Minutes()) / minInH
	return spentCalories, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 || height <= 0 {
		return 0
	}
	distanceKm := Distance(steps, height)
	meanSpeed := distanceKm / duration.Hours()

	return meanSpeed
}
func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	stepLength := height * stepLengthCoefficient
	totalDistance := stepLength * float64(steps)
	distanceInKm := totalDistance / mInKm

	return distanceInKm
}
