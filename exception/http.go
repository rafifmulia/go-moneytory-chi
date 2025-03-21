package exception

// Exception message should be a real error message (debug level message).
// RootMiddleware will filter out th message by message level before it reaches the client.

func NewBadRequestException(s string) *BadRequestException { return &BadRequestException{s} }

type BadRequestException struct{ s string }

func (e BadRequestException) Error() string { return e.s }

func NewUnauthorizedException(s string) *UnauthorizedException { return &UnauthorizedException{s} }

type UnauthorizedException struct{ s string }

func (e UnauthorizedException) Error() string { return e.s }

func NewNotFoundException(s string) *NotFoundException { return &NotFoundException{s} }

type NotFoundException struct{ s string }

func (e NotFoundException) Error() string { return e.s }

func NewUnprocessableEntityException(s string) UnprocessableEntityException {
	return UnprocessableEntityException{s}
}

type UnprocessableEntityException struct{ s string }

func (e UnprocessableEntityException) Error() string { return e.s }

func NewInternalServerErrorException(s string) InternalServerErrorException {
	return InternalServerErrorException{s}
}

type InternalServerErrorException struct{ s string }

func (e InternalServerErrorException) Error() string { return e.s }
