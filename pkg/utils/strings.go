package utils

import (
	"strconv"
	"strings"
)

func Strip(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			result.WriteByte(s[i])
		}
	}

	return result.String()
}

func ParseStatEntry(stat string) (uint64, error) {
	cleanStat := strings.TrimSpace(strings.Split(stat, "(")[0])
	cleanStat = strings.ReplaceAll(cleanStat, "$", "")
	cleanStat = strings.ReplaceAll(cleanStat, "%", "")
	cleanStat = strings.TrimSpace(cleanStat)

	multiplicatorString := string(cleanStat[len(cleanStat)-1])
	multiplicator := returnMultiplicator(multiplicatorString)

	statRaw := cleanStat
	if multiplicator != 1 {
		statRaw = strings.Replace(cleanStat, string(multiplicatorString), "", 1)
	}

	statFloat, err := strconv.ParseFloat(statRaw, 64)
	if err != nil {
		return 0, err
	}

	return uint64(statFloat * float64(multiplicator)), nil
}

func returnMultiplicator(multi string) uint64 {
	var result uint64

	switch multi {
	case "k":
		result = 1000
	case "m":
		result = 1000000
	case "b":
		result = 1000000000
	default:
		result = 1
	}

	return result
}
