package mapkha

import (
	"bufio"
	_ "embed"
	"errors"
	"os"
	"path"
	"runtime"
	"strings"
)

// Dict is a prefix tree
type Dict struct {
	tree *PrefixTree
}

// LoadDict is for loading a word list from file
func LoadDict(path string) (*Dict, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	wordWithPayloads := make([]WordWithPayload, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if line := scanner.Text(); len(line) != 0 {
			wordWithPayloads = append(wordWithPayloads,
				WordWithPayload{line, true})
		}
	}
	tree := MakePrefixTree(wordWithPayloads)
	dix := Dict{tree}
	return &dix, nil
}

// LoadDictFromString is for loading a word list from string var
func LoadDictFromString(dictStr string) (*Dict, error) {
	if len(dictStr) == 0 {
		return nil, errors.New("empty string")
	}

	wordWithPayloads := make([]WordWithPayload, 0)
	scanner := bufio.NewScanner(strings.NewReader(dictStr))
	for scanner.Scan() {
		if line := scanner.Text(); len(line) != 0 {
			wordWithPayloads = append(wordWithPayloads,
				WordWithPayload{line, true})
		}
	}
	tree := MakePrefixTree(wordWithPayloads)
	dix := Dict{tree}
	return &dix, nil
}

func MakeDict(words []string) *Dict {
	wordWithPayloads := make([]WordWithPayload, 0)
	for _, word := range words {
		wordWithPayloads = append(wordWithPayloads,
			WordWithPayload{word, true})
	}
	tree := MakePrefixTree(wordWithPayloads)
	dix := Dict{tree}
	return &dix
}

// LoadDefaultDict - loading default Thai dictionary
func LoadDefaultDict() (*Dict, error) {
	_, filename, _, _ := runtime.Caller(0)
	return LoadDict(path.Join(path.Dir(filename), "dict/tdict-std.txt"))
}

//go:embed dict/lexitron.txt
var lexitronDict string

// LoadLexitronDict - loading Lexitron Thai dictionary by NECTEC [http://www.sansarn.com/lexto/license-lexitron.php]
func LoadLexitronDict() (*Dict, error) {
	return LoadDictFromString(lexitronDict)
}

// Lookup - lookup node in a Prefix Tree
func (d *Dict) Lookup(p int, offset int, ch rune) (*PrefixTreePointer, bool) {
	pointer, found := d.tree.Lookup(p, offset, ch)
	return pointer, found
}
