// TODO this is garbage and panics
package main

import (
	"fmt"
	"reflect"
)

var cmds = map[string]interface{}{}

func Call(fn string) error {
	fmt.Printf(`%#v`, cmds)
	f := reflect.ValueOf(cmds[fn])
	fmt.Printf(`%#v`, f)
	v := f.Call([]reflect.Value{})
	res := v[0].Interface()
	err := res.(error)
	return err
}

func Register(fname string, fn interface{}) {
	cmds[fname] = fn
}