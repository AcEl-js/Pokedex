package main

import "testing"

func TestCleancli(t *testing.T){
cases := []struct {
    input string
    expected []string
    
}{
{
        input :"Hello World", 
        expected : []string{
            "hello",
            "world",
        },
    },
}
for _, cs :=range cases{
    actual := cleanCli(cs.input)
    if (len(actual) != len(cs.expected)){
       t.Errorf("The lengths are not equal: %v vs %v",
        len(actual),
        len(cs.input),
        )
            continue
    }
    
        for i := range actual {
            actualWord := actual[i]
            expectedWord := cs.expected[i]
            if actualWord != expectedWord {
                t.Errorf("%v does not equal to %v",
                    actualWord,
                    expectedWord)
            }
        }
}
}

