package templates

import "github.com/reaper47/recipya/internal/models"

// IndexData holds data to pass on to the index template.
type IndexData struct {
	Recipes []models.Recipe
}
