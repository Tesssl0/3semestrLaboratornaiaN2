#include "fullBinaryTree.h"
#include <iostream>
#include <string>   
using namespace std;

// ======== СТРУКТУРА УЗЛА ОЧЕРЕДИ ========
struct QueueNode {
    Node* treeNode;
    QueueNode* next;

    QueueNode(Node* node) : treeNode(node), next(nullptr) {}
};

// ======== ПРОСТАЯ РЕАЛИЗАЦИЯ ОЧЕРЕДИ ========
struct SimpleQueue {
    QueueNode* front;
    QueueNode* rear;
    int size;

    SimpleQueue() : front(nullptr), rear(nullptr), size(0) {}

    ~SimpleQueue() {
        while (!isEmpty()) {
            dequeue();
        }
    }

    void enqueue(Node* node) {
        QueueNode* newNode = new QueueNode(node);
        if (rear == nullptr) {
            front = rear = newNode;
        }
        else {
            rear->next = newNode;
            rear = newNode;
        }
        size++;
    }

    Node* dequeue() {
        if (front == nullptr) return nullptr;

        QueueNode* temp = front;
        Node* node = temp->treeNode;
        front = front->next;

        if (front == nullptr) {
            rear = nullptr;
        }

        delete temp;
        size--;
        return node;
    }

    bool isEmpty() {
        return front == nullptr;
    }

    int getSize() {
        return size;
    }
};
// ======== ВЫЧИСЛЕНИЕ ВЫСОТЫ ДЕРЕВА ========
int getTreeHeight(Node* node) {
    if (node == nullptr) {
        return 0;
    }

    int leftHeight = getTreeHeight(node->left);
    int rightHeight = getTreeHeight(node->right);

    return 1 + (leftHeight > rightHeight ? leftHeight : rightHeight);
}
// ======== ВСТАВКА В БИНАРНОЕ ДЕРЕВО ПОИСКА ========
void insertBST(Node*& node, const string& value) {
    if (!node) {
        node = new Node(value);
        return;
    }

    // Предполагаем, что значения - числа
    int newValue = stoi(value);
    int nodeValue = stoi(node->data);

    if (newValue < nodeValue)
        insertBST(node->left, value);
    else if (newValue > nodeValue)
        insertBST(node->right, value);
    // Если значения равны, можно игнорировать или обработать дубликаты
}

// Обертка для работы с деревом
void insertBST(fullBinaryTree* tree, const string& value) {
    if (!tree) return;

    try {
        // Проверяем, что значение можно преобразовать в число
        stoi(value);
        insertBST(tree->root, value);
    }
    catch (const exception& e) {
        cout << "Ошибка: значение '" << value << "' не является числом" << endl;
    }
}

// ======== ПОЛУЧЕНИЕ ОБХОДА В ШИРИНУ (BFS) В ВИДЕ СТРОКИ ========
string getBFSAsString(fullBinaryTree* tree) {
    string result;
    if (!tree->root) return result;

    SimpleQueue q;
    q.enqueue(tree->root);

    while (!q.isEmpty()) {
        Node* temp = q.dequeue();
        result += temp->data + " ";

        if (temp->left) q.enqueue(temp->left);
        if (temp->right) q.enqueue(temp->right);
    }

    if (!result.empty()) result.pop_back();
    return result;
}

// ======== РЕКУРСИВНОЕ УДАЛЕНИЕ ВСЕХ УЗЛОВ ДЕРЕВА ========
void clearTree(Node* node) {
    if (node == nullptr) return;
    clearTree(node->left);
    clearTree(node->right);
    delete node;
}

// ======== ОЧИСТКА ВСЕГО ДЕРЕВА ========
void clearFullBinaryTree(fullBinaryTree* tree) {
    clearTree(tree->root);
    tree->root = nullptr;
}

// ======== ОБХОДЫ ДЕРЕВА ========
void inorder(Node* node) {
    if (!node) return;
    inorder(node->left);
    cout << node->data << " ";  
    inorder(node->right);
}

void preorder(Node* node) {
    if (!node) return;
    cout << node->data << " "; 
    preorder(node->left);
    preorder(node->right);
}

void postorder(Node* node) {
    if (!node) return;
    postorder(node->left);
    postorder(node->right);
    cout << node->data << " ";  // ИСПРАВЛЕНО: data вместо value
}

// ======== ПОИСК В ШИРИНУ (BFS) ========
Node* BFS(fullBinaryTree* tree, const string& value) {
    if (!tree->root) return nullptr;

    SimpleQueue q;
    q.enqueue(tree->root);

    while (!q.isEmpty()) {
        Node* temp = q.dequeue();
        if (temp->data == value) return temp;  // ИСПРАВЛЕНО: data вместо value

        if (temp->left) q.enqueue(temp->left);
        if (temp->right) q.enqueue(temp->right);
    }
    return nullptr;
}

// ======== ВЫВОД ВСЕГО ДЕРЕВА В ШИРИНУ ========
void printBFS(fullBinaryTree* tree) {
    if (!tree->root) {
        cout << "Дерево пустое" << endl;
        return;
    }

    SimpleQueue q;
    q.enqueue(tree->root);

    cout << "Обход в ширину: ";
    while (!q.isEmpty()) {
        Node* temp = q.dequeue();
        cout << temp->data << " ";

        if (temp->left) q.enqueue(temp->left);
        if (temp->right) q.enqueue(temp->right);
    }
    cout << endl;
}

// ======== ПРОВЕРКА НА ПОЛНОЕ БИНАРНОЕ ДЕРЕВО ========
bool isCompleteBinaryTree(Node* root) {
    if (!root) return true;

    SimpleQueue q;
    q.enqueue(root);
    bool foundNull = false;

    while (!q.isEmpty()) {
        Node* temp = q.dequeue();

        if (temp->left) {
            if (foundNull) return false;
            q.enqueue(temp->left);
        }
        else {
            foundNull = true;
        }

        if (temp->right) {
            if (foundNull) return false;
            q.enqueue(temp->right);
        }
        else {
            foundNull = true;
        }
    }
    return true;
}

// ======== ПРОВЕРКА НА СТРОГО ПОЛНОЕ ДЕРЕВО ========
bool isFullBinaryTree(Node* root) {
    if (!root) return true;
    if (!root->left && !root->right) return true;
    if (root->left && root->right)
        return isFullBinaryTree(root->left) && isFullBinaryTree(root->right);
    return false;
}

// ======== ПОЛУЧЕНИЕ ОБХОДА INORDER В ВИДЕ СТРОКИ ========
void getInorderHelper(Node* node, string& result) {
    if (!node) return;
    getInorderHelper(node->left, result);
    result += node->data + " ";  
    getInorderHelper(node->right, result);
}

string getInorderAsString(Node* node) {
    string result;
    getInorderHelper(node, result);
    if (!result.empty()) result.pop_back();
    return result;
}

// ======== АНАЛИЗ ДЕРЕВА (ТИП, ВЫСОТА, УЗЛЫ) ========
void checkTreeType(fullBinaryTree* tree) {
    if (!tree->root) {
        cout << "The tree is empty" << endl;
        return;
    }

    bool complete = isCompleteBinaryTree(tree->root);
    bool full = isFullBinaryTree(tree->root);

    cout << "=== TREE ANALYSIS ===" << endl;

    // Подсчет количества узлов
    int count = 0;
    SimpleQueue q;
    if (tree->root) q.enqueue(tree->root);
    while (!q.isEmpty()) {
        Node* temp = q.dequeue();
        count++;
        if (temp->left) q.enqueue(temp->left);
        if (temp->right) q.enqueue(temp->right);
    }
    cout << "Number of nodes: " << count << endl;

    // Вычисление высоты дерева (по уровням)
    int height = 0;
    if (tree->root) {
        SimpleQueue q2;
        q2.enqueue(tree->root);

        while (!q2.isEmpty()) {
            int levelSize = q2.size;
            height++;
            for (int i = 0; i < levelSize; i++) {
                Node* temp = q2.dequeue();
                if (temp->left) q2.enqueue(temp->left);
                if (temp->right) q2.enqueue(temp->right);
            }
        }
    }

    cout << "Tree height: " << height << endl;

    
    cout << "Tree type: ";
    if (full) {
        cout << "Full Binary Tree" << endl;
    }
    else {
        cout << "Binary Tree" << endl;
        cout << "- Does not meet the criteria for a full binary tree" << endl;
    }

    
    cout << "Maximum possible nodes for height " << height << ": " << ((1 << height) - 1) << endl;
    cout << "Actual number of nodes: " << count << endl;

    if (complete && count == ((1 << height) - 1)) {
        cout << "The tree is a PERFECT binary tree!" << endl;
    }
}

// ======== ЗМЕЙКА ========
void printZigzag(fullBinaryTree* tree) {
    if (!tree->root) {
        cout << "Дерево пустое" << endl;
        return;
    }

    SimpleQueue q;
    q.enqueue(tree->root);
    bool leftToRight = true; // Направление вывода

    cout << "Обход змейкой: ";

    while (!q.isEmpty()) {
        int levelSize = q.size;

        // Временный массив для хранения узлов текущего уровня
        string* levelNodes = new string[levelSize];

        // Обрабатываем все узлы текущего уровня
        for (int i = 0; i < levelSize; i++) {
            Node* temp = q.dequeue();

            // Сохраняем узел в массив в зависимости от направления
            if (leftToRight) {
                levelNodes[i] = temp->data;
            }
            else {
                levelNodes[levelSize - 1 - i] = temp->data;
            }

            // Добавляем дочерние узлы в очередь
            if (temp->left) q.enqueue(temp->left);
            if (temp->right) q.enqueue(temp->right);
        }

        // Выводим узлы текущего уровня
        for (int i = 0; i < levelSize; i++) {
            cout << levelNodes[i];
            if (i < levelSize - 1) cout << " - ";
        }

        // Меняем направление для следующего уровня
        leftToRight = !leftToRight;

        // Освобождаем память
        delete[] levelNodes;

        // Добавляем разделитель между уровнями (кроме последнего)
        if (!q.isEmpty()) cout << " - ";
    }
    cout << endl;
}