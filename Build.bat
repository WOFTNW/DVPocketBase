@echo off
echo Building DVPocketBase for Windows...

REM Create builds directory if it doesn't exist
if not exist builds mkdir builds

REM Build the Windows executable
go build -v -o builds\dvpocketbase-windows.exe .\main.go

if %ERRORLEVEL% EQU 0 (
    echo Build completed successfully! Executable saved to builds\dvpocketbase-windows.exe
) else (
    echo Build failed with error code %ERRORLEVEL%
)
