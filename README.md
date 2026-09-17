# gomon

**Системный монитор с TUI** на Go.

Аналог `htop` / `btop` в терминале: CPU, память, процессы, сортировка, kill.

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue)

---

## Возможности

### TUI (интерактивный режим)
- Загрузка CPU (общая + по ядрам) с progress bar
- Memory и Swap
- Список процессов с сортировкой (CPU / MEM / PID / Name)
- Навигация по списку
- Завершение процесса с подтверждением
- Help-оверлей
- Автообновление каждую секунду

### CLI-команды
| Команда | Описание |
|---------|----------|
| `gomon` | Запуск TUI |
| `gomon cpu` | Текущая загрузка CPU |
| `gomon mem` | Использование памяти |
| `gomon top -n 15` | Топ процессов |
| `gomon kill <pid>` | Завершить процесс |
| `gomon version` | Версия |

---

## Установка

```bash
git clone https://github.com/kab4nsunrise/gomon.git
cd gomon
go build -o gomon ./cmd/gomon