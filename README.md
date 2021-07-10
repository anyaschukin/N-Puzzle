# N-Puzzle

solve [N-puzzles](https://en.wikipedia.org/wiki/Sliding_puzzle) using the A* search algorithm with several heuristics.

The goal is to solve quickly, with a target of under 10 seconds for puzzle size 3. This project solves size 3 in at worst a few milliseconds.

See the [subject](https://github.com/anyaschukin/N-Puzzle/blob/master/subject.pdf) for more details.

*Final Score 114/100*


## Getting Started

First you need to have your golang workspace set up on your machine.
Then clone this repo into your go-workspace/src/ folder. <br>
```git clone https://github.com/anyaschukin/N-Puzzle.git; cd N-Puzzle```

Download dependencies. <br>
```go get -d ./...; pip install -r requirements.txt```

To run with puzzle size 3. <br>
```go run main.go -s 3```

Alternatively, build & run the binary. <br>
```go build; ./N-Puzzle -s 3```

*N-Puzzle* first generates a random puzzle.
If the puzzle is solvable *N-Puzzle* prints the solution from inital state to solved.
*N-Puzzle* then prints the number of moves required, size & time complexity, and solve time.

<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/3a.png" width="520">

... intermediate states ...

<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/3b.png" width="520">


## Puzzle

Find a valid sequence of moves to reach the solved state, a.k.a the "snail solution".

<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/puzzle.png" width="320">


## Heuristics

### Manhattan

### Hamming

### Euclidian

### Nilsson

### OutRowCol


## test.sh

<img src="https://github.com/anyaschukin/N-Puzzle/blob/master/img/test.png" width="420">


## Links

https://tristanpenman.com/demos/n-puzzle/

https://qiao.github.io/PathFinding.js/visual/
