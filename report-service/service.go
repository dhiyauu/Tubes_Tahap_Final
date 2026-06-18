package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

// ==================== INTERFACE & SERVICE ====================

type ReportRepository interface {
	GetDailyReport(ctx context.Context, date string) (*DailyReport, error)
	GetProblems(ctx context.Context) ([]ProblemPackage, error)
	GetCourierPerformance(ctx context.Context, courierID int) (*CourierPerformance, error)
}

type ReportService struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

// Method lama (tetap pakai repo)
func (s *ReportService) GetDailyReport(ctx context.Context, date string) (*DailyReport, error) {
	return s.repo.GetDailyReport(ctx, date)
}

func (s *ReportService) GetProblems(ctx context.Context) ([]ProblemPackage, error) {
	return s.repo.GetProblems(ctx)
}

func (s *ReportService) GetCourierPerformance(ctx context.Context, courierID int) (*CourierPerformance, error) {
	return s.repo.GetCourierPerformance(ctx, courierID)
}

// ==================== FUNGSI PANGGIL API ====================

var httpClient = &http.Client{Timeout: 5 * time.Second}

func fetchOrders() ([]OrderItem, error) {
	url := os.Getenv("ORDER_SERVICE_URL") + "/orders"
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var orders []OrderItem
	if err := json.NewDecoder(resp.Body).Decode(&orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func fetchTrackings() ([]TrackingItem, error) {
	url := os.Getenv("TRACKING_SERVICE_URL") + "/trackings"
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var trackings []TrackingItem
	if err := json.NewDecoder(resp.Body).Decode(&trackings); err != nil {
		return nil, err
	}
	return trackings, nil
}

// ==================== METHOD LAPORAN BARU ====================

func (s *ReportService) GetStatusReport(ctx context.Context) ([]StatusReport, error) {
	trackings, err := fetchTrackings()
	if err != nil {
		return nil, err
	}
	count := make(map[string]int)
	for _, t := range trackings {
		count[t.Status]++
	}
	var result []StatusReport
	for status, total := range count {
		result = append(result, StatusReport{Status: status, Total: total})
	}
	return result, nil
}

func (s *ReportService) GetMonthlyReport(ctx context.Context, year int) ([]MonthlyReport, error) {
	orders, err := fetchOrders()
	if err != nil {
		return nil, err
	}
	monthly := make(map[string]int)
	for _, o := range orders {
		// Sesuaikan format waktu dengan response dari order-service
		t, err := time.Parse(time.RFC3339, o.CreatedAt)
		if err != nil {
			// Coba format lain kalau perlu
			t, err = time.Parse("2006-01-02 15:04:05", o.CreatedAt)
			if err != nil {
				continue
			}
		}
		if t.Year() == year {
			key := t.Format("2006-01")
			monthly[key]++
		}
	}
	var result []MonthlyReport
	for month, total := range monthly {
		result = append(result, MonthlyReport{Month: month, Total: total})
	}
	return result, nil
}