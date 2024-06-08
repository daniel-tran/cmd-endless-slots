package main
import (
    "fmt"
    "time"
    "os/exec"
    "os"
    "bufio"
    "math/rand/v2"
)

func cls() {
    // For Windows, commands have to be spawned in a separate cmd instance, refer to https://stackoverflow.com/a/19290028
    var cls = exec.Command("cmd", "/c", "cls")
    cls.Stdout = os.Stdout
    cls.Run()
}

// This function is predominantly sourced from https://stackoverflow.com/a/31146687
func getInput(input chan string) {
    for {
        var in = bufio.NewReader(os.Stdin)
        var result, err = in.ReadString('\n')
        if err != nil {
            fmt.Println("You decided to stop gambling.")
        }
        input <- result
    }
}

// Returns a list of indexes corresponding to positions of items in a slice from a given starting point
func getSlotResult(list []string, rows int, index int) []int {
    var result = []int{}
    for i := 0; i < rows; i++ {
        // Prepend the element to the results, as the expected ordering is rightmost = first
        // This is using the "Push Front/Unshift" trick from https://go.dev/wiki/SliceTricks
        result = append([]int{ (len(list) + (index - i)) % len(list) }, result...)
    }
    return result
}

func main() {
    var speed = "100"
    if (len(os.Args) >= 2) {
        speed = os.Args[1]
    }
    duration, _ := time.ParseDuration(speed + "ms")
    
    // These are organised from the rightmost element being the top to leftmost being the bottom
    // Columns are organised from the first being the leftmost column to last being the rightmost column
    var items = [][]string{
        {"♣", "♠", "♦", "♥"},
        {"♥", "♦", "♠", "♣"},
        {"♦", "♥", "♣", "♠"},
        {"♣", "♠", "♦", "♥"},
        {"♥", "♦", "♠", "♣"} }
    var columns = len(items)
    // This will loop around the values for a particular column if it is larger than a column's item count
    var rows = 3
    var values = [][]int{}
    var index = 0
    var randIndex = rand.IntN(len(items[index]))
    var previousResult = ""
    
    var input = make(chan string, 1)
    go getInput(input)
    
    for ;;{
        select {
            case i := <-input:
                fmt.Println(i)
                values = append(values, getSlotResult(items[index], rows, randIndex))
                randIndex = rand.IntN(len(items[index])) // Re-randomise the index to prevent the users from predicting the next column easily
                index++
                
                if index >= columns {
                    // Round has completed, should reset for the next round
                    index = 0

                    // Inform the user of what the last slot was, as pausing for any fixed amount of time on a command line is difficult as is
                    previousResult = ""
                    for y := 0; y < rows; y++ {
                        for x := 0; x < len(values); x++ {
                            previousResult += items[x][values[x][rows - 1 - y]]
                        }
                        previousResult += "\n"
                    }
                    
                    // Finally, clear the slot results
                    values = nil
                }
            case <-time.After(duration):
                cls()
                randIndex = (randIndex + 1) % len(items[index])
                if len(previousResult) > 0 {
                    fmt.Printf("Previous result:\n%v\n", previousResult)
                    fmt.Println("-----")
                }
                
                // Print the value of the user's current slot results as well as the current active column
                for y := 0; y < rows; y++ {
                    // Print row by row for each column that has a fixed value from previous pulls
                    for x := 0; x < len(values); x++ {
                        fmt.Print(items[x][values[x][rows - 1 - y]])
                    }
                    // Slots are drawn from index N...0 in the direction of top to bottom (in the slice, it's represented as right to left)
                    // This ensures the slots are going "downward"
                    fmt.Println(items[index][(len(items[index]) + (randIndex - y)) % len(items[index])])
                }
        }
    }
}
