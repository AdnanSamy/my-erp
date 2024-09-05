package repository

import (
	"sam/my-erp/config"
	"sam/my-erp/model"
)

type ContactRepo struct {
	db *config.DB
}

func GetContactRepo(db *config.DB) *ContactRepo {
	return &ContactRepo{db}
}

func (repo *ContactRepo) Create(r *model.Contact) error {
	query := `
		INSERT INTO contact (
			"name", "type", company_name, "address", vat, job_position, phone, mobile, email, website, title_id, active, created_uid, updated_uid
		) VALUES
		(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12,
			$13,
			$14 
		)
	`

	_, err := repo.db.Exec(query, r.Name, r.Type, r.CompanyName, r.Address, r.VAT, r.JobPosition, r.Phone, r.Mobile, r.Email, r.Website, r.TitleID, true, 0, 0)

	return err
}

func (repo *ContactRepo) All() ([]*model.Contact, error) {
	var contacts []*model.Contact

	query := `
		SELECT 
			id,
			name,
			type,
			company_name,
			address,
			vat,
			job_position,
			phone,
			mobile,
			email,
			website,
			title_id,
			active,
			created_at,
			created_uid,
			updated_at,
			updated_uid
		FROM 
			contact
		WHERE 
			active = true
	`

	err := repo.db.Select(&contacts, query)

	return contacts, err
}
