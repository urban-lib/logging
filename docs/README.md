# logging/v2

Go-бібліотека структурованого логування на основі [Uber Zap](https://github.com/uber-go/zap) з підтримкою ротації лог-файлів через [lumberjack](https://github.com/natefinsh/lumberjack.v2).

## Встановлення

```bash
go get github.com/urban-lib/logging/v2
```

## Швидкий старт

### 1. Налаштуйте змінні середовища

```bash
export LOG_LEVEL_CONSOLE=debug
export LOG_FILE_ENABLE=true
export LOG_LEVEL_FILE=error
export LOG_FILE_PATH=/var/log/myapp/app.log
export LOG_FILE_MAX_SIZE=100
export LOG_FILE_MAX_BACKUPS=7
export LOG_FILE_MAX_AGE=5
```

### 2. Ініціалізуйте логер

```go
package main

import "github.com/urban-lib/logging/v2"

func main() {
    logging.GetLogger()
    
    logging.Debugf("Application started")
    logging.Infof("Listening on port %d", 8080)
    logging.Warnf("Cache miss for key %s", "user:123")
    logging.Errorf("Failed to connect: %v", err)
}
```

### 3. Structured logging з полями

```go
logging.WithFields(logging.Fields{
    "user_id":    42,
    "request_id": "abc-123",
    "method":     "POST",
}).Infof("Request processed in %dms", 150)
```

## Змінні середовища

| Змінна | Опис | Значення за замовчуванням |
|--------|------|--------------------------|
| `LOG_LEVEL_CONSOLE` | Рівень логування в консоль (`debug`, `info`, `warn`, `error`) | `debug` |
| `LOG_FILE_ENABLE` | Увімкнути запис у файл | `false` |
| `LOG_LEVEL_FILE` | Рівень логування у файл | `debug` |
| `LOG_FILE_PATH` | Шлях до файлу логів | `logs/example.log` |
| `LOG_FILE_MAX_SIZE` | Максимальний розмір файлу в МБ (до ротації) | `100` |
| `LOG_FILE_MAX_BACKUPS` | Кількість ротованих файлів для зберігання | `7` |
| `LOG_FILE_MAX_AGE` | Максимальний вік файлу в днях | `5` |

## API

### Глобальні функції

```go
// Форматований вивід
logging.Debugf(format string, args ...interface{})
logging.Infof(format string, args ...interface{})
logging.Warnf(format string, args ...interface{})
logging.Errorf(format string, args ...interface{})
logging.Fatalf(format string, args ...interface{})
logging.Panicf(format string, args ...interface{})

// Простий вивід
logging.Debug(args ...interface{})
logging.Info(args ...interface{})
logging.Warn(args ...interface{})
logging.Error(args ...interface{})
logging.Fatal(args ...interface{})
logging.Panic(args ...interface{})

// Structured logging
logging.WithFields(fields logging.Fields) *zap.SugaredLogger
```

### Ініціалізація

```go
// Ініціалізує глобальний логер (singleton, thread-safe)
sugared, raw, err := logging.GetLogger()
```

### Діагностика

```go
// Вивести поточні значення всіх env-змінних
logging.CheckEnvironments()
```

## Формат виводу

**Консоль** — human-readable (development encoder):
```
2026-03-09T12:00:00.000+0200	INFO	myapp/main.go:15	Request processed in 150ms	{"user_id": 42}
```

**Файл** — JSON (production encoder):
```json
{"level":"info","ts":"2026-03-09T12:00:00.000+0200","caller":"myapp/main.go:15","msg":"Request processed in 150ms","user_id":42}
```

## Ротація файлів

При `LOG_FILE_ENABLE=true` лог-файли автоматично ротуються:
- При досягненні `LOG_FILE_MAX_SIZE` МБ
- Старі файли стискаються (gzip)
- Зберігається не більше `LOG_FILE_MAX_BACKUPS` архівів
- Файли старші за `LOG_FILE_MAX_AGE` днів видаляються

## Рівні логування

| Рівень | Призначення |
|--------|-------------|
| `debug` | Детальна діагностична інформація |
| `info` | Загальна інформація про роботу |
| `warn` | Попередження, що не блокують роботу |
| `error` | Помилки, що потребують уваги |
| `fatal` | Критичні помилки → `os.Exit(1)` |
| `panic` | Критичні помилки → `panic()` |

## Contributing

Див. [CONTRIBUTING.md](CONTRIBUTING.md).
