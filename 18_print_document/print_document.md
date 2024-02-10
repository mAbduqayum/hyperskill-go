# Print Document

## Task Conditions

It is necessary to print a document of 1 to k pages. The pages are numbered from 1 to k.

Some of the pages were already preliminarily printed. It is known that the printer had one last task, which contained a list of pages to print.

This list contains at least one page number, and although one page number or a range from 1 to k is not included in this list, the list of pages consists of elements separated by commas, where each element is:

- either the specific number of a single page (a whole number from 1 to k),
- or a range of pages, written in the format “l-r”, where l is the beginning of the range and r is the end of the range (l and r are whole numbers that satisfy 1 ≤ l ≤ r ≤ k).

The page can be mentioned multiple times in the list but will be printed only once.

In other words, the list of pages has a format similar to what is used in Microsoft Word® or other similar programs.

For example, if k = 8, then the permissible lists of pages are:

- 7 (only page 7 was printed),
- 1,7,1 (pages 1 and 7 were printed),
- 1-5,1,7-7 (pages 1,2,3,4,5,7 were printed).

Examples of incorrect ranges for k = 8 (these entries are not allowed):

- 1-8 (although one page of the document should still be unprinted),
- 1,,3 (commas cannot go in a row),
- 7-9 (a non-existent page cannot be sent to print),
- 1-5, (each entry must separate the elements of the list, she cannot finish the list),
- 1,2,3-5 (each comma must separate two elements, she cannot start the list),
- 3-4-7 (broken element format, so it is not allowed).

Output a shortened list of pages in the correct format that needs to be additionally sent to print so that in the end all pages from 1 to k, not printed earlier, are printed.

## Input Data

- In the first line of the input data is written a whole number t (1 ≤ t ≤ 100) — the number of sets of input data.

- The sets of input data in the test are independent. One set does not affect the others in any way.

- In the first line of each set is written a whole number k (2 ≤ k ≤ 100) — the number of pages in the document.

- The second line of each set contains a list of pages that have already been sent to print. This list is correct and formatted according to the rules described above. It contains only pages from 1 to k, contains at least one page, and although one page or a range from 1 to k is not included in it, the line contains from 1 to 400 characters, inclusive.

## Output Data

For each set of input data, output a line — for each set of input data output a shortened line, containing the correct list of pages that need to be printed and only them. If there are several optimal answers, then output any of them.

## Test Example 1

### Input Data

```
7
8
7
8
1,7,1
8
1-5,1,7-7
10
1-5
10
1,2,3,4,5,6,8,9,10
3
1-2
100
1-2,3-7,10-20,100
```


### Output Data
```
1-6,8
2-6,8
6,8
6-10
7
3
8-9,21-99
```