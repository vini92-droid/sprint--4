package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"log"

    "github.com/Yandex-Practicum/tracker/internal/spentcalories"
	
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	//преобразовывет строку в слайс
	parts := strings.Split(data, ",")// Парсинг в int
	// проверяет слайс на наличие двух значений
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("неверный формат данных: %s", data)
	}
	steps := parts[0] //количество шагов
	durations := parts[1] //продолжительность прогулки
	
	step, err := strconv.Atoi(steps)
	// для проверки на ошибку при преобразовании
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка int: %v", err)
	}
	// для проверки количества шагов на ноль
	if step <= 0 {
    return 0, 0, fmt.Errorf("ошибка: число шагов равно %d", step)
	}

	// Парсинг Duration 
	duration, err := time.ParseDuration(durations)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка длительности: %v", err)
	}
	if duration <= 0 {
    return 0, 0, fmt.Errorf("ошибка: число шагов равно %d", duration)
	}

	return step, duration, nil
}// TODO: реализовать функцию


func DayActionInfo(data string, weight, height float64) string {
	step, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	ckal := 0.0
	distance := float64(step) * stepLength / mInKm
	ckal, err = spentcalories.WalkingSpentCalories(step, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", step, distance, ckal) // TODO: реализовать функцию
}
