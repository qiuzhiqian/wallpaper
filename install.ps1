If (-NOT ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator"))
{   
	$arguments = "& '" + $myinvocation.mycommand.definition + "'"
	Start-Process powershell -Verb runAs -ArgumentList $arguments
	Break
}

Function AppRegisted(){
    $exist=0
    $key = Get-Item HKLM:\Software\Microsoft\Windows\CurrentVersion\Run
    $values = Get-ItemProperty $key.PSPath
    foreach ($value in $key.Property) 
    {
      if($value -eq "Wallpaper"){
        $exist=1
      }
    }
    return $exist
}

Function AppRunning(){
    $running=0
    $state=Get-Process
    foreach ($process in $state){
        $process.ProcessName
        if($process.ProcessName -eq "wallpaper"){
            $running=1
        }
    }
    return $running
}

Function Register(){
    $exist=AppRegisted
    if($exist -eq 0){
        $appPath=""""+$currentPth+"\start.bat"+""""
        $appPath
        New-ItemProperty -Path HKLM:\Software\Microsoft\Windows\CurrentVersion\Run -Name Wallpaper -PropertyType String -Value $appPath
    }
}

Function Unregister(){
    $exist=AppRegisted
    if($exist -eq 1){
        Remove-ItemProperty -Path HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Run -Name Wallpaper
    }
}

Function CmdMenu(){
    While($InNumber -ne 5)
    {
        Write-Host "#######################################################" -ForegroundColor Green
        Write-Host "# 1. Install Wallpaper App;                           #"
        Write-Host "# 2. Uninstall Wallpaper App;                         #"
        Write-Host "# 3. Start Wallpaper App;                             #"
        Write-Host "# 4. Stop Wallpaper App;                              #"
        Write-Host "# 5. Exit                                             #"
        Write-Host "#######################################################" -ForegroundColor Green
    
        $InNumber = Read-Host "Please Input The Number to Operate:"
    
        switch($InNumber)
        {
        1 {
            Write-Host "1. Install Wallpaper App`n" -ForegroundColor Green
            Register
            Write-Host "#######################################################" -ForegroundColor Green
        }
        2 {
            Write-Host "2. Uninstall Wallpaper App`n" -ForegroundColor Green
            $state=AppRunning
            if($state -eq 1){
                Get-Process -Name wallpaper | Stop-Process
            }
            
            Unregister
            Write-Host "#######################################################" -ForegroundColor Green
        }
        3 {
            Write-Host "3. Start Wallpaper App`n" -ForegroundColor Green
            $state=AppRunning
            if($state -eq 1){
                Get-Process -Name wallpaper | Stop-Process
            }
            $appPath="'"+$currentPth+"\start.ps1"+"'"
            Write-Host $appPath
            $appPath
            #. $appPath
			#& {set-executionpolicy Remotesigned -Scope Process; $appPath }
            Write-Host "#######################################################" -ForegroundColor Green
        }
        4 {
            Write-Host "4. Stop Wallpaper App`n" -ForegroundColor Green
            $state=AppRunning
            if($state -eq 1){
                Get-Process -Name wallpaper | Stop-Process
            }
            Write-Host "#######################################################" -ForegroundColor Green
        }
        5 {}
        Default { Write-Error "Please Input Number between 1 and 5"}
    
        }
        Start-Sleep 3 
        #Invoke-Command {cls}
    }
}

$currentPth=Split-Path -Parent $MyInvocation.MyCommand.Definition
#RunAdmin
CmdMenu
