$goarch = 'amd64'
$goos = 'windows'

# All applications must be build to a central folder (./bin)
$runPath = (Resolve-Path .\).Path
$appPaths = "./app/", "./api/"

For($i=0; $i -lt $appPaths.Count; $i++){
    $currPath = $appPaths[$i]

    Foreach($folder in Get-ChildItem $currPath){
        $appName= $folder.name
        Write-Host 'Attempting to build' $appName -ForegroundColor "green"

        if($appName -ne 'www'){
            Set-Location $folder.FullName

            if(Test-Path (".\main.go")){
                $outPath = $runPath + "\bin\" + $appName
                $exeName = $outPath + '\' + $appName

                if($goos -eq 'windows'){
                    $exeName += '.exe'
                }

                go build -o $exeName

                # if a conf folder is present, that should also be copied to the output folder.
                if(Test-Path (".\conf")){
                    $source = ".\conf\*"
                    $target = $outPath + '\conf\'

                    if(!(Test-Path -path $target)){
                        New-Item $target -Type Directory
                    }

                    Copy-Item -Path $source -Destination $target
                }
            }
            else {
                Write-Host 'No main.go found in' $appName -ForegroundColor "red"
            }
        }
    }

    Set-Location $runPath
}

# build the subdomain and website application
# copy contents of www application
#TODO