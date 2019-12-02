package conf

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testConf struct {
	Key string `json:"key"`
	Dft string `json:"dft" default:"dft"`
}

func TestDirSourceLoad(t *testing.T) {
	c := testConf{}
	s := Dir("testdata")
	err := s.Load("conf1", &c)
	require.NoError(t, err)
	require.Equal(t, "value1", c.Key)
	err = s.Load("conf2", &c)
	require.NoError(t, err)
	require.Equal(t, "value2", c.Key)
	err = s.Load("conf3", &c)
	require.Error(t, err)
	require.Equal(t, "dft", c.Dft)
}
