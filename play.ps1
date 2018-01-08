function startPlay() {
    $progs = getPrograms
    $wd = Convert-Path .

    Foreach($prog in $progs){
        $progPath = $wd + "/" + $prog.path
        if(Test-Path($progPath)){
            $cmd = "Write-Host 'Starting' $prog -ForegroundColor 'red'; cd $progPath; $prog.cmd; Read-Host"

            if($prog -eq 'gate'){
                Write-Host 'Gate must wait for all applications to register with the router before starting.' -ForegroundColor "green"
                Start-Sleep -s 2
            }

            Start-Process powershell -argument $cmd

            if($prog -eq 'router'){
                Write-Host 'Giving router some time to spin up.' -ForegroundColor "green"
                Start-Sleep -s 1
            }
        } else {
            Write-Host 'Directory' $progPath 'not Found. Please ensure "./build" has been run.' -ForegroundColor "red"
        }
    }
}

function getPrograms() {
    $json = Get-Content -Raw -Path "play.json"
    $obj = ConvertFrom-Json $json

    return $obj.programs
}

# Start
startPlay