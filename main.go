package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	switch len(os.Args) {
	case 2:
		// it's a 64-bit integer timestamp
		tsint, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			fmt.Printf("Need a number, not %s", os.Args[1])
			break
		}
		time, incr := fromTimestamp(uint64(tsint))
		fmt.Printf("Time (seconds since epoch): %d Increment %d\n", time, incr)

	case 3:
		timeint, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Need a number, not %s", os.Args[1])
			break
		}
		incrint, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Need a number, not %s", os.Args[1])
			break
		}
		ts := toTimestamp(uint32(timeint), uint32(incrint))
		fmt.Printf("Timestamp: %d\n", ts)
	default:
		fmt.Println("Please use one number for conversion from 64-bit timestamp, two numbers for conversion to 64-bit timestamp")
	}
}

// returns 64-bit unsigned integer which is the BSON timestamp
func toTimestamp(time uint32, incr uint32) uint64 {
	ts := uint64(time) << 32
	ts |= uint64(incr)
	return ts
}

// returns time and increment components of a BSON timestamp
func fromTimestamp(ts uint64) (uint32, uint32) {
	time := uint32(ts >> 32)
	incr := uint32(ts & 0x00000000FFFFFFFF)
	return time, incr
}
