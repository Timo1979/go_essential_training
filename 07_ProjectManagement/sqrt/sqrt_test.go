// run:
// go test -v -run TestSimple

// go test -v -bench . -run FAKETEST -cpuprofile=prof.out
// go tool pprof prof.out
// pprof> list Sqrt
package sqrt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func almostEquals(v1, v2 float64) bool {
	return Abs(v1-v2) < 0.001
}

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)
	if err != nil {
		t.Fatalf("error in calculation - %s\n", err)
	}
	if !almostEquals(val, 1.414214) {
		t.Fatalf("bad value %f", val)
	}
}

type Pair struct {
	val            float64
	expectedResult float64
}

func TestMany(t *testing.T) {
	testCases := []Pair{
		{0.0, 0.0},
		{1.0, 1.0},
		{2.0, 1.414214},
		{9.0, 3.0},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f", tc.val), func(t *testing.T) {
			out, err := Sqrt(tc.val)
			if err != nil {
				t.Fatal("error")
			}
			if !almostEquals(out, tc.expectedResult) {
				t.Fatalf("%f != %f", out, tc.expectedResult)
			}
		})
	}
}

func TestFromCsv(t *testing.T) {
	file, err := os.Open("sqrt_cases.csv")
	if err != nil {
		t.Fatalf("open file error %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for lnum := 1; scanner.Scan(); lnum++ {
		fields := strings.Split(scanner.Text(), ",")
		if len(fields) != 2 {
			t.Logf("%d bad line", lnum)
			continue
		}
		val, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			t.Fatalf("Wrong float format for %s", fields[0])
		}
		expected, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			t.Fatalf("Wrong float format for %s", fields[1])
		}
		t.Run(fmt.Sprintf("%f", val), func(t *testing.T) {

			actual, err := Sqrt(val)
			if err != nil {
				t.Fatalf("%s", err)
			}
			if !almostEquals(actual, expected) {
				t.Fatalf("%f != %f", actual, expected)
			}
		})

	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("%s", err)
	}
}

func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := Sqrt(float64(i)); err != nil {
			b.Fatal(err)
		}
	}
}
