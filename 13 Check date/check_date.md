# Date Verification

## Problem Statement

A date is given in the format "day month year" as three integers. It is guaranteed that:

- `day` is an integer from 1 to 31;
- `month` is an integer from 1 to 12;
- `year` is an integer from 1950 to 2300.

Verify that the given three numbers correspond to a correct date (in the modern Gregorian calendar).

Remember, according to the modern calendar, a year is considered a leap year if at least one of the following statements is true:

- divisible by 4, but not divisible by 100;
- divisible by 400.

For example, the years 2012 and 2000 are leap years, but the years 1999, 2022, and 2100 are not.


## Input
The first line contains an integer `t` (1 ≤ `t` ≤ 1000) — the number of sets of input data in the test.

The sets of input data in the test are independent. They do not affect each other.

Each set of input data is specified by one line, which contains three integers `d`, `m`, `y` (1 ≤ `d` ≤ 31, 1 ≤ `m` ≤ 12, 1950 ≤ `y` ≤ 2300) — the day, month, and year of the date to be verified.
## Output
For each set of input data, output `YES` if the corresponding date is correct (i.e., such a date exists in the modern calendar). Output `NO` otherwise.

You can output the answer in any case (for example, output `yEs`, `yes`, `Yes`, and `YES` will still be considered correct).
## Example Test 1

### Input
```
10
10 9 2022
21 9 2022
29 2 2022
31 2 2022
29 2 2000
29 2 2100
31 11 1999
31 12 1999
29 2 2024
29 2 2023
```

### Output
```
YES
YES
NO
NO
YES
NO
NO
YES
YES
NO

```