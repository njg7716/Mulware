if (!([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) { Start-Process -WindowStyle hidden powershell "-NoProfile -ExecutionPolicy Bypass -File `"$PSCommandPath`"" -Verb RunAs; exit }
$port="200"
netsh advfirewall firewall add rule name="Open Port $port" dir=in action=allow protocol=TCP localport=$port | Out-Null #adds Firewall Rule
Start-Process -WindowStyle hidden powershell {$port="200"; $x = [System.Net.Sockets.TcpListener]$port; $x.Start(); Start-Sleep -Seconds (30)} | Out-Null #opens port and sleeps to keep it open
