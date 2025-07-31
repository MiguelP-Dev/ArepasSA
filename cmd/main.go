package main

import (
	"ArepasSA/internal/config"
	"ArepasSA/internal/handlers"
	"ArepasSA/internal/models"
	"ArepasSA/internal/repositories"
	"ArepasSA/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Configuración inicial
	cfg, err := config.NewConfig("sales.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Migraciones automáticas
	if err := cfg.DB.AutoMigrate(
		&models.User{},
		&models.Supplier{},
		&models.Product{},
		&models.Combo{},
		&models.ComboItem{},
		&models.Sale{},
		&models.SaleItem{},
		&models.Alert{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Inicializar repositorios
	productRepo := repositories.NewProductRepository(cfg.DB)
	saleRepo := repositories.NewSaleRepository(cfg.DB)
	comboRepo := repositories.NewComboRepository(cfg.DB)
	alertRepo := repositories.NewAlertRepository(cfg.DB)
	reportRepo := repositories.NewReportRepository(cfg.DB) // Nuevo

	// Inicializar servicios
	productService := services.NewProductService(productRepo)
	saleService := services.NewSaleService(saleRepo, productRepo, comboRepo)
	comboService := services.NewComboService(comboRepo, productRepo)
	reportService := services.NewReportService(reportRepo) // Actualizado
	alertService := services.NewAlertService(alertRepo, productRepo)

	// Inicializar handlers
	productHandler := handlers.NewProductHandler(productService)
	saleHandler := handlers.NewSaleHandler(saleService)
	comboHandler := handlers.NewComboHandler(comboService)
	reportHandler := handlers.NewReportHandler(reportService) // Actualizado
	alertHandler := handlers.NewAlertHandler(alertService)

	// Iniciar monitoreo de alertas en segundo plano
	alertHandler.StartAlertMonitor()

	// Configurar router
	r := gin.Default()

	// Rutas de productos
	productRoutes := r.Group("/products")
	{
		productRoutes.POST("/", productHandler.CreateProduct)
		productRoutes.GET("/", productHandler.GetAllProducts)
		productRoutes.GET("/:id", productHandler.GetProduct)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
		productRoutes.PATCH("/:id/deactivate", productHandler.SoftDeleteProduct)
	}

	// Rutas de ventas
	saleRoutes := r.Group("/sales")
	{
		saleRoutes.POST("/", saleHandler.CreateSale)
		saleRoutes.GET("/", saleHandler.GetAllSales)
		saleRoutes.GET("/:id", saleHandler.GetSale)
		saleRoutes.POST("/:id/comments", saleHandler.AddComment)
		saleRoutes.GET("/:id/comments", saleHandler.GetComments)
	}

	// Rutas de combos
	comboRoutes := r.Group("/combos")
	{
		comboRoutes.POST("/", comboHandler.CreateCombo)
		comboRoutes.GET("/", comboHandler.GetAllCombos)
		comboRoutes.GET("/:id", comboHandler.GetCombo)
		comboRoutes.PUT("/:id", comboHandler.UpdateCombo)
		comboRoutes.DELETE("/:id", comboHandler.DeleteCombo)
		comboRoutes.POST("/:id/sell", comboHandler.SellCombo)
		comboRoutes.POST("/:id/sell-partial", comboHandler.SellPartialCombo)
	}

	// Rutas de reportes
	reportRoutes := r.Group("/reports")
	{
		reportRoutes.GET("/daily", reportHandler.GetDailySalesReport)
		reportRoutes.GET("/peak-hours", reportHandler.GetPeakHoursReport)
		reportRoutes.GET("/top-products", reportHandler.GetTopProductsReport)
		reportRoutes.GET("/least-sold", reportHandler.GetLeastSoldProductsReport)
		reportRoutes.GET("/price-ranges", reportHandler.GetPriceRangeReport)
	}

	// Rutas de alertas
	alertRoutes := r.Group("/alerts")
	{
		alertRoutes.GET("/active", alertHandler.GetActiveAlerts)
		alertRoutes.POST("/:id/resolve", alertHandler.ResolveAlert)
		alertRoutes.GET("/resolved", alertHandler.GetResolvedAlerts)
	}

	// Rutas de clientes
	clientRoutes := r.Group("/clients")
	{
		// ... (rutas CRUD existentes)
		clientRoutes.POST("/:id/comments", clientHandler.AddComment)
		clientRoutes.GET("/:id/preferences", clientHandler.GetPreferences)
	}

	// Iniciar servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
