# TODO

- Where should []byte be converted into specific message types for the game loop?
  - Where should the `broker` package live
- Messages should only be sent to the specific zone in which they are applicable, clients should only subscribe to one zone at a time
  - Exception: if you are on a planet (child zone of system), you should be notified when something occurs in your system
  - Global messages can be broadcast to all players
  - Group messages can be broadcast to specific groups of zones (i.e. some kind of event occurring nearby)
