# Poker 3

## Task Conditions

A deck consists of 52 cards. Each card is designated by one of thirteen values (2, 3, 4, 5, 6, 7, 8, 9, Ten, Jack, Queen, King, Ace) and one of four suits (Spades, Clubs, Diamonds, Hearts).

The drawn game of 3-Poker occurs as follows:

1. Initially, all players receive two cards from the deck.
2. After that, one card from the same deck is laid out on the table.
3. The players who collected the oldest combination win.

To determine the oldest combination that the i-th player has collected, the following rules are used:

- If both cards in the player's hand and the card on the table have the same value, the player collected a combination of 'Set with the value x';
- If from the two cards in the player's hand and the card on the table you can choose two cards with the same value x, the player collected a combination of 'Pair with the value x';
- Otherwise, the card with the oldest value from the two cards in the player's hand and the card on the table is retained, then the player collected the combination 'Oldest card x'.

Any set is older than any pair, and any pair is older than the combination of the oldest card. From identical combinations, the one with the oldest value wins. If the identical oldest combination is available to several players, they are all declared winners.

You are the first player. You know what cards each player has received in their hand. Determine which card can be laid out on the table so that you end up among the winners.

# Input Data

Each test consists of several sets of input data. The first line contains an integer t (1 ≤ t ≤ 10^3) — the number of sets of input data. The description of the sets of input data follows.

The first line of each set of input data contains an integer n (2 ≤ n ≤ 25) — the number of players.

The following n lines of each set of input data contain descriptions of two cards, separated by a space — the cards that the i-th player has received.

The description of the card consists of two symbols, written in sequence: the value and the suit.

# Output Data

For each set of input data, output in the first line the number of cards k, which can be laid out on the table to win. In the following k lines, output descriptions of these cards. The descriptions can be output in any order.

Let's take the first example.

In the first set of input data, for the first player to win, you can lay out a ten (T) on the table, then the first player wins with a combination of set with the value T.

In the second set of input data, it is impossible for the first player to win.

In the third set of input data, for the first player to win, you can lay out a seven (7) on the table, then the first player wins with a combination of pair with the value 7.

In the fourth set of input data, it is impossible for the first player to win.


# Test Example 1

## Input Data
```
4
2
TS TC
AD AH
3
2H 3H
9S 9C
4D QS
3
4C 7H
4H 4D
6S 6H
3
2S 3H
2C 2D
3C 3D
```

## Output Data
```
2
TD
TH
0
3
7S
7C
7D
0
```


# Test Example 2

## Input Data
```
1
7
AS AC
AD AH
KS JH
9D 9C
5H 5D
3C 3S
TC TH
```

## Output Data
```
30
2S
2C
2D
2H
4S
4C
4D
4H
6S
6C
6D
6H
7S
7C
7D
7H
8S
8C
8D
8H
JS
JC
JD
QS
QC
QD
QH
KC
KD
KH
```