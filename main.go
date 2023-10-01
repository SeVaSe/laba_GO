package main

import (
	"fmt"
	"math"
)

func main() {
	// ПУНКТ1
	fmt.Println("Задание 8\nПункт: 1.\nПротабулировать функцию, на интервале [-10;20] с шагом 2\nФункция: Y(\nLg(x+7), если x<-4\nx*Sqrt(x+2) + 1/ (x*2-2x+1),  если  -4<=x<=1\nx/tg(2-x), если 1<x<=10\n5x-2tg^2 * x / (x-17), в остальных случаях\n)\n")

	xStartPoint := -10.0
	xEndPoint := 20.0
	step := 2.0

	// Проходим по интервалу от xStartPoint до xEndPoint с шагом step
	for x := xStartPoint; x <= xEndPoint; x += step {
		var y float64

		if x < -4 {
			// Если x меньше -4, вычисляем y как логарифм от x + 7
			if x+7 <= 0 {
				// Проверка на отрицательный аргумент логарифма
				fmt.Printf("x = %.2f, y = NaN (логарифм от отрицательного числа)\n", x)
				continue
			}
			y = math.Log(x + 7)
		} else if -4 <= x && x <= 1 {
			// Если x находится между -4 и 1, вычисляем y по формуле
			if x == 0 || 2*x-2*x+1 == 0 {
				// Проверка на деление на ноль
				fmt.Printf("x = %.2f, y = NaN (деление на ноль)\n", x)
				continue
			}
			y = x*math.Sqrt(x+2) + 1/(2*x-2*x+1)
		} else if 1 < x && x <= 10 {
			// Если x больше 1 и меньше либо равен 10, вычисляем y по формуле
			if math.Abs(2-x) < 1e-10 {
				// Проверка на деление на ноль
				fmt.Printf("x = %.2f, y = NaN (деление на ноль)\n", x)
				continue
			}
			y = x / math.Tan(2-x)
		} else {
			// Если x больше 10, вычисляем y по формуле
			if x-17 == 0 {
				// Проверка на деление на ноль
				fmt.Printf("x = %.2f, y = NaN (деление на ноль)\n", x)
				continue
			}
			y = 5*x - 2*math.Pow(math.Tan(x), 2)/(x-17)
		}

		// Выводим x и соответствующее значение y
		fmt.Printf("x = %.2f, y = %.2f\n", x, y)
	}

	fmt.Println("\n\n")

	// ПУНКТ2
	// Выводим описание задания
	fmt.Println("Задание 8\nПункт: 2. \nДана последовательность a1,a2,...,a7. Найти наименьший  член последовательности и его номер и напечатать.  Вывести количество равных ему элементов и их номера. Вывести сообщение, если таких чисел нет\n")

	// Инициализируем переменные для хранения чисел
	var a1, a2, a3, a4, a5, a6, a7 int

	// Запрашиваем у пользователя ввод семи чисел
	fmt.Println("Введите семь чисел:")
	fmt.Print("a1: ")
	fmt.Scan(&a1)
	fmt.Print("a2: ")
	fmt.Scan(&a2)
	fmt.Print("a3: ")
	fmt.Scan(&a3)
	fmt.Print("a4: ")
	fmt.Scan(&a4)
	fmt.Print("a5: ")
	fmt.Scan(&a5)
	fmt.Print("a6: ")
	fmt.Scan(&a6)
	fmt.Print("a7: ")
	fmt.Scan(&a7)

	// Инициализируем текущий минимум и его номер
	min := a1
	minIndex := 1

	// Находим минимум и его номер в последовательности
	if a2 < min {
		min = a2
		minIndex = 2
	}
	if a3 < min {
		min = a3
		minIndex = 3
	}
	if a4 < min {
		min = a4
		minIndex = 4
	}
	if a5 < min {
		min = a5
		minIndex = 5
	}
	if a6 < min {
		min = a6
		minIndex = 6
	}
	if a7 < min {
		min = a7
		minIndex = 7
	}

	// Выводим наименьший член последовательности и его номер
	fmt.Printf("Первый наименьший член последовательности: %d\n", min)
	fmt.Printf("Его номер: %d\n\n", minIndex)

	// Подсчитываем количество равных ему элементов и выводим их номера
	count := 0
	fmt.Printf("Номера равных элементов: ")
	if a1 == min {
		count++
		fmt.Printf("1 ")
	}
	if a2 == min {
		count++
		fmt.Printf("2 ")
	}
	if a3 == min {
		count++
		fmt.Printf("3 ")
	}
	if a4 == min {
		count++
		fmt.Printf("4 ")
	}
	if a5 == min {
		count++
		fmt.Printf("5 ")
	}
	if a6 == min {
		count++
		fmt.Printf("6 ")
	}
	if a7 == min {
		count++
		fmt.Printf("7 ")
	}

	fmt.Println()

	// Выводим количество равных элементов или сообщение, если их нет
	if count > 1 {
		fmt.Printf("Количество равных ему элементов: %d\n", count)
	} else {
		fmt.Println("Таких чисел нет")
	}
	fmt.Println("\n\n")

	// ПУНКТ3
	fmt.Println("Задание 8 \nПункт: 3. \nВычислить: Y=Σ (-1)^к+1*Sin^2*(к+1)/(2*к)!\nn=12\nк=6\n")

	n := 12 // Верхний предел суммы
	k := 6  // Нижний предел суммы
	Y := 0.0

	// Проверяем, что k не больше чем n
	if k > n {
		fmt.Println("Ошибка: k больше чем n")
		return
	}

	fact := 1
	// Вычисляем факториал для использования в формуле
	for i := 1; i <= 2*n; i++ {
		fact *= i
	}

	// Вычисляем сумму согласно заданной формуле
	for i := k; i <= n; i++ {
		Y += math.Pow(-1, float64(i+1)) * math.Pow(math.Sin(float64(i+1)), 2) / float64(fact)
		fact /= (2 * i) * (2*i + 1)
	}

	// Выводим результат с точностью до 6 знаков после запятой
	fmt.Printf("Y = %.6f\n", Y)
	fmt.Println("Нажмите 'q' для завершения программы.")
	var aa string
	fmt.Scan(&aa)
	fmt.Println(aa)
}
