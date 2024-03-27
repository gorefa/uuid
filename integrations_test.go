package uuid_test

import (
	. "github.com/gorefa/uuid"
	"gopkg.in/stretchr/testify.v1/assert"
	"testing"
	"log"
	"io/ioutil"
)

func TestInit(t *testing.T) {
	assert.Panics(t, didRegisterPanic, "Should panic")
}

func didRegisterPanic() {
	config := &GeneratorConfig{
		Logger: log.New(ioutil.Discard, "", 0),
	}
	RegisterGenerator(config)
	RegisterGenerator(config)
}
