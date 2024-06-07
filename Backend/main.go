package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JonnierMB/DW_S5/Backend/controllers"
	"github.com/JonnierMB/DW_S5/Backend/handlers"
	"github.com/JonnierMB/DW_S5/Backend/models"
	"github.com/JonnierMB/DW_S5/Backend/repository"
	"github.com/gorilla/mux"  //go get -u github.com/gorilla/mux
	"github.com/jmoiron/sqlx" //go get -u github.com/jmoiron/sqlx
	"github.com/lib/pq"       //go get -u github.com/lib/pq
)

/*
Implementación de servidor para una aplicación ficticia
que toma los vehiculo que se hacen en una red social
*/

func main() {
	//Creando un objeto de conexión a postgresSQL
	// conn, err := ConectarDB("url", "postgres") //Conexión vacia a la base de datos, depende de donde estará en la nube Elephant
	db, err := ConectarDB(fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable", os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")), "postgres")

	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
		return
	}

	/*
		Creando una instancia del tipo Repository del paquete repository
		se debe especificar el tipo de struct que va a manejar la base de datos
	*/
	//Handler para vehiculo
	repoVehicle, err := repository.NewRepository[models.Vehiculo](db) //Crear la base de datos
	if err != nil {
		log.Fatalln("error conectando a la base de datos vehiculo", err.Error())
		return
	}
	controllerVehicle, err := controllers.NewController(repoVehicle)
	if err != nil {
		log.Fatalf("fallo al crear una instancia de controller vehiculo: %s", err.Error())
		return
	}
	handlerVehicle, err := handlers.NewHandlerVehiculo(controllerVehicle)
	if err != nil {
		log.Fatalf("fallo al crear una instancia de handler vehiculo: %s", err.Error())
		return
	}
	//Handler para usuarios
	repoUser, err := repository.NewRepository[models.Usuario](db) //Crear la base de datos
	if err != nil {
		log.Fatalln("error conectando a la base de datos usuario", err.Error())
		return
	}
	controllerUser, err := controllers.NewControllerUser(repoUser)
	if err != nil {
		log.Fatalf("fallo al crear una instancia de controller usuario: %s", err.Error())
		return
	}
	handlerUser, err := handlers.NewHandlerUsuarios(controllerUser)
	if err != nil {
		log.Fatalf("fallo al crear una instancia de handler usuario: %s", err.Error())
		return
	}
	/********************************MULTIPLEXOR Y ENRUTADOR***********************************/
	//Permiten asociar una ruta a un metodo y a un Handler que atiende peticiones que vienen con el metodo
	/* router (multiplexador) a los endpoints de la API (implementado con el paquete gorilla/mux) */
	router := mux.NewRouter() //Definir objeto de tipo multiplexor Cuando la petición llegue, la petición llegará al mux
	// Rutas para vehículos
	router.HandleFunc("/vehicles", handlerVehicle.ListarVehiculos).Methods(http.MethodGet)
	router.HandleFunc("/vehicles/{id}", handlerVehicle.TraerVehiculos).Methods(http.MethodGet)       //Un unico vehiculo
	router.HandleFunc("/vehicles/{id}", handlerVehicle.ActualizarVehiculo).Methods(http.MethodPatch) //Modificar un parametro en especifico

	// Rutas para usuarios
	router.HandleFunc("/users", handlerUser.ListarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/users", handlerUser.CrearUsuario).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", handlerUser.TraerUsuarios).Methods(http.MethodGet)       //Un unico usuario
	router.HandleFunc("/users/{id}", handlerUser.ActualizarUsuario).Methods(http.MethodPatch) //Modificar un parametro en especifico

	//********************************************************************************************************************
	//*ASOCIAR EL SERVIDOR AL MULTIPLEXOR: EL SERVIDOR ES LA INSTANCIA QUE PERMITE ABRIR UN PUERTO Y QUEDARSE ESCUCHANDO**
	//********************************************************************************************************************
	//definir servidor que estará escuchando
	log.Fatal(http.ListenAndServe(":8080", router)) //Si se utilzia nil en vez de mux, se utiliza el multiplexor para los handler por defecto de GO
}

func ConectarDB(url, driver string) (*sqlx.DB, error) { //INSTALACIÓN go get -u github.com/lib/pq
	pgUrl, _ := pq.ParseURL(url)           //paquete pq, permite implementar funciones adicionales con el fin de interactuar con bases de datos de tipo postgres
	db, err := sqlx.Connect(driver, pgUrl) // driver: postgres
	if err != nil {
		log.Printf("fallo la conexion a PostgreSQL, error: %s", err.Error())
		return nil, err
	}

	log.Printf("Nos conectamos bien a la base de datos db: %#v", db)
	return db, nil
}
