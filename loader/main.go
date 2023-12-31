package main

import (
	"fmt"
	"io"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"github.com/varunnayal/migrations/src/entity"
	"github.com/varunnayal/migrations/src/enum"
)

func main() {
	stmts, err := New("postgres").Load(
		enum.Electricity,
		enum.UserStatusActive,
		&entity.Category{},
		&entity.User{},
		&entity.Friends{},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}

	if _, err := io.WriteString(os.Stdout, stmts); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write to stdout: %v\n", err)
		os.Exit(1)
	}
}
