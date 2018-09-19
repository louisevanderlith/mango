#Build file custom to mango for windows (able to compile to linux)
function doBuild() {
    $goos = 'windows'
    
    # All applications must be built to a central folder (./bin)
    $runPath = (Resolve-Path .\).Path
    $appPaths = "./app/", "./api/"
    
    For ($i = 0; $i -lt $appPaths.Count; $i++) {
        $currPath = $appPaths[$i]
    
        Foreach ($folder in Get-ChildItem $currPath) {
            $appName = $folder.name
            Write-Host 'Attempting to build' $appName -ForegroundColor "green"
    
            $outPath = $runPath + "\bin\" + $appName
            $exeName = $outPath + '\' + $appName
    
            if ($goos -eq 'windows') {
                $exeName += '.exe'
            }
    
            Set-Location $folder.FullName
    
            if (Test-Path (".\main.go")) {

                $env:GOARCH = 'amd64'
                $env:GOOS = $goos
                go build -i -o $exeName
    
                copyFolder $outPath "conf"
                copyFolder $outPath "static"
                copyFolder $outPath "views"

                #key
                copyFile $outPath "privkey.pem" 
                #cert
                copyFile $outPath "fullchain.pem"
            }
            else {
                Write-Host 'No main.go found in' $appName -ForegroundColor "red"
            }
    
            Write-Host 'Finished building' $appName -ForegroundColor "green"
        }
    
        Set-Location $runPath
    }
}

function copyFolder($outPath, $folder){
    if (Test-Path (".\" + $folder)) {
        $source = ".\" + $folder
        $target = $outPath + '\'

        Copy-Item -Path $source -Destination $target -Recurse -Force
    }
}

function copyFile($outPath, $file){
    if(Test-Path (".\" + $file)){
        $source = ".\" + $file
        $target = $outPath

        Copy-Item $source $target
    }
}

#Start
doBuild