package main

//M2.1.1 Processes
/* Processes are an instance of a running program, several thing such as memory, registers, etc are unique to a process
typically called a context. An OS is what multiple processes to execute concurrently. From the process point of view, they have
the whole machine to themself (from CSE).

Processes are switched (scheduled) quickly which gives the illusion of parallelism, context needs to be switched as well
This context switching is long as it is saved in RAM - threads solve this by sharing parts of context - e.g. virtual mem, file descriptors
However they also have their own stack, data registers, code, etc

Goroutines are basically a thread in go, execute within a single OS (main) thread
Go Runtime Scheduler is responsible for scheduling the goroutines within the Go runtime */
