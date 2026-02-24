package daysteps

import (
	"errors"
	"time"
	"strings"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

var (
	ErrInvalidData = errors.New("invalid data format")
)

func parsePackage(data string) (int, time.Duration, error) {
	slice := strings.Split(data, ",")
	if len(slice) != 2 {
		return 0, 0, ErrInvalidData
	}
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
