package service

// Service defines the interface for the content service.
type Service interface {
	GetUpdatedContent() (string, error)
}

// ContentService implements the Service interface.
type ContentService struct{}

// NewContentService creates a new instance of ContentService.
func NewContentService() Service {
	return &ContentService{}
}

// GetUpdatedContent returns the updated content for the /update endpoint.
func (cs *ContentService) GetUpdatedContent() (string, error) {
	// In a real application, you would fetch the updated content from a database,
	// an external API, or another source. For this example, we'll return static text.
	return "<p>Content has been updated!</p>", nil
}
