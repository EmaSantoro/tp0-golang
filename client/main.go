package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()
	log.Println("Hola soy un log") // loggear "Hola soy un log" usando la biblioteca log

	globals.ClientConfig = utils.IniciarConfiguracion("client/config.json")

	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}

	//LOGGEAR EL VALOR DE LA CONFIGURACION MENSAJE
	log.Printf("Mensaje: %s", globals.ClientConfig.Mensaje)

	// loggeamos el valor de la config
	log.Printf("Config: %+v", globals.ClientConfig)

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)
	//EnviarMensaje(ip string, puerto int, mensajeTxt string)

	// leer de la consola el mensaje
	utils.LeerConsola()

	// generamos un paquete y lo enviamos al servidor
	utils.GenerarYEnviarPaquete()
}
