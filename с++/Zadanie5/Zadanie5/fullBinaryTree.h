#pragma once
#ifndef FULLBINARYTREE_H
#define FULLBINARYTREE_H
#include <iostream>
#include <string>

using namespace std;

// Структура узла дерева
struct Node {
    string data;  // ИСПРАВЛЕНО: было value, должно быть data
    Node* left;
    Node* right;

    Node(const string& val) : data(val), left(nullptr), right(nullptr) {}  // ИСПРАВЛЕНО: data вместо value
};

// Структура бинарного дерева
struct fullBinaryTree {
    Node* root = nullptr;

    // Добавляем конструктор по умолчанию
    fullBinaryTree() : root(nullptr) {}
};

// Основные операции с деревом
void insertBST(fullBinaryTree* tree, const string& value);
Node* BFS(fullBinaryTree* tree, const string& value);
void printZigzag(fullBinaryTree* tree);
// Обходы дерева
void inorder(Node* node);
void preorder(Node* node);
void postorder(Node* node);
void printBFS(fullBinaryTree* tree);
int getTreeHeight(Node* node);
// Очистка дерева
void clearTree(Node* node);
void clearFullBinaryTree(fullBinaryTree* tree);

// Проверка типов деревьев
bool isCompleteBinaryTree(Node* root);
bool isFullBinaryTree(Node* root);
void checkTreeType(fullBinaryTree* tree);
string getBFSAsString(fullBinaryTree* tree);
string getInorderAsString(Node* node);

#endif