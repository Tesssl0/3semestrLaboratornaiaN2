#include <iostream>
#include <string>
#include <algorithm>
#include <climits>
#include "hash_table.h"

using namespace std;

// Хеш-функция: вычисляет индекс по ключу
int ChainingHashTable::hash(int key) const {
    return abs(key) % size;  // Используем модуль для равномерного распределения
}

// Получение цепочки по индексу
ForwardListTwo* ChainingHashTable::getList(int index) const {
    if (index < 0 || index >= size) return nullptr;  // Проверка границ

    string ptrStr = get(*table, index);
    if (ptrStr.empty()) return nullptr;

    // Преобразование строки обратно в указатель
    return reinterpret_cast<ForwardListTwo*>(stoll(ptrStr));
}

// Очистка списка (освобождение памяти всех узлов)
void ChainingHashTable::clearList(ForwardListTwo* list) {
    if (!list) return;

    // Последовательное удаление всех узлов
    while (list->head != nullptr) {
        deleteNodeHeadTwo(list);
    }
    delete list;  // Освобождение самой структуры списка
}

// Конструктор хеш-таблицы
ChainingHashTable::ChainingHashTable(int tableSize) : size(tableSize) {
    table = new Array;
    init(*table, size);  // Инициализация массива

    // Создание пустых списков для каждой ячейки
    for (int i = 0; i < size; i++) {
        ForwardListTwo* newList = new ForwardListTwo;
        newList->head = nullptr;
        newList->tail = nullptr;
        add(*table, "");

        // Сохранение указателя на список как строку в массиве
        set(*table, i, to_string(reinterpret_cast<long long>(newList)));
    }
}

// Деструктор хеш-таблицы
ChainingHashTable::~ChainingHashTable() {
    // Очистка всех цепочек
    for (int i = 0; i < size; i++) {
        ForwardListTwo* list = getList(i);
        clearList(list);
    }

    // Освобождение памяти массива
    destroy(*table);
    delete table;
}

// Вставка элемента в хеш-таблицу
void ChainingHashTable::insert(int key) {
    int index = hash(key);  // Вычисление индекса
    ForwardListTwo* list = getList(index);
    if (list) {
        addNodeTailTwo(list, to_string(key));  // Добавление в конец цепочки
    }
}

// Поиск элемента в хеш-таблице
bool ChainingHashTable::search(int key) const {
    int index = hash(key);
    ForwardListTwo* list = getList(index);
    if (!list) return false;

    // Поиск в цепочке
    return findNodeIndexTwo(list, to_string(key));
}

// Получение статистики по хеш-таблице
void ChainingHashTable::getStatistics(int& minLength, int& maxLength, double& avgLength) const {
    minLength = INT_MAX;
    maxLength = 0;
    int totalElements = 0;
    int nonEmptyBuckets = 0;

    // Анализ всех цепочек
    for (int i = 0; i < size; i++) {
        ForwardListTwo* list = getList(i);
        if (list) {
            int listSize = countNodesTwo(*list);
            if (listSize > 0) {
                // Обновление минимальной и максимальной длины
                if (listSize < minLength) minLength = listSize;
                if (listSize > maxLength) maxLength = listSize;
                totalElements += listSize;
                nonEmptyBuckets++;
            }
        }
    }

    // Корректировка значений для пустой таблицы
    if (minLength == INT_MAX) minLength = 0;

    // Вычисление средней длины непустых цепочек
    avgLength = nonEmptyBuckets > 0 ? static_cast<double>(totalElements) / nonEmptyBuckets : 0;
}

// Удаление элемента из хеш-таблицы
void ChainingHashTable::remove(int key) {
    int index = hash(key);
    ForwardListTwo* list = getList(index);
    if (list) {
        deleteNodeIndexTwo(list, to_string(key));  // Удаление из цепочки
    }
}

// Вывод детальной статистики хеш-таблицы
void ChainingHashTable::printDetailedStatistics() const {
    int totalElements = 0;
    int emptyBuckets = 0;
    int maxChain = 0;
    int minChain = INT_MAX;

    cout << "Распределение цепочек:\n";

    // Анализ всех цепочек
    for (int i = 0; i < size; i++) {
        ForwardListTwo* list = getList(i);
        int chainLength = list ? countNodesTwo(*list) : 0;

        // Сбор статистики
        if (chainLength == 0) emptyBuckets++;
        if (chainLength > maxChain) maxChain = chainLength;
        if (chainLength < minChain && chainLength > 0) minChain = chainLength;

        totalElements += chainLength;

        // Вывод первых 10 цепочек для примера
        if (i < 10) {
            cout << "Ячейка " << i << ": " << chainLength << " элементов\n";
        }
    }

    // Вывод общей статистики
    cout << "\nОбщая статистика:\n";
    cout << "Минимальная длина непустой цепочки: " << (minChain == INT_MAX ? 0 : minChain) << "\n";
    cout << "Максимальная длина цепочки: " << maxChain << "\n";
    cout << "Пустых ячеек: " << emptyBuckets << " из " << size << "\n";
    cout << "Общее количество элементов: " << totalElements << "\n";
    cout << "Коэффициент заполнения: " << static_cast<double>(totalElements) / size << "\n";
}