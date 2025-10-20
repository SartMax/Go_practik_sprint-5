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
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("неверный формат данных")
	}

	// Парсинг шагов
	stepsStr := strings.TrimSpace(parts[0])
	if stepsStr == "" {
		return fmt.Errorf("неверный формат шагов")
	}
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("ошибка парсинга шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	ds.Steps = steps

	// Парсинг продолжительности
	durationStr := strings.TrimSpace(parts[1])
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
	info += fmt.Sprintf("Вы сожгли %.2f ккал.\n", calories) // Добавляем перевод строки в конце

	return info, nil
}
