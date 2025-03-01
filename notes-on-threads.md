
OS Threads -- Hardware Threads

- Context Switch

- Scheduling 

- Process 
- ID, Status, Resources(Files, IO)
- Memory (Code/Text Segment, Data Segment, Stack and Heap memory)
- PC,SP,Registers
- Context Switch
- 


a,b:=10,20 

c:= a * a + b * b +  2 * a * b

MOV r3 a
MUL r3 a 
MUL r4 b b
MUL r5 2 a
MUL r6 b r5
ADD r3 r3 r4
ADD r3 r3 r5
ADD r3 r3 r6


- Green Threads -- go routines
- 1. go routines very small in size
- 2. they are maintained/managed by go runtime(not by OS)
- 3. they are scheduled by go scheduler (not by OS)
- 4. each go routine is for about 2kb
- 5. memory is given to the go routine is based on the requirement instead of a chunk of memory
- 6. context switch is managed by go runtime


- Creates number of threads which are equal to number of cores
- each thread has some additional information so it is called as P (Process)
- when new G is created


-- G means Go routine
-- P means A thread + additional information
-- Global Queue 
-- every time a new Go routine is created , it is placed on a Global Queue
-- Every P contains a local queue
-- Work stealing 



