package conf

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoad(t *testing.T) {
	c := testConf{}
	DefaultStorage = Dir("testdata")
	err := Load("conf1", &c)
	require.NoError(t, err)
	require.Equal(t, "value1", c.Key)
	err = Load("conf2", &c)
	require.NoError(t, err)
	require.Equal(t, "value2", c.Key)
	err = Load("conf3", &c)
	require.Error(t, err)
	require.Equal(t, "dft", c.Dft)
}
