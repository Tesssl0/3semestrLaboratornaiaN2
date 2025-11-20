// taskscheduler.go
package main

import "strings"

func hasCycle(taskIndex int, tasks *Stack, dependencies *Stack,
	visited []bool, inStack []bool, currentPath *Stack) bool {
	if inStack[taskIndex] {
		return true // обнаружен цикл
	}
	if visited[taskIndex] {
		return false // уже посещена, циклов нет
	}

	visited[taskIndex] = true                // отмечаем как посещенную
	inStack[taskIndex] = true                // добавляем в текущий путь
	currentPath.Push(tasks.GetAt(taskIndex)) // добавляем в стек для отслеживания пути

	// Получаем имя текущей задачи
	currentTask := tasks.GetAt(taskIndex)

	// Проверяем все зависимости, где текущая задача является основной
	for i := 0; i <= dependencies.TopIndex(); i++ {
		dep := dependencies.GetAt(i)
		comma := strings.Index(dep, ",")
		if comma != -1 {
			dependentTask := dep[:comma] // задача, которая зависит
			mainTask := dep[comma+1:]    // задача, от которой зависят

			// Если текущая задача является основной в этой зависимости
			if mainTask == currentTask {
				// Находим индекс зависимой задачи
				dependentIndex := -1
				for j := 0; j <= tasks.TopIndex(); j++ {
					if tasks.GetAt(j) == dependentTask {
						dependentIndex = j
						break
					}
				}
				// Проверяем наличие циклов для зависимой задачи
				if dependentIndex != -1 && hasCycle(dependentIndex, tasks, dependencies,
					visited, inStack, currentPath) {
					return true
				}
			}
		}
	}

	currentPath.Pop()          // убираем задачу из текущего пути
	inStack[taskIndex] = false // снимаем отметку о нахождении в стеке
	return false
}

func CanFinish(tasks *Stack, dependencies *Stack) bool {
	taskCount := tasks.Size()

	if taskCount == 0 {
		return true
	}

	// Массивы для отслеживания посещенных вершин и вершин в текущем пути
	visited := make([]bool, taskCount)
	inStack := make([]bool, taskCount)

	currentPath := NewStack(taskCount, true) // стек для отслеживания текущего пути в DFS

	// Проверяем наличие циклов для всех непосещенных задач
	hasCycleFound := false
	for i := 0; i < taskCount; i++ {
		if !visited[i] {
			if hasCycle(i, tasks, dependencies, visited, inStack, currentPath) {
				hasCycleFound = true
				break
			}
		}
	}

	return !hasCycleFound
}
