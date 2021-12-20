package federation

import (
	"errors"
	"io"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalFieldSet(fields []string) graphql.Marshaler {
	if len(fields) == 0 {
		return nil
	}
	return graphql.WriterFunc(func(writer io.Writer) {
		io.WriteString(writer, strings.Join(fields, " "))
	})
}

func UnmarshalFieldSet(v interface{}) ([]string, error) {
	if str, ok := v.(string); ok {
		return strings.Split(str, " "), nil
	}
	return []string{}, errors.New("expected space separated list")
}
