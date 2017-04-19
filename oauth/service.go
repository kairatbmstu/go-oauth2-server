package oauth

import (
	"github.com/adam-hanna/go-oauth2-server/config"
	"github.com/adam-hanna/go-oauth2-server/oauth/roles"
	"github.com/jinzhu/gorm"
)

// Service struct keeps objects to avoid passing them around
type Service struct {
	cnf          *config.Config
	db           *gorm.DB
	allowedRoles []string
}

// InitService starts a new Service instance
func (s *Service) InitService(cnf *config.Config, db *gorm.DB) {
	s = &Service{
		cnf:          cnf,
		db:           db,
		allowedRoles: []string{roles.Superuser, roles.User},
	}
}

// GetConfig returns config.Config instance
func (s *Service) GetConfig() *config.Config {
	return s.cnf
}

// RestrictToRoles restricts this service to only specified roles
func (s *Service) RestrictToRoles(allowedRoles ...string) {
	s.allowedRoles = allowedRoles
}

// IsRoleAllowed returns true if the role is allowed to use this service
func (s *Service) IsRoleAllowed(role string) bool {
	for _, allowedRole := range s.allowedRoles {
		if role == allowedRole {
			return true
		}
	}
	return false
}
