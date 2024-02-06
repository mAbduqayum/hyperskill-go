# Sea Battle

## Task Description

This is the simplest task of the contest. We hope you will get acquainted with the Ozon Techpoint testing system.

Correct solutions to the tasks must pass all the tests prepared in advance by the jury and fit within the time/memory
limits on each test.

Below are the technical requirements for the solutions:

- The solution reads input data from the standard input (screen);
- The solution writes output data to the standard output (screen);
- The solution does not interact in any way with other computer resources (network, hard disk, processes, etc.);
- The solution uses only the standard library of the language;
- The solution is placed in the default package (or its equivalent for your language) and has a standard entry point for
  console programs;
- It is guaranteed that all tests meet the constraints contained in the task description - there is no need to check the
  input data for correctness, all tests strictly correspond to the described format in the task;
- Display the answer exactly in the format described in the task (do not display "explanatory" comments
  like `enter number` or `the answer is`);

Let's move on to the task.

You are participating in the development of a subsystem for checking the field for the game "Sea Battle". You need to
write a check for the correct number of ships on the field, taking into account their sizes. Let's recall that the field
must have:

- Four single-deck ships,
- Three double-deck ships,
- Two triple-deck ships,
- One four-deck ship.

You are given 10 integers from 1 to 4. Check that the given sizes meet the requirements above.

## Input Data

The first line contains an integer t (1 ≤ t ≤ 1000) — the number of sets of input data in the test.

The sets of input data in the test are independent. They do not affect each other in any way.

Each set of input data consists of one line, which contains 10 integers a1, a2, …, a10 (1 ≤ ai ≤ 4) — the sizes of the
ships on the field in any order.

Note that it is already guaranteed that there are exactly 10 ships on the field and their sizes are from 1 to 4,
inclusive. You need to check that the number of ships of each type corresponds to the rules of the game.

## Output Data

For each set of input data, output in a separate line:

- `YES`, if the given sizes of the ships on the field comply with the rules of the game;
- `NO`, otherwise.

You can output `YES` and `NO` in any case (for example, the strings `yEs`, `yes`, `Yes`, and `YES` will be recognized as
a positive answer).

## Example Test 1

### Input Data
```
5
2 1 3 1 2 3 1 1 4 2
1 1 1 2 2 2 3 3 3 4
1 1 1 1 2 2 2 3 3 4
4 3 3 2 2 2 1 1 1 1
4 4 4 4 4 4 4 4 4 4
```

### Output Data
```
YES
NO
YES
YES
NO
```
