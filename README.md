# network-stuff
Currently just a basic network logger lib for go

## Example

```
func (p *Handler)startChannel(conn net.Conn){
	defer handlers.HandlePanic()

	loggingconn := logger.NewIOLogger(conn, conn)
	decoder := json.NewDecoder(loggingconn)

	...

	err := decoder.Decode(&someobject)

	...
}

