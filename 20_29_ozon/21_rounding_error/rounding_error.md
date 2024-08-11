# Rounding Error

## Problem Statement

A marketplace operator earns a commission of `p%` on each sale of goods. If a product that costs `ai` rubles was sold, the commission should be `ai * p / 100` rubles. Due to a technical requirement, the commission is supposed to be rounded down to the nearest whole number of kopecks (i.e., to two decimal places). However, due to a programmer's error, the commission was always rounded down to the nearest whole number of rubles, which means the kopecks were always discarded.

Find the total amount of money lost due to this error after the sale of goods with a cost of `ai` rubles by the marketplace operator.

## Input

Each test consists of several sets of input data.

The first line contains an integer `t (1 ≤ t ≤ 10^5)` — the number of sets of input data.

This is followed by the description of the sets of input data.

The first line of each set of input data contains two integers `n` and `p (1 ≤ n ≤ 10^5, 1 ≤ p < 100)` — the number of sold goods and the percentage of the marketplace's commission.

The following `n` lines of each set contain one integer `ai (1 ≤ ai ≤ 10^9)` — the costs of the sold goods in rubles.

It is guaranteed that the sum of `n` over all sets of input data does not exceed `10^5`.

## Output

For each set of input data, output your answer on a separate line — the sum in rubles, lost by the marketplace due to the error. This number should contain exactly two decimal places (the number of kopecks).

## Comments to the first example:

In the first set of input data, the marketplace should have received 1, 2, 3, 4, and 5 kopecks of commission, but due to the error, 0 kopecks were received for each product.

In the second set of input data, the marketplace should have received:

- For the first product 50 * `5 / 100` = 2.5 rubles of commission, due to the error 50 kopecks were lost.
- For the second product 1 * `5 / 100` = 0.05 rubles of commission, due to the error 5 kopecks were lost.

In total, 50 + 5 kopecks were lost, which is 0.55 rubles.

## Test Example 1

### Input Data

```
3
5 1
1
2
3
4
5
1 5
40
2 99
50
1
```

### Output Data

```
0.15
0.00
1.49
```

## Test Example 2

### Input Data

```
10
3 53
639400267
173476987
342345118
1 49
536319315
2 95
677700176
240772006
3 45
740699623
683869018
209846795
4 95
143260951
482687878
124413839
971541706
5 75
818780933
777335599
240588375
882121733
614724075
4 78
984757266
436156900
653946032
333706697
4 1
93379217
973328560
118456146
720364705
3 10
672163500
657084180
902366997
1 76
461539143
```

### Output Data

```
1.16
0.35
0.90
1.20
1.30
2.25
2.10
1.28
0.70
0.68
```
