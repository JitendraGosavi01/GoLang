Large programs are divided into smaller sub-programs. Programs which run their smaller components at the same time is known as concurrency.

#Goroutines
The parts of an application that run concurrently are called goroutines. Goroutines and channels are used for structuring concurrent programs.

A process is an independently executing entity running in a machine which runs in its own address space in memory. A process has threads which are simultaneously executing entities. Threads share the same address space of the process.

Goroutines are lightweight, much lighter than a thread. Goroutines run in the same address space, so access to shared memory must be synchronized; This can be done by sync package, but it is recommended to use channels to synchronize goroutines.

A goroutine is implemented as a function or method. It is called (invoked) with the 'go' keyword. When the goroutine finishes, nothing is returned to the caller function.