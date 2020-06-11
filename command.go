package main

type CommandId int

const (
	CMD_NICK CommandId =iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

type command struct {
	id CommandId
	client *client
	args []string
}