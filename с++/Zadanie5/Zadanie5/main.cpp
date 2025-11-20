#include "fullBinaryTree.h"
#include <iostream>
#include <string>

using namespace std;

int main() {
    setlocale(LC_ALL, "Ru");
    fullBinaryTree tree;
    string input;

    cout << "Введите последовательность целых чисел (для завершения введите 'q'):" << endl;

    while (true) {
        cin >> input;

        if (input == "q" || input == "Q") {
            break;
        }

        try {
            stoi(input); // Проверяем, что введено число

            // Если числа нет в дереве, добавляем его
            if (BFS(&tree, input) == nullptr) {
                insertBST(&tree, input);
            }
        }
        catch (const exception& e) {
            cout << "Ошибка: '" << input << "' не является числом" << endl;
            break;
        }
    }

    // Вычисляем высоту дерева с помощью функции из библиотеки
    int height = getTreeHeight(tree.root);
    cout << "Высота получившегося дерева: " << height << endl;

    // Очищаем память
    clearFullBinaryTree(&tree);

    return 0;
}