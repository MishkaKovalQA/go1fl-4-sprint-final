package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

var (
	ErrInvalidData     = errors.New("invalid data format")
	ErrZeroSteps       = errors.New("steps must be greater than zero")
	ErrZeroWeight      = errors.New("weight must be greater than zero")
	ErrZeroHeight      = errors.New("height must be greater than zero")
	ErrZeroDuration    = errors.New("duration must be greater than zero")
	ErrUnknownActivity = errors.New("unknown activity type")
)

func parseTraining(data string) (int, string, time.Duration, error) {
	slice := strings.Split(data, ",")
	if len(slice) != 3 {
		return 0, "", 0, ErrInvalidData
	}

	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, "", 0, ErrInvalidData
	}

	activity := slice[1]

	duration, err := time.ParseDuration(slice[2])
	if err != nil {
		return 0, "", 0, ErrInvalidData
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	return float64(steps) * height * stepLengthCoefficient / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	durationInHours := duration.Hours()
	if durationInHours == 0 {
		return 0
	}

	return distance(steps, height) / durationInHours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	switch activity {
	case "Бег":
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

		meanSpeed := meanSpeed(steps, height, duration)
		distance := distance(steps, height)

		return fmt.Sprintf("Тип тренировки: Бег\nДлительность: %v ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", duration, distance, meanSpeed, calories), nil
	case "Ходьба":
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

		meanSpeed := meanSpeed(steps, height, duration)
		distance := distance(steps, height)

		return fmt.Sprintf("Тип тренировки: Бег\nДлительность: %v ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", duration, distance, meanSpeed, calories), nil
	default:
		return "", ErrUnknownActivity
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, ErrZeroSteps
	}
	if weight <= 0 {
		return 0, ErrZeroWeight
	}
	if height <= 0 {
		return 0, ErrZeroHeight
	}
	if duration <= 0 {
		return 0, ErrZeroDuration
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, ErrZeroSteps
	}
	if weight <= 0 {
		return 0, ErrZeroWeight
	}
	if height <= 0 {
		return 0, ErrZeroHeight
	}
	if duration <= 0 {
		return 0, ErrZeroDuration
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories * walkingCaloriesCoefficient, nil
}
