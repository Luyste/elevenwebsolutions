package handlers

import (
	"github.com/labstack/echo/v4"
)

type pageData struct {
	Stack []string
}

func Home(c echo.Context) error {
	data := pageData{
		Stack: []string{
			"Next.js · TypeScript",
			"Postgres · Prisma",
			"RAG · Embeddings",
			"Stripe · Auth",
			"Figma · Linear",
			"Edge · Vercel",
		},
	}
	return c.Render(200, "layout", data)
}
