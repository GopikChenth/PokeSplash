$ErrorActionPreference = "Stop"

$goExe = "C:\Program Files\Go\bin\go.exe"

Write-Host "Building PokéSplash..."
& $goExe build -ldflags "-s -w" -o pokesplash.exe .\cmd\pokesplash\main.go

if ($LASTEXITCODE -eq 0) {
    Write-Host "Build successful! Binary created: pokesplash.exe" -ForegroundColor Green
    
    # Get file size
    $fileSize = (Get-Item pokesplash.exe).Length / 1MB
    Write-Host ("Binary size: {0:N2} MB" -f $fileSize) -ForegroundColor Cyan
} else {
    Write-Host "Build failed!" -ForegroundColor Red
    exit 1
}
