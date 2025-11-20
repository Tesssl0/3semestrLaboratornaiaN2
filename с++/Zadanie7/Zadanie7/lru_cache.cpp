#include "LRU_Cache.h"
#include <iostream>
#include <sstream>

using namespace std;

// Функция для обработки пользовательских запросов к LRU кэшу
void processLRUQueries() {
    int cap, Q;
    cout << "Введите емкость кэша: ";
    cin >> cap;
    cout << "Введите количество запросов: ";
    cin >> Q;
    cin.ignore();  // очищаем буфер ввода

    LRUCache cache(cap);  // создаем кэш заданной емкости

    cout << "Введите запросы (SET x y или GET x):" << endl;

    // Счетчик для результатов GET запросов
    int getCount = 0;
    const int MAX_RESULTS = 100;
    string results[MAX_RESULTS];  // массив для хранения результатов

    // Обрабатываем все запросы
    for (int i = 0; i < Q; i++) {
        string query;
        getline(cin, query);  // читаем запрос

        istringstream iss(query);
        string command;
        iss >> command;  // извлекаем команду

        // Простая конвертация в верхний регистр
        for (size_t j = 0; j < command.length(); j++) {
            if (command[j] >= 'a' && command[j] <= 'z') {
                command[j] = command[j] - 'a' + 'A';  // преобразуем в верхний регистр
            }
        }

        // Обработка команды SET
        if (command == "SET") {
            int key, value;
            if (iss >> key >> value) {
                cache.set(key, value);  // устанавливаем значение в кэше
                cout << "SET " << key << " " << value << " выполнено" << endl;
            }
        }
        // Обработка команды GET
        else if (command == "GET") {
            int key;
            if (iss >> key) {
                int result = cache.get(key);  // получаем значение из кэша
                if (getCount < MAX_RESULTS) {
                    results[getCount] = to_string(result);  // сохраняем результат
                    getCount++;
                }
                cout << "GET " << key << " = " << result << endl;
            }
        }
        cache.printCache();  // выводим текущее состояние кэша
    }

    // Выводим все результаты GET запросов
    cout << "Результаты: ";
    for (int i = 0; i < getCount; i++) {
        cout << results[i];
        if (i < getCount - 1) {
            cout << " ";
        }
    }
    cout << endl;
}