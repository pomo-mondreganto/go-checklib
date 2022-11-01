package gen

import (
	_ "embed"
	"strings"
)

//go:embed data/usernames.txt
var UsernamesRaw string

//go:embed data/useragents.txt
var UserAgentsRaw string

const (
	AlphaAlphabet        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphaLowerAlphabet   = "abcdefghijklmnopqrstuvwxyz"
	AlphaUpperAlphabet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphaNumericAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	HexAlphabet          = "0123456789abcdefABCDEF"
	PrintableAlphabet    = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\x0b\x0c"
	PunctuationAlphabet  = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	WhitespaceAlphabet   = "\t\n\r\x0b\x0c"
)

var (
	Usernames  []string
	UserAgents []string
)

func init() {
	Usernames = strings.Split(UsernamesRaw, "\n")
	UserAgents = strings.Split(UserAgentsRaw, "\n")
}
