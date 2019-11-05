## Flights Company

### API
```
URI: http://localhost:8080/api/v1/destinations?search={name}
Method: GET

Example: http://localhost:8080/api/v1/destinations?search=Buenos Aires, Argentina 
```
### Instalacion

Pasos para la instalacion:   
    1. Clonar el proyecto  https://github.com/davidnastasi/flights-company.git
    2. Crear ejecutable con el comando go build -o ./bin/app ./cmd/flights-company/main.go  
    3. Ejecutar el script schema en una base de datos PostgreSQL
    4. Correr el ejectutable app en la carpeta bin
       Variables de entorno en caso de ser requeridas:
            - APP_SERVER_URL: url del servidor [ default: "localhost:8080" ]
            - APP_DB_DSN: dsn para postgres [ default "host=localhost port=5432 user=postgres dbname=flights password=s3cret sslmode=disable" ]  
            - APP_CLIENT_ID: id cliente foursquare [ defaul "HACVIHTUOMFKVK5HWQ0J0JCOKQAA2CSAVFS0LFQVN14EESS2" ] 
            - APP_CLIENT_SECRET: pass cliente foursquare [ default "50ITRVSKRB1GH2YWOBBQWZS5BEDVEIWN3Z2YABJEI454V2JZ" ] 
            - APP_RESERVATION_ENDPOINT: endpoint reservas [ default "https://brubank-flights.herokuapp.com/flight-reservations" ]
    
   
### Decisiones de arquitectura


- Se ha decidido utilizar una base de datos relacional (PostgreSQL) para almacenar cada una de las reservas que obtenidas del endpoint. 
- Para obtener las reservas se ha utilizado un cron que se ejecuta cada 15 segundos y almacenarlas en la base de datos.
- La tabla en la base de datos posee una estructura de dos tablas en la primera se almacena cada una de las ciudades y en la 
segunda tabla las reservas (fecha y id)
- No se ha decidido utilizar una tabla para almacenar los hoteles sino que se utilizo directamente la api de 
foursquare para obtener los hoteles
- En ambos caso se utilizo una cache en memoria con diferentes tiempos sindo el tiempo de expiracion de foursquare mas alto que 
que el de la base de datos. 
- En una primera iteracion se le pegaba tanto a la api como a la db. En una segunda iteracion se agrego las caches y en una tercera
se le agrego gorutinas para que sea mas performante a la hora de ejecutar las acciones y no sea bloqueante en el caso que 
se deba ir a la base de datos o salir a buscar los hoteles a la api.

### Futuras mejoras al sistema:
- En el caso de que la empresa agregue mas destinos (muchos mas destinos) habria que dividir por pais y ciudad, de momento son pocos destinos.
- El incremento de las reservas a medida que pasa el tiempo puede perjudicar la  performance de la base de datos, se podria pasar los 
registros con fecha menor a la actual pasarlos a un historico mediante algun proceso nocturno. (si es que el negocio lo permite)
- Dado el incremento mencionado anteriormente es necesario que a medida se agrege un paginado.
- Otra opcion podria ser recurrir a una base de datos NoSQL
- Se podria considerar utilizar otra chache mas robusta como ser redis. 
- En caso de que no se pueda tener disponibilidad hacia la base de datos el sistema deberia poder almacenar las reservas en algun tipo de medio
de almacenamiento y una vez que se posea disponibilidad poder cargarlas una vez que se posea disponibilidad.
- Actualmente las busquedas son realizadas con un match exacto deberian ser mas flexibles para poder realizar las busquedas. 
(Por ejemplo en caso de que no se encuentre retornar un error con posibles opciones, hacerlo case-insensitive , etc) 


### Consideraciones:
No se realizaron test sobre la api de reservas.





  
 


