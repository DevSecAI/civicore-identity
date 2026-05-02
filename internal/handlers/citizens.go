// CIV-SAST-001: SQL injection via fmt.Sprintf.
// CIV-SAST-009: no auth check / scope check on by-id route.
package handlers

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func ListCitizens(c *gin.Context) {
	region := c.Query("region")
	conn, _ := pgx.Connect(context.Background(), "postgres://civicore:Civicore2024!@db:5432/civicore?sslmode=disable")
	defer conn.Close(context.Background())

	// CIV-SAST-001: region interpolated into SQL.
	query := fmt.Sprintf("SELECT id, surname FROM citizens WHERE region = '%s'", region)
	rows, err := conn.Query(context.Background(), query)
	if err != nil { c.JSON(500, gin.H{"error": err.Error()}); return }
	defer rows.Close()

	var out []gin.H
	for rows.Next() {
		var id, surname string
		_ = rows.Scan(&id, &surname)
		out = append(out, gin.H{"id": id, "surname": surname})
	}
	c.JSON(200, out)
}

func GetCitizen(c *gin.Context) {
	// CIV-SAST-009: no JWT/role check; returns full PII for any id.
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "national_id": "TODO-fetch", "address": "TODO-fetch"})
}
