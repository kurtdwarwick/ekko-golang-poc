package core_test

import (
	// 	"context"
	"testing"
	// "go.uber.org/mock/gomock"
)

func TestOrganisationRepository_ValidateUniqueness(t *testing.T) {
	// 	controller := gomock.NewController(t)
	// 	organisationDao := access.NewMockOrganisationDAO(controller)

	// 	existingOrganisation := entities.Organisation{LegalName: "Existing Organisation"}

	// 	t.Run("should return an error if the organisation already exists", func(t *testing.T) {
	// 		newOrganisation := entities.Organisation{LegalName: existingOrganisation.LegalName}
	// 		organisationDao.EXPECT().Count(&newOrganisation, nil, context.TODO()).Return(int32(1), nil)

	// 		err := repositories.ValidateUniqueness(newOrganisation, organisationDao, nil, context.TODO())

	// 		if err == nil {
	// 			t.Errorf("expected an error")
	// 		}
	// 	})

	// 	t.Run("should not return an error if the organisation does not exist", func(t *testing.T) {
	// 		newOrganisation := entities.Organisation{LegalName: "New Organisation"}
	// 		organisationDao.EXPECT().Count(&newOrganisation, nil, context.TODO()).Return(int32(0), nil)

	// 		err := repositories.ValidateUniqueness(newOrganisation, organisationDao, nil, context.TODO())

	//		if err != nil {
	//			t.Errorf("expected no error, got %v", err)
	//		}
	//	})
}
