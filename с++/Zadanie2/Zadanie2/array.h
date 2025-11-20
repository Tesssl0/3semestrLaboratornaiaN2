#ifndef ARRAY_H
#define ARRAY_H

#include <string>

struct Array {
    std::string* data;    // указатель на массив
    int size;     // текущий размер
    int capacity; // максимальная вместимость
};

void init(Array& arr, int initialCapacity = 10);
void destroy(Array& arr);
void resize(Array& arr);
void add(Array& arr, const std::string& value);
void addAt(Array& arr, int index, const std::string& value);
std::string get(const Array& arr, int index);
void remove(Array& arr, int index);
void set(Array& arr, int index, const std::string& value);
int length(const Array& arr);
void print(const Array& arr);
int findInArray(const Array& arr, const std::string& value);
#endif  