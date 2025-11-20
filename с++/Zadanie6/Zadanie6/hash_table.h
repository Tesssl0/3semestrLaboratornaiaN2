#ifndef HASH_TABLE_H
#define HASH_TABLE_H

#include "array.h"
#include "listTwo.h"

// Перечисление методов разрешения коллизий
enum class CollisionResolution {
    CHAINING,       // Метод цепочек
    LINEAR_PROBING, // Линейное пробирование
    QUADRATIC_PROBING // Квадратичное пробирование
};

class HashTable {
private:
    Array* table;           // Основная таблица (массив)
    int size;               // Размер таблицы
    CollisionResolution method; // Метод разрешения коллизий
    int prime;              // Простое число для двойного хеширования

    // Хеш-функции
    int hash1(int key) const;  // Основная хеш-функция
    int hash2(int key) const;  // Вторая хеш-функция для двойного хеширования

    // Методы для цепочек
    ForwardListTwo* getList(int index) const; // Получить список по индексу
    void clearList(ForwardListTwo* list);     // Очистить список

    // Методы для открытой адресации
    bool isDeleted(int index) const;          // Проверить, удалена ли ячейка
    bool isEmpty(int index) const;            // Проверить, пуста ли ячейка
    int findSlotOpenAddressing(int key) const; // Найти слот для открытой адресации

public:
    // Конструктор и деструктор
    HashTable(int tableSize, CollisionResolution resolutionMethod);
    ~HashTable();

    // Основные операции хеш-таблицы
    void insert(int key);   // Вставка элемента
    bool search(int key) const; // Поиск элемента
    void remove(int key);   // Удаление элемента

    // Методы для получения статистики
    void getStatistics(int& minLength, int& maxLength, double& avgLength) const;
    void printDetailedStatistics() const;

    // Вспомогательные методы
    void setResolutionMethod(CollisionResolution method); // Установить метод разрешения коллизий
    int getNextPrime(int n) const; // Получить следующее простое число
};

#endif