package agent

import (
	"github.com/azme12/ai-agent-project/internal/api"
	"github.com/azme12/ai-agent-project/internal/config"
	"github.com/azme12/ai-agent-project/pkg/logger"
)

type Service struct {
	config    *config.Config
	logger    *logger.Logger
	handler   *Handler
	scheduler *Scheduler
	calendar  *api.CalendarService
	email     *api.EmailService
	nlp       *api.GeminiService
}

func NewService(cfg *config.Config, log *logger.Logger, cal *api.CalendarService, em *api.EmailService, nlp *api.GeminiService) *Service {
	handler := NewHandler(cfg, log, cal, em, nlp)
	scheduler := NewScheduler(cfg, log)

	// Set API services in scheduler
	scheduler.calendar = cal
	scheduler.email = em

	return &Service{
		config:    cfg,
		logger:    log,
		handler:   handler,
		scheduler: scheduler,
		calendar:  cal,
		email:     em,
		nlp:       nlp,
	}
}

func (s *Service) Start() error {
	s.logger.Info("Starting AI Agent Service")

	// Start scheduler
	if err := s.scheduler.Start(); err != nil {
		return err
	}

	s.logger.Info("AI Agent Service started successfully")
	return nil
}

func (s *Service) Stop() error {
	s.logger.Info("Stopping AI Agent Service")

	if err := s.scheduler.Stop(); err != nil {
		return err
	}

	s.logger.Info("AI Agent Service stopped")
	return nil
}

func (s *Service) ProcessTask(task string) error {
	return s.handler.ProcessTask(task)
}

func (s *Service) ProcessNLPCommand(command string) (string, error) {
	return s.nlp.ProcessCommand(command)
}
