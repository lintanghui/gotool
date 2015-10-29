package gotool

import (
    "fmt"
)

func TestStack(t *testing.T) {
    str := string(stack(2))
    if str == "" {
        t.Fail()
    }

}
