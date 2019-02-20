package viterbi

type prob struct {
    pr float64
    k int
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func maxProb(y []*prob) *prob {
    pr := &prob{}

    for _, z := range y {
        if z.pr > pr.pr {
            pr.pr = z.pr
            pr.k = z.k
        }
    }

    return pr
}

func reverse(s []string) []string {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return s
}

func (c *corpus) wordProb(w string) float64 {
    return float64(c.words[w]) / c.total
}

func (c *corpus) Predict(s string) []string {
    init := &prob{pr: 1.0, k: 0}
    T := []*prob{init}

    for i := 1; i < len(s) + 1; i++ {
        y := []*prob{}
        for j := max(0, i - c.maxlen); j < i; j++ {
            p := &prob{pr: T[j].pr * c.wordProb(s[j:i]), k: j}
            y = append(y, p)
        }

        pr := maxProb(y)
        T = append(T, pr)
    }

    words := []string{}
    i := len(s)

    for i > 0 {
        words = append(words, s[T[i].k:i])
        i = T[i].k
    }

    return reverse(words)
}
