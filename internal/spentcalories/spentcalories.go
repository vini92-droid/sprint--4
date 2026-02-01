package spentcalories

import (
	"fmt"
	"time"
	"errors"
	"log"
	"strings"
	"strconv"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)
//steps - количество шагов
//durations - продолжительность прогулки
//step - количество шагов
//distance - дистанция
//ckal - соженные калорий
//walking - вид активности ходьба
//stepLength - длина шага на от высоты пользователя и длины шага
//meanSpeed - средняя скорость
func parseTraining(data string) (int, string, time.Duration, error) {
	//"3456,Ходьба,3h00m"
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("неверный формат данных: %s", data)
	}
	steps := parts[0]// преобразование первого элемента слайс в int
	walkings := parts[1]
	durations := parts[2]
	step, err := strconv.Atoi(strings.TrimSpace(steps))
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка парсинга количество шагов.: %v", err)
	}
	if step <= 0 {
		return 0, "", 0, errors.New("количество шагов неположительная величина")
	}
	// извлечение второго элемента слайс в string
	walking := strings.TrimSpace(walkings)
	if walking == "" {
		return 0, "", 0, errors.New("не указан тип активности")
	}
	// преобразование третьего элемента слайс в string
	duration, err := time.ParseDuration(strings.TrimSpace(durations))
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка парсинга продолжительность активности: %v", err)
	}
	if duration <= 0 {
		return 0, "", 0, errors.New("время прогулки неположительное число")
	}
	return step, walking, duration, nil // TODO: реализовать функцию
}

func distance(steps int, height float64) float64 {
	//Функция принимает количество шагов и рост пользователя в метрах, 
	// а возвращает дистанцию в километрах.
	// для расчета длины шага
	stepLength := height * stepLengthCoefficient
	// для умножения пройденного количество шагов на длину шага и деления количества метров в километре
	return stepLength * float64(steps) / mInKm  // TODO: реализовать функцию
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
//Функция принимает количество шагов steps, рост пользователя height и продолжительность активности 
// duration  и возвращает среднюю скорость
// для проверки duration больше 0 
if duration <= 0 {
	return 0
} 
// для вычисления дистанции  
 distance := distance(steps, height)
 // Чтобы перевести продолжительность в часы, воспользуйтесь функцией из пакета
 hours := duration.Hours()
 //для вычисления средней скорости и ее возврата, разделив  дистанцию на продолжительность в часах. 
 return distance / hours // // TODO: реализовать функцию
}

func TrainingInfo(data string, weight, height float64) (string, error) {
//steps int — количество шагов.
//weight, height float64 — вес(кг.) и рост(м.) пользователя.
//duration time.Duration — продолжительность бега.
steps, walking, duration, err := parseTraining(data)
if err != nil {
	log.Println(err)
	return "", err
}
ckal := 0.0
switch walking{
case "Ходьба": 
	ckal, err = WalkingSpentCalories(steps, weight, height, duration)
case "Бег": 
	ckal, err = RunningSpentCalories(steps, weight, height, duration)
default:
	return "", errors.New("неизвестный тип тренировки")
}	
if err != nil {
	log.Println(err)
	return "", err
}// TODO: реализовать функцию
// Для каждого из видов тренировки рассчитываем дистанцию, среднюю скорость и выводим калории из функции
meanSpeed := meanSpeed(steps, height, duration)
distance := distance(steps, height)
return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", walking, float64(duration.Hours()), distance, meanSpeed, ckal), nil // TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
//steps int — количество шагов.
//weight, height float64 — вес(кг.) и рост(м.) пользователя.
//duration time.Duration — продолжительность бега.// TODO: реализовать функцию
//для проверки входные параметры на корректность 
if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
	return 0.0, errors.New("значение - неположительная величина")
}
//для рассчитать средней скорости  
meanSpeed := meanSpeed(steps, height, duration)
minutes := duration.Minutes()
	return (weight * meanSpeed * minutes) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
//steps int — количество шагов.
//weight, height float64 — вес(кг.) и рост(м.) пользователя.
//duration time.Duration — продолжительность ходьбы.// TODO: реализовать функцию
ckal, err := RunningSpentCalories(steps, weight, height, duration)
	return ckal * walkingCaloriesCoefficient, err
}
