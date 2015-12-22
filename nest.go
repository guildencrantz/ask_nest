package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"code.google.com/p/go-uuid/uuid"
	"github.com/jsgoecke/nest"
	"gopkg.in/yaml.v2"
)

var thermostat *nest.Thermostat

type NestConf struct {
	Productid     string
	Productsecret string
	Authorization string
	Token         string
}

func init() {
	f, err := ioutil.ReadFile("./nest.yml")
	if err != nil {
		panic(err)
	}

	var c NestConf
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		panic(err)
	}

	client := nest.New(c.Productid, uuid.NewUUID().String(), c.Productsecret, c.Authorization)
	client.Token = c.Token

	devices, apierr := client.Devices()
	if apierr != nil {
		panic(apierr)
	}

	// FIXME: If there's more than one thermostat to work with this is going to be frustrating.
	for _, thermostat = range devices.Thermostats {
	}

	fmt.Fprintln(os.Stderr, thermostat)
}
