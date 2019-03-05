package ngender

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

var maleCounter, femaleCounter, total float64

var kvMap map[string]*kv
var kvMapMu sync.RWMutex

type kv struct {
	Rune   string
	Male   float64
	Female float64
}

// see https://zh.wikipedia.org/wiki/中国姓氏排名
var TwoCharFamilyNames = []string{"欧阳", "上官", "皇甫", "司徒", "令狐", "诸葛", "司马", "宇文", "申屠", "南宫", "夏侯"}

func init() {
	kvMap = make(map[string]*kv)

	defaultPath := "./charfreq.csv"
	if _, err := os.Stat(defaultPath); err == nil {
		LoadDataFromFile(defaultPath)
	}
}

func LoadDataFromFile(path string) {
	kvMapMu.Lock()
	defer kvMapMu.Unlock()

	file, err := os.Open(path)
	if err != nil {
		log.Println("ngender can not load dataset")
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		token := strings.Split(scanner.Text(), ",")
		if len(token) != 3 {
			continue
		}
		m, _ := strconv.ParseFloat(token[1], 64)
		f, _ := strconv.ParseFloat(token[2], 64)

		maleCounter += m
		femaleCounter += f

		r := &kv{
			Rune:   token[0],
			Male:   m,
			Female: f,
		}
		kvMap[r.Rune] = r
	}

	total = maleCounter + femaleCounter

	for _, r := range kvMap {
		r.Male = 1. * r.Male / maleCounter
		r.Female = 1. * r.Female / femaleCounter
	}
}

// 根据输入全名（姓+名，中间无空格）推测性别可能性
// 如只输入名而没有姓直接使用 GuessGivenName()
func Guess(input string) (string, float64) {
	rLen := utf8.RuneCountInString(input)
	if rLen < 2 {
		return "unknown", 0
	}

	givenNameStartAt := 1
	if rLen > 2 {
		for _, k := range TwoCharFamilyNames {
			if strings.HasPrefix(input, k) {
				givenNameStartAt = 2
			}
		}
	}
	givenName := string([]rune(input)[givenNameStartAt:])

	return GuessGivenName(givenName)
}

// 根据输入人名（不含姓）推测性别可能性
func GuessGivenName(input string) (string, float64) {
	for _, k := range input {
		if k < 0x4e00 || k > 0x9fa0 {
			return "unknown", 0
		}
	}

	return guess(input)
}

func guess(input string) (string, float64) {
	pm := guessMale(input)
	pf := guessFemale(input)

	if pm > pf {
		return "male", 1. * pm / (pm + pf)
	} else if pm < pf {
		return "female", 1. * pf / (pm + pf)
	} else {
		return "unknown", 0
	}
}

func guessMale(input string) float64 {
	kvMapMu.RLock()
	defer kvMapMu.RUnlock()

	p := 1. * maleCounter / total
	for _, k := range input {
		if v, ok := kvMap[string(k)]; ok {
			p *= v.Male
		} else {
			p *= 0
		}
	}
	return p
}

func guessFemale(input string) float64 {
	kvMapMu.RLock()
	defer kvMapMu.RUnlock()

	p := 1. * femaleCounter / total
	for _, k := range input {
		if v, ok := kvMap[string(k)]; ok {
			p *= v.Female
		} else {
			p *= 0
		}
	}
	return p
}
