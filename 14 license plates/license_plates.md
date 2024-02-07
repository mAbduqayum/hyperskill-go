# Vehicle License Plates

## Problem Statement
In Berlyand, vehicle license plates consist of digits and uppercase letters of the Latin alphabet. There are two types of plates:

- Either the license plate is of the form `letter-digit-digit-letter-letter` (examples of correct first-type numbers: `R48FA`, `O00OO`, `A99OK`);
- Or the license plate is of the form `letter-digit-letter-letter` (examples of correct second-type numbers: `T7RR`, `A9PQ`, `O0OO`).

Thus, each vehicle license plate is a string of either the first or the second type.

You are given a string of digits and uppercase letters of the Latin alphabet. Is it possible to divide it with spaces into a sequence of correct vehicle license plates? In other words, check if the given string can be formed as a sequence of correct vehicle license plates, which are written consecutively without spaces. In case of a positive answer, output any such partition.

## Input
The first line contains an integer `t` (`1 ≤ t ≤ 1000`) — the number of sets of input data in the test.

The sets of input data in the test are independent. They do not affect each other.

Each set of input data is a non-empty string `s`, consisting of digits and uppercase letters of the Latin alphabet. The length of the string is from 1 to 50 characters.

## Output
Output `n` lines: the next line should contain the answer for the corresponding set of input data.

If the answer is negative — that is, the given string `s` cannot be presented as a sequence of car numbers — the line in the output should contain a single character `'-'` (minus, ASCII code 45).

In case of a positive answer, output any partition of the given string `s` into a sequence of correct numbers. Each number must correspond to one of the two types (see the condition). Separate the numbers with spaces. You can output an arbitrary number of spaces and even extra spaces after the last number.

## Test Example 1
### Input

```
6
R48FAO00OOO0OOA99OKA99OK
R48FAO00OOO0OOA99OKA99O
A9PQ
A9PQA
A99AAA99AAA99AAA99AA
AP9QA
```

### Output

```
R48FA O00OO O0OO A99OK A99OK
-
A9PQ
-
A99AA A99AA A99AA A99AA
-
```