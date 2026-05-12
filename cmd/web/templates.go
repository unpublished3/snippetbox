package main

import "snippetbox.unpublished3/internal/models"

type templateData struct {
	Snippet models.Snippet
	Snippets []models.Snippet
}