
### Architecture Diagram


![locus architecture](images/architecture.png)


---

[Web Server](https://github.com/itsjwala/locus/tree/master/web) serving the HTTP request runs on containers **alongside** the language containers thats gonna being spawned based from Web container depending on incoming request from client. For this to work we need to have that environment setup inside **web container** to interact with Docker host, we do this by mounting [Docker Socket](https://medium.com/better-programming/about-var-run-docker-sock-3bfd276e12fd) inside our web container in [rw](https://github.com/itsjwala/locus/blob/master/Makefile#L6) mode.


---


### Moving Parts

#### Locus Runner

#### Docker Images for Languages which holds runtime

#### Web Server 

#### Frontend Web Client






