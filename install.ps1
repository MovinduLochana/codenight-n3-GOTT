$ErrorActionPreference = "Stop"

$Repo = "Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO"

Write-Host "🦊 Downloading Gostlings Go Exercises..." -ForegroundColor Cyan

# 1. Determine download asset for Windows
$FileName = "gostlings-windows-amd64.zip"

# 2. Get latest release tag
try {
    $LatestRelease = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest"
    $Tag = $LatestRelease.tag_name
} catch {
    $Tag = "v1.0.0"
}

$DownloadUrl = "https://github.com/$Repo/releases/download/$Tag/$FileName"
$TempZip = [System.IO.Path]::GetTempFileName() + ".zip"
$TempExtract = Join-Path ([System.IO.Path]::GetTempPath()) ([System.Guid]::NewGuid().ToString())

Write-Host "Fetching $FileName ($Tag)..." -ForegroundColor Yellow

Invoke-WebRequest -Uri $DownloadUrl -OutFile $TempZip

Expand-Archive -Path $TempZip -DestinationPath $TempExtract -Force

$SubDir = Get-ChildItem -Path $TempExtract | Select-Object -First 1
Get-ChildItem -Path $SubDir.FullName | Copy-Item -Destination . -Recurse -Force

Remove-Item -Path $TempZip -Force -ErrorAction SilentlyContinue
Remove-Item -Path $TempExtract -Recurse -Force -ErrorAction SilentlyContinue

Write-Host "✅ Setup complete! Run .\gostlings.exe to start learning Go!" -ForegroundColor Green
