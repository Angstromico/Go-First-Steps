param (
    [Parameter(Position=0, Mandatory=$true)]
    [string]$Search,

    [Parameter(ValueFromRemainingArguments=$true)]
    [string[]]$GoArgs
)

# Search for main.go files recursively
$mainFiles = Get-ChildItem -Path . -Filter "main.go" -Recurse | Where-Object { $_.FullName -notlike "*\.git\*" }

# Filter by directory path containing the search string
$matches = $mainFiles | Where-Object { $_.DirectoryName -like "*$Search*" }

if ($matches.Count -eq 0) {
    Write-Host "Error: No project found matching '$Search'." -ForegroundColor Red
    return
}

if ($matches.Count -gt 1) {
    Write-Host "Multiple matches found for '$Search':" -ForegroundColor Yellow
    $matches | ForEach-Object { 
        $relPath = Resolve-Path $_.FullName -Relative
        Write-Host "  $relPath" 
    }
    Write-Host "Please provide a more specific name."
    return
}

# Exactly one match
$targetPath = Resolve-Path $matches[0].FullName -Relative
$targetDir = Split-Path $targetPath

Write-Host "Executing: go run $targetDir $($GoArgs -join ' ')" -ForegroundColor Cyan
go run $targetDir @GoArgs
