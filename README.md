# ChordProtocol_Go
Project for NYU Distributed Systems Course. Fall 2022


Whole implementation could be divided into following different parts
//1. Implement the  structs and function for base chord protocol without considering stablization and joining of new nodes.
/2. Mechanism to contact, send and fetch data from distributed nodes using google protocol buffer messaging system
3. Mechanism to test the implementation
//4. Functionality for join and removing of nodes
5. Mechanism to test the node crashes and other things using simulation
6. Implementation of addion of data to one of the machine

#  Testing mechanism -> If we hardcode the hash value as whatever bits we want, and even change m corresponding to it, we should see pretty accurate results. Let's see


//7. Crashing the machine used for addition of the data and check if that data can still be fetched or not correctly
8. Report creation


Details of each of the part explained above:
1. Part 1 -> Chord Ring, chord nodes, finger tables, keys etc
2. RPCs using TCP Protocol using go net library functionality
3. Testing the implementation by fetching the data from each of the node and checking if it is correct. Keys and values can be stored in a CSV that are accesible by that particular node and manually writing a diff to check if they are fetching the data correctly. Also, test the O(.) of data stored on that particular machine and whether it is being scaled similarly.
4. Implemention of join and crash
5. Simulating the crashes by manually crashing on of the appication and seeing the O(.) storage at particular node
6. Add data to one of the node and stablize that data that is make sure it is get transfered to in ordered fashion to rest of the machine.
7. Crash the node though which data was added and make sure one is still able to access that data
8. Write a report with details of all the implementation and graphs and simulation results.


Each processor will be of the type chordNode which will have a finger table, predecessor and successor and currPointer field.

All three of the predecessor, successor and currPointer will be of type node, which is basically the abstract for a different process, it will be having an ip address and port to contact it

Finger table will be having enties of type node

Pre


Github Reference Repo
https://github.com/cbocovic/chord

#Not implemented yet


Termil