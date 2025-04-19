# Financial & Utility Calculators Hub

## ℹ️ Как запускать api-app через терминал
- в терминале написать команду make dev
- запуститься сервер на http://localhost:8080/

```markdown
Проблемы с make не установлен:

make обычно поставляется с пакетом GNU build tools (MinGW, Cygwin, WSL в Windows или стандартно в Linux/macOS).
Windows: Установите MinGW или Cygwin, либо используйте WSL (Windows Subsystem for Linux).
Linux/macOS: Обычно make уже установлен. Если нет, установите через пакетный менеджер (например, sudo apt install make для Ubuntu).

Make не в PATH:
Если make установлен, но система его не находит, добавьте путь к нему в переменную окружения PATH.
Вы не в той директории:
Убедитесь, что вы запускаете команду в папке, где есть Makefile.

Альтернатива для Windows:
Если вы используете Visual Studio, попробуйте nmake (вместо make).
Для CMake проектов используйте cmake --build.
```

```markdown
Пакет для автоматического обновления сервера, когда изменяем файлы go

go install github.com/air-verse/air@latest

После установки бинарник появится в:
- Linux/Mac: $HOME/go/bin/air
- Windows: %USERPROFILE%\go\bin\air.exe

💡 Если команда не работает(ma:
Убедитесь, что $HOME/go/bin добавлен в PATH:

- bash
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc  # или ~/.zshrc
source ~/.bashrc
  
Или используйте полный путь:
- bash
~/go/bin/air
```


# Как запустить на Windows 
- Установить make на windows.
- Установить air пакет на go
go install github.com/cosmtrek/air@latest
или
go install github.com/air-verse/air@latest
- Проверить переменные среды на windows, должно быть %USERPROFILE%\go\bin\air.exe 
https://habr.com/ru/posts/823682/

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
