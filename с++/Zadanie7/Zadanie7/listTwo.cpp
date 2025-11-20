#include<iostream>
#include"listTwo.h"
using namespace std;

void addNodeHeadTwo(ForwardListTwo* flist, const string& num) {
    DoublyNode* newNode = new DoublyNode;
    newNode->node = num;
    newNode->next = flist->head;
    newNode->prev = nullptr;

    if (flist->head != nullptr) {
        flist->head->prev = newNode;
    }
    flist->head = newNode;

    if (flist->tail == nullptr) {
        flist->tail = newNode;
    }
}

void addNodeTailTwo(ForwardListTwo* flist, const string& num) {
    DoublyNode* newNode = new DoublyNode;
    newNode->node = num;
    newNode->next = nullptr;
    newNode->prev = flist->tail;

    if (flist->tail != nullptr) {
        flist->tail->next = newNode;
    }

    flist->tail = newNode;

    if (flist->head == nullptr) {
        flist->head = newNode;
    }
}

void addNodeAfterTwo(ForwardListTwo* flist, DoublyNode* target, const string& num) {
    if (target != nullptr && flist->head != nullptr) {
        DoublyNode* newNode = new DoublyNode;
        newNode->node = num;
        newNode->next = nullptr;
        newNode->prev = nullptr;

        // Проверяем, что target действительно находится в списке
        DoublyNode* current = flist->head;
        bool targetFound = false;

        while (current != nullptr) {
            if (current == target) {
                targetFound = true;
                break;
            }
            current = current->next;
        }

        if (targetFound) {
            // Вставляем после target
            newNode->next = target->next;
            newNode->prev = target;

            if (target->next != nullptr) {
                target->next->prev = newNode;
            }
            else {
                flist->tail = newNode;
            }
            target->next = newNode;
        }
        else {
            // target не принадлежит списку, освобождаем память
            delete newNode;
            cout << "Ошибка: указанный узел не найден в списке" << endl;
        }
    }
}

void addNodeBeforeTwo(ForwardListTwo* flist, DoublyNode* target, const string& num) {
    if (target != nullptr && flist->head != nullptr) {
        DoublyNode* newNode = new DoublyNode;
        newNode->node = num;
        newNode->next = nullptr;
        newNode->prev = nullptr;

        // Проверяем, что target действительно находится в списке
        DoublyNode* current = flist->head;
        bool targetFound = false;

        while (current != nullptr) {
            if (current == target) {
                targetFound = true;
                break;
            }
            current = current->next;
        }

        if (targetFound) {
            // Вставляем перед target
            newNode->next = target;
            newNode->prev = target->prev;

            if (target->prev != nullptr) {
                target->prev->next = newNode;
            }
            else {
                flist->head = newNode;
            }
            target->prev = newNode;
        }
        else {
            // target не принадлежит списку, освобождаем память
            delete newNode;
            cout << "Ошибка: указанный узел не найден в списке" << endl;
        }
    }
}



void deleteNodeHeadTwo(ForwardListTwo* flist) {
    if (flist->head != nullptr) {
        DoublyNode* toDelete = flist->head;
        flist->head = flist->head->next;
        if (flist->head != nullptr) {
            flist->head->prev = nullptr;
        }
        else {
            flist->tail = nullptr;
        }
        delete toDelete;
    }
}

void deleteNodeTailTwo(ForwardListTwo* flist) {
    if (flist->tail != nullptr) {
        DoublyNode* toDelete = flist->tail;
        flist->tail = flist->tail->prev;
        if (flist->tail != nullptr) {
            flist->tail->next = nullptr;
        }
        else {
            flist->head = nullptr;
        }
        delete toDelete;
    }
}

void deleteNodeAfterTwo(ForwardListTwo* flist, DoublyNode* target) {
    if (target != nullptr && flist->head != nullptr) {
        // Проверяем, что target действительно находится в списке
        DoublyNode* current = flist->head;
        bool targetFound = false;

        while (current != nullptr) {
            if (current == target) {
                targetFound = true;
                break;
            }
            current = current->next;
        }

        if (targetFound && target->next != nullptr) {
            DoublyNode* toDelete = target->next;
            target->next = toDelete->next;

            if (toDelete->next != nullptr) {
                toDelete->next->prev = target;
            }
            else {
                flist->tail = target;
            }
            delete toDelete;
        }
    }
}

void deleteNodeBeforeTwo(ForwardListTwo* flist, DoublyNode* target) {
    if (target != nullptr && flist->head != nullptr) {
        // Проверяем, что target действительно находится в списке
        DoublyNode* current = flist->head;
        bool targetFound = false;

        while (current != nullptr) {
            if (current == target) {
                targetFound = true;
                break;
            }
            current = current->next;
        }

        if (targetFound && target->prev != nullptr) {
            DoublyNode* toDelete = target->prev;
            target->prev = toDelete->prev;

            if (toDelete->prev != nullptr) {
                toDelete->prev->next = target;
            }
            else {
                flist->head = target;
            }
            delete toDelete;
        }
    }
}

bool deleteNodTwo(ForwardListTwo* flist, const string& num) {
    if (flist->head == nullptr) return false;

    DoublyNode* current = flist->head;

    while (current != nullptr && current->node != num) {
        current = current->next;
    }

    if (current == nullptr) return false;

    if (current->prev != nullptr) {
        current->prev->next = current->next;
    }
    else {
        flist->head = current->next;
    }

    if (current->next != nullptr) {
        current->next->prev = current->prev;
    }
    else {
        flist->tail = current->prev;
    }

    delete current;
    return true;
}

bool findNodeIndexTwo(ForwardListTwo* flist, const string& num) {
    DoublyNode* current = flist->head;
    while (current != nullptr) {
        if (current->node == num) return true;
        current = current->next;
    }
    return false;
}

void printListTwo(const ForwardListTwo& flist) {
    DoublyNode* current = flist.head;
    cout << "Список: ";
    while (current != nullptr) {
        cout << current->node << " ";
        current = current->next;
    }
    cout << endl;
}

int countNodesTwo(const ForwardListTwo& flist) {
    int count = 0;
    DoublyNode* current = flist.head;
    while (current != nullptr) {
        count++;
        current = current->next;
    }
    return count;
}

DoublyNode* getNodeByIndexTwo(ForwardListTwo& flist, int index) {
    if (index < 0) return nullptr;

    DoublyNode* current = flist.head;
    int currentIndex = 0;

    while (current != nullptr) {
        if (currentIndex == index) return current;
        current = current->next;
        currentIndex++;
    }

    return nullptr;
}

bool deleteNodeIndexTwo(ForwardListTwo* flist, const string& num) {
    if (flist->head == nullptr) return false;

    DoublyNode* current = flist->head;

    while (current != nullptr && current->node != num) {
        current = current->next;
    }

    if (current == nullptr) return false;

    if (current->prev != nullptr) {
        current->prev->next = current->next;
    }
    else {
        flist->head = current->next;
    }

    if (current->next != nullptr) {
        current->next->prev = current->prev;
    }
    else {
        flist->tail = current->prev;
    }

    delete current;
    return true;
}