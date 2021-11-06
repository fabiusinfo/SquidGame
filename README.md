# Laboratorio Sistemas Distribuidos
# SquidGame

### Integrantes
Fabián Arancibia 201573102-0

Javier Olivares 201373069-8

Katherine Salgado 201610515-8
	
### Desiciones tomadas en el Laboratorio:

- La organización de los procesos es la siguiente:
	- Máquina 1, 10.6.43.41: Líder (5672, 8080) y un Datanode (9000)
	- Máquina 2, 10.6.43.42: Namenode (8080)
	- Máquina 3, 10.6.43.43: Pozo (5672, 8080) y un Datanode (9000)
	- Máquina 4, 10.6.43.44: Jugadores y un Datanode (9000)
- Debido a una confusión con los términos "ronda" y "nivel", los archivos están identificados con el id de ronda que definimos como 1rv, 2tc y 3tn, y no con 1, 2, y 3.
- El jugador con id: 1 juega mediante la interfaz, los 15 jugadores restantes son bots manejados en la misma máquina virtual. 
- El monto del pozo se puede ver en cualquier momento de la ejecución dentro de un juego con -check-, aunque no se muestre al final de una ronda, se puede ver igual.

### Instrucciones de ejecución:

dist161:

	> cd SquidGame
	
	> make lider
dist161:

	> cd SquidGame
	
	> make datanode
dist162:

	> cd SquidGame
	> make namenode
dist163:
	> cd SquidGame
	> make pozo
dist163:
	> cd SquidGame
	> make datanode
dist164:
	> cd SquidGame
	> make jugador
dist164:
	> cd SquidGame
	> make datanode

Instrucciones para jugar:

Jugador -join-
Lider -start-
Lider -start-
Jugador -next-
x4 {
	Jugador -send- or -check-
	Jugador if -send- 1,10
	Lider -next-
}
Lider -start-
Jugador -next-
Jugador -send- or -check-
Jugador if -send- 1,4
Lider -start-
Jugador -next-
Jugador -send- or -check-
Jugador if -send- 1,10
Lider -finish-

