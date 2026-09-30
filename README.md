# DSA & Problem Solving in Go

Welcome to my Data Structures and Algorithms repository.

This repository is a collection of coding problems that I am solving to improve my problem-solving skills, strengthen my understanding of data structures and algorithms, and become better at writing clean and efficient Go code.

The main focus of this repository is not only to collect solutions, but also to understand **how and why** a particular solution works.

---

## About This Repository

I will be solving different coding problems step by step using Go.

For every problem, the goal is to document:

* Problem statement
* Examples
* Understanding of the problem
* Approach used to solve it
* Step-by-step explanation
* Complete Go solution
* Time complexity
* Space complexity
* Important Go concepts used
* Key problem-solving lessons

The explanations are written in simple, practical Indian English so that they are easy to understand for beginners and students.

The intention is to make this repository useful not only for me, but also for anyone who is learning Data Structures and Algorithms with Go.

---

## Repository Structure

The repository will follow a structure similar to this:

```text
dsa-go/
│
├── README.md
│
├── 001-two-sum/
│   ├── README.md
│   └── main.go
│
├── 002-valid-parentheses/
│   ├── README.md
│   └── main.go
│
├── 003-best-time-to-buy-and-sell-stock/
│   ├── README.md
│   └── main.go
│
├── 004-binary-search/
│   ├── README.md
│   └── main.go
│
└── ...
```

Each problem will have its own folder.

Inside each folder:

```text
README.md
main.go
```

The `README.md` explains the problem and solution, while `main.go` contains the complete Go implementation.

---

# Problems Solved

The list below will be updated as new problems are added.

| #   | Problem                   | Topic           | Difficulty |
| --- | ------------------------- | --------------- | ---------- |
| 001 | [Two Sum](./001-two-sum/) | Array, Hash Map | Easy       |
| 002 | Coming Soon               | -               | -          |
| 003 | Coming Soon               | -               | -          |

---

# How Each Problem Is Documented

Every problem will follow a similar format.

## 1. Problem Statement

First, the original problem will be explained in simple words.

For example:

> Given an array of integers and a target value, find two numbers whose sum is equal to the target.

---

## 2. Examples

Examples will be provided so that the problem becomes clear before looking at the code.

```text
Input:
nums = [2, 7, 11, 15]
target = 9

Output:
[0, 1]
```

---

## 3. Understanding the Problem

Before jumping directly into the code, the problem will be broken down into smaller parts.

The main goal is to understand:

* What is given?
* What is expected?
* What are the constraints?
* What edge cases should we consider?

---

## 4. Approach

The solution approach will be explained step by step.

If there are multiple possible approaches, they may be discussed separately.

For example:

```text
Brute Force
O(n²)

Optimized
O(n)
```

This helps understand not only the final solution but also why one approach can be better than another.

---

## 5. Complete Go Code

Every problem will contain a complete, runnable Go program whenever possible.

Example:

```go
package main

import "fmt"

func main() {
	fmt.Println("Solution")
}
```

The code will be kept simple and readable, especially for beginners.

---

## 6. Complexity Analysis

Every solution will include:

```text
Time Complexity: O(n)

Space Complexity: O(n)
```

The reason behind the complexity will also be explained.

Understanding complexity is an important part of learning DSA.

---

## 7. Go Concepts Used

When required, the problem README will also explain relevant Go concepts.

For example:

* Arrays
* Slices
* Maps
* Structs
* Pointers
* Interfaces
* Recursion
* Sorting
* `range`
* Input handling
* Functions

This is useful because the repository is not only for learning algorithms, but also for improving Go programming skills.

---

# Learning Approach

The main idea behind this repository is:

```text
Understand the problem
        ↓
Think about a simple solution
        ↓
Check its complexity
        ↓
Look for a better approach
        ↓
Implement in Go
        ↓
Test with examples
        ↓
Understand the complexity
        ↓
Document what was learned
```

I am trying to avoid simply copying solutions.

The focus is on understanding the reasoning behind the solution.

---

# Topics Covered

As the repository grows, problems will cover topics such as:

### Data Structures

* Arrays
* Slices
* Strings
* Hash Maps
* Hash Sets
* Stack
* Queue
* Linked List
* Trees
* Binary Search Tree
* Heap
* Graphs
* Trie

### Algorithms

* Searching
* Sorting
* Two Pointers
* Sliding Window
* Binary Search
* Recursion
* Backtracking
* Greedy
* Dynamic Programming
* Graph Traversal
* Breadth First Search
* Depth First Search

### Problem-Solving Patterns

* Frequency Counting
* Hash Map Lookup
* Prefix Sum
* Two Pointer Technique
* Sliding Window
* Fast and Slow Pointer
* Divide and Conquer
* Backtracking
* Dynamic Programming

---

# Goals

The main goals of this repository are:

1. Improve problem-solving skills.
2. Build a strong foundation in Data Structures and Algorithms.
3. Learn how to analyze time and space complexity.
4. Improve Go programming skills.
5. Prepare for technical interviews.
6. Build consistency by solving problems regularly.
7. Maintain notes that can be revised later.
8. Help other beginners understand the same problems.

---

# How to Run the Solutions

Make sure Go is installed:

```bash
go version
```

Clone the repository:

```bash
git clone <repository-url>
```

Move into the repository:

```bash
cd dsa-go
```

Go to any problem directory:

```bash
cd 001-two-sum
```

Run the program:

```bash
go run main.go
```

---

# Running Tests

As the repository grows, test cases may be added for individual problems.

For Go test files:

```bash
go test
```

To run tests with more information:

```bash
go test -v
```

---

# Code Style

The code in this repository will try to follow normal Go conventions.

The focus will be on:

* Readable variable names
* Simple functions
* Proper formatting
* Avoiding unnecessary complexity
* Writing understandable code
* Using appropriate data structures
* Explaining important logic

Go code will be formatted using:

```bash
gofmt -w main.go
```

---

# Difficulty Levels

Problems may be categorized as:

| Level  | Meaning                                                |
| ------ | ------------------------------------------------------ |
| Easy   | Good for beginners and basic concepts                  |
| Medium | Requires multiple concepts or stronger problem-solving |
| Hard   | Requires advanced algorithms or deeper understanding   |

Difficulty is mainly used as a reference for organizing the problems.

---

# Progress

This repository is an ongoing learning project.

The number of problems will increase over time.

```text
Problems Solved: 1
Language: Go
Focus: Data Structures & Algorithms
```

This section can be updated as new problems are solved.

---

# Why Go?

Go is being used for this repository because it is:

* Simple to read
* Fast
* Statically typed
* Widely used for backend and systems development
* Good for understanding fundamental programming concepts
* Useful for learning how data structures are implemented and used

Learning DSA in Go also helps improve general programming and problem-solving skills.

---

# Notes for Students

If you are using this repository while learning DSA, don't directly jump to the code.

A better approach is:

```text
1. Read the problem
2. Try solving it yourself
3. Think about the brute-force solution
4. Check the constraints
5. Look for a better approach
6. Compare your solution with the documented solution
7. Understand the complexity
8. Implement it yourself
9. Try different test cases
```

The goal is not to solve the maximum number of problems.

The goal is to understand the maximum amount from each problem.

---

# Contribution

This is primarily a personal learning repository, but suggestions and improvements are welcome.

If you find:

* A bug
* A better explanation
* A better approach
* An edge case
* A mistake in complexity analysis
* A Go-specific improvement

feel free to open an issue or submit a pull request.

---

# Progress Over Perfection

DSA becomes easier with consistent practice.

Some problems may be easy, while others may take more time to understand.

It is completely fine to struggle with a problem.

The important part is to:

```text
Try
Understand
Implement
Make mistakes
Learn
Improve
Repeat
```

This repository will grow gradually as more problems are solved.

---


