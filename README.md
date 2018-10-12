# Selpg_GO

`Selpg_GO` is a `Go` version rewritten a `Linux` CLI named `selpg` and the original [source code](https://www.ibm.com/developerworks/cn/linux/shell/clutil/selpg.c) of `selpg` is written in `c`.

## Introduction

`Selgo_GO` can be used to extract only a specified range of pages from an input text file and then print the output file in the command line, a file or anywhere you like.

## Code

### 

```
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	flag "pflag"
)
```

## Test

1. `$ selpg -s1 -e1 input_file`

2. `$ selpg -s1 -e1 < input_file`

3. `$ selpg -s1 -e1 input_file >output_file`

4. `$ selpg -s1 -e1 input_file 2>error_file`

5. `$ selpg -s1 -e10 -l1 input_file 2>error_file`

## Reference

1. [开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html#artrelatedtopics)

2. [Golang reference](https://golang.org/pkg/)

3. [琦哥正义的帮♂助](https://github.com/SiskonEmilia/Selpg)