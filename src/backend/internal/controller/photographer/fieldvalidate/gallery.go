package fieldvalidate

import (
	"errors"

	"github.com/Roongkun/software-eng-ii/internal/model"
)

func CreateGallery(input model.GalleryInput) []error {
	fieldErrs := []error{}

	if input.Name == nil {
		fieldErrs = append(fieldErrs, errors.New(
			"the name of the new gallery must be provided",
		))
	}
	if input.Location == nil {
		fieldErrs = append(fieldErrs, errors.New(
			"the location of the new gallery must be provided",
		))
	}
	if input.Price == nil {
		fieldErrs = append(fieldErrs, errors.New(
			"the price of the new gallery must be provided",
		))
	}

	return fieldErrs
}

func UpdateGallery(input model.GalleryInput) []error {
	fieldErrs := []error{}

	if input.Name == nil && input.Price == nil && input.Location == nil {
		fieldErrs = append(fieldErrs, errors.New(
			"one of the gallery fields must be changed",
		))
	}

	return fieldErrs
}
