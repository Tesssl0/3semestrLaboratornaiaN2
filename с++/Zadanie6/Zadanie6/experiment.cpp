#include "experiment.h"
#include <iostream>
#include <iomanip>

// Проведение одиночного эксперимента для N элементов
ExperimentResult HashTableExperiment::runSingleExperiment(int N, CollisionResolution method) {
    ExperimentResult result;
    result.method = method;

    // Оптимальные коэффициенты заполнения для разных методов
    if (method == CollisionResolution::CHAINING) {
        result.tableSize = std::max(1, N / 100);  // a = 1.0 - нормально для цепочек
    }
    else {
        // Для открытой адресации используем меньший коэффициент заполнения
        result.tableSize = std::max(1, N * 2);    // a = 0.5 - оптимально для открытой адресации
    }

    result.loadFactor = static_cast<double>(N) / result.tableSize;

    // Генерация случайных чисел и замер времени
    Timer timer;
    timer.start();
    Array* randomNumbers = RandomGenerator::generateRandomNumbers(N);
    result.genTime = timer.elapsedMillis();

    // Создание хеш-таблицы с выбранным методом
    HashTable hashTable(result.tableSize, method);

    // Заполнение таблицы и замер времени вставки
    timer.start();
    for (int i = 0; i < length(*randomNumbers); ++i) {
        int num = std::stoi(get(*randomNumbers, i));
        hashTable.insert(num);
    }
    result.insertTime = timer.elapsedMillis();

    // Сбор статистики по цепям
    hashTable.getStatistics(result.minLength, result.maxLength, result.avgLength);

    // Очистка памяти
    destroy(*randomNumbers);
    delete randomNumbers;

    return result;
}

// Проведение сравнительного эксперимента с разными размерами таблицы
void HashTableExperiment::runComparisonExperiment(int N) {
    std::cout << "=== Сравнительный эксперимент для N = " << N << " ===\n";

    // Генерация общих данных для всех тестов
    Array* randomNumbers = RandomGenerator::generateRandomNumbers(N);

    // Различные размеры таблицы для сравнения
    int tableSizes[] = { N / 200, N / 100, N / 50, N / 20 };

    // Тестирование каждого размера таблицы
    for (int tableSize : tableSizes) {
        if (tableSize < 1) tableSize = 1;

        HashTable hashTable(tableSize, CollisionResolution::CHAINING);

        // Заполнение таблицы
        for (int i = 0; i < length(*randomNumbers); ++i) {
            int num = std::stoi(get(*randomNumbers, i));
            hashTable.insert(num);
        }

        int minLength, maxLength;
        double avgLength;
        hashTable.getStatistics(minLength, maxLength, avgLength);

        // Вывод результатов сравнения
        std::cout << "Размер таблицы: " << std::setw(6) << tableSize
            << " | Заполнение: " << std::setw(8) << std::fixed << std::setprecision(2)
            << static_cast<double>(N) / tableSize
            << " | Цепочки: мин=" << std::setw(3) << minLength
            << ", макс=" << std::setw(3) << maxLength
            << ", средняя=" << std::setw(6) << std::setprecision(2) << avgLength << "\n";
    }

    // Очистка
    destroy(*randomNumbers);
    delete randomNumbers;
    std::cout << "\n" << std::string(50, '=') << "\n\n";
}

// Сравнение разных методов разрешения коллизий
void HashTableExperiment::runMethodComparisonExperiment(int N) {
    std::cout << "=== Сравнение методов разрешения коллизий для N = " << N << " ===\n";

    CollisionResolution methods[] = {
        CollisionResolution::CHAINING,
        CollisionResolution::LINEAR_PROBING,
        CollisionResolution::QUADRATIC_PROBING
    };

    const char* methodNames[] = {
        "Метод цепочек",
        "Линейное пробирование",
        "Квадратичное пробирование"
    };

    for (int i = 0; i < 3; ++i) {
        ExperimentResult result = runSingleExperiment(N, methods[i]);

        std::cout << methodNames[i] << ":\n";
        std::cout << "  Время вставки: " << result.insertTime << " мс\n";
        std::cout << "  Цепочки: мин=" << result.minLength
            << ", макс=" << result.maxLength
            << ", средняя=" << std::fixed << std::setprecision(2) << result.avgLength << "\n";
        std::cout << "  Коэф. заполнения: " << result.loadFactor << "\n\n";
    }

    std::cout << std::string(50, '=') << "\n\n";
}

// Вывод результатов одиночного эксперимента
void HashTableExperiment::printResult(const ExperimentResult& result, int N) {
    const char* methodName = "";
    switch (result.method) {
    case CollisionResolution::CHAINING:
        methodName = "Метод цепочек";
        break;
    case CollisionResolution::LINEAR_PROBING:
        methodName = "Линейное пробирование";
        break;
    case CollisionResolution::QUADRATIC_PROBING:
        methodName = "Квадратичное пробирование";
        break;
    }

    std::cout << "=== Эксперимент для N = " << N << " (" << methodName << ") ===\n";
    std::cout << "Размер хеш-таблицы: " << result.tableSize << "\n";
    std::cout << "Коэффициент заполнения: " << std::fixed << std::setprecision(2)
        << result.loadFactor << "\n";
    std::cout << "Время генерации чисел: " << result.genTime << " мс\n";
    std::cout << "Время вставки: " << result.insertTime << " мс\n";
    std::cout << "\nРезультаты:\n";
    std::cout << "Самая короткая цепочка: " << result.minLength << "\n";
    std::cout << "Самая длинная цепочка: " << result.maxLength << "\n";
    std::cout << "Средняя длина непустой цепочки: " << std::fixed << std::setprecision(2)
        << result.avgLength << "\n";
}

// Вывод детальной статистики хеш-таблицы
void HashTableExperiment::printDetailedStatistics(const HashTable& hashTable) {
    hashTable.printDetailedStatistics();
}