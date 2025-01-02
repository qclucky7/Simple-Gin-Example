@echo off
SETLOCAL

:: 设置项目目录
SET "PROJECT_DIR=%cd%"

:: 清理之前的编译结果
del /s /q %PROJECT_DIR%\dist\*

:: 创建dist目录
mkdir %PROJECT_DIR%\dist

:: 编译Windows x64
echo Compiling for Windows x64...
go build -o dist\windows_x64\server.exe -v

:: 编译macOS Intel
echo Compiling for macOS Intel...
set GOOS=darwin
set GOARCH=amd64
go build -o dist\mac_x64\server -v

:: 编译macOS ARM
echo Compiling for macOS ARM...
set GOOS=darwin
set GOARCH=arm64
go build -o dist\mac_arm64\server -v

:: 编译Linux
echo Compiling for Linux...
set GOOS=linux
set GOARCH=amd64
go build -o dist\linux_x64\server -v

echo Done.
pause
ENDLOCAL