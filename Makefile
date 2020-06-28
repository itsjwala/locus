#!/usr/bin/env make -f

run:
	cd runner && ./gen_images.sh && cd ..
	cd web && docker build . -t itsjwala/locus_web && cd ..
	docker run --rm --name locus_web -v /var/run/docker.sock:/var/run/docker.sock:rw itsjwala/locus_web

clear: 
	docker image prune --filter "label=stage=locus_web-builder" --force 
	docker rmi --force `docker images --filter=reference="itsjwala/*" -q | uniq` > /dev/null 2>&1  ||  :
