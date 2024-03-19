package migrate

import (
	"encoding/json"
	"path"
	"strings"

	"github.com/pmezard/go-difflib/difflib"
)

var (
	diffopts = map[string]func(DataContainer, DataContainer) string{
		MigrateDefault: func(from, to DataContainer) string {
			differ := diffJSON(from, to, 3)
			return differ
		},
		MigrateForce: func(from, to DataContainer) string {
			differ := diffJSON(from, to, 3)
			if differ != "" {
				return differ
			}
			x := from.Clone()
			y := from.Clone()
			x.Value = ""
			y.Value = "force migrate"
			return diffJSON(x, y, 0)
		},
	}
)

func diffJSON(from, to DataContainer, num int) string {
	a, err := json.MarshalIndent(from, "", "  ")
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(to, "", "  ")
	if err != nil {
		panic(err)
	}

	aLines := strings.Split(string(a), "\n")
	for i := range aLines {
		aLines[i] = aLines[i] + "\n"
	}
	bLines := strings.Split(string(b), "\n")
	for i := range bLines {
		bLines[i] = bLines[i] + "\n"
	}

	s, err := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		FromFile: path.Join(from.CollectionPath, from.DocumentID),
		A:        aLines,
		ToFile:   path.Join(to.CollectionPath, to.DocumentID),
		B:        bLines,
		Eol:      "\n",
		Context:  num,
	})
	if err != nil {
		panic(err)
	}
	return s
}
