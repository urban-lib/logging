# Contributing to logging/v2

Дякуємо за інтерес до проекту! Нижче описано правила та процес контрибуції.

## Вимоги

- **Go 1.22+**
- `golangci-lint` для лінтингу
- `git` з підтримкою [Conventional Commits](https://www.conventionalcommits.org/)

## Початок роботи

```bash
# Клонування
git clone https://github.com/urban-lib/logging.git
cd logging

# Встановлення залежностей
go mod download

# Запуск тестів
go test ./...

# Лінтинг
golangci-lint run ./...
```

## Git Workflow

### Гілки

| Гілка | Призначення |
|-------|-------------|
| `main` | Стабільна версія, захищена від прямих push |
| `develop` | Основна гілка розробки |
| `feature/*` | Нова функціональність |
| `fix/*` | Виправлення багів |
| `docs/*` | Зміни в документації |

### Процес

1. Створіть гілку від `develop`:
   ```bash
   git checkout develop
   git pull origin develop
   git checkout -b feature/my-feature
   ```
2. Внесіть зміни та напишіть тести.
3. Переконайтесь що всі тести проходять: `go test ./...`
4. Переконайтесь що лінтер не має зауважень: `golangci-lint run ./...`
5. Створіть Pull Request в `develop`.

## Conventional Commits

Всі коміти **мають** відповідати формату [Conventional Commits](https://www.conventionalcommits.org/). Це необхідно для автоматичної генерації тегів та changelog.

### Формат

```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

### Типи

| Тип | Опис | Вплив на версію |
|-----|------|-----------------|
| `feat` | Нова функціональність | **minor** (0.X.0) |
| `fix` | Виправлення бага | **patch** (0.0.X) |
| `docs` | Зміни в документації | — |
| `refactor` | Рефакторинг без зміни поведінки | — |
| `test` | Додавання / зміна тестів | — |
| `chore` | Оновлення залежностей, CI та ін. | — |
| `perf` | Покращення продуктивності | **patch** |
| `BREAKING CHANGE` | Зміна з порушенням зворотної сумісності (в footer або `!` після type) | **major** (X.0.0) |

### Приклади

```bash
# Нова функціональність
git commit -m "feat(logger): add WithContext method for trace propagation"

# Виправлення бага
git commit -m "fix(env): correct default level comparison in getLogLevelConsole"

# Breaking change
git commit -m "feat(api)!: replace GetLogger with New constructor"

# Документація
git commit -m "docs: update README with env variables table"
```

## Автоматичне тегування

При мержі в `main` GitHub Actions автоматично:
1. Аналізує коміти з моменту останнього тегу.
2. Визначає тип версії (major/minor/patch) на основі Conventional Commits.
3. Створює новий git tag (`vX.Y.Z`).
4. Генерує GitHub Release з changelog.

**Не створюйте теги вручну** — це робить CI автоматично.

## Code Style

- Дотримуйтесь `gofmt` / `goimports`.
- Публічні функції, типи та методи **мають** мати GoDoc коментарі.
- Назви змінних та функцій — англійською.
- Уникайте глобального стану де можливо.
- Помилки обробляйте явно, не ігноруйте.

## Тести

- Кожна нова функція або виправлення **має** супроводжуватись тестами.
- Використовуйте table-driven tests де доцільно.
- Мінімальне покриття: **80%**.
- Файли тестів: `*_test.go` поруч з кодом що тестується.

```bash
# Запуск тестів з покриттям
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Pull Request

### Чеклист перед створенням PR

- [ ] Код компілюється без помилок (`go build ./...`)
- [ ] Всі тести проходять (`go test ./...`)
- [ ] Лінтер не має зауважень (`golangci-lint run`)
- [ ] Додані тести для нової функціональності / виправлення
- [ ] Коміти відповідають Conventional Commits
- [ ] Оновлена документація (якщо змінюється публічний API)

### Review

- PR потребує мінімум **1 approve** для мержу.
- CI перевірки мають бути зелені.
- Squash merge в `develop`, merge commit в `main`.

## Повідомлення про баги

Створіть Issue з:
- Версія Go та ОС
- Мінімальний код для відтворення
- Очікувана та фактична поведінка
- Логи / stack trace (якщо є)

## Ліцензія

Вносячи зміни, ви погоджуєтесь що вони будуть опубліковані під тою ж ліцензією що і проект.
