package libs

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	once     *sync.Once  = &sync.Once{}
	mu       *sync.Mutex = &sync.Mutex{}
)

func init() {
	once.Do(func() {
		validate = validator.New(validator.WithRequiredStructEnabled())
	})
}

func ExportValidator() *validator.Validate {
	mu.Lock()
	defer mu.Unlock()
	return validate
}
