package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчётов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в одном часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчёта длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчёта калорий при ходьбе.
)

// Distance возвращает дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	stepLength := height * stepLengthCoefficient
	return (float64(steps) * stepLength) / mInKm
}

// MeanSpeed возвращает среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	return dist / duration.Hours()
}

// RunningSpentCalories возвращает количество калорий, потраченных при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные параметры для расчёта калорий при беге")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories, nil
}

// WalkingSpentCalories возвращает количество калорий, потраченных при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные параметры для расчёта калорий при ходьбе")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient
	return calories, nil
}
