package location

import (
	mesg "github.com/mesg-foundation/go-service"
	maxminddb "github.com/oschwald/maxminddb-golang"
)

// LocationService is a MESG service to find locations of IP addresses.
type LocationService struct {
	s *mesg.Service

	dbPath string
	db     *maxminddb.Reader
}

// New creates a new location service with given mesg service and maxmindDBPath.
func New(service *mesg.Service, maxmindDBPath string) *LocationService {
	return &LocationService{
		s:      service,
		dbPath: maxmindDBPath,
	}
}

// Start starts the location service.
func (s *LocationService) Start() error {
	if err := s.setupLocationDB(); err != nil {
		return err
	}
	return s.listenTasks()
}

const (
	locateTaskKey = "locate"
)

func (s *LocationService) listenTasks() error {
	return s.s.Listen(
		mesg.Task(locateTaskKey, s.locateHandler),
	)
}

// Close gracefully closes the location service.
func (s *LocationService) Close() error {
	defer s.db.Close()
	return s.s.Close()
}

func (s *LocationService) setupLocationDB() error {
	var err error
	s.db, err = maxminddb.Open(s.dbPath)
	return err
}
