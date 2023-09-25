#include <stddef.h>
typedef struct ValueListNode_s
{
    int key;
    int value;
    int counter;
    struct ValueListNode_s *next;
    struct ValueListNode_s *prev;
} ValueListNode;

typedef struct CounterListNode_s
{
    ValueListNode *head;
    struct CounterListNode_s *next;
    struct CounterListNode_s *prev;
} CounterListNode;

typedef struct
{
    int capacity;
    int currentCounter;
    CounterListNode *head;
    CounterListNode **counterHash;
    ValueListNode **keyHash;
} LFUCache;

extern void removeValueNode(CounterListNode *counterNode, ValueListNode *valueNode);
extern void insertValueNode(CounterListNode *counterNode, ValueListNode *valueNode);
extern void removeCounterNode(LFUCache *obj, CounterListNode *counterNode);
extern void insertCounterNode(LFUCache *obj, CounterListNode *counterPrev, CounterListNode *counterNode);

LFUCache *lFUCacheCreate(int capacity)
{
    LFUCache *obj = (LFUCache *)malloc(sizeof(LFUCache));
    obj->capacity = capacity;
    obj->currentCounter = 0;
    obj->keyHash = (ValueListNode **)calloc(100001, sizeof(CounterListNode));
    obj->counterHash = (CounterListNode **)calloc(200001, sizeof(CounterListNode));
    obj->head = NULL;
    return obj;
}

int lFUCacheGet(LFUCache *obj, int key)
{
    int value = -1;
    ValueListNode *valueNode = obj->keyHash[key];
    CounterListNode *counterNode = NULL, *counterNew = NULL;
    if (valueNode != NULL)
    {
        value = valueNode->value;
        counterNode = obj->counterHash[valueNode->counter];
        valueNode->counter++;
        counterNew = obj->counterHash[valueNode->counter];
        removeValueNode(counterNode, valueNode);
        if (counterNew == NULL)
        {
            counterNew = (CounterListNode *)calloc(1, sizeof(CounterListNode));
            obj->counterHash[valueNode->counter] = counterNew;
            insertCounterNode(obj, counterNode, counterNew);
        }
        if (counterNode->head == NULL)
        {
            removeCounterNode(obj, counterNode);
            free(counterNode);
            obj->counterHash[valueNode->counter - 1] = NULL;
        }
        insertValueNode(counterNew, valueNode);
    }
    return value;
}

void lFUCachePut(LFUCache *obj, int key, int value)
{
    int keyRemove = 0;
    ValueListNode *valueNode = obj->keyHash[key], *valueRemove = NULL;
    CounterListNode *counterNode = NULL, *counterNew = NULL;
    if (obj->capacity == 0)
    {
        return;
    }
    if (valueNode != NULL)
    {
        valueNode->value = value;
        counterNode = obj->counterHash[valueNode->counter];
        valueNode->counter++;
        counterNew = obj->counterHash[valueNode->counter];
        removeValueNode(counterNode, valueNode);
        if (counterNew == NULL)
        {
            counterNew = (CounterListNode *)calloc(1, sizeof(CounterListNode));
            obj->counterHash[valueNode->counter] = counterNew;
            insertCounterNode(obj, counterNode, counterNew);
        }
        if (counterNode->head == NULL)
        {
            removeCounterNode(obj, counterNode);
            free(counterNode);
            obj->counterHash[valueNode->counter - 1] = NULL;
        }

        insertValueNode(counterNew, valueNode);
    }
    else
    {
        if (obj->capacity > obj->currentCounter)
        {
            obj->currentCounter++;
        }
        else
        {
            counterNode = obj->head;
            valueRemove = counterNode->head->prev;
            keyRemove = valueRemove->key;
            removeValueNode(counterNode, valueRemove);

            if (counterNode->head == NULL)
            {
                removeCounterNode(obj, counterNode);
                free(counterNode);
                obj->counterHash[valueRemove->counter] = NULL;
            }

            free(valueRemove);
            obj->keyHash[keyRemove] = NULL;
        }

        valueNode = (ValueListNode *)calloc(1, sizeof(ValueListNode));
        valueNode->key = key;
        valueNode->value = value;
        valueNode->counter = 1;
        obj->keyHash[key] = valueNode;
        counterNew = obj->counterHash[1];
        if (counterNew == NULL)
        {
            counterNew = (CounterListNode *)calloc(1, sizeof(CounterListNode));
            obj->counterHash[1] = counterNew;
            insertCounterNode(obj, NULL, counterNew);
        }

        insertValueNode(counterNew, valueNode);
    }

    return;
}

void lFUCacheFree(LFUCache *obj)
{
    CounterListNode *counterNode = obj->head, *counterNext = NULL;
    ValueListNode *valueNode = NULL, *valueNext = NULL;
    while (counterNext != NULL)
    {
        counterNext = counterNode->next;
        valueNode = counterNode->head;
        do
        {
            valueNext = valueNode->next;
            free(valueNode);
            valueNode = valueNext;
        } while (counterNode->head != NULL);
        free(counterNode);
        counterNode = counterNext;
    }
    free(obj->keyHash);
    free(obj->counterHash);
    free(obj);
    return;
}

void removeValueNode(CounterListNode *counterNode, ValueListNode *valueNode)
{
    if (valueNode->next == valueNode)
    {
        counterNode->head = NULL;
    }
    else
    {
        valueNode->prev->next = valueNode->next;
        valueNode->next->prev = valueNode->prev;
        if (counterNode->head == valueNode)
        {
            counterNode->head = valueNode->next;
        }
    }

    return;
}

void insertValueNode(CounterListNode *counterNode, ValueListNode *valueNode)
{
    ValueListNode *tail = NULL;
    if (counterNode->head == NULL)
    {
        valueNode->prev = valueNode;
        valueNode->next = valueNode;
    }
    else
    {
        tail = counterNode->head->prev;
        valueNode->prev = tail;
        valueNode->next = counterNode->head;
        counterNode->head->prev = valueNode;
        tail->next = valueNode;
    }
    counterNode->head = valueNode;
    return;
}

void removeCounterNode(LFUCache *obj, CounterListNode *counterNode)
{
    if (obj->head == counterNode)
    {
        obj->head = counterNode->next;
        if (NULL != obj->head)
        {
            obj->head->prev = NULL;
        }
    }
    else
    {
        counterNode->prev->next = counterNode->next;
        if (NULL != counterNode->next)
        {
            counterNode->next->prev = counterNode->prev;
        }
    }
    return;
}

void insertCounterNode(LFUCache *obj, CounterListNode *counterPrev, CounterListNode *counterNode)
{
    if (NULL == counterPrev)
    {
        counterNode->prev = NULL;
        counterNode->next = obj->head;
        if (NULL != obj->head)
        {
            obj->head->prev = counterNode;
        }
        obj->head = counterNode;
    }
    else
    {
        counterNode->prev = counterPrev;
        counterNode->next = counterPrev->next;
        if (NULL != counterPrev->next)
        {
            counterPrev->next->prev = counterNode;
        }
        counterPrev->next = counterNode;
    }
    return;
}