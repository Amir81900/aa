param(
  [Parameter(Mandatory=$true)][string]$ZipPath,
  [Parameter(Mandatory=$false)][string]$ExePath,
  [Parameter(Mandatory=$false)][string]$Tag = "v1.3"
)

$ErrorActionPreference = "Stop"
$repo = "Amir81900/aa"

if (-not (Get-Command gh -ErrorAction SilentlyContinue)) {
  throw "GitHub CLI (gh) is not installed. Install it from https://cli.github.com/ and reopen PowerShell."
}
& gh auth status | Out-Null
if ($LASTEXITCODE -ne 0) {
  throw "GitHub CLI is not authenticated. Run: gh auth login"
}
if (-not (Test-Path -LiteralPath $ZipPath -PathType Leaf)) {
  throw "ZIP file not found: $ZipPath"
}
$assets = @((Resolve-Path -LiteralPath $ZipPath).Path)
if ($ExePath) {
  if (-not (Test-Path -LiteralPath $ExePath -PathType Leaf)) {
    throw "EXE file not found: $ExePath"
  }
  $assets += (Resolve-Path -LiteralPath $ExePath).Path
}
Write-Host "Uploading assets to $repo release $Tag ..."
& gh release create $Tag @assets --repo $repo --title "JADDEH 1405 $Tag" --notes "Development snapshot. Review README_FA.md for known limitations and asset fidelity notes."
if ($LASTEXITCODE -ne 0) {
  throw "GitHub release upload failed. Check repository permission, release tag, and asset size."
}
Write-Host "Release created successfully:"
& gh release view $Tag --repo $repo --web
