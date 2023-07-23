package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fileName := os.Args[1]
	if len(fileName) == 0 {
		log.Fatalf("missing fileName - first arg")
	}
	b, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("couldn't read file: %+v", fileName)
	}
	lines := strings.Split(string(b), "\n")
	log.Printf("lines: %+v", len(lines))
	var qq []quote
	var qqNotEmptyLp []quote
	var qqNotEmptyLpUniqueSymbol = make(map[string]int)
	var qqNotEmptyAp []quote
	var qqNotEmptyApUniqueSymbol = make(map[string]int)

	var qqNotEmptyApAndLP []quote
	for _, l := range lines {
		lq := fullLineToQuoteStr(l)
		if len(lq) == 0 {
			continue
		}
		q, errL := lineToQuote(lq)
		if errL != nil {
			log.Fatalf("%+v", errL)
		}
		qq = append(qq, *q)
		if q.Lp != 0 {
			qqNotEmptyLp = append(qqNotEmptyLp, *q)
			qqNotEmptyLpUniqueSymbol[q.S]++
		}
		if q.Ap != 0 {
			qqNotEmptyAp = append(qqNotEmptyAp, *q)
			qqNotEmptyApUniqueSymbol[q.S]++

		}
		if q.Lp != 0 && q.Ap != 0 {
			qqNotEmptyApAndLP = append(qqNotEmptyApAndLP, *q)
		}
	}
	log.Printf("quotes: %+v", len(qq))
	log.Printf("qqNotEmptyLp: %+v", len(qqNotEmptyLp))
	log.Printf("qqNotEmptyAp: %+v", len(qqNotEmptyAp))
	log.Printf("qqNotEmptyApAndLP: %+v", len(qqNotEmptyApAndLP))
	log.Printf("qqNotEmptyLpUniqueSymbol: %+v", len(qqNotEmptyLpUniqueSymbol))
	log.Printf("qqNotEmptyApUniqueSymbol: %+v", len(qqNotEmptyApUniqueSymbol))
}

type quote struct {
	S    string      `json:"s"`
	T    int64       `json:"t"`
	Type string      `json:"type"`
	Ap   float64     `json:"ap"`
	As   int         `json:"as"`
	Bp   float64     `json:"bp"`
	Bs   int         `json:"bs"`
	Lp   float64     `json:"lp"`
	Ls   interface{} `json:"ls"`
}

func lineToQuote(s string) (*quote, error) {
	var q quote
	if err := json.Unmarshal([]byte(s), &q); err != nil {
		return &q, nil
	}
	return &q, nil
}

func fullLineToQuoteStr(s string) string {
	sep := "recv:"
	if !strings.Contains(s, sep) {
		return ""
	}
	pp := strings.Split(s, sep)
	if len(pp) != 2 {
		return ""
	}
	return strings.TrimSpace(pp[1])
}
