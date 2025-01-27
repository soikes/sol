# Sol

Sol is a web browser-based, isometric action RPG set in space. Like many others in the 24th centry you are an expatriate from Earth who sets out to explore the galaxy in search of fortune, opportunity and new uninhabited planets to settle.

Sol was also inteded to be a vehicle for learning the process of creating a fully web-based multiplayer game from scratch and gaining more hands-on experience with technologies like TypeScript, Three.js, CockroachDB and Go. It was also a good excuse to practice DSA, 3D Modeling and attempt to understand architectures like Entity Component System.

## Dependencies

- Go >= 1.20 (https://go.dev/doc/install)
- Bun >= 1.0 (https://bun.sh)

## Running Sol Locally

To run Sol on a local server, cd to `server` and run the following:

```
  bin/atom build all && bin/atom start crdb sol
```

By default Sol listens on http://localhost:9999 and is playable in a web browser.
