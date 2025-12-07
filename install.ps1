# PokeSplash Installer for Windows

$InstallDir = "$env:LOCALAPPDATA\PokeSplash"
$BinaryName = "pokesplash.exe"
$SourceBinary = ".\pokesplash.exe"

Write-Host "Installing PokeSplash..." -ForegroundColor Cyan

# Check if binary exists in current dir
if (-not (Test-Path $SourceBinary)) {
    Write-Error "pokesplash.exe not found! Please build it first or download the release."
    exit 1
}

# Create install directory
if (-not (Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
}

# Copy binary
Copy-Item -Path $SourceBinary -Destination "$InstallDir\$BinaryName" -Force
Write-Host "Binary installed to $InstallDir" -ForegroundColor Green

# Add to PATH if not already present
$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -notlike "*$InstallDir*") {
    $NewPath = "$UserPath;$InstallDir"
    [Environment]::SetEnvironmentVariable("Path", $NewPath, "User")
    Write-Host "Added $InstallDir to User PATH." -ForegroundColor Green
    Write-Warning "You may need to restart your terminal for changes to take effect."
}
else {
    Write-Host "Path already configured." -ForegroundColor Gray
}

Write-Host "Installation Complete! ⚡" -ForegroundColor Green
Write-Host "Try running: pokesplash"
