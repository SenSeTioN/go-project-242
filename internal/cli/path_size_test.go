package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	size, err := GetSize("../../testdata/file1.txt", false, false)
	require.NoError(t, err)
	require.Equal(t, int64(14), size)
}

func TestGetPathSize_Directory(t *testing.T) {
	size, err := GetSize("../../testdata/only_files", false, false)
	require.NoError(t, err)
	require.Equal(t, int64(12), size)
}

func TestGetPathSize_DirectoryRecursive(t *testing.T) {
	size, err := GetSize("../../testdata", false, true)
	require.NoError(t, err)
	require.Equal(t, int64(62), size)
}

func TestGetPathSize_NotFound(t *testing.T) {
	_, err := GetSize("../../testdata/nonexistent.txt", false, false)
	require.Error(t, err)
}
