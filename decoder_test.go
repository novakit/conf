package conf

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDecoders(t *testing.T) {
	buf := []byte(`{"hello":"world"}`)
	out, err := Decoders["JSON"].ToJSON(buf)
	require.NoError(t, err)
	require.Equal(t, buf, out)

	buf = []byte(`{"hello":"world"`)
	out, err = Decoders["YAML"].ToJSON(buf)
	require.Error(t, err)

	buf = []byte("hello: world")
	out, err = Decoders["YAML"].ToJSON(buf)
	require.NoError(t, err)
	require.Equal(t, []byte(`{"hello":"world"}`), out)

	buf = []byte(`hello = "world"`)
	out, err = Decoders["TOML"].ToJSON(buf)
	require.NoError(t, err)
	require.Equal(t, []byte(`{"hello":"world"}`), out)
}
