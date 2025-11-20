#pragma once
#ifndef EXPERIMENT_H
#define EXPERIMENT_H

#include "hash_table.h"
#include "util.h"

// Структура для хранения результатов эксперимента
struct ExperimentResult {
    int minLength;          // Минимальная длина цепи
    int maxLength;          // Максимальная длина цепи
    double avgLength;       // Средняя длина непустой цепи
    long long genTime;      // Время генерации чисел (мс)
    long long insertTime;   // Время вставки в таблицу (мс)
    int tableSize;          // Размер хеш-таблицы
    double loadFactor;      // Коэффициент заполнения
    CollisionResolution method; // Метод разрешения коллизий
};

// Класс для проведения экспериментов с хеш-таблицей
class HashTableExperiment {
public:
    static ExperimentResult runSingleExperiment(int N, CollisionResolution method = CollisionResolution::CHAINING);
    static void runComparisonExperiment(int N);
    static void runMethodComparisonExperiment(int N);
    static void printResult(const ExperimentResult& result, int N);
    static void printDetailedStatistics(const HashTable& hashTable);
};

#endif