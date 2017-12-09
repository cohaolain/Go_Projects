// Turing machine simulator, checks for a sequence of 1010... on the tape. Machine and initial tape provided in binary.
// Includes force method for finding Turing machines which output record length repetition of a given sequence

// Written for specific specifications given in assignment, with intent of generalisation in the future.

package main

import (
	"fmt"
	"github.com/simplepush/simplepush-go"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func runTuringMachine(machineBinary, initialTape string, numStates, numSymbols, tapeStartPos int, tapeMoves []int8,
	chanOut chan string) (output string, steps int) {

	backup := machineBinary

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
	var passedVerification bool
	for i := range states {
		for j := range states[i] {
			if states[i][j][0] > byte(numStates) {
				chanOut <- ""
				chanOut <- ""
				machineBinary = ""
				return
			}
			if states[i][j][0] == byte(numStates) {
				passedVerification = true
			}
		}
	}
	if !passedVerification {
		chanOut <- ""
		chanOut <- ""
		machineBinary = ""

		return
	}

	// fmt.Println(states)

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

	// Run the Turing Machine
	currentState := 0
	tapePosition := tapeStartPos
	lim := 0 // This is the max number of steps the Turing Machine will take before assuming it doesn't halt
	for ; lim <= 1000000; lim++ {
		nextState := int(states[currentState][tape[tapePosition]][0])
		nextWrite := states[currentState][tape[tapePosition]][1]
		nextDirection := int(tapeMoves[int(states[currentState][tape[tapePosition]][2])])

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

		// Check for halt instruction
		if currentState == numStates {
			for _, val := range tape {
				output += strconv.Itoa(int(val))
			}
			steps = lim + 1
			chanOut <- backup
			chanOut <- output

			return
		}

	}

	// If the Turing Machine doesn't halt, return null values
	chanOut <- ""
	chanOut <- ""
	machineBinary = ""

	return

}

func main() {

	// Set up a random seed based on the current time
	rand.Seed(time.Now().UnixNano())

	// Set the minimum record for the program to start printing out results at
	// seqRec := 242
	seqRec := 0

	// Set up machine testing parameters
	initTape := "101010101010101010101010101010" // Sets the first bits on the tape
	maxBitsTapeSpecified := 30
	numStates := 7
	numSymbols := 2
	tapeStartPos := 0
	tapeMoves := []int8{-1, 1}
	sequence := "10"
	forcePositionZero := true
	forceToEnd := false
	allowTail := true

	// Remove IDE boolean warnings
	if false {
		forcePositionZero, forceToEnd, allowTail = false, false, false
	}

	// Set up multi-threading concurrency parameters
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	// Set the number of threads to use (default is the number of logical processors available)
	numConcur := numCPUs

	printTimes := false

	// Fix variable warning workaround
	if false {
		printTimes = true
	}

	t1 := time.Now()

	// Machine testing with concurrency
	for true {

		// Set up a channel to receive answers from the concurrent threads
		answers := make(chan string, 2*numConcur)

		// Run the specified number of Turing Machines simultaneously
		for i := 0; i < numConcur; i++ {
			go runTuringMachine(genXBits(70), initTape, numStates, numSymbols, tapeStartPos, tapeMoves, answers)
		}

		// Runs once for each instantiated Turing Machine
		for i := 0; i < numConcur; i++ {

			// Receive the Turing machine and its resultant tape
			binary := <-answers
			out := <-answers

			// Start running checks if the Turing machine halted
			if len(out) >= len(sequence) {

				// If the sequence appears for a longer period than the current record, set new record and report it
				if maxSeqRec := checkOutput(sequence, out, forcePositionZero, forceToEnd, allowTail); maxSeqRec > seqRec {
					seqRec = maxSeqRec
					reportRecord(binary, out, maxSeqRec)
				}

			}

		}

	}

	if printTimes {
		fmt.Printf("That took %v with concurrency\n", time.Since(t1))
	}

	// Machine testing with NO concurrency
	t2 := time.Now()
	for false {

		// Set up a channel
		answers := make(chan string, 2*numConcur)
		binary := genXBits(70)

		out, _ := runTuringMachine(binary, initTape, numStates, numSymbols, tapeStartPos, tapeMoves, answers)

		// Start running checks if the Turing machine halted
		if len(out) >= len(sequence) {

			// If the sequence appears for a longer period than the current record, set new record and report it
			if maxSeqRec := checkOutput(sequence, out, forcePositionZero, forceToEnd, allowTail); maxSeqRec > seqRec {
				seqRec = maxSeqRec
				reportRecord(binary, out, maxSeqRec)
			}

		}

	}
	if printTimes {
		fmt.Printf("That took %v without concurrency\n", time.Since(t2))
	}

	// Case test
	if false {
		alpha := make(chan string, 2)
		binary := ""
		initTape = ""
		o1, steps := runTuringMachine(binary, initTape, numStates, numSymbols, tapeStartPos, tapeMoves, alpha)
		fmt.Println("steps: ", steps)
		if len(o1) == 0 {
			fmt.Println("machine didn't halt")
		}
		if len(o1) >= len(sequence) {
			maxSeqRec := checkOutput(sequence, o1, forcePositionZero, forceToEnd, allowTail)
			reportRecord(binary, o1, maxSeqRec)
		} else {
			fmt.Println("invalid")
		}

	}

	// Tape testing
	initTapeInt := 0
	for initTapeInt >= 0 && initTapeInt <= int(math.Pow(float64(numSymbols), float64(maxBitsTapeSpecified))) && false {
		binary := ""

		// Progress print-out
		if initTapeInt%(numConcur*50000) == 0 {
			fmt.Printf("%.5f%%\n", (float64(initTapeInt)*100)/math.Pow(2, 30))
		}

		alpha := make(chan string, 2*numConcur)
		tried := make([]string, 4)
		var initTape string
		for i := 0; i < numConcur; i++ {
			initTape = toBase(genBase(2), initTapeInt)
			tried[i] = initTape

			// Ensure generation of proper tape length
			for len(initTape) < maxBitsTapeSpecified {
				initTape = "0" + initTape
			}
			if initTapeInt < 0 {
				os.Exit(2)
			}
			go runTuringMachine(binary, initTape, numStates, numSymbols, tapeStartPos, tapeMoves, alpha)
			initTapeInt++
		}
		for i := 0; i < numConcur; i++ {

			bin, output := <-alpha, <-alpha

			if len(output) >= len(sequence) {

				if maxSeqRec := checkOutput(sequence, output, forcePositionZero, forceToEnd, allowTail); maxSeqRec > seqRec {
					seqRec = maxSeqRec
					reportRecord(bin, output, maxSeqRec)
					fmt.Println(tried)
				}

			}
		}
	}

}

func genXBits(n int) (binary string) {
	for i := 0; i < n; i++ {
		binary += strconv.Itoa(rand.Intn(2))
	}
	return
}

func reportRecord(binary, out string, record int) {

	if _, err := os.Stat("out.txt"); os.IsNotExist(err) {
		f, _ := os.Create("out.txt")
		f.Close()
	}

	sendString := "\n" + binary + "\n" + out + "\n" + strconv.Itoa(record) + "\n"
	fmt.Print(sendString)
	f, _ := os.OpenFile("out.txt", os.O_APPEND|os.O_WRONLY, 0644)
	f.WriteString(sendString)
	f.Close()
	// Send a dud notification if running on windows (bug workaround)
	if runtime.GOOS == "windows" {
		simplepush.Send(simplepush.Message{"5JiGEp", "null", "null", "", false, "", ""})
	}
	simplepush.Send(simplepush.Message{"5JiGEp", "Turing Machine: Length " + strconv.Itoa(record), sendString, "", false, "", ""})

}

func checkOutput(sequence, out string, forcePositionZero, forceToEnd, allowTail bool) (recordLen int) {

	if len(out)%len(sequence) != 0 && forceToEnd {
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
					if out[index:] == sequence[:len(out[index:])] && allowTail {
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
								recordLen = 0
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
