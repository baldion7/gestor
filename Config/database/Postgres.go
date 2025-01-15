package database

import (
	model "gestor/Model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var (
	db  *gorm.DB
	err error
)

// ConfigurarDB configura GORM y establece la conexión a la base de datos
func ConfigurarDB() error {
	// Cargar variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando el archivo .env")
	}
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME_DATABASE")
	dbPort := os.Getenv("DB_PORT")
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	AutomigrarDB()
	InsertDataInit()
	return nil
}

// AutomigrarDB realiza automáticamente las migraciones necesarias para los modelos
func AutomigrarDB() {
	err := db.AutoMigrate(
		&model.Brand{},
		&model.Carvings{},
		&model.Color{},
		&model.CutMovements{},
		&model.CutOrder{},
		&model.CutSize{},
		&model.Movement{},
		&model.Product{},
		&model.Reference{},
		&model.Role{},
		&model.Sizes{},
		&model.Suppliers{},
		&model.User{},
	)
	if err != nil {
		log.Fatal(err)
	}
}

// ObtenerDB devuelve la instancia de la base de datos
func ObtenerDB() *gorm.DB {
	return db
}

// InsertDataInit inserta datos iniciales en la base de datos
func InsertDataInit() {
	insertRoles()
	insertUser()
	inserterTablas()
}

// insertRoles inserta roles iniciales
func insertRoles() {
	role := model.Role{Name: "Administrador", Description: "Permisos totales"}
	db.FirstOrCreate(&role, model.Role{Name: role.Name})
}

// insertUser inserta un usuario inicial
func insertUser() {
	user := model.User{
		Name:      "admin",
		RoleId:    1,
		Password:  "$2a$10$0/ukBG8wUa4GefCH70yKzeCdVPtsoDmTicImURbOnPv2lwgCucPf2", // Contraseña cifrada
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.FirstOrCreate(&user, model.User{Name: user.Name, RoleId: user.RoleId, Password: user.Password, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt})
}

// inserterTablas inserta las tallas iniciales
func inserterTablas() {
	Sizes := []string{"U", "XXXS", "XXS", "XS", "S", "M", "L", "XL", "XXL", "XXXL", "2", "4", "6", "8", "10", "12", "14", "16", "18"}
	for _, talla := range Sizes {
		size := model.Sizes{Name: talla}
		db.FirstOrCreate(&size, model.Sizes{Name: talla})
	}
}
