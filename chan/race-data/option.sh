#! /bin/bash

# race的影响：
# NOTE: memory usage may increase by 5-10x and execution time by 2-20x.

# The race detector writes its report to a file named log_path.pid
# The special names stdout and stderr cause reports to be written to standard output and standard error, respectively.
# export GORACE="log_path=/tmp/gr.err"
# export GORACE="log_path=stdout"
# go run -race race.go


export GORACE="strip_path_prefix=/Users/dxyan06/git/github/golang-learning"
go run -race race.go
