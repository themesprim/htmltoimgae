@echo off
echo Setting up HTML to Image Converter...

REM Check if Go is installed
where go >nul 2>nul
if %ERRORLEVEL% neq 0 (
    echo Go is not installed. Please install Go first.
    exit /b 1
)

REM Check if Chrome is installed
where chrome >nul 2>nul
if %ERRORLEVEL% neq 0 (
    echo Chrome is not installed. Please install Chrome browser.
    exit /b 1
)

REM Install Go dependencies
echo Installing Go dependencies...
go mod download
go mod tidy

REM Build the application
echo Building the application...
go build -o html-to-image.exe

echo Setup completed successfully!
echo You can now run the application with: run.bat 