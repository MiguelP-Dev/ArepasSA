package repositories

import (
	"ArepasSA/internal/models"
	"time"

	"gorm.io/gorm"
)

type ReportRepository struct {
	*BaseRepository
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{NewBaseRepository(db)}
}

func (r *ReportRepository) GetDailySales(date time.Time) (float64, int, error) {
	start := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	end := start.Add(24 * time.Hour)

	var total struct {
		TotalAmount float64
		Count       int
	}

	err := r.db.Model(&models.Sale{}).
		Select("SUM(total_amount) as total_amount, COUNT(*) as count").
		Where("date BETWEEN ? AND ?", start, end).
		Scan(&total).Error

	return total.TotalAmount, total.Count, err
}

func (r *ReportRepository) GetSalesByHour() (map[int]int, error) {
	var results []struct {
		Hour  int
		Count int
	}

	err := r.db.Model(&models.Sale{}).
		Select("strftime('%H', date) as hour, count(*) as count").
		Group("hour").
		Order("count DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	hourlySales := make(map[int]int)
	for _, result := range results {
		hourlySales[result.Hour] = result.Count
	}

	return hourlySales, nil
}

func (r *ReportRepository) GetTopProducts(limit int) ([]models.ProductSales, error) {
	var topProducts []models.ProductSales

	err := r.db.Table("sale_items").
		Select("product_id, SUM(quantity) as total_quantity, SUM(total_price) as total_sales").
		Where("product_id IS NOT NULL").
		Group("product_id").
		Order("total_quantity DESC").
		Limit(limit).
		Scan(&topProducts).Error

	return topProducts, err
}

func (r *ReportRepository) GetLeastSoldProducts(limit int) ([]models.ProductSales, error) {
	var leastSold []models.ProductSales

	err := r.db.Table("sale_items").
		Select("product_id, SUM(quantity) as total_quantity, SUM(total_price) as total_sales").
		Where("product_id IS NOT NULL").
		Group("product_id").
		Order("total_quantity ASC").
		Limit(limit).
		Scan(&leastSold).Error

	return leastSold, err
}

func (r *ReportRepository) GetSalesByPriceRange() ([]models.PriceRangeSales, error) {
	var ranges []models.PriceRangeSales

	err := r.db.Raw(`
		SELECT 
			CASE
				WHEN total_amount < 50 THEN '0-50'
				WHEN total_amount BETWEEN 50 AND 100 THEN '50-100'
				WHEN total_amount BETWEEN 100 AND 200 THEN '100-200'
				ELSE '200+'
			END as price_range,
			COUNT(*) as count,
			SUM(total_amount) as total
		FROM sales
		GROUP BY price_range
		ORDER BY price_range
	`).Scan(&ranges).Error

	return ranges, err
}
