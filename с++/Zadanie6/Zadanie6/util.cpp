#include "util.h"
#include <iostream>

// Генерация массива случайных чисел
Array* RandomGenerator::generateRandomNumbers(int count, int minVal, int maxVal) {
    Array* numbers = new Array;
    init(*numbers);  // Инициализация массива

    // Инициализация генератора случайных чисел
    std::random_device rd;  // Источник энтропии
    std::mt19937 gen(rd()); // Генератор Mersenne Twister
    std::uniform_int_distribution<int> dis(minVal, maxVal); // Равномерное распределение

    // Генерация count случайных чисел
    for (int i = 0; i < count; ++i) {
        add(*numbers, std::to_string(dis(gen))); // Добавление числа как строки в массив
    }

    return numbers;
}

// Запуск таймера
void Timer::start() {
    startTime = std::chrono::high_resolution_clock::now(); // Запись текущего времени
}

// Получение прошедшего времени в миллисекундах
unsigned long  Timer::elapsedMillis() const {
    auto endTime = std::chrono::high_resolution_clock::now(); // Текущее время
    // Вычисление разницы и преобразование в миллисекунды
    return std::chrono::duration_cast<std::chrono::milliseconds>(endTime - startTime).count();
}