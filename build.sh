#!/bin/bash
export GOOS=linux
export GOARCH=amd64

#colors
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0;37m\n'

function doBuild {
    runPath=$(pwd)
    printf "${BLUE}Runpath:${RED} ${runPath} ${NC}"

    for currPath in "app" "api"
    do
        printf "${BLUE}Current Path:${RED}  ${currPath} ${NC}"
        workingPath=$runPath"/"$currPath
        printf "${BLUE}Working Path:${RED}  ${workingPath} ${NC}"

        for folder in $(ls $currPath)
        do
            slashPath=$currPath"/"
            appName="${folder/$slashPath/\"\"}"
            printf "${BLUE}Attempting to build:${RED}  ${appName} ${NC}"

            outPath=$runPath"/bin/"$appName
            exeName=$outPath"/"$appName

            sourcePath=$workingPath"/"$appName
            cd $sourcePath

            maingo="main.go"
            if [ -f $maingo ]
            then
                buildRes=$(go build -i -o $exeName)
                printf "${RED}${buildRes}${NC}"

                copyFolder $outPath "conf"
                copyFolder $outPath "static"
                copyFolder $outPath "views"

                # For development, copy test certs to "bin"
                copyFile $outPath "fullchain.pem" #Cert
                copyFile $outPath "privkey.pem"
            fi

            printf "${GREEN}Finsished building:${RED} ${appName} ${NC}"
        done
        cd $runPath
    done
}

function copyFolder {
    if [ -d $2 ];    then
        cp -r $2 $1
    fi
}

function copyFile {
    if [ -f $2 ];    then
        cp $2 $1
    fi
}

#Start
doBuild