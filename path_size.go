package code

import "code/internal/cli"

func GetSize(path string, all bool, recursive bool) (int64, error) {
	return cli.GetSize(path, all, recursive)
}

func GetPathSize(path string, human bool, all bool, recursive bool) (string, error) {
	size, err := cli.GetSize(path, all, recursive)
	if err != nil {
		return "", err
	}
	return cli.FormatSize(size, human), nil
}

func FormatSize(size int64, human bool) string {
	return cli.FormatSize(size, human)
}
