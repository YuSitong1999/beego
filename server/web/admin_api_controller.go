// Copyright 2020
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"text/template"
)

type adminAPIController struct {
	Controller
	servers []*HttpServer
}

// QpsIndex write qps statistics map result info in json.
// it's registered with url pattern "/qps" in adminAPI module.
func (a *adminAPIController) QpsIndex() {
	a.Data["json"] = StatisticsMap.GetMap()
	a.ServeJSON()
}

// Conf write system conf info in json.
// it's registered with url pattern "/conf" in adminAPI module.
func (a *adminAPIController) Conf() {
	// Config without func
	type Config struct {
		AppName             string
		RunMode             string
		RouterCaseSensitive bool
		ServerName          string
		RecoverPanic        bool
		CopyRequestBody     bool
		EnableGzip          bool
		MaxMemory           int64
		MaxUploadSize       int64
		EnableErrorsShow    bool
		EnableErrorsRender  bool
		Listen              Listen
		WebConfig           WebConfig
		Log                 LogConfig
	}
	m := make(M)
	var config Config
	config.AppName = BConfig.AppName
	config.RunMode = BConfig.RunMode
	config.RouterCaseSensitive = BConfig.RouterCaseSensitive
	config.ServerName = BConfig.ServerName
	config.RecoverPanic = BConfig.RecoverPanic
	config.CopyRequestBody = BConfig.CopyRequestBody
	config.EnableGzip = BConfig.EnableGzip
	config.MaxMemory = BConfig.MaxMemory
	config.MaxUploadSize = BConfig.MaxUploadSize
	config.EnableErrorsShow = BConfig.EnableErrorsShow
	config.EnableErrorsRender = BConfig.EnableErrorsRender
	config.Listen = BConfig.Listen
	config.WebConfig = BConfig.WebConfig
	config.Log = BConfig.Log
	list("BConfig", config, m)
	m["appConfigPath"] = template.HTMLEscapeString(appConfigPath)
	m["appConfigProvider"] = template.HTMLEscapeString(appConfigProvider)
	a.Data["json"] = m
	a.ServeJSON()
}

// Router write routers info in main app in json.
// it's registered with url pattern "/router" in adminAPI module.
func (a *adminAPIController) Router() {
	a.Data["json"] = BeeApp.PrintTree()
	a.ServeJSON()
}
