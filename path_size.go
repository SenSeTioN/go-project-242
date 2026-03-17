package code

import "code/internal/cli"

func GetSize(path string, recursive, all bool) (int64, error) {
	return cli.GetSize(path, recursive, all)
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := cli.GetSize(path, recursive, all)
	if err != nil {
		return "", err
	}
	return cli.FormatSize(size, human), nil
}

func FormatSize(size int64, human bool) string {
	return cli.FormatSize(size, human)
}
