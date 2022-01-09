package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

// default
var C = Configuration{
	CLIENT_HEALTH_TIME_OUT: 1 * time.Second,
}

type Configuration struct {
	CLIENT_ID     string `mapstructure:"CLIENT_ID"`
	CLIENT_SECRET string `mapstructure:"CLIENT_SECRET"`

	CLIENT_HEALTH_TIME_OUT time.Duration `mapstructure:"CLIENT_HEALTH_TIME_OUT"`
}

func LoadConfiguration(filePath string) {
	if _, err := os.Stat(filePath); !errors.Is(err, os.ErrNotExist) {
		viper.SetConfigFile(filePath)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	bindEnvs(C)

	err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to decode into struct, %v", err))
		panic("[CONFIG] unable to decode into struct")
	}

	e := validate(C)
	if e != nil {
		for _, s := range e {
			log.Fatal(fmt.Sprintf("[CONFIG] %s: %s\n", s.FailedField, s.Tag))
		}
		panic("[CONFIG] invalid configuration")
	}
}

func validate(c Configuration) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func bindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, tv)...)
		default:
			err := viper.BindEnv(strings.Join(append(parts, tv), "."))
			if err != nil {
				fmt.Printf("can't bind config from ENV, %v\n", err)
			}
		}
	}
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}
