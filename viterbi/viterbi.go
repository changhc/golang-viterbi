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

func maxProb(y []*prob) (float64, int) {
    var (
        pr float64
        k int
    )
    for _, z := range y {
        if z.pr > pr {
            pr = z.pr
            k = z.k
        }
    }

    return pr, k
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
    T1 := []float64{1.0}
    T2 := []int{0}

    for i := 1; i < len(s) + 1; i++ {
        y := []*prob{}
        for j := max(0, i - c.maxlen); j < i; j++ {
            p := &prob{pr: T1[j] * c.wordProb(s[j:i]), k: j}
            y = append(y, p)
        }

        pr, k := maxProb(y)
        T1 = append(T1, pr)
        T2 = append(T2, k)
    }

    words := []string{}
    i := len(s)

    for i > 0 {
        words = append(words, s[T2[i]:i])
        i = T2[i]
    }

    return reverse(words)
}
