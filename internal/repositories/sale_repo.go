package repositories

import (
	"ArepasSA/internal/models"
	"time"

	"gorm.io/gorm"
)

type SaleRepository struct {
	*BaseRepository
}

func NewSaleRepository(db *gorm.DB) *SaleRepository {
	return &SaleRepository{NewBaseRepository(db)}
}

func (r *SaleRepository) Create(sale *models.Sale) error {
	return r.db.Create(sale).Error
}

func (r *SaleRepository) FindByID(id uint) (*models.Sale, error) {
	var sale models.Sale
	err := r.db.Preload("Items").First(&sale, id).Error
	return &sale, err
}

func (r *SaleRepository) FindAll(startDate, endDate time.Time) ([]models.Sale, error) {
	var sales []models.Sale
	query := r.db.Preload("Items")

	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("date BETWEEN ? AND ?", startDate, endDate)
	}

	err := query.Order("date DESC").Find(&sales).Error
	return sales, err
}

func (r *SaleRepository) GetSalesByHour() (map[int]int, error) {
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

func (r *SaleRepository) GetTopProducts(limit int) ([]models.ProductSales, error) {
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

func (r *SaleRepository) GetLeastSoldProducts(limit int) ([]models.ProductSales, error) {
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

func (r *SaleRepository) GetSalesByPriceRange() ([]models.PriceRangeSales, error) {
	var ranges []models.PriceRangeSales

	err := r.db.Raw(`
		SELECT
		CASE
			WHEN total_amount < 50 THEN '0-50'
			WHEN total_amount BETWEEN 50 AND 100 THEN '50-100'
			WHEN total_amount BETWEEn 100 AND 200 THEN '100-200'
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

func (r *SaleRepository) AddComment(saleID uint, comment *models.SaleComment) error {
	comment.SaleID = saleID
	comment.CommentAt = time.Now()
	return r.db.Create(comment).Error
}

func (r *SaleRepository) GetComments(saleID uint, includeInternal bool) ([]models.SaleComment, error) {
	var comments []models.SaleComment
	query := r.db.Where("sale_id = ?", saleID)

	if !includeInternal {
		query = query.Where("is_internal = ?", false)
	}

	err := query.Order("comment_at DESC").Find(&comments).Error
	return comments, err
}
