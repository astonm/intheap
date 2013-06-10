intheap
=======
Examples of int min-heaps implemented via container/heap and my generic/heap.

Advantages of generic/heap:
 - Less code to define the heap
 - No need to actually understand how to implement the heap (or the backing store)
 - Operations are attached to the heap, not calls with the heap as the first parameter.
 - Inputs and outputs of operations are statically type-checked.
 
Disadvantages of generic/heap:
 - Somewhere between a 5x and 50x slowdown relative to the interface{}-based implementation. (benchmark comparison_test.go for details)
