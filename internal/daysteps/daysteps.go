package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go-first-floor-sprint-five/internal/personaldata"
	"github.com/Yandex-Practicum/go-first-floor-sprint-five/internal/spentenergy"
)

// DaySteps хранит данные о дневной активности (прогулках) и пользователе.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse парсит строку формата "678,0h50m" и заполняет поля структуры DaySteps.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат строки данных о прогулке")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка преобразования количества шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("ошибка преобразования длительности прогулки: %w", err)
	}
	if duration <= 0 {
		return errors.New("длительность прогулки должна быть больше нуля")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

// ActionInfo формирует и возвращает строку с информацией о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		dist,
		calories,
	)

	return info, nil
}
