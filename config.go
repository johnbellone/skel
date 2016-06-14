package main

import (
	skel "github.com/johnbellone/skel/skel"
	_ "github.com/mitchellh/mapstructure"
)

type Config struct {
	Author    string
	Email     string
	Variables []skel.Variable
}
