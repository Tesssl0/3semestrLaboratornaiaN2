#include <iostream>
#include "array.h"
using namespace std;

// Инициализация массива
void init(Array& arr, int initialCapacity) {
    arr.capacity = initialCapacity;
    arr.size = 0;
    arr.data = new string[arr.capacity];
}

// Освобождение памяти
void destroy(Array& arr) {
    delete[] arr.data;
    arr.data = nullptr;
    arr.size = 0;
    arr.capacity = 0;
}

// Увеличение размера массива
void resize(Array& arr) {
    arr.capacity *= 2;
    string* newData = new string[arr.capacity];
    for (int i = 0; i < arr.size; i++) {
        newData[i] = arr.data[i];
    }
    delete[] arr.data;
    arr.data = newData;
}

// Добавление элемента в конец
void add(Array& arr, const string& value) {
    if (arr.size == arr.capacity) {
        resize(arr);
    }
    arr.data[arr.size++] = value;
}

// Добавление элемента по индексу
void addAt(Array& arr, int index, const string& value) {
    if (arr.size == arr.capacity) {
        resize(arr);
    }
    for (int i = arr.size; i > index; i--) {
        arr.data[i] = arr.data[i - 1];
    }
    arr.data[index] = value;
    arr.size++;
}

// Получение элемента по индексу
string get(const Array& arr, int index) {
    return arr.data[index];
}

// Удаление элемента по индексу
void remove(Array& arr, int index) {
    for (int i = index; i < arr.size - 1; i++) {
        arr.data[i] = arr.data[i + 1];
    }
    arr.size--;
}

// Замена элемента по индексу
void set(Array& arr, int index, const string& value) {
    arr.data[index] = value;
}

// Получение длины массива
int length(const Array& arr) {
    return arr.size;
}

// Чтение всех элементов
void print(const Array& arr) {
    for (int i = 0; i < arr.size; i++) {
        cout << arr.data[i] << " ";
    }
    cout << endl;
}

// Поиск элемента по значению
int findInArray(const Array& arr, const std::string& value) {
    for (int i = 0; i < arr.size; i++) {
        if (arr.data[i] == value) {
            return i;
        }
    }
    return -1;
}