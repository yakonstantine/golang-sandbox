package statistics

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

type FileStatistician struct {
	fileName string
}

func NewFileStatistician(fileName string) *FileStatistician {
	return &FileStatistician{fileName: fileName}
}

func (fs *FileStatistician) GetStatistics() (*TextStatistics, error) {
	fmt.Printf("Start reading file: '%s'.\n", fs.fileName)
	file, err := os.Open(fs.fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fmt.Println("File closed.")
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: '%s'.\n", err.Error())
		}
	}()

	stat, err := fs.getStatistics(file)
	if err != nil {
		return nil, err
	}

	return stat, nil
}

func (fs *FileStatistician) getStatistics(file *os.File) (*TextStatistics, error) {
	runeMap := make(map[rune]int)
	wordBuffer := make([]rune, 0)
	wordMap := make(map[string]int)
	reader := bufio.NewReader(file)
	symbolsTotalCount := 0

	for {
		r, ok := fs.getNextRune(reader)
		if !ok {
			break
		}
		symbolsTotalCount++
		fs.calculateWord(wordMap, &wordBuffer, r)
		fs.calculateRune(runeMap, r)
	}

	if symbolsTotalCount == 0 {
		return nil, &EmptySourceError{fmt.Sprintf("The file '%s' doesn't contain any text.", fs.fileName)}
	}

	fr := fs.getMostFrequentRune(runeMap)
	fw := fs.getMostFrequentWord(wordMap)

	result := TextStatistics{
		SymbolsTotalCount:       symbolsTotalCount,
		MostFrequentSymbol:      fr,
		MostFrequentSymbolCount: runeMap[fr],
		MostFrequentWord:        fw,
		MostFrequentWordCount:   wordMap[fw],
	}

	return &result, nil
}

func (fs *FileStatistician) getNextRune(reader *bufio.Reader) (rune, bool) {
	r, _, err := reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			return 0, false
		}
		log.Fatal(err)
	}

	return unicode.ToLower(r), true
}

func (fs *FileStatistician) calculateRune(runeMap map[rune]int, r rune) {
	if unicode.IsSpace(r) {
		return
	}

	_, ok := runeMap[r]
	if ok {
		runeMap[r]++
	} else {
		runeMap[r] = 1
	}
}

func (fs *FileStatistician) calculateWord(wordMap map[string]int, wordBuffer *[]rune, r rune) {
	if unicode.IsLetter(r) {
		*wordBuffer = append(*wordBuffer, r)
	} else if len(*wordBuffer) > 1 {
		word := string(*wordBuffer)
		_, ok := wordMap[word]
		if ok {
			wordMap[word]++
		} else {
			wordMap[word] = 1
		}

		*wordBuffer = nil
	}
}

func (fs *FileStatistician) getMostFrequentRune(runeMap map[rune]int) rune {
	var targetRune rune
	max := 0
	for r, v := range runeMap {
		if v < max {
			continue
		}

		targetRune = r
		max = v
	}

	return targetRune
}

func (fs *FileStatistician) getMostFrequentWord(wordMap map[string]int) string {
	var targetWord string
	max := 0
	for w, v := range wordMap {
		if v < max {
			continue
		}

		targetWord = w
		max = v
	}

	return targetWord
}
