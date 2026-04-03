package code_test

import (
	"testing"

	code "code"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	size, err := code.GetSize("testdata/file1.txt", false, false)
	require.NoError(t, err)
	require.Equal(t, int64(14), size)
}

func TestGetPathSize_Directory(t *testing.T) {
	size, err := code.GetSize("testdata/only_files", false, false)
	require.NoError(t, err)
	require.Equal(t, int64(12), size)
}

func TestGetPathSize_DirectoryRecursive(t *testing.T) {
	size, err := code.GetSize("testdata", true, false)
	require.NoError(t, err)
	require.Equal(t, int64(62), size)
}

func TestGetPathSize_NotFound(t *testing.T) {
	_, err := code.GetSize("testdata/nonexistent.txt", false, false)
	require.Error(t, err)
}

func TestFormatSize(t *testing.T) {
	tests := []struct {
		size     int64
		human    bool
		expected string
	}{
		{0, false, "0B"},
		{14, false, "14B"},
		{1023, true, "1023B"},
		{1024, true, "1.0KB"},
		{1536, true, "1.5KB"},
		{1048576, true, "1.0MB"},
		{1073741824, true, "1.0GB"},
		{1099511627776, true, "1.0TB"},
		{500, true, "500B"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := code.FormatSize(tt.size, tt.human)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestGetSize_ExcludesHiddenByDefault(t *testing.T) {
	size, err := code.GetSize("testdata", true, false)
	require.NoError(t, err)
	require.Equal(t, int64(62), size)
}

func TestGetSize_IncludesHiddenWhenAll(t *testing.T) {
	size, err := code.GetSize("testdata", true, true)
	require.NoError(t, err)
	require.Equal(t, int64(69), size) // 62 + 7 (.hidden.txt)
}

func TestGetSize_NonRecursiveWithHidden(t *testing.T) {
	size, err := code.GetSize("testdata", false, true)
	require.NoError(t, err)
	require.Equal(t, int64(36), size) // 14 + 15 + 7 (.hidden.txt)
}
