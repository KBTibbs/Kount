package happynumber

import (
    "testing"
    "fmt"
)

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
        if isHappy != test.out {
            t.Errorf("#%d: TestNumber(%d)=%t; want %t", i, test.in, isHappy, test.out)
        } else if (err == nil) != (test.err == nil) {
            // This isn't the best test, beacuse I'm only testing to see that AN error
            // happened, not necessarily that THE CORRECT error happened.
            // TODO: Figure out how to correctly compare errors.
            t.Errorf("#%d: TestNumber(%d)=%s; want %s", i, test.in, err, test.err)
        }
    }
}