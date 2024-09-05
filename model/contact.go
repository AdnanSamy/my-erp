package model

import "time"

type Contact struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Type        int       `db:"type" json:"type"`
	CompanyName string    `db:"company_name" json:"company_name"`
	Address     string    `db:"address" json:"address"`
	VAT         string    `db:"vat" json:"vat"`
	JobPosition string    `db:"job_position" json:"job_position"`
	Phone       string    `db:"phone" json:"phone"`
	Mobile      string    `db:"mobile" json:"mobile"`
	Email       string    `db:"email" json:"email"`
	Website     string    `db:"website" json:"website"`
	TitleID     *int      `db:"title_id" json:"title_id"`
	Active      bool      `db:"active" json:"active"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	CreatedUID  *int      `db:"created_uid" json:"created_uid"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	UpdatedUID  *int      `db:"updated_uid" json:"updated_uid"`
}
