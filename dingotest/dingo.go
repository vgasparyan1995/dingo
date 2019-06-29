package dingotest

import (
	go_sub_pkg "github.com/elliotchance/dingo/dingotest/go-sub-pkg"
	"github.com/jonboulle/clockwork"
	"net/http"
	"os"
	"time"
)

type Container struct {
	AFunc            func(int, int) (bool, bool)
	Clock            clockwork.Clock
	CustomerWelcome  *CustomerWelcome
	DependsOnTime    func(ParsedTime time.Time) time.Time
	HTTPSignerClient *HTTPSignerClient
	Now              func() time.Time
	OtherPkg         *go_sub_pkg.Person
	OtherPkg2        go_sub_pkg.Greeter
	OtherPkg3        *go_sub_pkg.Person
	ParsedTime       func(value string) time.Time
	SendEmail        EmailSender
	SendEmailError   *SendEmail
	Signer           func(req *http.Request) *Signer
	SomeEnv          *string
	WhatsTheTime     *WhatsTheTime
	WithEnv1         *SendEmail
	WithEnv2         *SendEmail
}

var DefaultContainer = NewContainer()

func NewContainer() *Container {
	return &Container{DependsOnTime: func(ParsedTime time.Time) time.Time {
		service := ParsedTime
		return service
	}, Now: func() time.Time {
		service := time.Now()
		return service
	}, ParsedTime: func(value string) time.Time {
		service, err := time.Parse(time.RFC822, value)
		if err != nil {
			return time.Now()
		}
		return service
	}, Signer: func(req *http.Request) *Signer {
		service := NewSigner(req)
		return service
	}}
}
func (container *Container) GetAFunc() func(int, int) (bool, bool) {
	if container.AFunc == nil {
		service := func(a, b int) (c, d bool) {
			c = (a + b) != 0
			d = container.GetSomeEnv() != ""

			return
		}

		container.AFunc = service
	}
	return container.AFunc
}
func (container *Container) GetClock() clockwork.Clock {
	if container.Clock == nil {
		service := clockwork.NewRealClock()
		container.Clock = service
	}
	return container.Clock
}
func (container *Container) GetCustomerWelcome() *CustomerWelcome {
	if container.CustomerWelcome == nil {
		service := NewCustomerWelcome(container.GetSendEmail())
		container.CustomerWelcome = service
	}
	return container.CustomerWelcome
}
func (container *Container) GetDependsOnTime() time.Time {
	return container.DependsOnTime(container.GetParsedTime("13 Jan 06 15:04 MST"))
}
func (container *Container) GetHTTPSignerClient() *HTTPSignerClient {
	if container.HTTPSignerClient == nil {
		service := &HTTPSignerClient{}
		service.CreateSigner = container.Signer
		container.HTTPSignerClient = service
	}
	return container.HTTPSignerClient
}
func (container *Container) GetNow() time.Time {
	return container.Now()
}
func (container *Container) GetOtherPkg() *go_sub_pkg.Person {
	if container.OtherPkg == nil {
		service := &go_sub_pkg.Person{}
		container.OtherPkg = service
	}
	return container.OtherPkg
}
func (container *Container) GetOtherPkg2() go_sub_pkg.Greeter {
	if container.OtherPkg2 == nil {
		service := go_sub_pkg.NewPerson()
		container.OtherPkg2 = service
	}
	return container.OtherPkg2
}
func (container *Container) GetOtherPkg3() go_sub_pkg.Person {
	if container.OtherPkg3 == nil {
		service := go_sub_pkg.Person{}
		container.OtherPkg3 = &service
	}
	return *container.OtherPkg3
}
func (container *Container) GetParsedTime(value string) time.Time {
	return container.ParsedTime(value)
}
func (container *Container) GetSendEmail() EmailSender {
	if container.SendEmail == nil {
		service := &SendEmail{}
		service.From = "hi@welcome.com"
		container.SendEmail = service
	}
	return container.SendEmail
}
func (container *Container) GetSendEmailError() *SendEmail {
	if container.SendEmailError == nil {
		service, err := NewSendEmail()
		if err != nil {
			panic(err)
		}
		container.SendEmailError = service
	}
	return container.SendEmailError
}
func (container *Container) GetSigner(req *http.Request) *Signer {
	return container.Signer(req)
}
func (container *Container) GetSomeEnv() string {
	if container.SomeEnv == nil {
		service := os.Getenv("ShouldBeSet")
		container.SomeEnv = &service
	}
	return *container.SomeEnv
}
func (container *Container) GetWhatsTheTime() *WhatsTheTime {
	if container.WhatsTheTime == nil {
		service := &WhatsTheTime{}
		service.clock = container.GetClock()
		container.WhatsTheTime = service
	}
	return container.WhatsTheTime
}
func (container *Container) GetWithEnv1() SendEmail {
	if container.WithEnv1 == nil {
		service := SendEmail{}
		service.From = os.Getenv("ShouldBeSet")
		container.WithEnv1 = &service
	}
	return *container.WithEnv1
}
func (container *Container) GetWithEnv2() *SendEmail {
	if container.WithEnv2 == nil {
		service := &SendEmail{}
		service.From = "foo-" + os.Getenv("ShouldBeSet") + "-bar"
		container.WithEnv2 = service
	}
	return container.WithEnv2
}
