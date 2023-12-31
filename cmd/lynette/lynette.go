/*
Copyright © 2023 Józef Daniecki jozek.daniecki@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"os"

	"log/slog"

	"github.com/jdaniecki/lynette/internal/cmd"
)

var programLevel = new(slog.LevelVar)

func init() {
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	programLevel.Set(slog.LevelDebug)

}

func main() {
	if err := cmd.Execute(); err != nil {
		slog.Error("Root command failed", "error", err)
		os.Exit(1)
	}
}
