# chatRoomCli-go

TCP chat server using Go, which enables clients to communicate with each other
Once the user connects to the chat server using telnet command line program, they can use the following commands to talk to the server:

- `/nick <name>` - get a name, otherwise user will stay anonymous.
- `/join <name>` - join a room, if room doesn't exist, the new room will be created. User can be only in one room at the same time.
- `/rooms` - show list of available rooms to join.
- `/msg	<msg>` - broadcast message to everyone in a room.
- `/quit` - disconnects from the chat server.


I would like to highlight that this program is not final yet and it misses few very important items.
- Validation of message body: commands, arguments, body size.
- State: current server is stateless, meaning if it shuts down - all connections will be closed. 
