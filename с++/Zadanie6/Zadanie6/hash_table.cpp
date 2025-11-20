#include <iostream>
#include <string>
#include <algorithm>
#include <cmath>
#include "hash_table.h"

using namespace std;

// Основная хеш-функция - метод деления
int HashTable::hash1(int key) const {
    return abs(key) % size;
}

// Вторая хеш-функция для двойного хеширования
int HashTable::hash2(int key) const {
    return prime - (abs(key) % prime);
}

// Получение следующего простого числа после n
int HashTable::getNextPrime(int n) const {
    while (true) {
        bool isPrime = true;
        // Проверка на простоту
        for (int i = 2; i <= sqrt(n); i++) {
            if (n % i == 0) {
                isPrime = false;
                break;
            }
        }
        if (isPrime) return n;
        n++;
    }
}

// Методы для работы с цепочками (chaining)

// Получить список по индексу таблицы
ForwardListTwo* HashTable::getList(int index) const {
    if (index < 0 || index >= size) return nullptr;
    string ptrStr = get(*table, index);
    if (ptrStr.empty()) return nullptr;
    // Преобразование строки обратно в указатель
    return reinterpret_cast<ForwardListTwo*>(stoll(ptrStr));
}

// Очистка списка и освобождение памяти
void HashTable::clearList(ForwardListTwo* list) {
    if (!list) return;
    // Удаление всех узлов списка
    while (list->head != nullptr) {
        deleteNodeHeadTwo(list);
    }
    delete list;
}

// Методы для открытой адресации

// Проверка, помечена ли ячейка как удаленная
bool HashTable::isDeleted(int index) const {
    if (method == CollisionResolution::CHAINING) return false;
    string value = get(*table, index);
    return value == "DELETED";
}

// Проверка, пуста ли ячейка
bool HashTable::isEmpty(int index) const {
    if (method == CollisionResolution::CHAINING) {
        // Для метода цепочек - проверяем пустоту списка
        ForwardListTwo* list = getList(index);
        return !list || countNodesTwo(*list) == 0;
    }
    else {
        // Для открытой адресации - проверяем значение ячейки
        string value = get(*table, index);
        return value.empty() || value == "DELETED";
    }
}

// Поиск слота для вставки/поиска при открытой адресации
int HashTable::findSlotOpenAddressing(int key) const {
    int index = hash1(key);
    string keyStr = to_string(key);

    // Линейное пробирование
    if (method == CollisionResolution::LINEAR_PROBING) {
        for (int i = 0; i < size; i++) {
            int currentIndex = (index + i) % size;
            string currentValue = get(*table, currentIndex);

            // Пустая ячейка или удаленная - можем использовать
            if (isEmpty(currentIndex) || isDeleted(currentIndex)) {
                return currentIndex;
            }
            // Нашли ключ - возвращаем для поиска/удаления
            if (currentValue == keyStr) {
                return currentIndex;
            }
        }
    }
    // Квадратичное пробирование
    else if (method == CollisionResolution::QUADRATIC_PROBING) {
        for (int i = 0; i < size; i++) {
            int currentIndex = (index + i * i) % size;
            string currentValue = get(*table, currentIndex);

            if (isEmpty(currentIndex) || isDeleted(currentIndex)) {
                return currentIndex;
            }
            if (currentValue == keyStr) {
                return currentIndex;
            }
        }
    }
    return -1; // Таблица заполнена
}

// Конструктор хеш-таблицы
HashTable::HashTable(int tableSize, CollisionResolution resolutionMethod)
    : size(tableSize), method(resolutionMethod) {
    table = new Array;
    init(*table, size);
    prime = getNextPrime(size / 2); // Для двойного хеширования

    // Инициализация в зависимости от метода разрешения коллизий
    if (method == CollisionResolution::CHAINING) {
        // Инициализация цепочек - создаем список для каждой ячейки
        for (int i = 0; i < size; i++) {
            ForwardListTwo* newList = new ForwardListTwo;
            newList->head = nullptr;
            newList->tail = nullptr;
            add(*table, "");
            // Сохраняем указатель на список как строку
            set(*table, i, to_string(reinterpret_cast<long long>(newList)));
        }
    }
    else {
        // Инициализация для открытой адресации - пустые ячейки
        for (int i = 0; i < size; i++) {
            add(*table, ""); // Пустые ячейки
        }
    }
}

// Деструктор хеш-таблицы
HashTable::~HashTable() {
    // Очистка памяти в зависимости от метода
    if (method == CollisionResolution::CHAINING) {
        for (int i = 0; i < size; i++) {
            ForwardListTwo* list = getList(i);
            clearList(list);
        }
    }
    destroy(*table);
    delete table;
}

// Вставка элемента в хеш-таблицу
void HashTable::insert(int key) {
    if (method == CollisionResolution::CHAINING) {
        // Метод цепочек - добавляем в конец списка
        int index = hash1(key);
        ForwardListTwo* list = getList(index);
        if (list) {
            addNodeTailTwo(list, to_string(key));
        }
    }
    else {
        // Открытая адресация - находим свободный слот
        int index = findSlotOpenAddressing(key);
        if (index != -1) {
            set(*table, index, to_string(key));
        }
    }
}

// Поиск элемента в хеш-таблице
bool HashTable::search(int key) const {
    if (method == CollisionResolution::CHAINING) {
        // Метод цепочек - ищем в списке
        int index = hash1(key);
        ForwardListTwo* list = getList(index);
        if (!list) return false;
        return findNodeIndexTwo(list, to_string(key));
    }
    else {
        // Открытая адресация - используем поиск слота
        int index = findSlotOpenAddressing(key);
        return index != -1 && get(*table, index) == to_string(key);
    }
}

// Удаление элемента из хеш-таблицы
void HashTable::remove(int key) {
    if (method == CollisionResolution::CHAINING) {
        // Метод цепочек - удаляем из списка
        int index = hash1(key);
        ForwardListTwo* list = getList(index);
        if (list) {
            deleteNodeIndexTwo(list, to_string(key));
        }
    }
    else {
        // Открытая адресация - помечаем как удаленное
        int index = findSlotOpenAddressing(key);
        if (index != -1 && get(*table, index) == to_string(key)) {
            set(*table, index, "DELETED");
        }
    }
}

// Получение статистики хеш-таблицы
void HashTable::getStatistics(int& minLength, int& maxLength, double& avgLength) const {
    minLength = INT_MAX;
    maxLength = 0;
    int totalElements = 0;
    int nonEmptyBuckets = 0;

    if (method == CollisionResolution::CHAINING) {
        // Логика для метода цепочек
        for (int i = 0; i < size; i++) {
            ForwardListTwo* list = getList(i);
            if (list) {
                int listSize = countNodesTwo(*list);
                if (listSize > 0) {
                    if (listSize < minLength) minLength = listSize;
                    if (listSize > maxLength) maxLength = listSize;
                    totalElements += listSize;
                    nonEmptyBuckets++;
                }
            }
        }
    }
    else {
        // Логика для открытой адресации - анализ кластеров
        Array visitedArray;
        init(visitedArray, size);

        // Инициализируем массив посещенных ячеек
        for (int i = 0; i < size; i++) {
            add(visitedArray, "0"); // "0" = false, "1" = true
        }

        // Анализ кластеров для линейного пробирования
        for (int i = 0; i < size; i++) {
            if (!isEmpty(i) && !isDeleted(i) && get(visitedArray, i) == "0") {
                int clusterSize = 0;
                int current = i;

                // Измеряем размер кластера
                while (!isEmpty(current) && get(visitedArray, current) == "0") {
                    if (!isDeleted(current)) {
                        clusterSize++;
                    }
                    set(visitedArray, current, "1"); // помечаем как посещенный
                    current = (current + 1) % size; // для линейного пробирования
                }

                if (clusterSize > 0) {
                    if (clusterSize < minLength) minLength = clusterSize;
                    if (clusterSize > maxLength) maxLength = clusterSize;
                    totalElements += clusterSize;
                    nonEmptyBuckets++;
                }
            }
        }

        destroy(visitedArray);
    }

    // Корректировка значений по умолчанию
    if (minLength == INT_MAX) minLength = 0;
    avgLength = nonEmptyBuckets > 0 ? static_cast<double>(totalElements) / nonEmptyBuckets : 0;
}

// Вывод подробной статистики хеш-таблицы
void HashTable::printDetailedStatistics() const {
    int totalElements = 0;
    int emptyBuckets = 0;
    int maxChain = 0;
    int minChain = INT_MAX;

    // Вывод информации о методе разрешения коллизий
    cout << "Метод разрешения коллизий: ";
    switch (method) {
    case CollisionResolution::CHAINING:
        cout << "Метод цепочек\n";
        break;
    case CollisionResolution::LINEAR_PROBING:
        cout << "Линейное пробирование\n";
        break;
    case CollisionResolution::QUADRATIC_PROBING:
        cout << "Квадратичное пробирование\n";
        break;
    }

    cout << "Распределение:\n";

    // Сбор статистики по всем ячейкам
    for (int i = 0; i < size; i++) {
        int chainLength = 0;

        if (method == CollisionResolution::CHAINING) {
            ForwardListTwo* list = getList(i);
            chainLength = list ? countNodesTwo(*list) : 0;
        }
        else {
            chainLength = (!isEmpty(i) && !isDeleted(i)) ? 1 : 0;
        }

        // Обновление статистики
        if (chainLength == 0) emptyBuckets++;
        if (chainLength > maxChain) maxChain = chainLength;
        if (chainLength < minChain && chainLength > 0) minChain = chainLength;

        totalElements += chainLength;

        // Вывод первых 10 ячеек для примера
        if (i < 10) {
            cout << "Ячейка " << i << ": " << chainLength << " элементов\n";
        }
    }

    // Вывод итоговой статистики
    cout << "\nОбщая статистика:\n";
    cout << "Минимальная длина непустой цепочки: " << (minChain == INT_MAX ? 0 : minChain) << "\n";
    cout << "Максимальная длина цепочки: " << maxChain << "\n";
    cout << "Пустых ячеек: " << emptyBuckets << " из " << size << "\n";
    cout << "Общее количество элементов: " << totalElements << "\n";
    cout << "Коэффициент заполнения: " << static_cast<double>(totalElements) / size << "\n";
}

// Установка метода разрешения коллизий
void HashTable::setResolutionMethod(CollisionResolution method) {
    this->method = method;
}