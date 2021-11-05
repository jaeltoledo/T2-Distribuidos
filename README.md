# T2-Distribuidos
## Squid-Game

Las máquinas virtuales están asociadas de la siguiente manera

| Máquina | Procesos |
|-------|-------- |
| dist69| Lider, DataNode |
| dist70| Pozo, DataNode|
| dist71| NameNode|
| dist72| DataNode, DataNode|

Es de suma importancia ejecutar en el siguiente orden las máquinas

| Lugar | Máquina |
|-------|-------- |
| 1| dist69|
| 2| dist70|
| 3| dist71|
| 4| dist72|

En cada máquina aparecerán las instrucciones correspondientes sobre como continuar.

### IMPORTANTE:

- Se asumirá que los jugadores son inteligentes, esto quiere decir que no eligiran valores extraños o valores fuera de los límites especificados en la tarea.
- El lider deberá contestar uno por uno a cada jugagador, para que así ninguno se quede atrás en el juego y tenga sentido la competencia, por favor, realizar este paso es de suma importancia.

