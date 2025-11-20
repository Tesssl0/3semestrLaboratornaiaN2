#include "Set.h"
#include <fstream>
#include <iostream>
#include <sstream>

using namespace std;

// Конструктор множества
Set::Set(int size) : table(size), tableSize(size) {}

// Добавление элемента в множество
bool Set::add(int value) {
    if (table.search(value))
        return false;  // Элемент уже существует
    table.insert(value);
    return true;  // Элемент успешно добавлен
}

// Удаление элемента из множества
bool Set::remove(int value) {
    if (!table.search(value))
        return false;  // Элемент не найден
    table.remove(value);
    return true;  // Элемент успешно удален
}

// Проверка наличия элемента в множестве
bool Set::contains(int value) {
    return table.search(value);
}

// Загрузка данных из файла
void Set::loadFromFile(const string& filename) {
    ifstream file(filename);
    if (!file.is_open()) {
        cerr << "Ошибка: не удалось открыть файл " << filename << endl;
        return;
    }

    string line;
    while (getline(file, line)) {
        if (!line.empty()) {
            try {
                int value = stoi(line);  // Преобразование строки в число
                table.insert(value);     // Добавление в таблицу
            }
            catch (...) {
                cerr << "Неверный формат данных в файле: " << line << endl;
            }
        }
    }
    file.close();
}

// Сохранение данных в файл
void Set::saveToFile(const string& filename) {
    ofstream file(filename);
    if (!file.is_open()) {
        cerr << "Ошибка: не удалось открыть файл для записи " << filename << endl;
        return;
    }

    // Обход всех цепочек хеш-таблицы
    for (int i = 0; i < tableSize; i++) {
        ForwardListTwo* list = table.getList(i);
        if (list && list->head) {
            DoublyNode* current = list->head;
            while (current) {
                file << current->node << endl;  // Запись элемента в файл
                current = current->next;
            }
        }
    }
    file.close();
}