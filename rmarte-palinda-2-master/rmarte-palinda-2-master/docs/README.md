# Task 1

## Bug01.go
The program will get stuck when sending data to the channel. This is since go is waiting for an other goroutine to receive the value. 

This can either be fixed with buffered channels or as shown in the code by sending the data through a goroutine. This way the code wont get stuck on sending the data since it continues in parallell to the part where it's recieving data. 

## Bug02.go
The problem here is that the main loop is closed before the goroutine is done printing all the numbers. 

This can be fixed by adding another channel which the function sends a message through when done. The main function then waits for the message that the goroutine is done and only exists when it has received the message. 

# Task 2

* What happens if you switch the order of the statements
  `wgp.Wait()` and `close(ch)` in the end of the `main` function?

The channel will close before the producers are done writing to it, causing the program to panic since the producers are trying to write to a closed channel. 

* What happens if you move the `close(ch)` from the `main` function
  and instead close the channel in the end of the function
  `Produce`?

The channel will be closed after the first produce goroutine is done, resulting in the program panicing since the next producer will try to write to a closed channel. 

* What happens if you remove the statement `close(ch)` completely?

In this case nothing will happen since it will be garbage collected. It is only necessary to close a channel if it's important for the receiver goroutines to know when all data have been sent.

* What happens if you increase the number of consumers from 2 to 4?

Since prducers and consumers are about as fast the current limiting factor to the speed are the number of consumers. When increasing the number of consumers from 2 to 4 (same as the number of producers) this part of the code should run about twice as fast. 

* Can you be sure that all strings are printed before the program
  stops?

No, the program only waits for all the consumer routines to finish and for the for loop to finish. When the main function finishes it kills the remaining consumer routines.