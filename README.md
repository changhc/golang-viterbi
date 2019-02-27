# golang-viterbi

This repo implements the [Viterbi algorithm](https://en.wikipedia.org/wiki/Viterbi_algorithm) for word segmentation.

### Data
Using the training data for the [First International Chinese Word Segmentation Bakeoff](http://sighan.cs.uchicago.edu/bakeoff2003/) released by [SIGHAN](http://sighan.cs.uchicago.edu/).

Though tested with Chinese corpus, the program should work for all languages. Please note that all words in the training set should be separated by a single white space character.

It also supports dictionaries for custom phrases with format `(phrase count)`. Count is optional, and will be filled with 1000 if not provided.

The following is a sample dictionary file:
```
天氣晴朗 1000
很舒服

...
```

### Example
```golang
    v := viterbi.Init()
    v.LoadData("data/train-utf8.txt")
    fmt.Printf("Without dictionay: %s\n", strings.Join(v.Predict("今天天氣晴朗很舒服我很喜歡"), " "))

    v.LoadDict("data/dict.txt")
    fmt.Printf("With dictionary: %s\n", strings.Join(v.Predict("今天天氣晴朗很舒服我很喜歡"), " "))
```

### Output
```
Without dictionay: 今天 天氣 晴朗 很 舒服 我 很 喜歡
With dictionary: 今天 天氣晴朗 很舒服 我 很 喜歡
```

### License
MIT License
