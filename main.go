package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type MagicSquare struct {
	size   *big.Int
	square [][]*big.Int
	unique bool
	power  *big.Int // 2 for squares, 3 for cubes, etc.
}

// NewMagicSquare initializes an empty MagicSquare of given size.
func NewMagicSquare(size *big.Int, unique bool, power *big.Int) *MagicSquare {
	square := make([][]*big.Int, size.Int64()) // NOTE: this will fail on size > max int64!!!
	for i := range square {
		square[i] = make([]*big.Int, size.Int64()) // NOTE: this will fail on size > max int64!!!
		for j := range square[i] {
			square[i][j] = new(big.Int)
		}
	}
	return &MagicSquare{size: size, square: square, unique: unique, power: power}
}

// GenerateRandomBigInt generates a random big.Int between lowerBound and upperBound, optionally raising it to the specified power.
func GenerateRandomBigInt(lowerBound, upperBound, power *big.Int) (*big.Int, error) {
	diff := new(big.Int).Sub(upperBound, lowerBound)
	randNum, err := rand.Int(rand.Reader, new(big.Int).Add(diff, big.NewInt(1)))
	if err != nil {
		return nil, err
	}
	randNum.Add(randNum, lowerBound)
	if power.Cmp(big.NewInt(1)) > 0 {
		return new(big.Int).Exp(randNum, power, nil), nil
	}
	return randNum, nil
}

// PopulateSquareRandom fills the square with random big.Int values between lowerBound and upperBound.
func (ms *MagicSquare) PopulateSquareRandom(lowerBound, upperBound *big.Int) error {
	usedNumbers := make(map[string]bool)

	for i := big.NewInt(0); i.Cmp(ms.size) < 0; i.Add(i, big.NewInt(1)) {
		for j := big.NewInt(0); j.Cmp(ms.size) < 0; j.Add(j, big.NewInt(1)) {
			for {
				num, err := GenerateRandomBigInt(lowerBound, upperBound, ms.power)
				if err != nil {
					return err
				}
				if !ms.unique || !usedNumbers[num.String()] {
					ms.square[i.Int64()][j.Int64()].Set(num) // NOTE: this will fail on size > max int64!!!
					usedNumbers[num.String()] = true
					break
				}
			}
		}
	}
	return nil
}

// PopulateSquareDetermined partially fills the square and attempts to complete it as a magic square.
func (ms *MagicSquare) PopulateSquareDetermined(lowerBound, upperBound *big.Int) error {
	usedNumbers := make(map[string]bool)

	// Fill first row randomly - and track sum value, and used numbers
	rowSum := big.NewInt(0)
	for j := big.NewInt(0); j.Cmp(ms.size) < 0; j.Add(j, big.NewInt(1)) {
		num, err := GenerateRandomBigInt(lowerBound, upperBound, ms.power)
		if err != nil {
			return err
		}
		for usedNumbers[num.String()] {
			num, err = GenerateRandomBigInt(lowerBound, upperBound, ms.power)
			if err != nil {
				return err
			}
		}
		ms.square[0][j.Int64()].Set(num) // NOTE: this will fail on size > max int64!!!
		rowSum.Add(rowSum, num)
		usedNumbers[num.String()] = true
	}

	// Fill rest of square strategically
	tempColSums := make([]*big.Int, ms.size.Int64())
	for i := big.NewInt(0); i.Cmp(ms.size) < 0; i.Add(i, big.NewInt(1)) {
		tempRowSum := big.NewInt(0)
		for j := big.NewInt(0); j.Cmp(ms.size) < 0; j.Add(j, big.NewInt(1)) {
			if i.Cmp(new(big.Int).Sub(ms.size, big.NewInt(1))) == 0 || j.Cmp(new(big.Int).Sub(ms.size, big.NewInt(1))) == 0 { // If last row or column
				num := new(big.Int).Sub(rowSum, tempRowSum)
				ms.square[i.Int64()][j.Int64()] = num                                               // NOTE: this will fail on size > max int64!!!
				tempColSums[j.Int64()].Add(tempColSums[j.Int64()], ms.square[i.Int64()][j.Int64()]) // should always add to rowSum
				usedNumbers[num.String()] = true
			} else {
				for {
					newUpperBound := new(big.Int).Set(rowSum)

					if rowSum.Cmp(tempColSums[j.Int64()]) < 0 {
						newUpperBound.Sub(newUpperBound, tempColSums[j.Int64()])
					} else {
						newUpperBound.Sub(newUpperBound, rowSum)
					}

					num, err := GenerateRandomBigInt(lowerBound, upperBound, ms.power)
					if err != nil {
						return err
					}
					for usedNumbers[num.String()] {
						num, err = GenerateRandomBigInt(lowerBound, upperBound, ms.power)
						if err != nil {
							return err
						}
					}

					usedNumbers[num.String()] = true
				}
			}
		}
	}
	if ms.unique {
		if int64(len(usedNumbers)) < (ms.size.Int64() * ms.size.Int64()) {
			return fmt.Errorf("Failed to populate unique square with unique numbers")
		}
	}
	return nil
}

// IsMagic checks whether the square satisfies the magic square conditions.
func (ms *MagicSquare) IsMagic() bool {
	if ms.size == 0 {
		return false
	}

	sum := new(big.Int)
	tempSum := new(big.Int)
	for j := 0; j < ms.size; j++ {
		sum.Add(sum, ms.square[0][j])
	}

	// Check rows
	for i := 1; i < ms.size; i++ {
		tempSum.SetInt64(0)
		for j := 0; j < ms.size; j++ {
			tempSum.Add(tempSum, ms.square[i][j])
		}
		if tempSum.Cmp(sum) != 0 {
			return false
		}
	}

	// Check columns
	for j := 0; j < ms.size; j++ {
		tempSum.SetInt64(0)
		for i := 0; i < ms.size; i++ {
			tempSum.Add(tempSum, ms.square[i][j])
		}
		if tempSum.Cmp(sum) != 0 {
			return false
		}
	}

	// Check diagonals
	mainDiag := new(big.Int)
	antiDiag := new(big.Int)
	for i := 0; i < ms.size; i++ {
		mainDiag.Add(mainDiag, ms.square[i][i])
		antiDiag.Add(antiDiag, ms.square[i][ms.size-i-1])
	}

	return mainDiag.Cmp(sum) == 0 && antiDiag.Cmp(sum) == 0
}

// PrintSquare prints the magic square.
func (ms *MagicSquare) PrintSquare() {
	for _, row := range ms.square {
		for _, num := range row {
			fmt.Printf("%s\t", num.String())
		}
		fmt.Println()
	}
}

func main() {
	ms := NewMagicSquare(big.NewInt(3), true, big.NewInt(1))
	lowerBound, ok := new(big.Int).SetString("1", 10)
	if !ok {
		fmt.Println("Error creating lower bound")
		return
	}
	upperBound, ok := new(big.Int).SetString("1000000", 10)
	if !ok {
		fmt.Println("Error creating lower bound")
		return
	}

	if err := ms.PopulateSquareRandom(lowerBound, upperBound); err != nil {
		fmt.Println("Error populating square:", err)
		return
	}
	ms.PrintSquare()

	if ms.IsMagic() {
		fmt.Println("Random square IS a Magic Square!")
	} else {
		fmt.Println("Random square IS NOT a Magic Square.")
	}

	if err := ms.PopulateSquareDetermined(lowerBound, upperBound); err != nil {
		fmt.Println("Error populating determined square:", err)
		return
	}
	ms.PrintSquare()

	if ms.IsMagic() {
		fmt.Println("Determined square IS a Magic Square!")
	} else {
		fmt.Println("Determined square IS NOT a Magic Square.")
	}
}
