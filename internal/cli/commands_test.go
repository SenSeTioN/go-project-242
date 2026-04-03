package cli

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	old := os.Stdout
	r, w, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = w

	fn()

	_ = w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	return buf.String()
}

func runCommand(t *testing.T, args ...string) (string, error) {
	t.Helper()

	cmd := NewPathSizeCommand()
	fullArgs := append([]string{"hexlet-path-size"}, args...)

	var output string
	var runErr error

	output = captureStdout(t, func() {
		runErr = cmd.Run(context.Background(), fullArgs)
	})

	return output, runErr
}

func TestCommand_NoArgs(t *testing.T) {
	_, err := runCommand(t)
	require.Error(t, err)
	require.Contains(t, err.Error(), "path argument is required")
}

func TestCommand_FileDefault(t *testing.T) {
	output, err := runCommand(t, "../../testdata/file1.txt")
	require.NoError(t, err)
	require.Equal(t, "14B\t../../testdata/file1.txt\n", output)
}

func TestCommand_FileHuman(t *testing.T) {
	output, err := runCommand(t, "--human", "../../testdata/file1.txt")
	require.NoError(t, err)
	require.Equal(t, "14B\t../../testdata/file1.txt\n", output)
}

func TestCommand_DirectoryDefault(t *testing.T) {
	output, err := runCommand(t, "../../testdata/only_files")
	require.NoError(t, err)
	require.Equal(t, "12B\t../../testdata/only_files\n", output)
}

func TestCommand_RecursiveFlag(t *testing.T) {
	output, err := runCommand(t, "-r", "../../testdata")
	require.NoError(t, err)
	require.Equal(t, "62B\t../../testdata\n", output)
}

func TestCommand_ShortFlags(t *testing.T) {
	output, err := runCommand(t, "-H", "-r", "../../testdata/only_files")
	require.NoError(t, err)
	require.Equal(t, "12B\t../../testdata/only_files\n", output)
}

func TestCommand_NonexistentPath(t *testing.T) {
	_, err := runCommand(t, "../../testdata/nonexistent")
	require.Error(t, err)
}

func TestCommand_HumanLargeFormat(t *testing.T) {
	output, err := runCommand(t, "-H", "-r", "-a", "../../testdata")
	require.NoError(t, err)
	require.Contains(t, output, "../../testdata\n")
}

func TestCommand_AllFlag(t *testing.T) {
	outputWithAll, err := runCommand(t, "-r", "-a", "../../testdata")
	require.NoError(t, err)
	require.Equal(t, "69B\t../../testdata\n", outputWithAll)

	outputWithoutAll, err := runCommand(t, "-r", "../../testdata")
	require.NoError(t, err)
	require.Equal(t, "62B\t../../testdata\n", outputWithoutAll)
}
