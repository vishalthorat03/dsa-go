# Two Sum — Go Solution

This repository contains the solution for the **Two Sum** problem using **Go (Golang)**.

This is one of the most common beginner problems for learning:

* Arrays
* Hash Maps
* Searching
* Time and Space Complexity
* Problem-solving approach

---

##  Problem Statement

You are given an array of integers `nums` and an integer `target`.

You need to find **two different numbers** in the array whose sum is equal to the given `target`.

Return the **indices** of those two numbers.

### Important Points

* We cannot use the same element twice.
* There will be **exactly one valid answer**.
* The answer can be returned in any order.

---

##  Example 1

### Input

```text
nums = [2, 7, 11, 15]
target = 9
```

We can see:

```text
2 + 7 = 9
```

`2` is present at index `0`.

`7` is present at index `1`.

### Output

```text
[0, 1]
```

---

##  Example 2

### Input

```text
nums = [3, 2, 4]
target = 6
```

Here:

```text
2 + 4 = 6
```

Index of `2` = `1`

Index of `4` = `2`

### Output

```text
[1, 2]
```

---

##  Example 3

### Input

```text
nums = [3, 3]
target = 6
```

Here:

```text
3 + 3 = 6
```

The first `3` is at index `0`.

The second `3` is at index `1`.

### Output

```text
[0, 1]
```

---

#  How to Think About the Problem

Let's first think about the simple approach.

Suppose:

```text
nums = [2, 7, 11, 15]
target = 9
```

For every number, we need to find another number that makes the target.

For `2`:

```text
9 - 2 = 7
```

So we need `7`.

Instead of searching the whole array again and again, we can store the numbers that we have already seen in a **Hash Map**.

This makes the solution much faster.

---

#  Optimized Approach — Hash Map

We create a map:

```go
seen := make(map[int]int)
```

The map stores:

```text
number -> index
```

For example:

```text
2 -> 0
7 -> 1
```

Now, while traversing the array, we calculate:

```text
complement = target - current number
```

Then we check whether this complement already exists in our map.

---

#  Step-by-Step Example

Consider:

```text
nums = [2, 7, 11, 15]
target = 9
```

### Step 1

Current number:

```text
2
```

Calculate:

```text
9 - 2 = 7
```

We need `7`.

But `7` is not in the map yet.

So we store:

```text
2 -> 0
```

---

### Step 2

Current number:

```text
7
```

Calculate:

```text
9 - 7 = 2
```

Now we check the map.

We already have:

```text
2 -> 0
```

So we found the answer.

Current index is `1`.

Therefore:

```text
[0, 1]
```

---

#  Why Do We Check Before Inserting?

This is an important point.

We do:

```go
if j, ok := seen[complement]; ok {
    return []int{j, i}
}

seen[num] = i
```

We check the map **before** adding the current number.

This ensures that we don't accidentally use the same array element twice.

For example:

```text
nums = [3, 3]
target = 6
```

First `3`:

```text
complement = 6 - 3
            = 3
```

`3` is not in the map.

So we store:

```text
3 -> 0
```

Second `3`:

```text
complement = 6 - 3
            = 3
```

Now `3` already exists in the map.

So we return:

```text
[0, 1]
```

---

#  Complete Go Code

```go
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}

		seen[num] = i
	}

	return []int{}
}

func main() {
	var n int

	fmt.Print("Enter number of elements: ")
	fmt.Scan(&n)

	nums := make([]int, n)

	fmt.Println("Enter the elements:")

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	var target int

	fmt.Print("Enter target: ")
	fmt.Scan(&target)

	result := twoSum(nums, target)

	fmt.Println("Answer:", result)
}
```

---

#  How to Run

First, make sure Go is installed.

Check using:

```bash
go version
```

Create a file:

```text
main.go
```

Put the above code inside it.

Then run:

```bash
go run main.go
```

---

#  Sample Input

```text
Enter number of elements: 4
Enter the elements:
2 7 11 15
Enter target: 9
```

#  Sample Output

```text
Answer: [0 1]
```

---

#  Time Complexity

### Brute Force Approach

The simple approach uses two loops:

```text
O(n²)
```

Why?

For every element, we may have to check almost every other element.

For example:

```text
10,000 elements
```

can result in a very large number of comparisons.

---

### Hash Map Approach

Our solution takes:

```text
O(n)
```

on average.

We traverse the array only once.

Map lookup is approximately:

```text
O(1)
```

So overall:

```text
O(n)
```

---

#  Space Complexity

We are using a Hash Map:

```go
seen := make(map[int]int)
```

In the worst case, we may store almost every element.

Therefore:

```text
O(n)
```

space complexity.

---

#  Complexity Summary

| Approach    |         Time | Space |
| ----------- | -----------: | ----: |
| Brute Force |        O(n²) |  O(1) |
| Hash Map    | O(n) average |  O(n) |

For this problem, the **Hash Map approach** satisfies the follow-up requirement of less than `O(n²)` time.

---

#  Important Go Concepts Used

## 1. Slice

```go
nums []int
```

A slice is used to store the array of integers.

Example:

```go
nums := []int{2, 7, 11, 15}
```

---

## 2. Map

```go
seen := make(map[int]int)
```

This creates a map where:

```text
key   = number
value = index
```

Example:

```text
2 -> 0
7 -> 1
```

---

## 3. `range`

We use:

```go
for i, num := range nums
```

Here:

* `i` = index
* `num` = value

For:

```text
nums = [2, 7, 11]
```

we get:

```text
i = 0, num = 2
i = 1, num = 7
i = 2, num = 11
```

---

## 4. Checking a Map

This line:

```go
if j, ok := seen[complement]; ok
```

checks whether `complement` exists in the map.

`j` contains the stored index.

`ok` tells us whether the key exists.

For example:

```go
j, ok := seen[2]
```

If `2` exists:

```text
j  = its index
ok = true
```

If it doesn't exist:

```text
ok = false
```

---

#  General Problem-Solving Pattern

The most important thing to learn from this problem is not only the code.

The important pattern is:

```text
Current Value
      ↓
What value do I need?
      ↓
target - current value
      ↓
Have I already seen it?
      ↓
     Yes → Answer
      ↓
      No
      ↓
Store current value
```

This type of thinking is useful in many coding interview problems.

Whenever you see a problem like:

> "Find two numbers that satisfy some condition"

you should think about whether a **Hash Map / Hash Set** can help reduce the search time.

---

#  What I Learned From This Problem

After solving Two Sum, you should understand:

* How arrays/slices work in Go
* How to take input using `fmt.Scan`
* How Go maps work
* How to store a value with its index
* How to find a complement
* Why Hash Maps can reduce `O(n²)` to `O(n)`
* Difference between time and space complexity
* Why checking before inserting prevents using the same element twice

---

#  LeetCode

This problem is commonly known as **Two Sum** on LeetCode.

Problem number:

```text
#1
```

It is a good beginner problem for understanding Hash Maps and algorithmic thinking.

---

#  Quick Revision

Remember this:

```go
complement := target - num

if j, ok := seen[complement]; ok {
	return []int{j, i}
}

seen[num] = i
```

In simple words:

> "For the current number, calculate what number is required to reach the target. If I have already seen that number, return both indexes. Otherwise, remember the current number."

---

##  Final Takeaway

The brute-force solution is easy to understand, but it takes:

```text
O(n²)
```

The Hash Map solution is slightly more advanced but reduces the average time complexity to:

```text
O(n)
```

So this problem is a very good example of how choosing the right **data structure** can make an algorithm much faster.

**Don't just memorize the code. Understand the thought process:**

```text
target - current number = required number
```

Once this becomes natural, many Hash Map problems become easier to solve.
