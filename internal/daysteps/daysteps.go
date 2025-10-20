package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("неверный формат данных")
	}

	// Парсинг шагов с проверкой на пробелы
	stepsStr := parts[0]

	// Проверяем, есть ли пробелы в начале или конце
	if strings.TrimSpace(stepsStr) != stepsStr {
		return fmt.Errorf("неверный формат шагов")
	}

	if stepsStr == "" {
		return fmt.Errorf("неверный формат шагов")
	}

	// Проверяем, что строка состоит только из цифр (и возможного знака)
	for _, char := range stepsStr {
		if !unicode.IsDigit(char) && char != '+' && char != '-' {
			return fmt.Errorf("неверный формат шагов")
		}
	}

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("ошибка парсинга шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	ds.Steps = steps

	// Парсинг продолжительности с проверкой на пробелы
	durationStr := parts[1]

	// Проверяем, есть ли пробелы в начале или конце
	if strings.TrimSpace(durationStr) != durationStr {
		return fmt.Errorf("неверный формат продолжительности")
	}

	if durationStr == "" {
		return fmt.Errorf("продолжительность не может быть пустой")
	}
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("ошибка парсинга продолжительности: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)
	info += fmt.Sprintf("Дистанция составила %.2f км.\n", distance)
	info += fmt.Sprintf("Вы сожгли %.2f ккал.\n", calories)

	return info, nil
}
