package viterbi

import (
    "strings"
    "io/ioutil"
    "log"
)

type corpus struct {
    total float64
    words map[string]int
    maxlen int
}

// Init creates a new corpus object
func Init() *corpus {
    c := &corpus{}
    return c
}

// LoadData loads data from the given path and builds probability
func (c *corpus) LoadData(path string) {
    content, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }

    words := []string{}
    split := strings.Split(string(content), "\n")

    for i := range split {
        tokens := strings.Split(split[i], "  ")

        for j := range tokens {
            words = append(words, tokens[j])
        }
    }

    c.buildProb(words)
}

// buildProb records the count of each word to be used as probability
func (c *corpus) buildProb(data []string) {
    maxlen := 0
    dict := map[string]int{}
    total := 0
    for i  := range data {
        l := len(data[i])
        if l > maxlen {
            maxlen = l
        }
        dict[data[i]] += 1
        total += 1
    }

    c.total = float64(total)
    c.words = dict
    c.maxlen = maxlen
}


