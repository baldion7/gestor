package Config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // Modo de ejecución de GIN
	r := gin.Default()

	// Configurar CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"} // Permitir cualquier origen
	// Agrega todos los métodos permitidos
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 12 * 3600 // Tiempo de almacenamiento en caché de las opciones CORS preflight

	r.Use(cors.New(corsConfig))

	return r
}
