# Build for Android/Termux (ARM64)
$env:GOOS = "linux"
$env:GOARCH = "arm64"
$Output = "pokesplash-android"

Write-Host "Building for Android (Termux)..." -ForegroundColor Cyan
go build -ldflags "-s -w" -o $Output ./cmd/pokesplash

if ($LASTEXITCODE -eq 0) {
    Write-Host "Build successful! Binary: $Output" -ForegroundColor Green
    Write-Host "Transfer this file to your phone and run it in Termux!"
}
else {
    Write-Error "Build failed!"
}

# Reset env vars (optional, but good practice in script scope)
$env:GOOS = $null
$env:GOARCH = $null
