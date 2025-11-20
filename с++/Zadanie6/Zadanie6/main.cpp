#include <iostream>
#include "experiment.h"

int main() {
    std::setlocale(LC_ALL, "Ru");

    // Основные эксперименты
    const int sizes[] = { 5000, 10000, 20000 };

    std::cout << "=== ЭКСПЕРИМЕНТЫ С ХЕШ-ТАБЛИЦАМИ ===\n\n";

    // Сравнение методов для разных размеров
    for (int N : sizes) {
        HashTableExperiment::runMethodComparisonExperiment(N);
    }

    // Детальные эксперименты с методом цепочек
    std::cout << "=== ДЕТАЛЬНЫЕ ЭКСПЕРИМЕНТЫ (МЕТОД ЦЕПОЧЕК) ===\n\n";
    for (int N : sizes) {
        ExperimentResult result = HashTableExperiment::runSingleExperiment(N, CollisionResolution::CHAINING);
        HashTableExperiment::printResult(result, N);

        // Для демонстрации детальной статистики используем первую таблицу
        if (N == 5000) {
            HashTable demoTable(std::max(1, N / 100), CollisionResolution::CHAINING);
            Array* demoNumbers = RandomGenerator::generateRandomNumbers(100);
            for (int i = 0; i < length(*demoNumbers); ++i) {
                int num = std::stoi(get(*demoNumbers, i));
                demoTable.insert(num);
            }
            std::cout << "\nДетальная статистика для демо-таблицы:\n";
            HashTableExperiment::printDetailedStatistics(demoTable);
            destroy(*demoNumbers);
            delete demoNumbers;
        }

        std::cout << "\n" << std::string(50, '-') << "\n\n";
    }

    // Сравнительные эксперименты с разными размерами таблиц
    HashTableExperiment::runComparisonExperiment(10000);

    // Дополнительный эксперимент: сравнение методов для большой таблицы
    std::cout << "=== СРАВНЕНИЕ МЕТОДОВ ДЛЯ БОЛЬШОЙ ТАБЛИЦЫ ===\n";
    HashTableExperiment::runMethodComparisonExperiment(50000);

    return 0;
}