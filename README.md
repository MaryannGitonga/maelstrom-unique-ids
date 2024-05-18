### Challenge #2: Unique ID Generator

Tested using Maelstrom:
- It runs a 3-node cluster for 30 seconds,
- It requests new IDs at the rate of 1000 requests per second.
- It checks for total availability and induces network partitions during the test.
- It verifies that all IDs are unique.
