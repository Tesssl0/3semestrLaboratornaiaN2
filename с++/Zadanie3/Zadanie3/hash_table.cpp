#include <iostream>
#include <string>
#include <algorithm>
#include <climits>
#include "hash_table.h"

using namespace std;

// Хеш-функция: преобразует ключ в индекс таблицы
int ChainingHashTable::hash(int key) const {
    return abs(key) % size;  // Берем модуль от абсолютного значения ключа
}

// Получение списка по индексу из хеш-таблицы
ForwardListTwo* ChainingHashTable::getList(int index) const {
    if (index < 0 || index >= size) return nullptr;  // Проверка корректности индекса
    string ptrStr = get(*table, index);  // Получаем строку с указателем
    if (ptrStr.empty()) return nullptr;  // Если строка пустая - возвращаем nullptr
    return reinterpret_cast<ForwardListTwo*>(stoll(ptrStr));  // Преобразуем строку обратно в указатель
}

// Очистка списка (освобождение памяти всех узлов)
void ChainingHashTable::clearList(ForwardListTwo* list) {
    if (!list) return;  // Если список не существует - выходим
    while (list->head != nullptr) {
        deleteNodeHeadTwo(list);  // Удаляем все узлы с головы списка
    }
    delete list;  // Освобождаем память самой структуры списка
}

// Конструктор хеш-таблицы
ChainingHashTable::ChainingHashTable(int tableSize) : size(tableSize) {
    table = new Array;  // Создаем динамический массив
    init(*table, size);  // Инициализируем массив заданным размером

    // Инициализируем каждую ячейку таблицы пустым списком
    for (int i = 0; i < size; i++) {
        ForwardListTwo* newList = new ForwardListTwo;  // Создаем новый список
        newList->head = nullptr;  // Инициализируем указатели
        newList->tail = nullptr;
        add(*table, "");  // Добавляем пустую строку в массив
        set(*table, i, to_string(reinterpret_cast<long long>(newList)));  // Сохраняем указатель как строку
    }
}

// Деструктор хеш-таблицы
ChainingHashTable::~ChainingHashTable() {
    // Очищаем все списки в таблице
    for (int i = 0; i < size; i++) {
        ForwardListTwo* list = getList(i);  // Получаем список по индексу
        clearList(list);  // Очищаем список
    }
    destroy(*table);  // Уничтожаем массив
    delete table;  // Освобождаем память под структуру массива
}

// Вставка элемента в хеш-таблицу
void ChainingHashTable::insert(int key) {
    int index = hash(key);  // Вычисляем индекс с помощью хеш-функции
    ForwardListTwo* list = getList(index);  // Получаем список для этого индекса
    if (list) {
        addNodeTailTwo(list, to_string(key));  // Добавляем ключ в конец списка
    }
}

// Поиск элемента в хеш-таблице
bool ChainingHashTable::search(int key) const {
    int index = hash(key);  // Вычисляем индекс
    ForwardListTwo* list = getList(index);  // Получаем соответствующий список
    if (!list) return false;  // Если список не существует - элемент не найден
    return findNodeIndexTwo(list, to_string(key));  // Ищем ключ в списке
}

// Получение статистики о хеш-таблице
void ChainingHashTable::getStatistics(int& minLength, int& maxLength, double& avgLength) const {
    minLength = INT_MAX;  // Инициализируем минимальную длину максимальным значением
    maxLength = 0;        // Инициализируем максимальную длину нулем
    int totalElements = 0;  // Общее количество элементов
    int nonEmptyBuckets = 0;  // Количество непустых ячеек

    // Проходим по всем ячейкам таблицы
    for (int i = 0; i < size; i++) {
        ForwardListTwo* list = getList(i);  // Получаем список для текущей ячейки
        if (list) {
            int listSize = countNodesTwo(*list);  // Считаем количество элементов в списке
            if (listSize > 0) {
                // Обновляем минимальную и максимальную длину
                if (listSize < minLength) minLength = listSize;
                if (listSize > maxLength) maxLength = listSize;
                totalElements += listSize;  // Увеличиваем общее количество элементов
                nonEmptyBuckets++;  // Увеличиваем счетчик непустых ячеек
            }
        }
    }

    // Корректируем значения если нет непустых ячеек
    if (minLength == INT_MAX) minLength = 0;
    // Вычисляем среднюю длину цепочки для непустых ячеек
    avgLength = nonEmptyBuckets > 0 ? static_cast<double>(totalElements) / nonEmptyBuckets : 0;
}

// Удаление элемента из хеш-таблицы
void ChainingHashTable::remove(int key) {
    int index = hash(key);  // Вычисляем индекс
    ForwardListTwo* list = getList(index);  // Получаем соответствующий список
    if (list) {
        deleteNodeIndexTwo(list, to_string(key));  // Удаляем узел с заданным ключом
    }
}

// Детальная статистика хеш-таблицы
void ChainingHashTable::printDetailedStatistics() const {
    int totalElements = 0;    // Общее количество элементов
    int emptyBuckets = 0;     // Количество пустых ячеек
    int maxChain = 0;         // Максимальная длина цепочки
    int minChain = INT_MAX;   // Минимальная длина непустой цепочки

    cout << "Распределение цепочек:\n";
    // Анализируем каждую ячейку таблицы
    for (int i = 0; i < size; i++) {
        ForwardListTwo* list = getList(i);  // Получаем список
        int chainLength = list ? countNodesTwo(*list) : 0;  // Длина цепочки

        // Собираем статистику
        if (chainLength == 0) emptyBuckets++;
        if (chainLength > maxChain) maxChain = chainLength;
        if (chainLength < minChain && chainLength > 0) minChain = chainLength;

        totalElements += chainLength;  // Суммируем общее количество элементов

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