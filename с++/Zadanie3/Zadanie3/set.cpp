#include "Set.h"
#include <fstream>
#include <iostream>
#include <sstream>

using namespace std;

// Конструктор множества
Set::Set(int size) : table(size), tableSize(size) {}

// Добавление элемента в множество
bool Set::add(int value) {
    if (table.search(value))  // Если элемент уже существует
        return false;         // Возвращаем false (не добавлен)
    table.insert(value);      // Вставляем элемент в таблицу
    return true;              // Возвращаем true (успешно добавлен)
}

// Удаление элемента из множества
bool Set::remove(int value) {
    if (!table.search(value))  // Если элемент не существует
        return false;          // Возвращаем false (не удален)
    table.remove(value);       // Удаляем элемент из таблицы
    return true;               // Возвращаем true (успешно удален)
}

// Проверка наличия элемента в множестве
bool Set::contains(int value) {
    return table.search(value);  // Ищем элемент в хеш-таблице
}

// Загрузка данных из файла в множество
void Set::loadFromFile(const string& filename) {
    ifstream file(filename);  // Открываем файл для чтения
    if (!file.is_open()) {
        cerr << "Ошибка: не удалось открыть файл " << filename << endl;
        return;
    }

    string line;
    // Читаем файл построчно
    while (getline(file, line)) {
        if (!line.empty()) {  // Пропускаем пустые строки
            try {
                int value = stoi(line);  // Преобразуем строку в число
                table.insert(value);     // Вставляем в таблицу
            }
            catch (...) {  // Обрабатываем ошибки преобразования
                cerr << "Неверный формат данных в файле: " << line << endl;
            }
        }
    }
    file.close();  // Закрываем файл
}

// Сохранение данных множества в файл
void Set::saveToFile(const string& filename) {
    ofstream file(filename);  // Открываем файл для записи
    if (!file.is_open()) {
        cerr << "Ошибка: не удалось открыть файл для записи " << filename << endl;
        return;
    }

    // Проходим по всем ячейкам хеш-таблицы
    for (int i = 0; i < tableSize; i++) {
        ForwardListTwo* list = table.getList(i);  // Получаем список из ячейки
        if (list && list->head) {  // Если список существует и не пустой
            DoublyNode* current = list->head;  // Начинаем с головы списка
            while (current) {  // Проходим по всем узлам списка
                file << current->node << endl;  // Записываем значение в файл
                current = current->next;  // Переходим к следующему узлу
            }
        }
    }
    file.close();  // Закрываем файл
}