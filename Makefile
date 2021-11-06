lider:
	echo "Ejecutando lider"
	go run server/Lider.go

namenode:
	echo "Ejecutando namenode"
	go run client/nameNode.go

pozo:
	echo "Ejecutando pozo"
	go run client/pozo.go

datanode:
	echo "Ejecutando datanode"
	go run server/dataNode.go
	
jugador:
	echo "Ejecutando jugadores"
	go run client/jugador.go

abrir:
	cd SquidGame/
