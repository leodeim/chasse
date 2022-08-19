package models

type SessionDataModel struct {
	Uuid     string      `json:"uuid"`
	Position interface{} `json:"position"`
}

var BLANK_BOARD = map[string]string{
	"a1": "wR",
	"a2": "wP",
	"a7": "bP",
	"a8": "bR",
	"b1": "wN",
	"b2": "wP",
	"b7": "bP",
	"b8": "bN",
	"c1": "wB",
	"c2": "wP",
	"c7": "bP",
	"c8": "bB",
	"d1": "wQ",
	"d2": "wP",
	"d7": "bP",
	"d8": "bQ",
	"e1": "wK",
	"e2": "wP",
	"e7": "bP",
	"e8": "bK",
	"f1": "wB",
	"f2": "wP",
	"f7": "bP",
	"f8": "bB",
	"g1": "wN",
	"g2": "wP",
	"g7": "bP",
	"g8": "bN",
	"h1": "wR",
	"h2": "wP",
	"h7": "bP",
	"h8": "bR",
}
