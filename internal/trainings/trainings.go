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
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
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
	t.Steps = steps

	// Парсинг типа тренировки
	trainingType := strings.TrimSpace(parts[1])
	if trainingType == "" {
		return fmt.Errorf("тип тренировки не может быть пустым")
	}
	t.TrainingType = trainingType

	// Парсинг продолжительности
	durationStr := strings.TrimSpace(parts[2])
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
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("Тип тренировки: %s\n", t.TrainingType)
	info += fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours())
	info += fmt.Sprintf("Дистанция: %.2f км.\n", distance)
	info += fmt.Sprintf("Скорость: %.2f км/ч\n", speed)
	info += fmt.Sprintf("Сожгли калорий: %.2f\n", calories) // Добавляем перевод строки в конце

	return info, nil
}
