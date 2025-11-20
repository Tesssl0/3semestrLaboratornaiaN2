#include <iostream>
#include <string>
#include "TaskScheduler.h"

using namespace std;

// Функция очистки буфера ввода
void clearInputBuffer() {
    cin.clear();
    cin.ignore(10000, '\n');
}

int main() {
    setlocale(LC_ALL, "Ru");

    // Используем стеки вместо массивов для хранения задач и зависимостей
    Stack tasks(50, true);
    Stack dependencies(50, true);

    cout << "=== Проверка выполнения задач с зависимостями ===" << endl;

    // Ввод задач
    int taskCount;
    cout << "Введите количество задач: ";
    cin >> taskCount;
    clearInputBuffer();

    // Проверка корректности количества задач
    if (taskCount <= 0) {
        cout << "Ошибка: количество задач должно быть положительным числом." << endl;
        return 1;
    }

    cout << "Введите названия задач (каждое с новой строки):" << endl;
    for (int i = 0; i < taskCount; i++) {
        string task;
        cout << "Задача " << (i + 1) << ": ";
        getline(cin, task);

        // Проверяем, что задача не пустая
        if (task.empty()) {
            cout << "Ошибка: название задачи не может быть пустым. Попробуйте снова." << endl;
            i--;
            continue;
        }

        // Проверяем уникальность задачи
        if (findInStack(tasks, task) != -1) {
            cout << "Ошибка: задача '" << task << "' уже существует. Введите уникальное название." << endl;
            i--;
            continue;
        }

        tasks.push(task);  // добавляем задачу в стек
    }

    // Ввод зависимостей
    int depCount;
    cout << "\nВведите количество зависимостей: ";
    cin >> depCount;
    clearInputBuffer();

    if (depCount > 0) {
        cout << "Введите зависимости в формате 'ЗАВИСИМАЯ_ЗАДАЧА,ОСНОВНАЯ_ЗАДАЧА':" << endl;
        cout << "Пример: A,B  (задача A зависит от выполнения задачи B)" << endl;

        for (int i = 0; i < depCount; i++) {
            string dep;
            cout << "Зависимость " << (i + 1) << ": ";
            getline(cin, dep);

            // Проверяем формат ввода
            if (dep.find(',') == string::npos) {
                cout << "Ошибка: используйте формат 'A,B'. Попробуйте снова." << endl;
                i--;
                continue;
            }

            // Разбираем зависимость на части
            size_t comma = dep.find(',');
            string after = dep.substr(0, comma);   // задача, которая зависит
            string before = dep.substr(comma + 1); // задача, от которой зависит

            // Убираем возможные пробелы по краям
            after.erase(0, after.find_first_not_of(" \t"));
            after.erase(after.find_last_not_of(" \t") + 1);
            before.erase(0, before.find_first_not_of(" \t"));
            before.erase(before.find_last_not_of(" \t") + 1);

            // Проверяем существование задач
            if (findInStack(tasks, after) == -1) {
                cout << "Ошибка: задача '" << after << "' не найдена. Попробуйте снова." << endl;
                i--;
                continue;
            }

            if (findInStack(tasks, before) == -1) {
                cout << "Ошибка: задача '" << before << "' не найдена. Попробуйте снова." << endl;
                i--;
                continue;
            }

            // Проверяем самозависимость
            if (after == before) {
                cout << "Ошибка: задача не может зависеть от самой себя. Попробуйте снова." << endl;
                i--;
                continue;
            }

            // Проверяем дублирование зависимости
            string existingDep = after + "," + before;
            bool duplicate = false;
            for (int j = 0; j <= dependencies.getTopIndex(); j++) {
                if (dependencies.getAt(j) == existingDep) {
                    duplicate = true;
                    break;
                }
            }

            if (duplicate) {
                cout << "Ошибка: эта зависимость уже существует. Попробуйте снова." << endl;
                i--;
                continue;
            }

            dependencies.push(existingDep);  // добавляем зависимость
        }
    }

    // Вывод введенных данных
    cout << "\n=== Введенные данные ===" << endl;
    cout << "Задачи: ";
    printStack(tasks);
    cout << "Зависимости: ";
    printStack(dependencies);

    // Проверка возможности выполнения
    cout << "\n=== Результат проверки ===" << endl;
    if (canFinish(tasks, dependencies)) {
        cout << "ВОЗМОЖНО выполнить все задачи" << endl;
        cout << "Порядок выполнения: все задачи могут быть выполнены последовательно" << endl;
    }
    else {
        cout << " НЕВОЗМОЖНО выполнить все задачи" << endl;
        cout << "Обнаружен цикл зависимостей!" << endl;
    }

    cout << "\nНажмите Enter для выхода...";
    clearInputBuffer();
    getchar();

    return 0;
}