# N-Puzzle

A [sliding puzzle](https://en.wikipedia.org/wiki/Sliding_puzzle) solver using the A* search algorithm with several heuristics.

The goal is to solve quickly, with a target of under 10 seconds for puzzle size 3. This project solves size 3 in at worst a few milliseconds.

See the [subject](https://github.com/anyaschukin/N-Puzzle/blob/master/subject.pdf) for more details.

*Final Score 114/100*


## Getting Started

First you need to have your golang workspace set up on your machine.
Then clone this repo into your go-workspace/src/ folder. <br>
```git clone https://github.com/anyaschukin/N-Puzzle.git; cd N-Puzzle```

Download dependencies. <br>
```go get -d ./...; pip install -r requirements.txt```

To run. <br>
```go run main.go```

Alternatively, build & run the binary. <br>
```go build; ./N-Puzzle```

*N-Puzzle* first generates a random puzzle.
If the puzzle is solvable *N-Puzzle* prints the solution from inital state to solved.
*N-Puzzle* then prints the number of moves required, size & time complexity, and solve time.

<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/3a.png" width="520">

... intermediate states ...

<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/3b.png" width="520">


## Puzzle

Find a valid sequence of moves to reach the solved state, a.k.a the "snail solution".

<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/puzzle.png" width="320">


## Flags

### -s Size

Set puzzle size, default 3.

```go run main.go -s 4```


### -h Heuristic

Set heuristic to guide the A* search. Default manhattan. Options:

* manhattan
* hamming
* euclidean
* nilsson
* outRowCol

```go run main.go -h hamming```


## test.sh

The *boards/* folder contains 169 unit tests, solvable and unsolvable, depth 3 to 9. *boards/* also contains [generator.py](https://github.com/anyaschukin/N-Puzzle/blob/master/boards/generator.py), a random boards generator.

*test.sh* runs static unit tests from the *boards/* folder, & random tests using *generator.py*.

```./test.sh```

<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/test.png" width="420">

For each size *test.sh* then plots solve time, moves, size & time complexity by heuristic. For size 3:

<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/3solve_time.png" width="420">
<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/3moves.png" width="420">
<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/3size_complexity.png" width="420">
<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/3time_complexity.png" width="420">

These plots show for size 3 the manhattan heuristic performs best, it solves fastest while still providing a low number of moves.

the nilsson heuristic is almost as quick as manhattan, but usually takes more moves. Heuristics outRowCol, hamming, & euclidean take progressively more time to solve compared to manhattan for no improvement in number of moves.

For size 4:

<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/4solve_time.png" width="420">
<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/4moves.png" width="420">
<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/4size_complexity.png" width="420">
<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/4time_complexity.png" width="420">


## Links

[Path finding visualizer](https://qiao.github.io/PathFinding.js/visual/)

[Tristan Penman n-puzzle app](https://tristanpenman.com/demos/n-puzzle/)

