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

$currentPth=Split-Path -Parent $MyInvocation.MyCommand.Definition
$appPath=""""+$currentPth+"\wallpaper.exe"+""""
$appPath
$state=AppRunning
if($state -eq 1){
    Get-Process -Name wallpaper | Stop-Process
}
Start-Process $appPath  -WindowStyle Hidden