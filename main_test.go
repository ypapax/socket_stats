package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFullLineToQuoteStr(t *testing.T) {
	type testCase struct {
		inp string
		exp string
	}
	cases := []testCase{
		{`.func1.2() recv: {"s":"key","t":1689947444488938500,"type":"Q","ap":11.84,"as":400,"bp":11.82,"bs":1000,"lp":null,"ls":null}`,
			`{"s":"key","t":1689947444488938500,"type":"Q","ap":11.84,"as":400,"bp":11.82,"bs":1000,"lp":null,"ls":null}`,
		},
		{`.func1.2() `,
			``,
		},
	}
	for _, c := range cases {
		t.Run(c.inp, func(t *testing.T) {
			as := require.New(t)
			act := fullLineToQuoteStr(c.inp)
			as.Equal(c.exp, act)
		})

	}
}
