Mockable is a quick and dirty package that lets you know if your user wants
you to use mocked interfaces and such.

## WHY?!?!?!?! ##

Well, because I write a lot of API clients and such, and then I end up using
them in my apps. When I run tests against those apps, I find it easiest if
the library itself provides a transparent mocked interface.

I prefer to decide if mocking is desired by checking an environment variable,
because I write a lot of my app-level tests in languages other than Go. As
I do this sort of thing quite a lot, occasionallly needing to enable mocking
for several clients at once, it finally made sense to just go ahead and
formalize the process and use a uniform environment variable, MOCKABLE,
for everything.

## Installation ##

I suggest the use of [glide](https://glide.sh/) for managing your Go deps, but
you should be able to install it directly without much issue:

```
go get github.com/ess/mockable
```

As mentioned, though, a better idea is to use glide (or another package manager
that supports SemVer).

## Usage ##

Here's a quasi-real-world example. Let's say your library generates widgets.
Now, in production, you really want to make good, real, quality widgets. Under
test, though, the app that uses your widget generator may just need to ensure
that a widget gets generated. Here's how I like to handle that:

```go
package widget

import "github.com/ess/mockable"

type Widget struct {
  Name string
}

type Generator interface {
  Generate(name string) *Widget
}

type realGenerator struct {}

func (generator *realGenerator) Generate(name string) *Widget {
  // do a bunch of stuff that hits the real widget service out there in space
  // ...
  return generatedWidget
}

type fakeGenerator struct {}

func (generator *fakeGenerator) Generate(name string) *Widget {
  return &Widget{Name: name}
}

func NewGenerator() Generator {
  if mockable.Mocked() {
    return &fakeGenerator{}
  }

  return &realGenerator{}
}
```

Now, when the above is used and `widget.NewGenerator()` is called, a fake
generator is returned if mocking is enabled (by setting the `MOCKABLE`
environment variable to pretty much anything), but a real generator otherwise.

## Test Helpers ##

In addition to the main functionality, there are a few helpers available for
writing your unit tests.

In the main `mockable` package, there is `Enable()` and `Disable()`, which are
used to explicitly set the mocking state.

Additionally, in `mockable/mocking`,we have `EnabledDo` and `DisabledDo`. Both
of these functions take a `*testing.T` and a `func(*testing.T)` and are used 
like so:

```go
package main

import (
  "testing"

  "github.com/ess/mockable/mocking"
)

func TestExpecationsMet(t *testing.T) {
  mocking.EnabledDo(t, func(t *testing.T) {
    if !ExpectationsMet() {
      t.Error("Expectations not met when mocking is enabled!!!")
    }
  })
}
```
## History ##

* v0.2.0 - Now with test helpers
* v0.1.3 - I like y'all, so I've removed the vendor from this repo
* v0.1.2 - Added Enable()/Disable() for use in consumer unit tests
* v0.1.1 - Documentation added
* v0.1.0 - Initial release
