#pragma once
#ifndef TASKSCHEDULER_H
#define TASKSCHEDULER_H

#include "stack.h"


// Функция проверки возможности выполнения всех задач с учетом зависимостей
bool canFinish(Stack& tasks, Stack& dependencies);
bool hasCycle(int taskIndex, Stack& tasks, Stack& dependencies, bool visited[], bool inStack[], Stack& currentPath);
#endif // TASKSCHEDULER_H