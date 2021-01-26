package main

// M1.1.1 Parallel Execution
/* Programs run in parallel if they execute at exactly the same time - different from concurrency
Von Neumann Bottleneck is the delay created by accessing and utilising RAM which is slower than the CPU
Moore's law (or rather observation) is also increasingly untrue as clock speeds are beginning to be limited
due to factors such as power consumption, volt leakage and lower tolerance for noise (Dennard voltage scaling of chips), etc
Code can be made to execute on multiple cores to achieve speedup, however the impetus is on the programmer to split tasks
*/

// M1.1.2 Concurrent Execution
/* Two tasks are concurrent if the start and end times overlap - different from parallel (executing together at same time)
Parallel tasks must be executed on replicated hardware, concurrent tasks may be executed on the same hardware
Mapping of tasks to hardware is usually not handled by programmer except in certain exotic languages
Concurrency improves performance even without parallelism by hiding latency (utilising wait time during RAM access, I/O inputs, etc)
Concurrency is implicitly implemented by Go, no need to import package to handle
Difference between defining parallel tasks vs dictating it*/
