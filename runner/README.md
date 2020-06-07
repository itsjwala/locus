# locus_runner

this is cli runner for locus

usage

```sh
docker build -t runner:dev -f Dockerfile-runner . 

docker run runner:dev '{"Code":"\nfor i in range(100):\n    print(f\"This is number :-{i}\")\n","Language":"python"}'
```



