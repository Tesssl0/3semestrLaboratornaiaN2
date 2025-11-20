#include "TaskScheduler.h"
#include <iostream>
#include <string>

using namespace std;

// ¬спомогательна€ функци€ дл€ обнаружени€ циклов в графе зависимостей
bool hasCycle(int taskIndex, Stack& tasks, Stack& dependencies,
    bool visited[], bool inStack[], Stack& currentPath) {
    if (inStack[taskIndex]) return true;  // обнаружен цикл
    if (visited[taskIndex]) return false; // уже посещена, циклов нет

    visited[taskIndex] = true; // отмечаем как посещенную
    inStack[taskIndex] = true; // добавл€ем в текущий путь
    currentPath.push(tasks.getAt(taskIndex));  // добавл€ем в стек дл€ отслеживани€ пути

    // ѕолучаем им€ текущей задачи
    string currentTask = tasks.getAt(taskIndex);

    // ѕровер€ем все зависимости, где текуща€ задача €вл€етс€ основной
    for (int i = 0; i <= dependencies.getTopIndex(); i++) {
        string dep = dependencies.getAt(i);
        size_t comma = dep.find(',');
        if (comma != string::npos) {
            string dependentTask = dep.substr(0, comma);  // задача, котора€ зависит
            string mainTask = dep.substr(comma + 1); // задача, от которой завис€т

            // ≈сли текуща€ задача €вл€етс€ основной в этой зависимости
            if (mainTask == currentTask) {
                // Ќаходим индекс зависимой задачи
                int dependentIndex = -1;
                for (int j = 0; j <= tasks.getTopIndex(); j++) {
                    if (tasks.getAt(j) == dependentTask) {
                        dependentIndex = j;
                        break;
                    }
                }
                // ѕровер€ем наличие циклов дл€ зависимой задачи
                if (dependentIndex != -1 && hasCycle(dependentIndex, tasks, dependencies,
                    visited, inStack, currentPath)) {
                    return true;
                }
            }
        }
    }

    currentPath.pop(); // убираем задачу из текущего пути
    inStack[taskIndex] = false;  // снимаем отметку о нахождении в стеке
    return false;
}

// ќсновна€ функци€ проверки возможности выполнени€ всех задач
bool canFinish(Stack& tasks, Stack& dependencies) {
    int taskCount = tasks.getTopIndex() + 1;

    if (taskCount == 0) return true;  

    // ћассивы дл€ отслеживани€ посещенных вершин и вершин в текущем пути
    bool* visited = new bool[taskCount];
    bool* inStack = new bool[taskCount];
    for (int i = 0; i < taskCount; i++) {
        visited[i] = false;
        inStack[i] = false;
    }

    Stack currentPath(taskCount, true);  // стек дл€ отслеживани€ текущего пути в DFS

    // ѕровер€ем наличие циклов дл€ всех непосещенных задач
    bool hasCycleFound = false;
    for (int i = 0; i < taskCount; i++) {
        if (!visited[i]) {
            if (hasCycle(i, tasks, dependencies, visited, inStack, currentPath)) {
                hasCycleFound = true;
                break;
            }
        }
    }

    delete[] visited;
    delete[] inStack;

    return !hasCycleFound;  
}