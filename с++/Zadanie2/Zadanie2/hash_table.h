#ifndef HASH_TABLE_H
#define HASH_TABLE_H

#include "array.h"
#include "listTwo.h"

// Класс хеш-таблицы с методом цепочек для разрешения коллизий
class ChainingHashTable {
private:
    Array* table;  // Массив двусвязных списков (цепочек)
    int size;      // Размер хеш-таблицы

    // Хеш-функция для вычисления индекса
    int hash(int key) const;

    // Очистка списка (освобождение памяти)
    void clearList(ForwardListTwo* list);

public:
    ChainingHashTable(int tableSize);  // Конструктор
    ~ChainingHashTable();              // Деструктор

    // Основные операции
    void insert(int key);              // Вставка элемента
    bool search(int key) const;        // Поиск элемента
    void remove(int key);              // Удаление элемента

    // Статистика
    void getStatistics(int& minLength, int& maxLength, double& avgLength) const;
    void printDetailedStatistics() const;

    // Вспомогательные методы
    ForwardListTwo* getList(int index) const;  // Получение цепочки по индексу
};

#endif