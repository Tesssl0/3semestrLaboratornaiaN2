#include "pattern_matcher.h"

bool matchPattern(const std::string& str, const std::string& pattern) {
    int n = str.size(), m = pattern.size();
    int i = 0, j = 0, star = -1, match = 0;

    while (i < n) {
        if (j < m && (pattern[j] == '?' || pattern[j] == str[i])) {
            i++; j++;
        }
        else if (j < m && pattern[j] == '*') {
            star = j++;
            match = i;
        }
        else if (star != -1) {
            j = star + 1;
            i = ++match;
        }
        else {
            return false;
        }
    }
    while (j < m && pattern[j] == '*') j++;
    return j == m;
}