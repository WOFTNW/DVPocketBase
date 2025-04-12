@echo off
echo Running DVPocketBase for Windows...

REM Check if the executable exists, if not build it
if not exist builds\dvpocketbase-windows.exe (
    echo Executable not found. Building first...
    call Build.bat
)

REM Run the application
echo Starting DVPocketBase...
builds\dvpocketbase-windows.exe %*

echo Application exited with code %ERRORLEVEL%
