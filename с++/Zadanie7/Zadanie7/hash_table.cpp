#include <iostream>
#include <string>
#include <algorithm>
#include <climits>
#include "hash_table.h"

using namespace std;

// Хеш-функция: вычисляет индекс в таблице по ключу
int ChainingHashTable::hash(int key) const {
    return abs(key) % size;  // используем модуль для равномерного распределения
}

// Получение списка по индексу из таблицы
ForwardListTwo* ChainingHashTable::getList(int index) const {
    if (index < 0 || index >= size) return nullptr;  // проверка валидности индекса
    string ptrStr = get(*table, index);              // получаем строку с указателем
    if (ptrStr.empty()) return nullptr;              // если строка пустая - возвращаем nullptr
    return reinterpret_cast<ForwardListTwo*>(stoll(ptrStr));  // преобразуем строку обратно в указатель
}

// Очистка списка: удаление всех узлов и самого списка
void ChainingHashTable::clearList(ForwardListTwo* list) {
    if (!list) return;  // проверка на nullptr
    // Последовательно удаляем все узлы из головы списка
    while (list->head != nullptr) {
        deleteNodeHeadTwo(list);
    }
    delete list;  // освобождаем память самого списка
}

// Конструктор хеш-таблицы
ChainingHashTable::ChainingHashTable(int tableSize) : size(tableSize) {
    table = new Array;        // создаем массив для таблицы
    init(*table, size);       // инициализируем массив

    // Инициализируем каждую ячейку таблицы пустым списком
    for (int i = 0; i < size; i++) {
        ForwardListTwo* newList = new ForwardListTwo;  // создаем новый список
        newList->head = nullptr;  // инициализируем голову
        newList->tail = nullptr;  // инициализируем хвост
        add(*table, "");          // добавляем пустую строку в массив
        // Сохраняем указатель на список как строку в массиве
        set(*table, i, to_string(reinterpret_cast<long long>(newList)));
    }
}

// Деструктор хеш-таблицы
ChainingHashTable::~ChainingHashTable() {
    // Очищаем все списки в таблице
    for (int i = 0; i < size; i++) {
        ForwardListTwo* list = getList(i);  // получаем список по индексу
        clearList(list);                    // очищаем список
    }
    destroy(*table);  // уничтожаем массив таблицы
    delete table;     // освобождаем память таблицы
}

// Вставка элемента в хеш-таблицу
void ChainingHashTable::insert(int key) {
    int index = hash(key);           // вычисляем индекс
    ForwardListTwo* list = getList(index);  // получаем список для этого индекса
    if (list) {
        addNodeTailTwo(list, to_string(key));  // добавляем ключ в конец списка
    }
}

// Поиск элемента в хеш-таблице
bool ChainingHashTable::search(int key) const {
    int index = hash(key);           // вычисляем индекс
    ForwardListTwo* list = getList(index);  // получаем список для этого индекса
    if (!list) return false;         // если списка нет - элемент не найден
    return findNodeIndexTwo(list, to_string(key));  // ищем ключ в списке
}

// Получение статистики по хеш-таблице
void ChainingHashTable::getStatistics(int& minLength, int& maxLength, double& avgLength) const {
    minLength = INT_MAX;  // инициализируем минимальную длину
    maxLength = 0;        // инициализируем максимальную длину
    int totalElements = 0;    // общее количество элементов
    int nonEmptyBuckets = 0;  // количество непустых ячеек

    // Проходим по всем ячейкам таблицы
    for (int i = 0; i < size; i++) {
        ForwardListTwo* list = getList(i);  // получаем список
        if (list) {
            int listSize = countNodesTwo(*list);  // считаем количество элементов в списке
            if (listSize > 0) {
                // Обновляем минимальную и максимальную длину
                if (listSize < minLength) minLength = listSize;
                if (listSize > maxLength) maxLength = listSize;
                totalElements += listSize;  // увеличиваем общее количество элементов
                nonEmptyBuckets++;          // увеличиваем счетчик непустых ячеек
            }
        }
    }

    // Корректируем значения если нет непустых ячеек
    if (minLength == INT_MAX) minLength = 0;
    // Вычисляем среднюю длину цепочек
    avgLength = nonEmptyBuckets > 0 ? static_cast<double>(totalElements) / nonEmptyBuckets : 0;
}

// Удаление элемента из хеш-таблицы
void ChainingHashTable::remove(int key) {
    int index = hash(key);           // вычисляем индекс
    ForwardListTwo* list = getList(index);  // получаем список для этого индекса
    if (list) {
        deleteNodeIndexTwo(list, to_string(key));  // удаляем ключ из списка
    }
}

// Детальная статистика по хеш-таблице
void ChainingHashTable::printDetailedStatistics() const {
    int totalElements = 0;  // общее количество элементов
    int emptyBuckets = 0;   // количество пустых ячеек
    int maxChain = 0;       // максимальная длина цепочки
    int minChain = INT_MAX; // минимальная длина непустой цепочки

    cout << "Распределение цепочек:\n";
    // Анализируем каждую ячейку таблицы
    for (int i = 0; i < size; i++) {
        ForwardListTwo* list = getList(i);  // получаем список
        int chainLength = list ? countNodesTwo(*list) : 0;  // вычисляем длину цепочки

        // Собираем статистику
        if (chainLength == 0) emptyBuckets++;
        if (chainLength > maxChain) maxChain = chainLength;
        if (chainLength < minChain && chainLength > 0) minChain = chainLength;

        totalElements += chainLength;  // увеличиваем общее количество элементов

        // Показываем первые 10 цепочек для примера
        if (i < 10) {
            cout << "Ячейка " << i << ": " << chainLength << " элементов\n";
        }
    }

    // Выводим общую статистику
    cout << "\nОбщая статистика:\n";
    cout << "Минимальная длина непустой цепочки: " << (minChain == INT_MAX ? 0 : minChain) << "\n";
    cout << "Максимальная длина цепочки: " << maxChain << "\n";
    cout << "Пустых ячеек: " << emptyBuckets << " из " << size << "\n";
    cout << "Общее количество элементов: " << totalElements << "\n";
    cout << "Коэффициент заполнения: " << static_cast<double>(totalElements) / size << "\n";
}