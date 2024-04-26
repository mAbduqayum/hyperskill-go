# Problem Statement

Data compression is an algorithmic (usually reversible) transformation of data, produced with the purpose of reducing their occupied volume.

You are to implement a simple algorithm for compressing a sequence of numbers, which assumes that the sequence consists of subsequences of consecutive increasing or decreasing numbers.

Formally, we will represent the result of compression as a new sequence of even length in the form `b1, c1, b2, c2, ..., bk, ck`, where the pair of consecutive numbers `bi, ci` indicates that:
- if `ci ≥ 0`, then in the original sequence there was an increasing run of numbers `bi, bi + 1, ..., bi + ci`;
- if `ci < 0`, then in the original sequence there was a decreasing run of numbers `bi, bi - 1, ..., bi - |ci|`.

For example, if the original sequence was `[3, 2, 1, 0, -1, 0, 6, 6, 7]`, the result of compression could look like `[3, -4, 0, 0, 6, 0, 6, 1]`.

For a given sequence, output the result of compression that has the smallest length (the number of elements in the sequence-result). If there are several ways to compress the sequence to achieve the smallest number of elements, output any of them.

## Input Data

The first line of the input data contains an integer `t (1 ≤ t ≤ 100)` — the number of sets of input data.

The input data sets in the test are independent. They do not affect each other.

Each set of input data consists of two lines.

The first contains an integer `n (1 ≤ n ≤ 50)` — the length of the given sequence `a`. The second contains the sequence of integers `a1, a2, ..., an (-1000 ≤ ai ≤ 1000)` — the elements of the sequence `a`.

## Output Data

For each set of input data, output the answer in a pair of lines. The first line should contain the length of the result (an even positive number), and the second should contain the result of the optimal compression — a sequence of integers.

Note that during the compression process, the length of the sequence may increase.

If there are several ways to compress the given sequence so that the number of elements in the result is the smallest, choose any of them.

## Example Test 1

### Input Data

```
5
9
3 2 1 0 -1 0 6 6 7
1
1000
7
1 2 3 4 5 6 7
7
1 3 5 7 9 11 13
11
100 101 102 103 19 20 21 22 42 41 40
```


### Output Data

```
8
3 -4 0 0 6 0 6 1
2
1000 0
2
1 6
14
1 0 3 0 5 0 7 0 9 0 11 0 13 0
6
100 3 19 3 42 -2
```