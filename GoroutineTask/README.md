# Parallel Task Executor

This project is a Go application designed to execute a list of tasks in parallel using a worker pool. The application supports specifying the maximum number of parallel tasks and a maximum error threshold, after which the processing is halted.

## Features

- Execute tasks in parallel using goroutines.
- Limit the number of concurrent tasks.
- Stop processing when a specified number of errors is reached.

## Installation

1. Make sure you have Go installed. You can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```sh
git clone https://github.com/yourusername/parallel-task-executor.git