#!/bin/bash

# This scripts extract all routes from https://github.com/Binaryify/NeteaseCloudMusicApi/tree/master/module to a route.txt file
dir="NeteaseCloudMusicApi/module"
targetDir=$PWD


pushd $dir
for file in `ls`
do
    route=$(sed -n '/https\:\/\/.*/p' $file | cut -d '`' -f 2)
    striped=`echo $file | cut -d '.' -f 1`
    echo "$striped: $route" >> $targetDir/route.txt
done
popd