package content

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type ContentItem struct {
	ID        uuid.UUID
	AuthorID  int
	Title     string
	Body      string
	Status    Status
	Version   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewContentItem(authorID int, title, body string) (*ContentItem, error) {
	if authorID <= 0 {
		return nil, errors.New("authorID is required")
	}
	if title == "" || body == "" {
		return nil, errors.New("title and body are required")
	}
	now := time.Now().UTC()
	return &ContentItem{
		ID:        uuid.New(),
		AuthorID:  authorID,
		Title:     title,
		Body:      body,
		Status:    StatusDraft,
		Version:   1,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (c *ContentItem) MarkEnrichmentRequested() {
	c.Status = StatusEnrichmentRequested
	c.UpdatedAt = time.Now().UTC()
	c.Version += 1
}

func (c *ContentItem) MarkEnriched() {
	c.Status = StatusEnriched
	c.UpdatedAt = time.Now().UTC()
	c.Version += 1
}

func (c *ContentItem) MarkPublished() {
	c.Status = StatusPublished
	c.UpdatedAt = time.Now().UTC()
	c.Version += 1
}

func (c *ContentItem) MarkFailed() {
	c.Status = StatusFailed
	c.UpdatedAt = time.Now().UTC()
	c.Version += 1
}
