package main

import (
	"embed"
	"github.com/glebarez/sqlite"
	"github.com/wailsapp/wails/v2"
	"gorm.io/gorm"
	"log"
	"net/http"
	"tcp_crm/controllers"
	"tcp_crm/models"
	"time"

	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist/spa
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func ServeHTTP() {
	fs := http.FileServer(http.Dir("storage"))

	// Manejar todas las solicitudes a / con el FileServer
	http.Handle("/", fs)

	// Iniciar el servidor en el puerto 3000
	log.Println("Servidor escuchando en :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	go ServeHTTP()
	db, err := gorm.Open(sqlite.Open("tcp_crm_db.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Product{}, &models.Supplier{}, &models.Purchase{}, &models.PurchaseDetail{})
	db.Exec("DELETE FROM products")
	db.Exec("DELETE FROM suppliers")
	db.Exec("DELETE FROM purchases")
	db.Exec("DELETE FROM purchase_details")
	// Create
	db.Create(&models.Product{
		ID:            1,
		Name:          "Zapatos",
		Description:   "Zapatos negros",
		PurchasePrice: "0",
		SalePrice:     "0",
		ImagePath:     "products/product.png",
	})
	db.Create(&models.Supplier{
		ID:      1,
		Name:    "Periquito Perez",
		Address: "Cienfuegos",
		Phone:   "5030258",
		Email:   "periquito@gmail.com",
	})
	db.Create(&models.Purchase{
		ID:         1,
		Date:       time.Time{},
		SupplierID: 1,
		TotalUnits: 5,
		TotalPrice: 100,
	})

	controllers.DB = db
	productController := controllers.NewProductController()
	supplierController := controllers.NewSupplierController()
	purchaseController := controllers.NewPurchaseController()

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:             "TCP-CRM",
		Width:             1000,
		Height:            700,
		MinWidth:          800,
		MinHeight:         600,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: NewFileLoader(),
		},
		Menu:             nil,
		Logger:           nil,
		LogLevel:         logger.DEBUG,
		OnStartup:        app.startup,
		OnDomReady:       app.domReady,
		OnBeforeClose:    app.beforeClose,
		OnShutdown:       app.shutdown,
		WindowStartState: options.Maximised,
		Bind: []interface{}{
			app,
			&models.Product{},
			productController,
			&models.Supplier{},
			supplierController,
			&models.Purchase{},
			purchaseController,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
			ZoomFactor:          1.0,
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: false,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "tcp_crm",
				Message: "",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
