# Go Collection Library

The **Go Collection Library** is a robust set of data structures implemented in Go. It is designed to provide developers with efficient and reusable components for building complex applications.

## Modules

### 1. **Heap**
   - **Description**: A generic heap implementation.
   - **Features**:
     - Min-Heap or Max-Heap functionality based on a customizable comparator.
     - Operations: Push, Pop, Size.
   - **Usage**:
     ```go
     heap := NewHeap(func(a, b int) bool { return a < b }) // Min-Heap
     heap.Push(10)
     heap.Push(5)
     top := heap.Pop() // Returns 5
     ```

### 2. **List**
   - **Description**: A doubly linked list implementation.
   - **Features**:
     - Operations: Push, Append, PopFront, PopBack, InsertAtPosition, IterateForward, IterateBackward.
   - **Usage**:
     ```go
     list := NewList[int]()
     list.Push(1)
     list.Append(2)
     front := list.Front().Element() // Returns 1
     ```

### 3. **Node**
   - **Description**: Node abstraction for use in data structures like lists.
   - **Features**:
     - Encapsulates element data with next and previous pointers.

### 4. **Queue**
   - **Description**: A generic queue implementation.
   - **Features**:
     - FIFO (First In First Out) behavior.
     - Operations: Enqueue, Dequeue, Peek.
   - **Usage**:
     ```go
     queue := NewQueue[int]()
     queue.Enqueue(1)
     queue.Enqueue(2)
     front := queue.Dequeue() // Returns 1
     ```

### 5. **Set**
   - **Description**: A generic set implementation.
   - **Features**:
     - Operations: Add, Remove, Contains, Size.
   - **Usage**:
     ```go
     set := NewSet[int]()
     set.Add(1)
     set.Add(2)
     exists := set.Contains(1) // Returns true
     ```

### 6. **Stack**
   - **Description**: A generic stack implementation.
   - **Features**:
     - LIFO (Last In First Out) behavior.
     - Operations: Push, Pop, Top, IsEmpty.
   - **Usage**:
     ```go
     stack := NewStack[int]()
     stack.Push(10)
     top := stack.Top() // Returns 10
     ```

---

## Contribution Guide

We welcome contributions to enhance the Go Collection Library. Please follow the steps below:

### 1. Fork the Repository
   - Fork this repository to your GitHub account and clone it locally.

### 2. Set Up the Environment
   - Install [Go](https://golang.org/dl/) (ensure version >= 1.19).
   - Run `go mod tidy` to install dependencies.

### 3. Code Standards
   - Adhere to the Go style guide and naming conventions.
   - Use appropriate comments for exported functions.

### 4. Testing
   - Ensure all new code is covered by unit tests.
   - Run tests using `go test ./... -cover` to verify functionality and maintain code coverage.

### 5. Submit a Pull Request
   - Push your changes to your fork and create a pull request to the main repository.
   - Include a detailed description of your changes and the problem they solve.

---

## License

This project is licensed under the terms of the [MIT License](LICENSE).

## Contact

For any questions or suggestions, feel free to raise an issue or reach out.
