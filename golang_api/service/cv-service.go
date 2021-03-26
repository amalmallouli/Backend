package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"github.com/ydhnwb/golang_api/repository"
)

//CVService is a ....
type CVService interface {
	Insert(b dto.CVCreateDTO) entity.CV
	Update(b dto.CVUpdateDTO) entity.CV
	Delete(b entity.CV)
	All() []entity.CV
	FindByID(cvID uint64) entity.CV
	IsAllowedToEdit(userID string, cvID uint64) bool
}

type cvService struct {
	cvRepository repository.CVRepository
}

//NewCVService .....
func NewCVService(cvRepo repository.CVRepository) CVService {
	return &cvService{
		cvRepository: cvRepo,
	}
}

func (service *cvService) Insert(b dto.CVCreateDTO) entity.CV {
	cv := entity.CV{}
	err := smapping.FillStruct(&cv, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.cvRepository.InsertCV(cv)
	return res
}

func (service *cvService) Update(b dto.CVUpdateDTO) entity.CV {
	cv := entity.CV{}
	err := smapping.FillStruct(&cv, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.cvRepository.UpdateCV(cv)
	return res
}

func (service *cvService) Delete(b entity.CV) {
	service.cvRepository.DeleteCV(b)
}

func (service *cvService) All() []entity.CV {
	return service.cvRepository.AllCV()
}

func (service *cvService) FindByID(cvID uint64) entity.CV {
	return service.cvRepository.FindCVByID(cvID)
}

func (service *cvService) IsAllowedToEdit(userID string, cvID uint64) bool {
	b := service.cvRepository.FindCVByID(cvID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
