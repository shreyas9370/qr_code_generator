package models

import "qr-code-generator/db"

type QRCode struct {
	ID     int64  `json:"id"`
	URL    string `json: "url"`
	QRCode []byte `json: "qr_code"`
}

func (q *QRCode) Save() (int64, error) {
	result, err := db.DB.Exec("INSERT INTO qr_codes (url, qr_code) VALUES (?, ?)", q.URL, q.QRCode)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	q.ID = id
	return id, nil
}

func (q *QRCode) FindByID(id string) error {
	return db.DB.QueryRow("SELECT id, url, qr_code FROM qr_codes WHERE id = ?", id).Scan(&q.ID, &q.URL, &q.QRCode)
}
