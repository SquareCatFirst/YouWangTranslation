@echo off
chcp 65001 >nul
REM ================================
REM 获取 bat 文件所在目录
REM ================================
SET BAT_DIR=%~dp0
REM 去掉末尾的反斜杠
SET BAT_DIR=%BAT_DIR:~0,-1%

echo Bat 所在目录: %BAT_DIR%

REM ================================
REM 指定 Go 源码目录
REM ================================
SET SRC_DIR=%BAT_DIR%\cmd\server
echo Go 源码目录: %SRC_DIR%

REM ================================
REM 输出可执行文件到 bat 所在目录
REM ================================
SET OUTPUT=%BAT_DIR%\server.exe
echo 输出路径: %OUTPUT%

REM ================================
REM 构建可执行文件
REM ================================
go build -o "%OUTPUT%" "%SRC_DIR%"

REM ================================
REM 检查构建结果
REM ================================
IF %ERRORLEVEL% NEQ 0 (
    echo 构建失败！
    pause
    exit /b 1
)

echo 构建成功: %OUTPUT%
pause
