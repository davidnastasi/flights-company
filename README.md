## Flights Company

### API
```
URI: http://localhost:8080/api/v1/destinations?search={name}
Method: GET

Example: http://localhost:8080/api/v1/destinations?search=Buenos Aires, Argentina 
```
### Instalacion

Pasos para la instalacion:   
    1. Clonar el proyecto  
    2. Ejecutar el script de inicializacion  


### Decisiones de arquitectura

Se ha utilizado la base de datos PostgreSQL como base de datos relacional
Se ha decidido utilizar una base de datos relacional para almacenar cada una de las reservas que obtenidas del endpoint. 
Para obtener las reservas se ha utilizado un cron que se ejecuta cada 15 segundos y almacenarlas en la base de datos.
La tabla en la base de datos posee una estructura de dos tablas en la primera se almacena cada una de las ciudades y en la 
segunda tabla las reservas (fecha y id)
No se ha decidido utilizar una tabla para almacenar los hoteles sino que se utilizo directamente la api de 
foursquare para obtener los hoteles
En ambos caso se utilizo una cache en memoria con diferentes tiempos sindo el tiempo de expiracion de foursquare mas alto que 
que el de la base de datos
En una primera iteracion se le pegaba tanto a la api como a la db. En una segunda iteracion se agrego las caches y en una tercera
se le agrego gorutinas para que sea mas performante a la hora de ejecutar las acciones y no sea bloqueante en el caso que 
se deba ir a la base de datos o salir a buscar los hoteles a la api.




  
 


