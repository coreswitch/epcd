// Copyright 2019 epcd project.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/coreswitch/component"
	"github.com/coreswitch/epcd/pkg/pgw"
	"github.com/coreswitch/epcd/pkg/sgw"
	"github.com/coreswitch/log"
)

func main() {
	log.Info("epcd started")

	pgwComponent := &pgw.Component{}
	sgwComponent := &sgw.Component{}

	systemMap := component.SystemMap{
		"pwg": pgwComponent,
		"swg": sgwComponent,
	}
	systemMap.Start()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Ignore(syscall.SIGWINCH)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigs
		systemMap.Stop()
		done <- true
	}()

	<-done
	log.Info("epcd stopped")
}
