package hotels

import (
	"context"
	"hotel-booking/internal/models"

	"github.com/google/uuid"
)

// Hotels postgres repository
type PGRepository interface {
	CreateHotel(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error)
	UpdateHotel(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error)
	UpdateHotelImage(ctx context.Context, hotelID uuid.UUID, imageURL string) error
	GetHotelByID(ctx context.Context, hotelID uuid.UUID) (*models.Hotel, error)
	// GetHotels(ctx context.Context, query *utils.PaginationQuery) (*models.HotelsList, error)
}
