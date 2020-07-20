
#### Links

1. Docker Golang SDK [godoc](https://godoc.org/github.com/docker/docker/client), [examples](https://docs.docker.com/engine/api/sdk/examples/)

2. [Docker best practices](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)

3. Docker tips and tricks from Docker Captain 👌

[![Watch the video](https://img.youtube.com/vi/woBI466WMR8/maxresdefault.jpg)](https://www.youtube.com/watch?v=woBI466WMR8)


4. [Golang learning resource](https://www.callicoder.com/categories/golang/)

5. my docker [daemon config](https://docs.docker.com/config/daemon/#configure-the-docker-daemon) file for reference

```json
{
  "debug": true,
  "experimental": false,
  "log-driver": "local",
  "log-opts": {
    "max-size": "10m",
    "max-file": "3"
  },
  "features": {
    "buildkit": true
  }
}
```
