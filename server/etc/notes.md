# Notes

- Every entity has an ID, looping through entire entity array to find proper ID will be too expensive: O(n) every loop
  - Alternatives:
    - ECS system - every system has an array of components it interacts with, and entity IDs are used as array indexes: O(1)
    - Map: entities are stored in a Map and thus can be referenced by their id in O(1) time


- What I have now:
  - Entity:
    - ID
    - Component list
  - Component:
    - State
    - Pending actions
    - Action functions on state
    - Stores reference to other components needed
  - System:
    - Not formalized, app holds single static input state, etc. Not suitable for multiplayer game...

- ECS:
  - Entity:
    - ID
  - Component:
    - State
    - Pending actions
  - System:
    - Holds lists of every component it can interact with
    - Communicates between components using entity ID as array index
    - Action functions on state in certain order


## Client Server Messaging

Types of inputs to server:
- Keypress

Types of outputs from server:
- All Components

1. Client: {"type":"keypress", "forward":true, "left":true}
2. 