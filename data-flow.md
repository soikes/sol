## Client / Sever data flow / handshake process

1. Client logs in
2. Client asks server for game assets
3. Client connects to websocket authenticated
4. Server creates unique client id - ties to session JWT
5. Client id returned to client and client saves ID
6. Player information is loaded from DB or initialized
  - Current Zone
  - Coordinates
  - Credits
  - Explored Systems / Waypoints
  - Ship type
  - Inventory
  - Colonies
7. Server subscribes to client ID inputs
8. Server spawns Entity for any client entities that will respond to inputs
9. Server subscribes entities to relevant player inputs
10. Server begins publishing world state to client
  - Location / behaviour of planets in zone (transform, physics components)
  - Other players
  - NPCs
  - Other player artifacts
    - Colonies
11. Client begins publishing messages (mostly inputs) to server

## Moving between zones

Every solar system is instanced with its own game loop and players can jump between solar systems (warp gate, drive, etc). Doing so will unsubscribe the client from the current zone updates and subscribe it to the new zone updates. Players can also land on planets, which are distinct zones as well.

A sol galaxy graph will look like:

- 1-to-1 relationship between solar systems
- 2 or 3 warp gates per solar system (each solar system is a zone)
- N planets per system (landing takes to new single zone)
- N stars per system (cannot land, high gravity, give damage)
- TBD

## Server start / setup process

Before a dedicated Sol server runs and serves the game, a galaxy must be created with a certain number of systems (zones). An `atom` target will be used with paramters to generate a galaxy with:

1. Number of zones
2. Number of planets
3. Connectivity between zones
4. TBD

This will populate a graph data structure (storage TBD) that describes the state of a galaxy and its zones/systems.
 - Storage: unique DB in cockroachdb?
 - What does the schema look like

After this stage, a Sol server can be run pointed to the generated and stored galaxy and players can join and begin to interact with the world.


## Message brokering

- Zones emit messsages that are relevant to the zone itself
- Clients subscribe to one zone at a time


- Client
  - chan incoming []byte
  - chan outgoing []byte

- Hub 
  - chan incoming []byte
  - chan broadcast []byte

- Marshaler
  - chan incoming []byte
  - chan inputsIn message
  - ...etc

  - chan outgoing []byte
  - chan transformsOut message
  - ...etc

- Zone
  - chan incoming message<Type>
  - chan outgoing message<type>

- Input Component
  - queue input message<Input>


- Game
  - Has Zones