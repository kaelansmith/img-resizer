@echo off
REM Image Resizer - Multi-platform build script for Windows
REM Builds executables for Windows, Mac, and Linux

echo 🔨 Building Image Resizer for all platforms...

REM Create builds directory if it doesn't exist
if not exist builds mkdir builds

REM Build for Windows (64-bit)
echo 📦 Building for Windows (64-bit)...
set GOOS=windows
set GOARCH=amd64
go build -o builds\img-resizer-windows-amd64.exe main.go

REM Build for Mac (64-bit)
echo 📦 Building for Mac (64-bit)...
set GOOS=darwin
set GOARCH=amd64
go build -o builds\img-resizer-darwin-amd64 main.go

REM Build for Mac (Apple Silicon)
echo 📦 Building for Mac (Apple Silicon)...
set GOOS=darwin
set GOARCH=arm64
go build -o builds\img-resizer-darwin-arm64 main.go

REM Build for Linux (64-bit)
echo 📦 Building for Linux (64-bit)...
set GOOS=linux
set GOARCH=amd64
go build -o builds\img-resizer-linux-amd64 main.go

REM Build for Linux (ARM64)
echo 📦 Building for Linux (ARM64)...
set GOOS=linux
set GOARCH=arm64
go build -o builds\img-resizer-linux-arm64 main.go

echo.
echo ✅ Build complete! Executables created in 'builds' directory:
echo.
dir builds
echo.
echo 📁 Files created:
echo   • img-resizer-windows-amd64.exe  (Windows 64-bit)
echo   • img-resizer-darwin-amd64       (Mac Intel)
echo   • img-resizer-darwin-arm64       (Mac Apple Silicon)
echo   • img-resizer-linux-amd64        (Linux 64-bit)
echo   • img-resizer-linux-arm64        (Linux ARM64)
echo.
echo 💡 Usage: Copy the appropriate executable to your image folder and run it!
pause
