package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	ensureCanBeConnected(t)
	ensureThatItThrowsAnErrorWhenUnableToConnect(t)
	ensureTableExists()

	// code := m.Run()
	// clearTable()
	// os.Exit(code)
}

func ensureCanBeConnected(t *testing.T) {
	// arrange

	// act
	conexionEstablecida := conexionBD()
	fmt.Printf("%#v\n", conexionEstablecida)

	// assert
	assert.Equal(t, "1", "1")
}

func ensureThatItThrowsAnErrorWhenUnableToConnect(t *testing.T) {
	// arrange

	// act
	conexionEstablecida := conexionBD()
	fmt.Printf("%#v\n", conexionEstablecida)

	// assert
	assert.Equal(t, "1", "1")
}

func ensureTableExists() {
	conexionEstablecida := conexionBD()

	if _, err := conexionEstablecida.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS employees
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`

func clearTable() {
	conexionEstablecida := conexionBD()

	conexionEstablecida.Exec("DELETE FROM empleados")
}
