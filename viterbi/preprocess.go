package viterbi

import (
    "strings"
    "io/ioutil"
    "log"
)

func LoadData(path string) [][][]rune {
    content, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }

    lines := [][][]rune{}
    split := strings.Split(string(content), "\n")

    for i := range split {
        tokens := strings.Split(split[i], "  ")
        lines = append(lines, [][]rune{})

        for j := range tokens {
            lines[i] = append(lines[i], []rune(tokens[j]))
        }
    }
    return lines
}

func BuildDict(data [][][]rune) map[rune]bool {
    dict := map[rune]bool{}
    for i  := range data {
        for j := range data[i] {
            for k := range data[i][j] {
                dict[data[i][j][k]] = true
            }
        }
    }
    return dict
}
