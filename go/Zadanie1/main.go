// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Проверка выполнения задач с зависимостями ===")

	// Ввод задач
	var taskCount int
	fmt.Print("Введите количество задач: ")
	_, err := fmt.Scan(&taskCount)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	clearInputBuffer()

	// Проверка корректности количества задач
	if taskCount <= 0 {
		fmt.Println("Ошибка: количество задач должно быть положительным числом.")
		return
	}

	tasks := NewStack(taskCount, true)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите названия задач (каждое с новой строки):")
	for i := 0; i < taskCount; i++ {
		fmt.Printf("Задача %d: ", i+1)

		// Читаем ввод пользователя
		if !scanner.Scan() {
			fmt.Println("Ошибка при чтении ввода")
			return
		}

		task := strings.TrimSpace(scanner.Text())

		// Проверяем, что задача не пустая
		if task == "" {
			fmt.Println("Ошибка: название задачи не может быть пустым. Попробуйте снова.")
			i--
			continue
		}

		// Проверяем уникальность задачи
		if tasks.Find(task) != -1 {
			fmt.Printf("Ошибка: задача '%s' уже существует. Введите уникальное название.\n", task)
			i--
			continue
		}

		tasks.Push(task) // добавляем задачу в стек
	}

	// Ввод зависимостей
	var depCount int
	fmt.Print("\nВведите количество зависимостей: ")
	_, err = fmt.Scan(&depCount)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	clearInputBuffer()

	dependencies := NewStack(depCount, true)

	if depCount > 0 {
		fmt.Println("Введите зависимости в формате 'ЗАВИСИМАЯ_ЗАДАЧА,ОСНОВНАЯ_ЗАДАЧА':")
		fmt.Println("Пример: A,B  (задача A зависит от выполнения задачи B)")

		for i := 0; i < depCount; i++ {
			fmt.Printf("Зависимость %d: ", i+1)

			if !scanner.Scan() {
				fmt.Println("Ошибка при чтении ввода")
				return
			}

			dep := strings.TrimSpace(scanner.Text())

			// Проверяем формат ввода
			if !strings.Contains(dep, ",") {
				fmt.Println("Ошибка: используйте формат 'A,B'. Попробуйте снова.")
				i--
				continue
			}

			// Разбираем зависимость на части
			parts := strings.Split(dep, ",")
			if len(parts) != 2 {
				fmt.Println("Ошибка: используйте формат 'A,B'. Попробуйте снова.")
				i--
				continue
			}

			after := strings.TrimSpace(parts[0])  // задача, которая зависит
			before := strings.TrimSpace(parts[1]) // задача, от которой зависит

			// Проверяем существование задач
			if tasks.Find(after) == -1 {
				fmt.Printf("Ошибка: задача '%s' не найдена. Попробуйте снова.\n", after)
				i--
				continue
			}

			if tasks.Find(before) == -1 {
				fmt.Printf("Ошибка: задача '%s' не найдена. Попробуйте снова.\n", before)
				i--
				continue
			}

			// Проверяем самозависимость
			if after == before {
				fmt.Println("Ошибка: задача не может зависеть от самой себя. Попробуйте снова.")
				i--
				continue
			}

			// Проверяем дублирование зависимости
			existingDep := after + "," + before
			duplicate := false
			for j := 0; j <= dependencies.TopIndex(); j++ {
				if dependencies.GetAt(j) == existingDep {
					duplicate = true
					break
				}
			}

			if duplicate {
				fmt.Println("Ошибка: эта зависимость уже существует. Попробуйте снова.")
				i--
				continue
			}

			dependencies.Push(existingDep) // добавляем зависимость
		}
	}

	// Вывод введенных данных
	fmt.Println("\n=== Введенные данные ===")
	fmt.Print("Задачи: ")
	tasks.Print()
	fmt.Print("Зависимости: ")
	dependencies.Print()

	// Проверка возможности выполнения
	fmt.Println("\n=== Результат проверки ===")
	if CanFinish(tasks, dependencies) {
		fmt.Println("ВОЗМОЖНО выполнить все задачи")
		fmt.Println("Порядок выполнения: все задачи могут быть выполнены последовательно")
	} else {
		fmt.Println("НЕВОЗМОЖНО выполнить все задачи")
		fmt.Println("Обнаружен цикл зависимостей!")
	}

	fmt.Println("\nНажмите Enter для выхода...")
	scanner.Scan()
}

func clearInputBuffer() {
	// Читаем все оставшиеся символы до новой строки
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}
