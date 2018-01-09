function startPlay() {
    $progs = getPrograms
    $wd = Convert-Path .

    Foreach($prog in $progs){
        $progName = $prog.name
        $progCmd = $prog.cmd
        $progPath = $wd + "\\" + $prog.type + "\\" + $progName

        if(Test-Path($progPath)){

            if($progName -eq 'gate'){
                Write-Host 'Gate must wait for all applications to register with the router before starting.' -ForegroundColor "green"
                Start-Sleep -s 5
            }

            $cmd = "Write-Host 'Starting' $progName -ForegroundColor 'red'; cd $progPath; $progCmd; Read-Host"
            Start-Process powershell -argument $cmd

            if($progName -eq 'router'){
                Write-Host 'Giving router some time to spin up.' -ForegroundColor "green"
                Start-Sleep -s 2
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