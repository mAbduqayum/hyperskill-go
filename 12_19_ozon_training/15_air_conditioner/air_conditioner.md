# Battle for the Air Conditioner

## Problem Statement

In an office, there's an air conditioner that can be set to a temperature ranging from 15 to 30 degrees.

Employees enter the office one after another, each with their own temperature preference. The i-th of them wishes for a
temperature no higher or no lower than a certain value.

After the arrival of each employee, determine if it's possible to set a temperature that satisfies everyone in the
office.

## Input

Each test consists of several sets of input data. The first line contains an integer t (1 ≤ t ≤ 10^3) — the number of
data sets. The description of the data sets follows.

The first line of each set contains an integer n (1 ≤ n ≤ 10^3) — the number of employees.

The i-th of the following n lines of each data set contains the temperature requirement from the i-th employee: either ≥
a_i or ≤ a_i (15 ≤ a_i ≤ 30, a_i — an integer). The requirement ≥ a_i means that the i-th employee wishes for a
temperature no lower than a_i; the requirement ≤ a_i means that the i-th employee wishes for a temperature no higher
than a_i.

It is guaranteed that the sum of n across all data sets does not exceed 10^3.

## Output

For each set of input data, output n lines, the i-th of which contains a temperature that satisfies all employees from 1
to i inclusive. If such a temperature does not exist, output -1. After outputting the answer for each data set, output
an empty line.

If there are multiple answers, output any of them.

### Explanation for the First Example:

1. The requirement ≥ 30 is added, the possible temperature range is [30, 30], so the only possible answer is 30 degrees.

### Explanation for the Second Example:

1. The requirement ≥ 18 is added, the possible temperature range is [18, 30], so 29 degrees is taken as an example;
2. The requirement ≤ 23 is added, the possible temperature range is [18, 23], so 19 degrees is taken as an example;
3. The requirement ≥ 20 is added, the possible temperature range is [20, 23], so 22 degrees is taken as an example;
4. The requirement ≤ 27 is added, the possible temperature range remains [20, 23], so 21 degrees is taken as an example;
5. The requirement ≤ 21 is added, the possible temperature range is [20, 21], so 20 degrees is taken as an example;
6. The requirement ≥ 28 is added, the possible temperature range becomes [28, 21], so there is no answer, and -1 should
   be output.

### Explanation for the Third Example:

1. The requirement ≤ 25 is added, the possible temperature range is [15, 25], so 23 degrees is taken as an example;
2. The requirement ≥ 20 is added, the possible temperature range is [20, 25], so 22 degrees is taken as an example;
3. The requirement ≥ 25 is added, the possible temperature range is [25, 25], so the only possible temperature is 25
   degrees.

### Explanation for the Fourth Example:

1. The requirement ≤ 15 is added, the possible temperature range is [15, 15], so the only possible temperature is 15
   degrees;
2. The requirement ≥ 30 is added, the possible temperature range becomes [30, 15], so there is no answer, and -1 should
   be output;
3. The requirement ≤ 24 is added, the possible temperature range remains [30, 15], so there is no answer, and -1 should
   be output.

## Test Example 1

### Input

```
4
1
>= 30
6
>= 18
<= 23
>= 20
<= 27
<= 21
>= 28
3
<= 25
>= 20
>= 25
3
<= 15
>= 30
<= 24
```

### Output

```
30

29
19
22
21
20
-1

23
22
25

15
-1
-1
```