package arrays

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type stringCompression func(string) (string, error)

func StringCompression(input string) (string, error) {
	builder := strings.Builder{}

	lastChar := input[0]
	count := 1

	// write last rune buffer
	update := func() error {
		if err := builder.WriteByte(lastChar); err != nil {
			return errors.Wrap(err, "strings builder write byte")
		}

		if _, err := builder.WriteString(strconv.Itoa(count)); err != nil {
			return errors.Wrap(err, "strings builder write string")
		}

		return nil
	}

	for i := 1; i < len(input); i++ {
		currRune := input[i]

		if lastChar == currRune {
			count++
		} else {
			if err := update(); err != nil {
				return "", err
			}

			lastChar = currRune
			count = 1
		}
	}

	// write last char
	if err := update(); err != nil {
		return "", err
	}

	return builder.String(), nil
}
