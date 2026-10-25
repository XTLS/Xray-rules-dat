Write-Host "Downloading Iran related lists..."
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/Chocolate4U/Iran-v2ray-rules/refs/heads/redundant/redundant-domains.txt" -OutFile ".\temp\pre-ir-redundant"
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/bootmortis/iran-hosted-domains/refs/heads/release/domains.txt" -OutFile ".\temp\pre-expand-ir"
$lines = Get-Content -LiteralPath ".\temp\pre-expand-ir"
$lines = $lines | Where-Object { $_ -notmatch '.+\.ir$' }
$lines = @($lines) + @("ir", "xn--mgba3a4f16a")
$set = [System.Collections.Generic.SortedSet[string]]::new([System.StringComparer]::Ordinal)
foreach ($line in $lines) {
    [void]$set.Add($line)
}
[System.IO.File]::WriteAllLines(
    (Join-Path (Get-Location) ".\temp\expand-ir"),
    [string[]]@($set),
    [System.Text.UTF8Encoding]::new($false)
)
$a = [System.Collections.Generic.SortedSet[string]]::new([System.StringComparer]::Ordinal)
$b = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
Get-Content -LiteralPath ".\temp\expand-ir" | ForEach-Object { [void]$a.Add($_) }
Get-Content -LiteralPath ".\temp\pre-ir-redundant" | ForEach-Object { [void]$b.Add($_) }
$result = @(
    foreach ($line in $a) {
        if (-not $b.Contains($line)) {
            $line
        }
    }
)
[System.IO.File]::WriteAllLines(
    (Join-Path (Get-Location) ".\temp\mainly-ir"),
    [string[]]$result,
    [System.Text.UTF8Encoding]::new($false)
)
Copy-Item -Path ".\temp\mainly-ir" -Destination ".\data\std\"
Copy-Item -Path ".\temp\mainly-ir" -Destination ".\data\exp\"
Copy-Item -Path ".\temp\expand-ir" -Destination ".\data\exp\"
