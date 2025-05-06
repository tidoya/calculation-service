# Financial & Utility Calculators Hub

## ℹ️ Как запускать api-app через терминал
- в терминале написать команду make run
- запуститься сервер на http://localhost:8080/

# Как запустить на Windows 
- Установить make на windows.
- Установить air пакет на go
go install github.com/cosmtrek/air@latest
или
go install github.com/air-verse/air@latest
- Проверить переменные среды на windows, должно быть %USERPROFILE%\go\bin\air.exe 
https://habr.com/ru/posts/823682/

# Запуск postgress локально на своей машине разворачиваем в Docker
-- docker pull postgres
-- docker run --name=calculation-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgress
// если нужны файлы миграции(не нужно они созданы)
-- migrate create -ext sql -dir ./schema -seq init
// поднять файлы миграции 
-- migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
// подключиться к бд
-- docker exec -it "container id(DOCKER PS)" /bin/bash
// посмотреть схему бд
-- psql -U postgres
-- \d

![Calculator Banner]() <!-- Consider adding a real banner image -->

## 📋 Оглавление
- [Финансовые калькуляторы](#-финансовые-калькуляторы)
- [Налоги и бухгалтерия](#-налоги-и-бухгалтерия)
- [Юридические инструменты](#-юридические-инструменты)
- [Автомобильные расчеты](#-автомобильные-расчеты)
- [Математические инструменты](#-математические-инструменты)
- [Здоровье](#-здоровье)
- [Для бизнеса](#-для-бизнеса)
- [Конвертеры](#-конвертеры)

## 🏦 Финансовые калькуляторы

### 💰 Кредиты и займы
| Калькулятор | Описание |
|-------------|----------|
| [Кредитный](#) | Расчет аннуитетных и дифференцированных платежей |
| [Ипотечный](#) | С учетом страховки и первоначального взноса |
| [Досрочное погашение](#) | Экономия при частичном/полном досрочном погашении |
| [Микрозаймы](#) | Расчет переплаты по краткосрочным займам |

### 📈 Инвестиции


### Packages: 
## Linters
- [golangci-lint](https://golangci-lint.run/welcome/install/)
