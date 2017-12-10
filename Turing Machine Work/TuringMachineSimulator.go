// Turing machine simulator, checks for a sequence of 1010... on the tape. Machine and initial tape provided in binary.
// Written for specific specifications given in assignment, with intent of generalisation in the future.
// Some functionality commented out for ease of use (tape and state access heat maps)

package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {

	// Set up machine testing parameters
	numStates := 7
	numSymbols := 2
	tapeStartPos := 0
	tapeMoves := []int8{-1, 1}
	sequence := "10"
	forcePositionZero := true
	forceToEnd := false
	allowTail := true

	pauseBetweenTapePrints := 0// Makes the machine output readable while it's running, tape won't print if set to -1
	printHeatMaps := false

	// Remove IDE boolean warnings
	if false {
		forcePositionZero, forceToEnd, allowTail = false, false, false
	}

	// Multiple test cases
	fmt.Println("\n")
	for true {
		fmt.Println("\tEnter 70 bit (7 state, 2 symbol, 2 directional) machine, and then enter the first 30 bits on the tape:\n")
		alpha := make(chan string, 2)
		var binary, initTape string
		// /*
		fmt.Print("\tMachine:\t")
		fmt.Scanf("%s\n", &binary)
		fmt.Print("\tTape:\t\t")
		fmt.Scanf("%s\n", &initTape)
		o1, steps, hMapTape, hMapStates := runTuringMachine(binary, initTape, numStates, numSymbols, tapeStartPos, tapeMoves, alpha, pauseBetweenTapePrints)
		if steps>0 {
			fmt.Println("\tSteps:\t\t" + strconv.Itoa(steps))
		}

		if printHeatMaps {
			fmt.Println("\tTape heat map:")
			for i := 0; i < len(o1); i++ {
				fmt.Printf("\t%d\t%d\n", i, hMapTape[i])
			}

			fmt.Println("\tState heat map:")

			for i := 0; i < numStates; i++ {
				fmt.Printf("\t%d\t%d\n", i, hMapStates[i])
			}
		}

		if len(o1) == 0 {
			fmt.Println("\tThe machine didn't halt.")
		} else {
			maxSeqRec := checkOutput(sequence, o1, forcePositionZero, forceToEnd, allowTail)
			reportRecord(binary, o1, maxSeqRec)
		}

		fmt.Println("\n\n\tAnd again:")

	}
}

func runTuringMachine(machineBinary, initialTape string, numStates, numSymbols, tapeStartPos int, tapeMoves []int8, chanOut chan string, pause int) (output string, steps int, heatMapRet map[int]int, heatMapRetStates map[int]int) {

	backup := machineBinary

	tapePrints := false

	bitsForStates := len(toBase(genBase(2), numStates))
	bitsForSymbols := len(toBase(genBase(2), numSymbols-1))
	bitsForMoves := len(toBase(genBase(2), len(tapeMoves)-1))

	// Add leading zeros if not present for proper length
	for len(machineBinary) < (bitsForMoves+bitsForStates+bitsForSymbols)*numSymbols*numStates {
		machineBinary = "0" + machineBinary
	}

	// Verify input
	properLength := (bitsForMoves + bitsForStates + bitsForSymbols) * numSymbols * numStates
	if len(machineBinary) != properLength {
		fmt.Println("Input error!")
		// os.Exit(1)
	}

	// Create Turing Machine definition
	states := make([][][]byte, numStates)
	for i := range states {
		states[i] = make([][]byte, numSymbols)
		for j := range states[i] {
			states[i][j] = make([]byte, 3)
		}
	}

	// Add Turing Machine values
	for i := range states {

		for j := range states[i] {

			for k := range states[i][j] {
				if k == 0 {
					states[i][j][k] = byte(toInt(machineBinary[0:bitsForStates], genBase(2)))
					machineBinary = machineBinary[bitsForStates:]
				} else if k == 1 {
					states[i][j][k] = byte(toInt(machineBinary[0:bitsForSymbols], genBase(2)))
					machineBinary = machineBinary[bitsForSymbols:]
				} else {
					states[i][j][k] = byte(toInt(machineBinary[0:bitsForMoves], genBase(2)))
					machineBinary = machineBinary[bitsForMoves:]
				}
			}

		}
	}

	// Restore machineBinary value
	machineBinary = backup

	// Verify machine
	for i := range states {
		for j := range states[i] {
			if states[i][j][0] > byte(numStates) {
				return
			}
		}
	}

	// fmt.Println(states)
	
	// Slightly neater
	for i:=0; i<len(states); i++ {

		fmt.Printf("\tState #%d:\t", i)
		fmt.Println(states[i])

	}

	/*
		// Print Turing machine
		for i:=0;i<len(states); i++ {
			for j:=0;j<len(states[i]); j++ {
				for k:=0;k<len(states[i][j]); k++ {
					if k==0 {
						stateOut := toBase(genBase(2), int(states[i][j][k]))
						for len(stateOut) != bitsForStates {
							stateOut = "0"+stateOut
						}
						// fmt.Print(stateOut)
					} else {
						// fmt.Print(states[i][j][k])
					}
				}
			}
		}
	*/

	// Set up tape and add predefined values
	tape := make([]uint8, len(initialTape))
	for i := range tape {
		val, _ := strconv.Atoi(initialTape[i : i+1])
		tape[i] = uint8(val)
	}

	// Set up tape heat maps
	heatMap := make(map[int]int)
	heatMapStates := make(map[int]int)

	// fmt.Print("\tMovements:\t")

	// Run the Turing Machine
	currentState := 0
	tapePosition := tapeStartPos
	lim := 0 // This is the max number of steps the Turing Machine will take before assuming it doesn't halt
	for ; lim <= 1000000; lim++ {
		nextState := int(states[currentState][tape[tapePosition]][0])
		nextWrite := states[currentState][tape[tapePosition]][1]
		nextDirection := int(tapeMoves[int(states[currentState][tape[tapePosition]][2])])

		// Heat maps
		heatMap[tapePosition]++
		heatMapStates[currentState]++

		tape[tapePosition] = nextWrite
		tapePosition += nextDirection
		currentState = nextState

		// Make the tape longer if required
		if tapePosition == len(tape) {
			tape = append(tape, 0)
		}

		// Prevent going left of the first position
		if tapePosition < 0 {
			tapePosition = 0
		}
		if tapePrints && pause >= 0 {
			for _, val := range tape {
				fmt.Print(val)
			}
			time.Sleep(time.Duration(pause) * time.Millisecond)
			fmt.Println()
		}

		// Check for halt instruction
		if currentState == numStates {
			for _, val := range tape {
				output += strconv.Itoa(int(val))
			}
			steps = lim + 1
			chanOut <- backup
			chanOut <- output
			heatMapRet = heatMap
			heatMapRetStates = heatMapStates

			return
		}

	}

	// If the Turing Machine doesn't halt, return null values
	chanOut <- ""
	chanOut <- ""
	machineBinary = ""

	return

}

func reportRecord(binary, out string, record int) {

	sendString := "\tOutput:\t\t" + out + "\n\tSequence length: " + strconv.Itoa(record) + "\n"
	fmt.Print(sendString)

}

func checkOutput(sequence, out string, forcePositionZero, forceToEnd, allowTail bool) (recordLen int) {

	if len(out)%len(sequence)!=0 && forceToEnd {
		return
	}

	for i := range out {
		currentRec := 0
		if forcePositionZero && i > 0 {
			break
		}
		if i+len(sequence) > len(out) {
			break
		}
		if out[i:i+len(sequence)] == sequence {
			currentRec += len(sequence)
			for index := i + len(sequence); index < len(out); index += len(sequence) {
				if index+len(sequence) > len(out) {
					if out[index:]==sequence[:len(out[index:])] && allowTail {
						currentRec += len(out[index:])
					}
					break
				}
				if out[index:index+len(sequence)] == sequence {
					currentRec += len(sequence)
				} else {
					if allowTail {
						for k := range out[index:] {
							if out[index+k:index+1+k] != "0" {
								recordLen=0
								return
							}
						}
					}
					break
				}
			}
		}

		if currentRec > recordLen {
			recordLen = currentRec
		}

	}

	return

}

func intPower(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func toInt(input, source string) (outputInt int) {
	baseSource := len(source)
	if baseSource == 1 {
		return len(input)
	}
	for j, indexInSource := len(input)-1, -1; j >= 0; j-- {
		indexInSource = strings.Index(source, string(input[j]))
		outputInt += indexInSource * intPower(baseSource, len(input)-j-1)
	}
	return
}

func toBase(target string, input int) (output string) {
	baseTarget := len(target)
	for remainder := input; true; {
		output += string(target[remainder%baseTarget])
		remainder = (remainder - remainder%baseTarget) / baseTarget
		if remainder == 0 {
			break
		}
	}
	output = reverseString(output)
	return
}

func reverseString(s string) (sNew string) {
	for i := len(s) - 1; i >= 0; i-- {
		sNew += string(s[i])
	}
	return
}

func genBase(base int) string {
	mostBases := "0123456789abcdefghijklmnopqrstuvwxyz"
	if base == 1 {
		return "1"
	} else {
		return mostBases[0:base]
	}
}
