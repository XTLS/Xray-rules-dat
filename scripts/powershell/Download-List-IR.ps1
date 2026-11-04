Write-Host "Downloading Iran white lists..."
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/Chocolate4U/Iran-v2ray-rules/refs/heads/redundant/redundant-domains.txt" -OutFile ".\temp\pre-ir-redundant"
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/bootmortis/iran-hosted-domains/refs/heads/release/domains.txt" -OutFile ".\temp\pre-iran-hosted-domains"
$lines = Get-Content -LiteralPath ".\temp\pre-iran-hosted-domains"
$lines = $lines | Where-Object { $_ -notmatch '.+\.ir$' }
$lines = @($lines) + @("ir", "xn--mgba3a4f16a")
$set = [System.Collections.Generic.SortedSet[string]]::new([System.StringComparer]::Ordinal)
foreach ($line in $lines) {
    [void]$set.Add($line)
}
[System.IO.File]::WriteAllLines(
    (Join-Path (Get-Location) ".\temp\iran-hosted-domains"),
    [string[]]@($set),
    [System.Text.UTF8Encoding]::new($false)
)
$a = [System.Collections.Generic.SortedSet[string]]::new([System.StringComparer]::Ordinal)
$b = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
Get-Content -LiteralPath ".\temp\iran-hosted-domains" | ForEach-Object { [void]$a.Add($_) }
Get-Content -LiteralPath ".\temp\pre-ir-redundant" | ForEach-Object { [void]$b.Add($_) }
$result = @(
    foreach ($line in $a) {
        if (-not $b.Contains($line)) {
            $line
        }
    }
)
[System.IO.File]::WriteAllLines(
    (Join-Path (Get-Location) ".\temp\iran-hosted-domains-lite"),
    [string[]]$result,
    [System.Text.UTF8Encoding]::new($false)
)
Copy-Item -Path ".\temp\iran-hosted-domains-lite" -Destination ".\data\std\"
Copy-Item -Path ".\temp\iran-hosted-domains-lite" -Destination ".\data\exp\"
Copy-Item -Path ".\temp\iran-hosted-domains" -Destination ".\data\exp\"
Write-Host "Downloading Iran black lists..."
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/MrDevAnony/DynX-AntiBan-Domains/refs/heads/main/DynX-AntiBan-list.lst" -OutFile ".\temp\dynx-antiban-domains"
Copy-Item -Path ".\temp\dynx-antiban-domains" -Destination ".\data\std\"
Copy-Item -Path ".\temp\dynx-antiban-domains" -Destination ".\data\exp\"
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/armanlife/iran-sanctioned-domains/refs/heads/main/data/ir-sanctioned-domain" -OutFile ".\temp\iran-sanctioned-domains"
Copy-Item -Path ".\temp\iran-sanctioned-domains" -Destination ".\data\std\"
Copy-Item -Path ".\temp\iran-sanctioned-domains" -Destination ".\data\exp\"
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/armanlife/iran-blocked-domains/refs/heads/main/data/ir-blocked" -OutFile ".\temp\pre-iran-blocked-domains"
.\sanitize.exe --inputpath ".\temp\pre-iran-blocked-domains" --outputname "iran-blocked-domains" --outputdir ".\data\exp\"
