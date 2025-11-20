#ifndef HASH_TABLE_H
#define HASH_TABLE_H

#include "array.h"
#include "listTwo.h"

// Класс хеш-таблицы с методом цепочек для разрешения коллизий
class ChainingHashTable {
private:
    Array* table;  // Массив для хранения указателей на списки
    int size;      // Размер таблицы (количество ячеек)

    int hash(int key) const;  // Хеш-функция
    void clearList(ForwardListTwo* list);  // Очистка списка

public:
    ChainingHashTable(int tableSize);  // Конструктор
    ~ChainingHashTable();              // Деструктор

    void insert(int key);      // Вставка элемента
    bool search(int key) const;  // Поиск элемента
    void remove(int key);      // Удаление элемента

    // Статистические функции
    void getStatistics(int& minLength, int& maxLength, double& avgLength) const;
    void printDetailedStatistics() const;

    ForwardListTwo* getList(int index) const;  // Получение списка по индексу
};

#endif