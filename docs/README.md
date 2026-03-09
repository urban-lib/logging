# logging/v3 — Documentation

Go-бібліотека структурованого логування на основі [Uber Zap](https://github.com/uber-go/zap) з підтримкою ротації лог-файлів через [lumberjack](https://github.com/natefinch/lumberjack.v2).

## Встановлення

```bash
go get github.com/urban-lib/logging/v3
```

## Конструктори

### `New(opts ...Option) (Logger, error)`

Створює логер із функціональними опціями поверх `DefaultConfig()`:

```go
logger, err := logging.New(
    logging.WithConsoleLevel("info"),
    logging.WithFileEnabled(true),
    logging.WithFilePath("logs/app.log"),
    logging.WithFileLevel("error"),
)
```

### `NewWithConfig(cfg Config) (Logger, error)`

Створює логер напряму зі структури `Config`:

```go
logger, err := logging.NewWithConfig(logging.Config{
    ConsoleLevel: "debug",
    FileEnabled:  true,
    FileLevel:    "info",
    FilePath:     "logs/app.log",
    FileMaxSize:  100,
    CallerSkip:   1,
})
```

### `NewFromEnv() (Logger, error)`

Створює логер із змінних середовища (перезаписує `DefaultConfig()`):

```go
logger, err := logging.NewFromEnv()
```

## Інтерфейс Logger

```go
type Logger interface {
    Debugf(format string, args ...any)
    Infof(format string, args ...any)
    Warnf(format string, args ...any)
    Errorf(format string, args ...any)
    Fatalf(format string, args ...any)
    Panicf(format string, args ...any)

    Debug(args ...any)
    Info(args ...any)
    Warn(args ...any)
    Error(args ...any)
    Panic(args ...any)
    Fatal(args ...any)

    Debugw(msg string, keysAndValues ...any)
    Infow(msg string, keysAndValues ...any)
    Warnw(msg string, keysAndValues ...any)
    Errorw(msg string, keysAndValues ...any)
    Fatalw(msg string, keysAndValues ...any)
    Panicw(msg string, keysAndValues ...any)

    WithFields(fields Fields) Logger
    WithContext(ctx context.Context) Logger
    Log() *zap.Logger
    Sync() error
}
```

- `WithFields` — повертає **новий** `Logger` з доданими полями (батьківський не змінюється).
- `WithContext` — повертає **новий** `Logger` з полями з контексту (див. `ContextWithFields`).
- `Debugw/Infow/...` — sugar key-value API: `logger.Infow("msg", "key", "val", "num", 42)`.
- `Log()` — повертає `*zap.Logger` для розширеного використання.
- `Sync()` — flush буферів. Завжди викликайте `defer logger.Sync()` перед виходом з програми.

## Config

```go
type Config struct {
    ConsoleLevel       string  // "debug", "info", "warn", "error" — default: "debug"
    FileEnabled        bool    // default: false
    FileLevel          string  // default: "debug"
    FilePath           string  // default: "logs/example.log"
    FileMaxSize        int     // MB — default: 100
    FileMaxBackups     int     // default: 7
    FileMaxAge         int     // days — default: 5
    FileCompress       bool    // default: true
    CallerSkip         int     // default: 1
    SamplingInitial    int     // messages/sec before sampling — default: 0 (disabled)
    SamplingThereafter int     // keep every Nth after initial — default: 0
}
```

`DefaultConfig()` повертає конфігурацію з усіма значеннями за замовчуванням.

## Функціональні опції

| Опція | Опис |
|-------|------|
| `WithConsoleLevel(level)` | Рівень логування в консоль |
| `WithFileEnabled(bool)` | Увімкнути/вимкнути файловий лог |
| `WithFileLevel(level)` | Рівень логування у файл |
| `WithFilePath(path)` | Шлях до лог-файлу |
| `WithFileMaxSize(mb)` | Максимальний розмір файлу |
| `WithFileMaxBackups(n)` | Кількість ротованих файлів |
| `WithFileMaxAge(days)` | Термін зберігання файлів |
| `WithFileCompress(bool)` | Gzip-стиснення ротованих файлів |
| `WithCallerSkip(n)` | Глибина caller skip |
| `WithSampling(initial, thereafter)` | Семплінг логів (rate limiting) |

## Context-aware логування

Для прокидання trace/request ID через `context.Context`:

```go
// Зберігти поля в контексті
ctx = logging.ContextWithFields(ctx, logging.Fields{
    "trace_id":   "abc-xyz",
    "request_id": "req-123",
})

// Кілька викликів накопичують поля:
ctx = logging.ContextWithFields(ctx, logging.Fields{"user_id": 42})

// Використати в логері
logger.WithContext(ctx).Infof("Processing request")

// Або глобально
logging.WithContext(ctx).Infof("Global with context")

// Дістати поля з контексту
fields := logging.FieldsFromContext(ctx)
```

## Sugar key-value API

Методи `Debugw/Infow/Warnw/Errorw/Fatalw/Panicw` приймають повідомлення та пари ключ-значення:

```go
logger.Infow("request handled",
    "method", "POST",
    "path", "/api/users",
    "latency_ms", 42,
)
```

Також доступні на рівні пакету: `logging.Infow(...)`, `logging.Errorw(...)` тощо.

## Типізовані поля (Field helpers)

Для zero-allocation полів при роботі з `*zap.Logger` через `Log()`:

```go
logger.Log().Info("structured",
    logging.String("service", "payment"),
    logging.Int("attempt", 3),
    logging.Duration("latency", 150*time.Millisecond),
    logging.Bool("cached", false),
    logging.Err(err),
    logging.NamedErr("cause", rootErr),
    logging.Float64("score", 0.95),
    logging.Time("started_at", startTime),
    logging.Any("payload", myStruct),
    logging.Stringer("addr", netAddr),
)
```

| Хелпер | Опис |
|--------|------|
| `String(key, val)` | Рядок |
| `Int(key, val)` | Ціле число |
| `Int64(key, val)` | Ціле число (int64) |
| `Float64(key, val)` | Дробове число |
| `Bool(key, val)` | Булеве значення |
| `Err(err)` | Помилка (ключ = `"error"`) |
| `NamedErr(key, err)` | Помилка з кастомним ключем |
| `Duration(key, val)` | Тривалість |
| `Time(key, val)` | Час |
| `Any(key, val)` | Довільне значення (reflection) |
| `Stringer(key, val)` | Значення з методом `String()` |

## Семплінг (Log Sampling)

Для high-throughput сервісів можна обмежити кількість однакових лог-рядків:

```go
logger, _ := logging.New(
    logging.WithSampling(100, 10), // 100 msg/sec початково, потім кожне 10-те
)
```

Коли `SamplingInitial = 0` — семплінг вимкнений (за замовчуванням).

## Глобальні функції

Пакет надає зручний шар глобальних функцій, які делегують до default-логера:

```go
logging.Infof("Server started on port %d", 8080)
logging.WithFields(logging.Fields{"user": "alice"}).Warnf("Slow request")
logging.Infow("event", "key", "value")
logging.WithContext(ctx).Infof("Traced request")
```

Глобальний логер ініціалізується лениво з env-змінних при першому виклику.

### `SetDefault(l Logger)`

Замінює глобальний логер на кастомний екземпляр. CallerSkip автоматично коригується (+1) для коректного відображення джерела виклику.

```go
logger, _ := logging.New(logging.WithConsoleLevel("warn"))
logging.SetDefault(logger)
logging.Warnf("Now using custom logger")
```

## Змінні середовища

| Змінна | Опис | Значення за замовчуванням |
|--------|------|--------------------------|
| `LOG_LEVEL_CONSOLE` | Рівень логування в консоль | `debug` |
| `LOG_FILE_ENABLE` | Увімкнути запис у файл | `false` |
| `LOG_LEVEL_FILE` | Рівень логування у файл | `debug` |
| `LOG_FILE_PATH` | Шлях до файлу логів | `logs/example.log` |
| `LOG_FILE_MAX_SIZE` | Максимальний розмір файлу (МБ) | `100` |
| `LOG_FILE_MAX_BACKUPS` | Кількість ротованих файлів | `7` |
| `LOG_FILE_MAX_AGE` | Максимальний вік файлу (днів) | `5` |

## Формат виводу

**Консоль** — human-readable (development encoder):
```
2026-03-09T12:00:00.000+0200  INFO  myapp/main.go:15  Request processed  {"user_id": 42}
```

**Файл** — JSON (production encoder):
```json
{"level":"info","ts":"2026-03-09T12:00:00.000+0200","caller":"myapp/main.go:15","msg":"Request processed","user_id":42}
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

## CallerSkip

- **Instance-based** (`New`, `NewWithConfig`) — `CallerSkip: 1` (за замовчуванням), caller вказує на код, що викликав `logger.Infof(...)`.
- **Глобальні функції** (`logging.Infof(...)`) — +1 автоматично (= 2), щоб пропустити обгортку в `formatting.go`.
- **`SetDefault`** — автоматично додає +1 до CallerSkip.
- **`WithFields` (глобальний)** — коригує CallerSkip(-1), щоб повернутий `Logger` правильно показував caller при прямому використанні.

## Діагностика

```go
logging.CheckEnvironments()
```

Виводить поточні значення всіх env-змінних через `log.Println`.

## Contributing

Див. [CONTRIBUTING.md](../CONTRIBUTING.md).
