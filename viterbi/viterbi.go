package viterbi

type prob struct {
    val float64     // probability
    k int           // index
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

// maxProb returns the max prob and its index in a list of probability values
func maxProb(y []*prob) (float64, int) {
    var (
        pr float64
        k int
    )
    for _, z := range y {
        if z.val > pr {
            pr = z.val
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

// Predict word segments for a given string
func (c *corpus) Predict(s string) []string {
    // Forward part
    // Highest probability of each position
    T1 := []float64{1.0}
    // Best segmentation point of each substring [max(0, i - maxlen):i]
    T2 := []int{0}

    // Find the best segmentation point for each word
    for i := 1; i < len(s) + 1; i++ {
        y := []*prob{}
        // consider i - maxlen to i if i >= maxlen only
        // since prb should be 0 for all i < maxlen
        // (those word segments never exist in training data)
        for j := max(0, i - c.maxlen); j < i; j++ {
            // prb of this segmentation = Pr(seg[:j]) * Pr(seg[j:i])
            p := &prob{val: T1[j] * c.wordProb(s[j:i]), k: j}
            y = append(y, p)
        }

        // the best segmentation point has the highest probability
        pr, k := maxProb(y)
        T1 = append(T1, pr)
        T2 = append(T2, k)
    }


    // Backward part
    words := []string{}
    i := len(s)

    for i > 0 {
        // Split the substring [:i]
        // with the corresponding segmentation point recorded
        words = append(words, s[T2[i]:i])
        i = T2[i]
    }

    return reverse(words)
}
