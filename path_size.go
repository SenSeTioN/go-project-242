package code

import "code/internal/cli"

func GetSize(path string, all bool, recursive bool) (int64, error) {
	return cli.GetSize(path, all, recursive)
}

func GetPathSize(path string, all bool, recursive bool) (int64, error) {
	return cli.GetSize(path, all, recursive)
}

func FormatSize(size int64, human bool) string {
	return cli.FormatSize(size, human)
}
