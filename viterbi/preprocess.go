package viterbi

import (
    "strings"
    "bufio"
    "log"
    "os"
    "strconv"
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
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()
    words := []string{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        tokens := strings.Split(strings.Trim(scanner.Text(), "\n"), " ")

        for j := range tokens {
            words = append(words, tokens[j])
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    c.buildProb(words)
}

// LoadDict loads customized dictionary from file
func (c *corpus) LoadDict(path string) {
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        count := 1000
        tokens := strings.Split(strings.Trim(scanner.Text(), "\n"), " ")
        if len(tokens) < 1 {
            continue
        }

        word := tokens[0]
        if len(tokens) == 2 {
            if c, err := strconv.Atoi(tokens[1]); err == nil {
                count = c
            }
        }
        c.words[word] += count
        c.total += float64(count)
        if len(word) > c.maxlen {
            c.maxlen = len(word)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

// buildPror records the count of each word to be used as probability
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


