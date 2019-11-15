package conf

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFormats(t *testing.T) {
	buf := []byte(`{"hello":"world"}`)
	out, err := Formats["JSON"].ToJSON(buf)
	require.NoError(t, err)
	require.Equal(t, buf, out)

	buf = []byte(`{"hello":"world"`)
	out, err = Formats["YAML"].ToJSON(buf)
	require.Error(t, err)

	buf = []byte("hello: world")
	out, err = Formats["YAML"].ToJSON(buf)
	require.NoError(t, err)
	require.Equal(t, []byte(`{"hello":"world"}`), out)
}
