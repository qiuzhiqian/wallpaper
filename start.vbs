Set ws = CreateObject("Wscript.Shell")

test = createobject("Scripting.FileSystemObject").GetFile(Wscript.ScriptFullName).ParentFolder.Path
appPath = test+"\start.ps1"
startCmd = "powershell "+appPath
ws.run startCmd,vbhide

