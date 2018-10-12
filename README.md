# Selpg_GO

`Selpg_GO` is a `Go` version rewritten a `Linux` CLI named `selpg` and the original [source code](https://www.ibm.com/developerworks/cn/linux/shell/clutil/selpg.c) of `selpg` is written in `c`.

## Introduction

`Selgo_GO` can be used to extract only a specified range of pages from an input text file and then print the output file in the command line, a file or anywhere you like.

## File Structure

- selpg.go: the Go version of selpg 

- input_file: the file used to input

- output_file: the file used to output result

- error_file: the file used to output error record



## Code Analysis

1. Import required packages in the project.

```go
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	flag "pflag"
)
```

2. A struct which is the core of selpg.

```go
type sp_args struct {
	start_page  int
	end_page    int
	in_filename string
	page_len    int
	page_type   bool
	print_dest  string
}
```

3. Input the arguments of command line, check them for validity and then construct a `sp_args` struct based on the user's input arguments.

```go
func (sa *sp_args) process_args()
```

4. Process the input and output the result and error record into the specified locations accroding to the `sp_args`.

```go
func (sa sp_args) process_input()
```

## Test Case

Typng an instruction `$ sh ./test.sh` in command line. 

```bash
// test.sh
for i in {1..7200}
do
  echo $i >> input_file
done
```

Then we can get a text file consisting of 7200 numbers from 1 to 7200 and each number occupies one line.

### Case 1

input: `$ selpg -s1 -e1 input_file`

output:
```
// location: command line
1
2
...
72
E:\大三上各科\服务计算\selpg.exe: done
```

### Case 2

input: `$ selpg -s1 -e1 < input_file`

output:
```
// location: command line
1
2
...
72
E:\大三上各科\服务计算\selpg.exe: done
```

### Case 3

input:  `$ selpg -s1 -e1 input_file >output_file`

output:
```
// location: output_file
1
2
...
72

// location: command line
E:\大三上各科\服务计算\selpg.exe: done
```

### Case 4

input: `$ selpg -s1 -e1 input_file 2>error_file`

output:
```
// location: command line
1
2
...
72

// location: error_file
E:\大三上各科\服务计算\selpg.exe: done
```

### Case 5

input: `$ selpg -s1 -e10 -l1 input_file 2>error_file`

output:
```
// location: command line
1
2
...
10

// location: error_file
E:\大三上各科\服务计算\selpg.exe: done
```

## Reference

1. [开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html#artrelatedtopics)

2. [Golang reference](https://golang.org/pkg/)

3. [琦哥正义的帮♂助](https://github.com/SiskonEmilia/Selpg)