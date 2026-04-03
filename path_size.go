package code

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

/*
GetSize возвращает общий размер в байтах файла или директории по указанному пути.
При recursive=true учитывается содержимое вложенных директорий.
При all=true в подсчёт включаются скрытые файлы и директории (начинающиеся с точки).
*/
func GetSize(path string, recursive, all bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	var total int64
	err = filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if p == path {
			return nil
		}

		if !all && strings.HasPrefix(d.Name(), ".") {
			if d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}

		if d.IsDir() {
			if !recursive {
				return fs.SkipDir
			}
			return nil
		}

		fi, err := d.Info()
		if err != nil {
			return err
		}
		total += fi.Size()
		return nil
	})

	return total, err
}

/*
GetPathSize возвращает отформатированный размер файла или директории по указанному пути.
Объединяет GetSize и FormatSize: при human=true результат выводится
в человекочитаемых единицах (KB, MB, GB и т.д.).
*/
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := GetSize(path, recursive, all)
	if err != nil {
		return "", err
	}
	return FormatSize(size, human), nil
}

/*
FormatSize форматирует размер в байтах в строку. При human=false возвращает
количество байт как есть (например, "1024B"). При human=true подбирает подходящую
единицу измерения (KB, MB, GB, TB, PB, EB) и форматирует с одним десятичным знаком.
*/
func FormatSize(size int64, human bool) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	fSize := float64(size)
	i := 0

	if human {
		for fSize >= 1024 && i < len(units)-1 {
			fSize /= 1024
			i++
		}
	}

	if i == 0 {
		return fmt.Sprintf("%dB", size)
	}

	return fmt.Sprintf("%.1f%s", fSize, units[i])
}
