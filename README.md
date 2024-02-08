## Producer-Consumer-Replication

This Go program demonstrates a simple implementation of the Producer-Consumer problem using goroutines and channels.

### Code Overview

- **Producer Function**: 
  - `Producer` function generates integers from 0 to 10 and sends them to the channel.
  - It accepts a channel (`chan int`) and a WaitGroup (`*sync.WaitGroup`) to synchronize goroutines.

- **Consumer Function**:
  - `Consumer` function receives integers from the channel and prints them.
  - It uses a `select` statement to handle channel communication and a default case to break the loop when the channel is closed or empty.
  - Similar to the `Producer` function, it also accepts a channel and a WaitGroup.

- **Main Function**:
  - In the `main` function, a `sync.WaitGroup` is initialized to synchronize goroutines.
  - A buffered channel of integers is created to exchange data between the producer and consumer.
  - Goroutines for the producer and consumer are launched concurrently using `go` keyword.
  - `WaitGroup.Wait()` is used to wait for all goroutines to finish execution before exiting.

### Running the Program

To run the program, execute the following command:

```shell
go run main.go
```

### Conclusion

The Producer-Consumer problem is a classic synchronization challenge in concurrent programming. This program demonstrates how Go's goroutines and channels can be used to implement a simple solution to this problem, ensuring safe and efficient communication between producers and consumers in a concurrent environment.
