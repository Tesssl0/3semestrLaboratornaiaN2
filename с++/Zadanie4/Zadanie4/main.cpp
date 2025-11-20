#include <iostream>
#include <string>
#include "array.h"
#include "pattern_matcher.h"

using namespace std;

int main() {
    setlocale(LC_ALL, "Ru");
    Array sequences;
    init(sequences, 5);

    Array patterns;
    init(patterns, 5);

    cout << "¬ведите строку: ";
    string s;
    cin >> s;
    add(sequences, s);

    cout << "¬ведите шаблон: ";
    string pattern;
    cin >> pattern;

    for (int i = 0; i < length(sequences); i++) {
        string current = get(sequences, i);
        cout << current << " - ";
        if (matchPattern(current, pattern))
            cout << "соответствует шаблону " << endl;
        else
            cout << "не соответствует шаблону " << endl;
    }

    destroy(sequences);
    destroy(patterns);
    return 0;
}