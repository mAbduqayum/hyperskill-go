# Terminal

## Task Conditions

Implement a functionality element of a basic terminal.

Initially, the terminal contains one empty line, with the cursor at the beginning.

Your program should be able to process a sequence of character inputs (input string). The processing of a character depends on its value:

- The letters `L` and `R` represent the left and right arrow keys, respectively. They move the cursor one position to the left or right. If there is no character in the corresponding direction, the operation is ignored. Note that the cursor, in any case, remains in the same line.
- The letters `U` and `D` represent the up and down arrow keys. They move the cursor up one line or down one line. If there is no line in the corresponding direction, the operation is ignored. If a line exists, but there is no required position in it, the cursor moves to the end of the line.
- The letters `H` and `E` represent the `Home` and `End` keys. They move the cursor to the beginning or the end of the current line.
- The letter `N` represents pressing the `Enter` key — it inserts a new line. If the cursor is not at the end of the current line, it splits, and part after the cursor moves to the new line. The cursor after this operation is at the beginning of the new line.

You can imagine that this simulates the sequence of keystrokes in a text editor where the cursor is positioned between two characters of the line (or at the beginning or end of the line).

For example, if the input string looks like `otLLLrRuEe256LLLN`, the result will be two lines:

```
route
256
```

## Input Data

In the first line of the input data, there is an integer (1 ≤ t ≤ 100) — the number of input data sets.

The sets of input data are independent. One does not affect the other.

Each set of input data consists of one non-empty line — a sequence of characters for processing. It is guaranteed that the length of this line does not exceed 50. The allowed characters in the line are lowercase letters of the Latin alphabet, digits, and the letters L, R, U, D, H, E, N.

## Output Data

For each set of input data, output the resulting sequence of lines. Output all lines, including empty ones. After each set of output data, output an additional line with a single character `-` (minus).

## Test Example 1

### Input Data

```
4
otLLLrRuEe256LLLN
LRLUUDE
itisatest
abNcdLLLeUfNxNx
```


### Output Data

```
route
256
-

-
itisatest
-
af
x
xb
ecd
-
```