package service

import (
	"toychart/model"
	"toychart/repository"
)

type IPSeriesService struct {
	ipSeriesRepo repository.IPSeriesRepository
}

func NewIPSeriesService(ipSeriesRepo repository.IPSeriesRepository) *IPSeriesService {
	return &IPSeriesService{ipSeriesRepo: ipSeriesRepo}
}

func (s *IPSeriesService) Create(IPSeries *model.IPSeries) error {
	return s.ipSeriesRepo.Create(IPSeries)
}

func (s *IPSeriesService) GetById(id string) (*model.IPSeries, error) {
	return s.ipSeriesRepo.GetById(id)
}

func (s *IPSeriesService) GetBySeries(series string) (*model.IPSeries, error) {
	return s.ipSeriesRepo.GetBySeries(series)
}

func (s *IPSeriesService) CheckIPSeriesExists(series string) (bool, error) {
	return s.ipSeriesRepo.CheckIPSeriesExists(series)
}

func (s *IPSeriesService) GetbyIpId(ipId string) ([]*model.IPSeries, error) {
	return s.ipSeriesRepo.GetbyIpId(ipId)
}

func (s *IPSeriesService) Update(IPSeries *model.IPSeries) error {
	return s.ipSeriesRepo.Update(IPSeries)
}

func (s *IPSeriesService) Delete(id string) error {
	return s.ipSeriesRepo.Delete(id)
}
