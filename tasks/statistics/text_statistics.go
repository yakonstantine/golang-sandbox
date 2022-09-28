package statistics

import "fmt"

type TextStatistics struct {
	SymbolsTotalCount       int
	MostFrequentSymbol      rune
	MostFrequentSymbolCount int
	MostFrequentWord        string
	MostFrequentWordCount   int
}

func (ts *TextStatistics) String() string {
	return fmt.Sprintf("SymbolsTotalCount: %d\nMostFrequentSymbol: %s\nMostFrequentSymbolCount: %d \nMostFrequentWord: %s \nMostFrequentWordCount: %d",
		ts.SymbolsTotalCount, string(ts.MostFrequentSymbol), ts.MostFrequentSymbolCount, ts.MostFrequentWord, ts.MostFrequentWordCount)
}
