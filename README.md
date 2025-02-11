# Magic Square Finder

## Introduction

Magic squares are mathematical structures where numbers are arranged in an: $ n \times n $ grid such that the sum of the numbers in each row, each column, and both main diagonals are equal. This sum is called the magic constant.

This project aims to develop a computational approach to finding magic squares, starting with basic random generation and progressing toward more efficient algorithms that can generate valid magic squares with constraints such as uniqueness and specific power conditions.

## Inspiration

This project is inspired by the properties of magic squares and their connections to various fields of mathematics and recreational number theory. The following videos by or with Matt Parker provide an excellent introduction to the topic and explore some of the challenges and peculiarities of magic squares:

- [Magic Squares Explained](https://www.youtube.com/watch?v=stpiBy6gWOA)
- [How to Solve a Magic Square](https://www.youtube.com/watch?v=aOT_bG-vWyg)
- [Recreating the Parker Square](https://www.youtube.com/watch?v=U9dtpycbFSY)
- [Investigating Magic Square Properties](https://www.youtube.com/watch?v=aQxCnmhqZko)
- [How Magic Squares are Constructed](https://www.youtube.com/watch?v=FCczHiXPVcA)
- [Solving an Unsolvable Magic Square](https://www.youtube.com/watch?v=uz9jOIdhzs0)
- [The Illusive Parker Square](https://www.youtube.com/watch?v=Kdsj84UdeYg)
- [History of Magic Squares](https://www.youtube.com/watch?v=T0U9ou0HOgY)
- [Connections Between Magic Squares and Other Maths Topics](https://www.youtube.com/watch?v=G1m7goLCJDY)

## Features

### Current Features:

- **Random Magic Square Generation:** Generates magic squares using random numbers within a specified range.
- **Unique Entries Constraint:** Ensures that each number appears only once in the square (optional).
- **Power Constraints:** Numbers can be forced to be perfect squares, cubes, or higher powers.
- **Magic Square Validation:** Checks if a given square meets the magic square conditions.

### Planned Features:

- **Guided Magic Square Generation:** Instead of random generation, attempt to strategically place numbers to ensure a valid magic square.
- **Optimized Searching Algorithms:** Implement backtracking and other methods to find magic squares more efficiently.
- **Larger and More Complex Magic Squares:** Extending support for large magic squares while maintaining computational feasibility (although... evidence is - finding them gets EASIER as they get larger...)

## Getting Started

### Prerequisites

- Go (latest version recommended)
- Git (for version control)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/magic-square-finder.git
   cd magic-square-finder
   ```
2. Build the project:
   ```sh
   go build
   ```
3. Run the program:
   ```sh
   go run main.go
   ```

## Usage

The program currently generates magic squares of a given size using random numbers within a specified range. Future updates will include more structured methods for finding magic squares.

## Contributions

Contributions are welcome! If you have ideas for improving the efficiency of the algorithm or implementing new techniques, feel free to submit a pull request.

## License

This project is licensed under the MIT License.
