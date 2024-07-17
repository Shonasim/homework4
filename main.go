package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 1. PrintGreeting
func PrintGreeting() {
	var dayType string
	fmt.Print("Введите время дня (утро, день, вечер): ")
	fmt.Scan(&dayType)
	switch dayType {
	case "утро":
		fmt.Println("Доброе утро!")
	case "день":
		fmt.Println("Добрый день!")
	case "вечер":
		fmt.Println("Добрый вечер!")
	default:
		fmt.Println("Неправильное время дня")
	}
}

// 2. PrintWeather
func PrintWeather() {
	var weatherType string
	fmt.Print("Введите тип погоды (солнечно, облачно, дождливо): ")
	fmt.Scan(&weatherType)
	switch weatherType {
	case "солнечно":
		fmt.Println("Солнечно")
	case "облачно":
		fmt.Println("Облачно")
	case "дождливо":
		fmt.Println("Дождливо")
	default:
		fmt.Println("Неправильный тип погоды")
	}
}

// 3. PrintTrafficLight
func PrintTrafficLight() {
	var lightColor string
	fmt.Print("Введите цвет светофора (красный, желтый, зеленый): ")
	fmt.Scan(&lightColor)
	switch lightColor {
	case "красный":
		fmt.Println("Стоп")
	case "желтый":
		fmt.Println("Внимание")
	case "зеленый":
		fmt.Println("Идите")
	default:
		fmt.Println("Неправильный цвет светофора")
	}
}

// 4. GetGrade
func GetGrade() string {
	var score int
	fmt.Print("Введите оценку (0-100): ")
	fmt.Scan(&score)
	switch {
	case score >= 90:
		return "A"
	case score >= 75:
		return "B"
	default:
		return "C"
	}
}

// 5. GetDiscount
func GetDiscount() string {
	var amount int
	fmt.Print("Введите сумму покупки: ")
	fmt.Scan(&amount)
	switch {
	case amount > 1000:
		return "10%"
	case amount > 500:
		return "5%"
	default:
		return "0%"
	}
}

// 6. GetTemperatureDescription
func GetTemperatureDescription() string {
	var temperature int
	fmt.Print("Введите температуру: ")
	fmt.Scan(&temperature)
	switch {
	case temperature < 10:
		return "Холодно"
	case temperature <= 25:
		return "Тепло"
	default:
		return "Жарко"
	}
}

// 7. CheckNumber
func CheckNumber(n int) {
	if n > 0 {
		fmt.Println("Положительное")
	} else if n < 0 {
		fmt.Println("Отрицательное")
	} else {
		fmt.Println("Ноль")
	}
}

// 8. CheckAge
func CheckAge(age int) {
	if age >= 18 {
		fmt.Println("Совершеннолетний")
	} else {
		fmt.Println("Несовершеннолетний")
	}
}

// 9. CheckPassword
func CheckPassword(password string) {
	if password == "1234" {
		fmt.Println("Пароль верный")
	} else {
		fmt.Println("Пароль неверный")
	}
}

// 10. Add
func Add(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a + b
}

// 11. CompareStrings
func CompareStrings(s1, s2 string) string {
	if s1 == s2 {
		return "равны"
	}
	return "не равны"
}

// 12. Max
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 13. Operation
var operation = func(a, b int) int {
	return a - b
}

// 14. Concat
var concat = func(s1, s2 string) string {
	if s1 == "" || s2 == "" {
		return s1 + s2
	}
	return s1 + " " + s2
}

// 15. Multiply
var multiply = func(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a * b
}

// 16. ApplyOperation
func ApplyOperation(a, b int, op func(int, int) int) {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	fmt.Println(op(a, b))
}

// 17. CheckCondition
func CheckCondition(n int, condition func(int) bool) {
	if condition(n) {
		fmt.Println("Условие выполнено")
	} else {
		fmt.Println("Условие не выполнено")
	}
}

// 18. FormatAndPrint
func FormatAndPrint(s string, formatter func(string) string) {
	if s == "" {
		fmt.Println("Пустая строка")
	} else {
		fmt.Println(formatter(s))
	}
}

// 19. CreateMultiplier
func CreateMultiplier(n int) func(int) int {
	return func(x int) int {
		if x < 0 {
			x = -x
		}
		return n * x
	}
}

// 20. CreateGreeter
func CreateGreeter(greeting string) func(string) string {
	return func(name string) string {
		if name == "" {
			return greeting + " гость"
		}
		return greeting + " " + name
	}
}

// 21. CreateValidator
func CreateValidator(password string) func(string) bool {
	return func(input string) bool {
		return input == password
	}
}

// 22. CreateFormatter
func CreateFormatter(separator string) func(string, string) string {
	return func(s1, s2 string) string {
		if s1 == "" || s2 == "" {
			return s1 + s2
		}
		return s1 + separator + s2
	}
}

// 23. CreateFilter
func CreateFilter(min, max int) func(int) bool {
	return func(n int) bool {
		return n >= min && n <= max
	}
}

// 24. CreateConverter
func CreateConverter(base int) func(int) string {
	return func(n int) string {
		return strconv.FormatInt(int64(n), base)
	}
}

// 25. CreateCounter
func CreateCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

func main() {
	PrintGreeting()
	PrintWeather()
	PrintTrafficLight()
	fmt.Println("Оценка:", GetGrade())
	fmt.Println("Скидка:", GetDiscount())
	fmt.Println("Температура:", GetTemperatureDescription())
	CheckNumber(5)
	CheckNumber(-3)
	CheckNumber(0)
	CheckAge(20)
	CheckAge(16)
	CheckPassword("1234")
	CheckPassword("4321")
	fmt.Println("Сумма модулей:", Add(-5, 3))
	fmt.Println("Сравнение строк:", CompareStrings("hello", "hello"))
	fmt.Println("Сравнение строк:", CompareStrings("hello", "world"))
	fmt.Println("Максимум:", Max(5, 10))
	fmt.Println("Разность:", operation(10, 5))
	fmt.Println("Конкатенация:", concat("hello", "world"))
	fmt.Println("Произведение модулей:", multiply(-4, 3))
	ApplyOperation(-4, 3, multiply)
	CheckCondition(5, func(n int) bool { return n > 0 })
	FormatAndPrint("hello", strings.ToUpper)
	multiplier := CreateMultiplier(2)
	fmt.Println("Умножение:", multiplier(-3))
	greeter := CreateGreeter("Привет")
	fmt.Println(greeter("Мир"))
	validator := CreateValidator("1234")
	fmt.Println("Пароль верный:", validator("1234"))
	formatter := CreateFormatter(", ")
	fmt.Println(formatter("Привет", "Мир"))
	filter := CreateFilter(10, 20)
	fmt.Println("Фильтр:", filter(15))
	converter := CreateConverter(16)
	fmt.Println("Конвертер:", converter(255))
	counter := CreateCounter(0)
	fmt.Println("Счетчик:", counter())
	fmt.Println("Счетчик:", counter())
}
