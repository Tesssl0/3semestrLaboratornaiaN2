#ifndef LRUCACHE_H
#define LRUCACHE_H

#include "hash_table.h"
#include "listTwo.h"
#include <string>
#include<iostream>

class LRUCache {
private:
    int capacity;                    // максимальная емкость кэша
    ForwardListTwo usageList;        // список использования (от недавнего к старому)
    ChainingHashTable* cacheMap;     // хеш-таблица для быстрого доступа

    // Перемещение элемента в начало списка (как недавно использованного)
    void moveToFront(int key) {
        string keyStr = to_string(key);
        DoublyNode* current = usageList.head;
        // Ищем узел с данным ключом
        while (current != nullptr) {
            if (current->node.find(keyStr + ":") == 0) {  // проверяем начало строки
                string nodeData = current->node;
                deleteNodeIndexTwo(&usageList, nodeData);  // удаляем узел
                addNodeHeadTwo(&usageList, nodeData);      // добавляем в начало
                break;
            }
            current = current->next;
        }
    }

    // Удаление наименее recently used элемента (последнего в списке)
    void removeLRU() {
        if (usageList.tail != nullptr) {
            string lruData = usageList.tail->node;  // данные последнего элемента
            size_t colonPos = lruData.find(':');
            if (colonPos != std::string::npos) {
                int key = stoi(lruData.substr(0, colonPos));  // извлекаем ключ
                cacheMap->remove(key);  // удаляем из хеш-таблицы
            }
            deleteNodeTailTwo(&usageList);  // удаляем из списка
        }
    }

public:
    // Конструктор: инициализирует кэш заданной емкостью
    LRUCache(int cap) : capacity(cap) {
        usageList.head = nullptr;
        usageList.tail = nullptr;
        cacheMap = new ChainingHashTable(100);  // создаем хеш-таблицу
    }

    // Деструктор: очищает память
    ~LRUCache() {
        while (usageList.head != nullptr) {
            deleteNodeHeadTwo(&usageList);  // удаляем все узлы списка
        }
        delete cacheMap;  // удаляем хеш-таблицу
    }

    // Установка значения в кэше
    void set(int key, int value) {
        string keyStr = to_string(key);
        string itemStr = keyStr + ":" + to_string(value);  // формат: "key:value"

        // Если ключ уже существует в кэше
        if (cacheMap->search(key)) {
            moveToFront(key);  // перемещаем в начало
            // Обновляем значение в списке
            DoublyNode* current = usageList.head;
            while (current != nullptr) {
                if (current->node.find(keyStr + ":") == 0) {
                    current->node = itemStr;  // обновляем данные
                    break;
                }
                current = current->next;
            }
        }
        else {
            // Если кэш заполнен, удаляем LRU элемент
            if (countNodesTwo(usageList) >= capacity) {
                removeLRU();
            }
            cacheMap->insert(key);              // добавляем ключ в хеш-таблицу
            addNodeHeadTwo(&usageList, itemStr); // добавляем элемент в начало списка
        }
    }

    // Получение значения из кэша
    int get(int key) {
        if (cacheMap->search(key)) {  // проверяем наличие ключа
            moveToFront(key);         // перемещаем элемент в начало
            string keyStr = to_string(key);
            DoublyNode* current = usageList.head;
            // Ищем элемент в списке и извлекаем значение
            while (current != nullptr) {
                if (current->node.find(keyStr + ":") == 0) {
                    size_t colonPos = current->node.find(':');
                    if (colonPos != std::string::npos) {
                        return stoi(current->node.substr(colonPos + 1));  // возвращаем значение
                    }
                }
                current = current->next;
            }
        }
        return -1;  // ключ не найден
    }

    // Печать текущего состояния кэша
    void printCache() {
        cout << "LRU Cache: ";
        DoublyNode* current = usageList.head;
        // Выводим все элементы от недавних к старым
        while (current != nullptr) {
            cout << "[" << current->node << "] ";
            current = current->next;
        }
        cout << endl;
    }
};

// Прототип функции для обработки запросов
void processLRUQueries();

#endif