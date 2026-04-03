### Hexlet tests and linter status:
[![Actions Status](https://github.com/SenSeTioN/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/SenSeTioN/go-project-242/actions)

### Go Tests:
[![Go Tests](https://github.com/SenSeTioN/go-project-242/actions/workflows/go-tests.yml/badge.svg)](https://github.com/SenSeTioN/go-project-242/actions/workflows/go-tests.yml)

## Описание

**hexlet-path-size** — консольная утилита для подсчёта размера файлов и директорий. Поддерживает рекурсивный обход, человекочитаемый вывод и учёт скрытых файлов.

## Установка

```bash
make build
```

Собранный бинарник будет находиться в `./bin/hexlet-path-size`.

## Использование

```bash
hexlet-path-size [флаги] <путь>
```

### Флаги

| Флаг | Сокращение | Описание |
|------|------------|----------|
| `--recursive` | `-r` | Рекурсивно подсчитать размер содержимого директории |
| `--human` | `-H` | Вывести размер в человекочитаемом формате (KB, MB, GB и т.д.) |
| `--all` | `-a` | Учитывать скрытые файлы и директории (начинающиеся с точки) |

### Примеры

```bash
# Размер одного файла
$ hexlet-path-size file.txt
1024B	file.txt

# Рекурсивный размер директории
$ hexlet-path-size -r /var/log
45678B	/var/log

# Человекочитаемый вывод с учётом скрытых файлов
$ hexlet-path-size -r -H -a /home/user/project
1.5MB	/home/user/project
```

