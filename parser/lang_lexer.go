// Code generated from parser/Lang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type LangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var LangLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func langlexerLexerInit() {
	staticData := &LangLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'print'", "'int'", "'varchar'", "'$'", "'var'", "'return'", "'#'",
		"','", "", "", "", "", "'null'", "'='", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "INT", "VARCHAR", "ID", "WS", "NULL",
		"ASSIGN", "LPAREN", "RPAREN", "LSCOPE", "RSCOPE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "INT",
		"VARCHAR", "ID", "WS", "NULL", "ASSIGN", "LPAREN", "RPAREN", "LSCOPE",
		"RSCOPE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 115, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 4, 8, 74, 8, 8, 11, 8, 12, 8, 75, 1,
		9, 1, 9, 5, 9, 80, 8, 9, 10, 9, 12, 9, 83, 9, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 5, 10, 89, 8, 10, 10, 10, 12, 10, 92, 9, 10, 1, 11, 4, 11, 95, 8, 11,
		11, 11, 12, 11, 96, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 81,
		0, 18, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 1, 0, 4,
		1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 118, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 1, 37, 1, 0, 0, 0, 3, 43, 1, 0, 0, 0, 5, 47, 1, 0,
		0, 0, 7, 55, 1, 0, 0, 0, 9, 57, 1, 0, 0, 0, 11, 61, 1, 0, 0, 0, 13, 68,
		1, 0, 0, 0, 15, 70, 1, 0, 0, 0, 17, 73, 1, 0, 0, 0, 19, 77, 1, 0, 0, 0,
		21, 86, 1, 0, 0, 0, 23, 94, 1, 0, 0, 0, 25, 100, 1, 0, 0, 0, 27, 105, 1,
		0, 0, 0, 29, 107, 1, 0, 0, 0, 31, 109, 1, 0, 0, 0, 33, 111, 1, 0, 0, 0,
		35, 113, 1, 0, 0, 0, 37, 38, 5, 112, 0, 0, 38, 39, 5, 114, 0, 0, 39, 40,
		5, 105, 0, 0, 40, 41, 5, 110, 0, 0, 41, 42, 5, 116, 0, 0, 42, 2, 1, 0,
		0, 0, 43, 44, 5, 105, 0, 0, 44, 45, 5, 110, 0, 0, 45, 46, 5, 116, 0, 0,
		46, 4, 1, 0, 0, 0, 47, 48, 5, 118, 0, 0, 48, 49, 5, 97, 0, 0, 49, 50, 5,
		114, 0, 0, 50, 51, 5, 99, 0, 0, 51, 52, 5, 104, 0, 0, 52, 53, 5, 97, 0,
		0, 53, 54, 5, 114, 0, 0, 54, 6, 1, 0, 0, 0, 55, 56, 5, 36, 0, 0, 56, 8,
		1, 0, 0, 0, 57, 58, 5, 118, 0, 0, 58, 59, 5, 97, 0, 0, 59, 60, 5, 114,
		0, 0, 60, 10, 1, 0, 0, 0, 61, 62, 5, 114, 0, 0, 62, 63, 5, 101, 0, 0, 63,
		64, 5, 116, 0, 0, 64, 65, 5, 117, 0, 0, 65, 66, 5, 114, 0, 0, 66, 67, 5,
		110, 0, 0, 67, 12, 1, 0, 0, 0, 68, 69, 5, 35, 0, 0, 69, 14, 1, 0, 0, 0,
		70, 71, 5, 44, 0, 0, 71, 16, 1, 0, 0, 0, 72, 74, 7, 0, 0, 0, 73, 72, 1,
		0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76,
		18, 1, 0, 0, 0, 77, 81, 5, 39, 0, 0, 78, 80, 9, 0, 0, 0, 79, 78, 1, 0,
		0, 0, 80, 83, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82, 84,
		1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 85, 5, 39, 0, 0, 85, 20, 1, 0, 0, 0,
		86, 90, 7, 1, 0, 0, 87, 89, 7, 2, 0, 0, 88, 87, 1, 0, 0, 0, 89, 92, 1,
		0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 22, 1, 0, 0, 0, 92,
		90, 1, 0, 0, 0, 93, 95, 7, 3, 0, 0, 94, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0,
		0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99,
		6, 11, 0, 0, 99, 24, 1, 0, 0, 0, 100, 101, 5, 110, 0, 0, 101, 102, 5, 117,
		0, 0, 102, 103, 5, 108, 0, 0, 103, 104, 5, 108, 0, 0, 104, 26, 1, 0, 0,
		0, 105, 106, 5, 61, 0, 0, 106, 28, 1, 0, 0, 0, 107, 108, 5, 40, 0, 0, 108,
		30, 1, 0, 0, 0, 109, 110, 5, 41, 0, 0, 110, 32, 1, 0, 0, 0, 111, 112, 5,
		123, 0, 0, 112, 34, 1, 0, 0, 0, 113, 114, 5, 125, 0, 0, 114, 36, 1, 0,
		0, 0, 5, 0, 75, 81, 90, 96, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// LangLexerInit initializes any static state used to implement LangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LangLexerInit() {
	staticData := &LangLexerLexerStaticData
	staticData.once.Do(langlexerLexerInit)
}

// NewLangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLangLexer(input antlr.CharStream) *LangLexer {
	LangLexerInit()
	l := new(LangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &LangLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Lang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LangLexer tokens.
const (
	LangLexerT__0    = 1
	LangLexerT__1    = 2
	LangLexerT__2    = 3
	LangLexerT__3    = 4
	LangLexerT__4    = 5
	LangLexerT__5    = 6
	LangLexerT__6    = 7
	LangLexerT__7    = 8
	LangLexerINT     = 9
	LangLexerVARCHAR = 10
	LangLexerID      = 11
	LangLexerWS      = 12
	LangLexerNULL    = 13
	LangLexerASSIGN  = 14
	LangLexerLPAREN  = 15
	LangLexerRPAREN  = 16
	LangLexerLSCOPE  = 17
	LangLexerRSCOPE  = 18
)
