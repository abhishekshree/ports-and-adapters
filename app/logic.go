package app

import (
	"errors"
	"time"

	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrorRedirectNotFound = errors.New("Redirect not found")
	ErrorRedirectInvalid  = errors.New("Redirect is invalid")
)

type redirectService struct {
	repo RedirectRepository
}

func NewRedirectService(repo RedirectRepository) RedirectService {
	return &redirectService{repo}
}

func (s *redirectService) Find(code string) (*Redirect, error) {
	return s.repo.Find(code)
}

func (s *redirectService) Store(redirect *Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		return errs.Wrap(ErrorRedirectInvalid, "service.Redirect.Store")
	}

	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().Unix()
	return s.repo.Store(redirect)
}
