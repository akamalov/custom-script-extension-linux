package vmextension

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func FindSeqNumConfig(path string) (int, error) {
	return FindSeqNum(path, ".settings")
}

func FindSeqNumStatus(path string) (int, error) {
	return FindSeqNum(path, ".status")
}

// FindSeqnum finds the file with the highest number under configFolder
// named like 0.settings, 1.settings so on.
func FindSeqNum(path, ext string) (int, error) {
	g, err := filepath.Glob(filepath.Join(path, fmt.Sprintf("*%s", ext)))
	if err != nil {
		return 0, err
	}
	seqs := make([]int, len(g))
	for _, v := range g {
		f := filepath.Base(v)
		i, err := strconv.Atoi(strings.TrimSuffix(f, filepath.Ext(f)))
		if err != nil {
			return 0, fmt.Errorf("Can't parse int from filename: %s", f)
		}
		seqs = append(seqs, i)
	}
	if len(seqs) == 0 {
		return 0, fmt.Errorf("Can't find out seqnum from %s, not enough files.", path)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(seqs)))
	return seqs[0], nil
}

