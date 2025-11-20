#include <iostream>
#include <string>
#include "Set.h"

using namespace std;

// Функция для преобразования пары символов в уникальное число
// Используется для кодирования пар оснований ДНК
int pairToInt(const string& pair) {
    // Преобразуем два символа в число от 0 до 675 (26*26)
    return (pair[0] - 'A') * 26 + (pair[1] - 'A');
}

int main() {
    string genome1, genome2;  // Входные геномные последовательности
    cin >> genome1 >> genome2;  // Чтение геномов из входного потока

    // Создаем множество для хранения пар оснований второго генома
    Set pairsSet(1000);  // Хеш-таблица размером 1000 ячеек

    // Добавляем все пары из второго генома в множество
    for (size_t i = 0; i + 1 < genome2.length(); i++) {
        string pair = genome2.substr(i, 2);  // Извлекаем пару символов
        pairsSet.add(pairToInt(pair));  // Преобразуем в число и добавляем в множество
    }

    int closeness = 0;  // Счетчик общих пар

    // Подсчитываем пары из первого генома, которые есть во втором
    for (size_t i = 0; i + 1 < genome1.length(); i++) {
        string pair = genome1.substr(i, 2);  // Извлекаем пару из первого генома
        int pairValue = pairToInt(pair);  // Преобразуем в число

        // Если пара присутствует во втором геноме
        if (pairsSet.contains(pairValue)) {
            closeness++;  // Увеличиваем счетчик близости
        }
    }

    cout << closeness << endl;  // Выводим результат
    return 0;
}