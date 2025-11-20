#include <iostream>
#include <fstream>
#include <sstream>
#include <windows.h>
#include "Set.h"

using namespace std;

void printFile(const string& filename) {
    ifstream file(filename);
    if (!file.is_open()) {
        cerr << "ќшибка: не удалось открыть файл " << filename << endl;
        return;
    }
    cout << "—одержимое файла " << filename << ":" << endl;
    cout << "----------------------------------------" << endl;
    string line;
    int lineNumber = 1;
    while (getline(file, line)) {
        if (!line.empty()) {
            cout << lineNumber++ << ": " << line << endl;
        }
    }
    if (lineNumber == 1)
        cout << "‘айл пуст" << endl;
    cout << "----------------------------------------" << endl;
    file.close();
}

void printUsage() {
    cout << "»спользование: ./program --file <файл> --query <операци€>" << endl;
    cout << "ƒоступные команды:" << endl;
    cout << "  SETADD <число>  - добавить элемент в множество" << endl;
    cout << "  SETDEL <число>  - удалить элемент из множества" << endl;
    cout << "  SET_AT <число>  - проверить наличие элемента" << endl;
    cout << "  SHOW_FILE       - вывести содержимое файла" << endl;
}

int main(int argc, char* argv[]) {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);

    if (argc != 5) {
        printUsage();
        return 1;
    }

    string filename, query;
    for (int i = 1; i < argc; i++) {
        string arg = argv[i];
        if (arg == "--file" && i + 1 < argc)
            filename = argv[++i];
        else if (arg == "--query" && i + 1 < argc)
            query = argv[++i];
    }

    if (filename.empty() || query.empty()) {
        printUsage();
        return 1;
    }

    Set mySet;
    mySet.loadFromFile(filename);

    istringstream iss(query);
    string cmd, valueStr;
    iss >> cmd >> valueStr;

    if (cmd == "SHOW_FILE") {
        printFile(filename);
        return 0;
    }

    if (valueStr.empty()) {
        cerr << "ќшибка: не указано значение дл€ операции" << endl;
        return 1;
    }

    try {
        int value = stoi(valueStr);
        if (cmd == "SETADD") {
            if (mySet.add(value)) {
                cout << "Ёлемент " << value << " добавлен" << endl;
                mySet.saveToFile(filename);
            }
            else {
                cout << "Ёлемент " << value << " уже существует" << endl;
            }
        }
        else if (cmd == "SETDEL") {
            if (mySet.remove(value)) {
                cout << "Ёлемент " << value << " удалЄн" << endl;
                mySet.saveToFile(filename);
            }
            else {
                cout << "Ёлемент " << value << " не найден" << endl;
            }
        }
        else if (cmd == "SET_AT") {
            cout << "Ёлемент " << value
                << (mySet.contains(value) ? " присутствует" : " отсутствует")
                << " в множестве" << endl;
        }
        else {
            cerr << "Ќеизвестна€ операци€: " << cmd << endl;
            printUsage();
            return 1;
        }
    }
    catch (...) {
        cerr << "ќшибка: некорректное значение числа" << endl;
        return 1;
    }

    return 0;
}
