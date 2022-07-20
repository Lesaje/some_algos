//I was defeated. After week of searching for plain explanation of algorithm,
//rewriting code from C and Java, I can not make this sort working. 
//Maybe in future I will finish this sorting, so I leave code here
//maybe you could sort out how to correct this code.

package main

import (
    "encoding/binary"
    "fmt"
    "math/rand"
    "time"
)

func sort(input []int64, start, length, depth uint32) {
    count := make([]uint32, 256)
    offset := make([]uint32, 256)
    buf := make([]uint8, 8)
    for i := start; i <= length; i++ {
        d := input[i]
        binary.PutVarint(buf, d)
        count[buf[depth]]++
    }
    offset[0] = start
    for i := 1; i < 256; i++ {
        offset[i] = offset[i-1] + count[i-1]
    }

    for i := 0; i < 256; i++ {
        for count[i] > 0 {
            origin := offset[i]
            from := origin
            num := input[from]
            input[from] = -1
            for {
                binary.PutVarint(buf, num)
                b := buf[depth]

                to := offset[b]
                offset[b]++
                count[b]--
                tmp := input[to]
                input[to] = num
                num = tmp
                from = to
                if from == origin {
                    break
                }
            }
        }
    }
    if depth < 6 {
        for i := 0; i < 256; i++ {
            var begin uint32
            if i > 0 {
                begin = offset[i-1]
            } else {
                begin = start
            }
            end := offset[i]
            if end-begin > 1 {
                sort(input, begin, end, depth+1)
            }
        }
    }
}

func AmericanFlagSort(input []int64) {
    sort(input, 0, uint32(len(input)-1), 0)
}

func main() {
    input := make([]int64, 300)
    testSlice := rand.Perm(300)
    for i := 300 - 1; i >= 0; i-- {
        input[i] = int64(testSlice[i])
    }
    start := time.Now()
    AmericanFlagSort(input)
    fmt.Println(time.Since(start))
    fmt.Println(input)
}
