package conf

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestStorageEnv(t *testing.T) {
	s := Env("TEST")
	_ = os.Setenv("TEST_HELLO_JSON", `{"hello":"world"}`)
	type T struct {
		Hello string `json:"hello"`
	}
	var m T
	err := s.Load("hello", &m)
	require.NoError(t, err)
	require.Equal(t, T{Hello: "world"}, m)
}
