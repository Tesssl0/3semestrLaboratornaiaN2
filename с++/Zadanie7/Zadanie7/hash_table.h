#ifndef HASH_TABLE_H
#define HASH_TABLE_H

#include "array.h"
#include "listTwo.h"

// Класс хеш-таблицы с методом цепочек для разрешения коллизий
class ChainingHashTable {
private:
    Array* table;  // массив для хранения указателей на списки
    int size;      // размер хеш-таблицы

    int hash(int key) const;  // хеш-функция

    void clearList(ForwardListTwo* list);  // очистка списка

public:
    ChainingHashTable(int tableSize);  // конструктор
    ~ChainingHashTable();              // деструктор

    void insert(int key);          // вставка элемента
    bool search(int key) const;    // поиск элемента
    void remove(int key);          // удаление элемента

    // Получение статистики по таблице
    void getStatistics(int& minLength, int& maxLength, double& avgLength) const;
    void printDetailedStatistics() const;  // детальная статистика

    ForwardListTwo* getList(int index) const;  // получение списка по индексу
};

#endif