Write-Host "Starting UNSCH Horarios Backend..." -ForegroundColor Cyan
Set-Location $PSScriptRoot/backend
go run ./cmd/api
