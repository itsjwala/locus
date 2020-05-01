# locus_runner

this is cli runner for locus

this is basically go binary which takes code of any language and spits out stderr and stdout 



```sh
# for golang
$ locus_runner code.go

# for java
$ locus_runner code.java

# for c++
$ locus_runner code.cpp
```

output should be writing into stdout and stderr files respectively depending on BASEPATH given(maybe via environment variable?) 


note:
1. will support stdin later

