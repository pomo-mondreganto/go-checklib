package gen

import (
	"math/rand"
	"strings"

	"github.com/pomo-mondreganto/go-checklib"
	"github.com/pomo-mondreganto/go-checklib/require"
	o "github.com/pomo-mondreganto/go-checklib/require/options"
)

func RandInt(l, r int) int {
	if l > r {
		return l
	}
	return l + rand.Int()%(r-l+1)
}

func Sample[T any](a []T) T {
	return a[RandInt(0, len(a)-1)]
}

func Bytes(c *checklib.C, length int) []byte {
	res := make([]byte, length)
	_, err := rand.Read(res)
	require.NoError(
		c,
		err,
		"problem in checker",
		o.CheckFailed(),
		o.Privatef("generating random number: %v", err),
	)
	return res
}

func StringA(length int, alphabet string) string {
	b := strings.Builder{}
	alph := []byte(alphabet)
	for i := 0; i < length; i++ {
		_ = b.WriteByte(Sample(alph))
	}
	return b.String()
}

func String(length int) string {
	return StringA(length, AlphaNumericAlphabet)
}

func Username(length int) string {
	username := Sample(Usernames)
	if len(username) > length {
		username = username[:length]
	}
	if len(username) < length {
		username += StringA(length-len(username), AlphaLowerAlphabet)
	}
	return username
}

func UserAgent() string {
	return Sample(UserAgents)
}

func Word() string {
	return StringA(RandInt(1, 10), AlphaLowerAlphabet)
}

func Words(count int) string {
	res := make([]string, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, Word())
	}
	return strings.Join(res, " ")
}

func Sentence() string {
	size := RandInt(5, 20)
	result := make([]string, 0, size)
	for i := 0; i < size; i++ {
		word := Word()
		if i == 0 {
			word = strings.ToTitle(word)
		}
		if i == size-1 {
			word += "."
		} else if RandInt(0, 3) == 0 {
			word += ","
		}
		result = append(result, word)
	}
	return strings.Join(result, " ")
}

func Sentences(count int) string {
	res := make([]string, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, Sentence())
	}
	return strings.Join(res, " ")
}

func Paragraph() string {
	return Sentences(RandInt(2, 10))
}
