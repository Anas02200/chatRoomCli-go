package main

import (
	"log"
	"net"
)

type server struct {
	rooms map[string]*room
	 commands chan command
}
func newServer()*server  {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}

}
func (s *server)run()  {
	for cmd:=range s.commands{
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args[1])
		case CMD_JOIN:
			s.join(cmd.client, cmd.args[1])
		case CMD_ROOMS:
			s.listRooms(cmd.client)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("new client has connected :%s", conn.RemoteAddr().String())
	c:=&client{
		conn:conn,
		nick:"anonymous",
		commands: s.commands,
	}
c.readInput()
}

