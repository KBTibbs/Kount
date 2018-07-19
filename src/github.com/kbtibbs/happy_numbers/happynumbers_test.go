package happynumber

import (
    "testing"
    "fmt"
    "os"
)

func TestMain(m *testing.M) {
    // open the out file for writing
    f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }

    defer f.Close()
    os.Stdout = f

    os.Exit(m.Run())
}

type Test struct {
    in  int
    out bool
    err error
}

var tests = []Test{
    {1, true, nil},
    {2147483647, false, nil},
    {-1, false, fmt.Errorf("TestNumber: cannot test numbers below 1: %d", -1)},
}

func TestHappyNumbers(t *testing.T) {
    for i, test := range tests {
        isHappy, err := TestNumber(test.in)
        
        // This isn't the best test, beacuse I'm only testing to see that AN error
        // happened, not necessarily that THE CORRECT error happened.
        // TODO: Figure out how to correctly compare errors.
        if (err == nil) != (test.err == nil) {
            // either an unexpected error or there should have been an error and wasn't
            t.Errorf("#%d: TestNumber(%d)=%s; want %s", i, test.in, err, test.err)
        } else if isHappy != test.out {
            t.Errorf("#%d: TestNumber(%d)=%t; want %t", i, test.in, isHappy, test.out)
        } 
    }
}

func BenchmarkHappNumbers(b *testing.B) {
    for n := 1; n < b.N; n++ {
        TestNumber(n)
    }
}