#pragma once
#include "hash_table.h"
#include <string>

// Класс множества на основе хеш-таблицы
class Set {
private:
    ChainingHashTable table;  // Хеш-таблица для хранения элементов
    int tableSize;            // Размер таблицы

public:
    Set(int size = 100);  // Конструктор с размером по умолчанию

    bool add(int value);      // Добавление элемента (SETADD)
    bool remove(int value);   // Удаление элемента (SETDEL)
    bool contains(int value); // Проверка наличия элемента (SET_AT)

    void loadFromFile(const std::string& filename);  // Загрузка из файла
    void saveToFile(const std::string& filename);    // Сохранение в файл
};