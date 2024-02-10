package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	// validate:"required" หมายถึงข้อมูลที่รับเข้ามาจะต้องไม่เป็นค่าว่างหรือ nil
	// db คือ การทำ ORM (Object-Relational Mapping)
	ID             uuid.UUID     `json:"_id" db:"id"`
	First_Name     *string       `json:"first_name" db:"first_name" validate:"required,min=2,max=30"`
	Last_Name      *string       `json:"last_name"  db:"last_name"  validate:"required,min=2,max=30"`
	Email          *string       `json:"email"      db:"email"      validate:"email,required"`
	Password       *string       `json:"password"   db:"password"   validate:"required,min=6"`
	Phone          *string       `json:"phone"      db:"phone"      validate:"required"`
	Token          *string       `json:"token"      db:"token"`
	Refresh_Token  *string       `json:"refresh_token" db:"refresh_token"`
	Created_At     time.Time     `json:"created_at" db:"created_at"`
	Updated_At     time.Time     `json:"updated_at" db:"updated_at"`
	UserCart       []ProductUser `json:"usercart" db:"usercart"`
	Address_Detail []Address     `json:"address" db:"address"`
}

type Product struct {
	// uint64 ชนิดข้อมูลที่ใช้เก็บค่าจำนวนเต็มบวก ซึ่งมีขนาด 64 บิตหรือ 8 ไบต์ มีขนาดสูงสุด 2^64 - 1 และต่ะสุด คือ 0 ใช้เก็บ จำนวนเงิน, จำนวนตำแหน่งของสินค้า, หรือขนาดของข้อมูลในไฟล์
	// uint8 ชนิดข้อมูลที่ใช้เก็บค่าจำนวนเต็มบวก มีขนาด 8 บิตหรือ 1 ไบต์ มีค่าสูงสุดเท่ากับ 255 (2^8 - 1) และค่าต่ำสุดเท่ากับ 0 ใช้เก็บ อายุ, จำนวนเงินขนาดเล็ก เป็นต้น
	Product_ID   uuid.UUID `db:"id"`
	Product_Name *string   `json:"product_name" db:"product_name"`
	Price        *uint64   `json:"price" db:"price"`
	Rating       *uint8    `json:"rating" db:"rating"`
	Image        *string   `json:"image" db:"image"`
}

type ProductUser struct {
	Product_ID   uuid.UUID `db:"id"`
	Product_Name *string   `json:"product_name" db:"product_name"`
	Price        *uint64   `json:"price" db:"price"`
	Rating       *uint8    `json:"rating" db:"rating"`
	Image        *string   `json:"image" db:"image"`
}

type Address struct {
	Address_ID uuid.UUID `db:"id"`
	House      *string   `json:"house_name" db:"house_name"`
	Street     *string   `json:"street_name" db:"street_name"`
	City       *string   `json:"city_name" db:"city_name"`
	PinCode    *string   `json:"pinCode" db:"pinCode"`
}

type Order struct {
	Order_ID      uuid.UUID     `db:"id"`
	Order_Cart    []ProductUser `json:"order_list" db:"order_list"`
	Orderered_At  time.Time     `json:"ordered_on" db:"ordered_on"`
	TotalPrice    int           `json:"total_price" db:"total_price"`
	Discount      *int          `json:"discount" db:"discount"`
	PaymentMethod Payment       `json:"payment_method" db:"payment_method"`
}

type Payment struct {
	Digital bool `json:"digital" db:"digital"`
	COD     bool `json:"cod" db:"cod"`
}
