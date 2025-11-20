#pragma once
#ifndef STACK_H
#define STACK_H

#include <string>

class Stack {
public:
    std::string* arr;       // массив для хранения элементов стека
    int capacity;   // максимальный размер стека
    int topIndex;   // индекс вершины стека
    bool silent;    // флаг тихого режима

    // Конструктор
    Stack(int size, bool silentMode = false);

    // Конструктор копирования
    Stack(const Stack& other);

    int getTopIndex() const { return topIndex; } // возвращает индекс вершины
    std::string getAt(int index) const { return arr[index]; } // возвращает элемент по индексу

    // Оператор присваивания
    Stack& operator=(const Stack& other);

    // Деструктор
    ~Stack();

    // Добавление элемента в стек
    void push(const std::string& value);

    // Удаление элемента из стека
    void pop();

    // Просмотр верхнего элемента
    std::string top();

    // Проверка, пуст ли стек
    bool isEmpty();
};

// Внешние функции для работы со стеком
int findInStack(const Stack& stack, const std::string& value);
void printStack(const Stack& stack);

#endif // STACK_H