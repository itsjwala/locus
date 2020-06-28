#!/usr/bin/env bash

#build runner image 
docker build -t itsjwala/locus_runner .

cd languages 

for language in `ls -d */ | cut -f1 -d'/'`
do  
    # build all languages image that contain our runner binary (from runner image)
    docker build -t "itsjwala/locus_runner-$language" ./$language 
done

# clean runner image
docker image rm itsjwala/locus_runner
