package entities

import (
	"time"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Email       string         `json:"email" gorm:"uniqueIndex;not null"`
	Username    string         `json:"username" gorm:"uniqueIndex;not null"`
	Password    string         `json:"-" gorm:"not null"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Phone       string         `json:"phone"`
	Avatar      string         `json:"avatar"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	IsVerified  bool           `json:"is_verified" gorm:"default:false"`
	Role        string         `json:"role" gorm:"default:user"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// Relationships
	Listings    []Listing      `json:"listings,omitempty" gorm:"foreignKey:UserID"`
	Favorites   []Favorite     `json:"favorites,omitempty" gorm:"foreignKey:UserID"`
}

// Category represents a listing category with specialized types
type Category struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Slug        string         `json:"slug" gorm:"uniqueIndex;not null"`
	Description string         `json:"description"`
	Icon        string         `json:"icon"`
	Type        string         `json:"type" gorm:"not null"` // "emlak", "arac", "ikinci_el"
	ParentID    *uint          `json:"parent_id"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// Relationships
	Parent      *Category      `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children    []Category     `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	Listings    []Listing      `json:"listings,omitempty" gorm:"foreignKey:CategoryID"`
}

// Listing represents a marketplace listing with category-specific fields
type Listing struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description" gorm:"type:text"`
	Price       float64        `json:"price" gorm:"not null"`
	Currency    string         `json:"currency" gorm:"default:TRY"`
	Location    string         `json:"location"`
	UserID      uint           `json:"user_id" gorm:"not null"`
	CategoryID  uint           `json:"category_id" gorm:"not null"`
	Status      string         `json:"status" gorm:"default:active"` // active, sold, inactive
	ViewCount   int            `json:"view_count" gorm:"default:0"`
	IsPromoted  bool           `json:"is_promoted" gorm:"default:false"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// Category-specific fields stored as JSON
	EmlakDetails *EmlakDetails `json:"emlak_details,omitempty" gorm:"type:json"`
	AracDetails  *AracDetails  `json:"arac_details,omitempty" gorm:"type:json"`
	IkinciElDetails *IkinciElDetails `json:"ikinci_el_details,omitempty" gorm:"type:json"`
	
	// Relationships
	User        User           `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Category    Category       `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Images      []ListingImage `json:"images,omitempty" gorm:"foreignKey:ListingID"`
	Favorites   []Favorite     `json:"favorites,omitempty" gorm:"foreignKey:ListingID"`
}

// EmlakDetails represents real estate specific fields
type EmlakDetails struct {
	PropertyType    string  `json:"property_type"`    // konut, is_yeri, arsa
	RoomCount       string  `json:"room_count"`       // 1+1, 2+1, 3+1, etc.
	Area            float64 `json:"area"`             // m2
	Floor           int     `json:"floor"`            // kat
	TotalFloors     int     `json:"total_floors"`     // toplam kat
	BuildingAge     int     `json:"building_age"`     // yapım yılı
	Furnished       bool    `json:"furnished"`        // eşyalı
	HasBalcony      bool    `json:"has_balcony"`      // balkon
	HasElevator     bool    `json:"has_elevator"`     // asansör
	HasParking      bool    `json:"has_parking"`      // otopark
	Heating         string  `json:"heating"`          // ısıtma
	DeedStatus      string  `json:"deed_status"`      // tapu durumu
}

// AracDetails represents vehicle specific fields
type AracDetails struct {
	Brand           string  `json:"brand"`            // marka
	Model           string  `json:"model"`            // model
	Year            int     `json:"year"`             // yıl
	Mileage         int     `json:"mileage"`          // km
	FuelType        string  `json:"fuel_type"`        // yakıt türü
	Transmission    string  `json:"transmission"`     // vites
	EngineSize      string  `json:"engine_size"`      // motor hacmi
	Color           string  `json:"color"`            // renk
	Condition       string  `json:"condition"`        // durumu
	BodyType        string  `json:"body_type"`        // kasa tipi
	HasAccident     bool    `json:"has_accident"`     // hasarlı
	IsChanged       bool    `json:"is_changed"`       // değişen
	OwnerCount      int     `json:"owner_count"`      // kaçıncı sahibi
	PlateCode       string  `json:"plate_code"`       // plaka il kodu
}

// IkinciElDetails represents second-hand items specific fields
type IkinciElDetails struct {
	Brand           string  `json:"brand"`            // marka
	Model           string  `json:"model"`            // model
	Condition       string  `json:"condition"`        // kullanım durumu
	WarrantyStatus  string  `json:"warranty_status"`  // garanti durumu
	Color           string  `json:"color"`            // renk
	Size            string  `json:"size"`             // beden/boyut
	Material        string  `json:"material"`         // malzeme
	Age             int     `json:"age"`              // yaş
	IsNegotiable    bool    `json:"is_negotiable"`    // pazarlık
}

// ListingImage represents images for a listing
type ListingImage struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ListingID uint           `json:"listing_id" gorm:"not null"`
	URL       string         `json:"url" gorm:"not null"`
	Alt       string         `json:"alt"`
	Order     int            `json:"order" gorm:"default:0"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	
	// Relationships
	Listing   Listing        `json:"listing,omitempty" gorm:"foreignKey:ListingID"`
}

// Favorite represents user favorites
type Favorite struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	ListingID uint           `json:"listing_id" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	
	// Relationships
	User      User           `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Listing   Listing        `json:"listing,omitempty" gorm:"foreignKey:ListingID"`
}

// Message represents messages between users
type Message struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	SenderID   uint           `json:"sender_id" gorm:"not null"`
	ReceiverID uint           `json:"receiver_id" gorm:"not null"`
	ListingID  *uint          `json:"listing_id"`
	Content    string         `json:"content" gorm:"type:text;not null"`
	IsRead     bool           `json:"is_read" gorm:"default:false"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
	
	// Relationships
	Sender     User           `json:"sender,omitempty" gorm:"foreignKey:SenderID"`
	Receiver   User           `json:"receiver,omitempty" gorm:"foreignKey:ReceiverID"`
	Listing    *Listing       `json:"listing,omitempty" gorm:"foreignKey:ListingID"`
}
