
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
