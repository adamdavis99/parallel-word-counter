## Parallel word counter

Simple Go application of a parallel word counter, which reads from a txt file (tried to read from PDF, but no good PDF readers in Go exist at the time of publishing this code.

Its a simple CLI application, where you can give the txt file as argument, and the application will print the count of words in it, and the time taken to run it

It prints both sequential and parallel runs, here is one of the sample results with a 10MB file as given in the repo

```
Total word count: 1501560
Time taken in parallel: 17.510194ms
Total word count: 1501560
Time taken in sequential: 30.733677ms
```

It uses waitGroup and channels to do the parallelization. This is just a simple exercise to understand how powerful go's concurrency techniques are
