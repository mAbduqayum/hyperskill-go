# 3 True Queue

# Problem Statement

In the system, there is a message queue, as well as 3 events that are logged into it - X, Y, and Z. It is known that there are external services that can make entries of events in the following pairs - XY, XZ, and YZ. Let's call them *correct pairs*. You were asked to collect analytics from the queue. For this, you waited until all services finished recording, and paused the queue operation.

Currently, the queue consists of `n` messages. Since recording occurs asynchronously, events from the service in the queue may not be in sequence, meaning after recording the first event by one service, an event from another service could have been recorded.

Your task is to determine whether it is possible to represent the current queue as a set of `n/2` correct pairs of events, where one event belongs to only one correct pair.

## Input
Each test consists of several sets of input data.

The first line contains an integer `t` (1 ≤ t ≤ 10^4) - the number of sets of input data. This is followed by a description of the sets of input data.

The first line of each set of input data contains an even integer `n` (2 ≤ n ≤ 2⋅10^5) - the number of messages in the queue.

The second line of each set of input data contains a string of `n` characters X, Y, and Z - messages from the queue, from the very first to the very last.

It is guaranteed that the sum of the values of `n` across all sets of input data does not exceed 2⋅10^5.

## Output
For each set of input data, print in a separate line `Yes` if it is possible. Otherwise, print `No`.

- For the first set of input data, one of the possible options is a set of correct pairs with indices:
  - (1, 5) - YZ,
  - (2, 4) - XY,
  - (3, 6) - YZ;

- For the second set, it can be proven that from this queue it is impossible to select two pairs so that both of them are correct;

- For the third set, one of the possible options is a set of correct pairs with indices:
  - (1, 3) - XZ,
  - (2, 5) - YZ,
  - (4, 6) - XY.


### Example Test 1

### Input
```
3
6
YXYYZZ
4
ZYXZ
6
XYZXZY
```

### Output
```
Yes
No
Yes
```


### Example Test 2

### Input
```
3
16
YYZYZZZZYXYZZXYX
20
YYXXYXYZZXYXZZZYZZYY
10
XYYXYZZZXY
```

### Output
```
No
Yes
Yes
```