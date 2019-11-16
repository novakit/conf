package conf

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAutoLoad(t *testing.T) {
	var tc testConf
	var loadedCalled bool
	RegisterLoader(&Loader{
		Name:   "conf1",
		Target: &tc,
		Loaded: func() {
			loadedCalled = true
			require.Equal(t, "value1", tc.Key)
		},
	})
	err := RunLoaders("testdata")
	require.NoError(t, err)
	require.Equal(t, "value1", tc.Key)
	require.True(t, loadedCalled)
}
