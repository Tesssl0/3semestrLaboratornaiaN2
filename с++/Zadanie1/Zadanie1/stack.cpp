#include <iostream>
#include "stack.h"
using namespace std;

// Конструктор
Stack::Stack(int size, bool silentMode) {
    capacity = size;
    arr = new string[capacity];
    topIndex = -1;
    silent = silentMode;

    // Инициализируем все элементы пустыми строками
    for (int i = 0; i < capacity; i++) {
        arr[i] = "";
    }
}

// Деструктор
Stack::~Stack() {
    if (arr != nullptr) {
        delete[] arr;
        arr = nullptr;
    }
}

// Поиск элемента в стеке по значению
int findInStack(const Stack& stack, const std::string& value) {
    for (int i = 0; i <= stack.topIndex; i++) {
        if (stack.arr[i] == value) {
            return i;
        }
    }
    return -1;
}

// Вывод всех элементов стека
void printStack(const Stack& stack) {
    for (int i = 0; i <= stack.topIndex; i++) {
        cout << stack.arr[i];
        if (i < stack.topIndex) {
            cout << ", ";
        }
    }
    cout << endl;
}

// Добавление элемента в стек
void Stack::push(const string& value) {
    if (topIndex >= capacity - 1) {
        if (!silent) {
            cout << "Стек переполнен!" << endl;
        }
        return;
    }
    topIndex++;
    arr[topIndex] = value;
    if (!silent) {
        cout << "Элемент " << value << " добавлен в стек" << endl;
    }
}

// Удаление элемента из стека
void Stack::pop() {
    if (topIndex < 0) {
        if (!silent) {
            cout << "Стек пуст!" << endl;
        }
        return;
    }
    if (!silent) {
        cout << "Элемент " << arr[topIndex] << " удалён из стека" << endl;
    }
    topIndex--;
}

// Просмотр верхнего элемента
string Stack::top() {
    if (topIndex < 0) {
        if (!silent) {
            cout << "Стек пуст!" << endl;
        }
        return "";
    }
    return arr[topIndex];
}

// Проверка, пуст ли стек
bool Stack::isEmpty() {
    return topIndex < 0;
}

// Конструктор копирования
Stack::Stack(const Stack& other) {
    capacity = other.capacity;
    topIndex = other.topIndex;
    silent = other.silent;
    arr = new string[capacity];
    for (int i = 0; i <= topIndex; i++) {
        arr[i] = other.arr[i];
    }
}

// Оператор присваивания
Stack& Stack::operator=(const Stack& other) {
    if (this != &other) {
        delete[] arr;
        capacity = other.capacity;
        topIndex = other.topIndex;
        silent = other.silent;
        arr = new string[capacity];
        for (int i = 0; i <= topIndex; i++) {
            arr[i] = other.arr[i];
        }
    }
    return *this;
}