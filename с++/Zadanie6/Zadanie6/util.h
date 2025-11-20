#ifndef UTIL_H
#define UTIL_H

#include <random>
#include <chrono>
#include "array.h"

// Класс для генерации случайных чисел
class RandomGenerator {
public:
    // Генерация массива случайных чисел
    static Array* generateRandomNumbers(int count, int minVal = 0, int maxVal = 1000000);
};

// Класс для измерения времени выполнения
class Timer {
private:
    std::chrono::high_resolution_clock::time_point startTime;

public:
    void start();                       // Запуск таймера
    unsigned long elapsedMillis() const;    // Получение прошедшего времени в миллисекундах
};

#endif