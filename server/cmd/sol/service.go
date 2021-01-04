package main

import (
	"context"
	// "fmt"

	"soikke.li/sol/config"
	// "soikke.li/sol/graph"
	// "soikke.li/sol/svc/universe"

	"github.com/rs/zerolog/log"
)

func cmdService(ctx context.Context, cfg *config.Config) error {
	log.Info().Msg(`running service`)

	// go func() {
	// 	// g := graph.Graph{}
	// 	// v1 := &graph.Vertex{Value: `hello`}
	// 	// g.Add(v1)
	// 	// v2 := &graph.Vertex{Value: `goodbye`}
	// 	// g.Add(v2)
	// 	// graph.Link(v1, v2, 1)
	// 	// v3 := &graph.Vertex{Value: `no`}
	// 	// g.Add(v3)
	// 	// graph.Link(v2, v3, 1)
	// 	// fmt.Println(`adjacent:`, graph.IsAdjacent(v1, v2))
	// 	// fmt.Println(`adjacent:`, graph.IsAdjacent(v2, v3))
	// 	// fmt.Println(`adjacent:`, graph.IsAdjacent(v1, v3))

	// 	// for _, n := range v1.Neighbours() {
	// 	// 	fmt.Println(`neighbour:`, n.Value)
	// 	// }

	// 	// fmt.Println(`graph:`)
	// 	// fmt.Println(g.ToString())

	// 	// found := g.Search(`no`, graph.MethodDFS)
	// 	// fmt.Println(`found:`, found)
	// 	// return
	// 	// u := universe.NewUniverse(100)
	// 	// fmt.Println(u.StarSystems.ToString())
	// }()

	go cfg.Sol.Game.Run(ctx)
	go cfg.Sol.Web.StartHTTP(ctx)
	return nil
}
