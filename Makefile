lider:
	echo "Ejecutando lider"
	go run repo/SquidGame/main.go

namenode:
	echo "Ejecutando namenode"
	go run SquidGame/client/pozo.go

pozo:
	echo "Ejecutando pozo"
	go run SquidGame/client/pozo.go

datanode:
	echo "Ejecutando datanode"
	go run SquidGame/server/dataNode.go
	
jugador:
	echo "Ejecutando jugadores"
	go run SquidGame/client/jugador.go

