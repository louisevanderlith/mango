# This file provides a way for us to launch all services & tasks required to run/debug this application

# Start APIs (callrouter and proxy must always be first)
# Beego API's use swagger, so we want to include all generation for that also.
$apis = @('callrouter', 'classifiedcore')

for ($i = 0; $i -lt $apis.Count; $i++){
    $apiPath = $apis[$i]
    $cmd = "cd ..\$apiPath; bee run -downdoc=true -gendoc=true;"

    Start-Process powershell $cmd
}

# Start Proxy
Start-Process powershell "cd ..\proxy; go build; .\proxy.exe"

# Start GULP
Start-Process powershell gulp

# Start the classifieds application
bee run


