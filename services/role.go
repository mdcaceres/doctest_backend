package services

import (
	"errors"
	"github.com/mdcaceres/doctest/models"
)

func validateRoles(userRoles []string) error {
	for _, ur := range userRoles {
		for _, r := range models.AvailableRoles {
			if ur != r {
				return errors.New("invalid role")
			}
		}
	}

	return nil
}
