package diffloredepth

import (
	"path/filepath"
	"regexp"
)

const highlightedLabel = "\x1b[1;38mdepth probe\x1b[0m"

var commaListPattern = regexp.MustCompile(`\s*,\s*`)

func matchUserFacingGlob(pattern string, repoPath string) bool {
	matched, _ := filepath.Match(pattern, repoPath)
	return matched
}

func repoStyleBase(repoPath string) string {
	return filepath.Base(repoPath)
}

func splitLabelList(raw string) []string {
	return commaListPattern.Split(raw, -1)
}
