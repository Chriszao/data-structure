# Linked List

- ### **Linear list:**
  
  It is a data structure in which elements of the same data type are arranged sequentially. Not necessarily, these elements are physically in sequence. But every element must have exactly one successor and one predecessor, with the exception of the first element, which has no predecessor, and the last, which has no successor.

  The usability of each depends on the application, but in simpler terms, if we need speed when accessing an index of an element, sequential list is best, but if we need something that is easier to add and remove items, or the number of elements can increase or decrease during execution, regardless of the maximum capacity, so it is better to use linked list.

- ### **Sequential linear list:**
  
  In a linear sequential (or contiguous) list, the elements, in addition to being in a logical sequence, are also physically in sequence. The simplest way to accommodate a linear list on a computer is through the use of a vector.

  ![Linear sequential](/_notes/assets/linked1.png)
  Sequential linear list example


  - To access a value at a given index, the program takes the index and multiplies it by the amount of bytes of the value and manages to access the position in memory where the value is stored. For example, to access position three of a list of integers (each integer is equivalent to 4 bytes in memory), the calculation will be as follows: $3.4 = 12$

  - Benefits:
      - Easy access to each element (in constant time);
  - Disadvantages:
      - You need to allocate enough space for all elements at once;
      - May have unused allocated memory;
      - It can be difficult to add or remove elements and need to do
      displacements.
    ![list with empty spaces](/_notes/assets/linked2.png)
    
    Example of a linear list with unused spaces

- ### **Linked list:**
  
  - Linked list or linked list is a sparse linear list,that is, the elements are not necessarily stored sequentially in memory, but the logical order between the elements that make up the list is maintained.
  
    ![linked list example](/_notes/assets/linked3.png)
    Linked list example

  - Benefits:
    - Easy insertion and removal of elements.
    - Numbers of elements can increase or decrease during the
    execution;
  - Disadvantages:
    - The access in each element is not in constant time, but
    depends on the number of elements.
  
> These annotations are for dynamic stacks and dynamic queues.

> Linear list = vectors 

> Linked list = pointers