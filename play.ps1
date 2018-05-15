function startPlay() {
    $progs = getPrograms
    $wd = Convert-Path .

    Foreach($prog in $progs){
        $progName = $prog.name
        $progCmd = $prog.cmd
        $progPath = $wd + "\\" + $prog.type + "\\" + $progName

        if(Test-Path($progPath)){

            if($progName -eq 'gate'){
                $waitTime = ($progs.Length * 1.5)
                $msg = 'Gate must wait ' + $waitTime + ' seconds for all applications to register with the router before starting.'
                Write-Host $msg -ForegroundColor "green"
                Start-Sleep -s $waitTime
            }

            $cmd = "Write-Host 'Starting' $progName -ForegroundColor 'red'; cd $progPath; $progCmd; Read-Host"
            spawnWindow $cmd $progName

            if($progName -eq 'router'){
                Write-Host 'Giving router some time to spin up.' -ForegroundColor "green"
                Start-Sleep -s 2
            }
        } else {
            Write-Host 'Directory' $progPath 'not Found. Please ensure "./build" has been run.' -ForegroundColor "red"
        }
    }
}

function spawnWindow($cmd, $name) {
    $process = new-object System.Diagnostics.Process
    $startInfo = new-object System.Diagnostics.ProcessStartInfo
    $startInfo.FileName = "$pshome\powershell.exe"
    $startInfo.Arguments = $cmd

    $process.StartInfo = $startInfo
    $process.Start()
}

function getPrograms() {
    $json = Get-Content -Raw -Path "play.json"
    $obj = ConvertFrom-Json $json

    return $obj.programs
}

# Start
startPlay