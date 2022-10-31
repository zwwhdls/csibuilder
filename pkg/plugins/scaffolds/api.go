/*
 Copyright 2022 CSIBuilder

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License
*/

package scaffolds

import (
	"csibuilder/pkg/machinery"
	"csibuilder/pkg/model"
	"csibuilder/pkg/plugins/scaffolds/internal"
	"csibuilder/pkg/plugins/scaffolds/internal/csi"
	"csibuilder/pkg/plugins/scaffolds/internal/deploy"
	"csibuilder/pkg/plugins/scaffolds/internal/hack"
	"fmt"
	"github.com/spf13/afero"
	"path/filepath"
)

// apiScaffolder contains configuration for generating scaffolding for Go type
// representing the API and controller that implements the behavior for the API.
type apiScaffolder struct {
	// fs is the filesystem that will be used by the scaffolder
	fs machinery.Filesystem

	// force indicates whether to scaffold controller files even if it exists or not
	force    bool
	resource model.Resource
	config   model.Config
}

// NewAPIScaffolder returns a new Scaffolder for API/controller creation operations
func NewAPIScaffolder(conf model.Config, res model.Resource, force bool) Scaffolder {
	return &apiScaffolder{
		force:    force,
		resource: res,
		config:   conf,
	}
}

func (s *apiScaffolder) InjectFS(filesystem machinery.Filesystem) {
	s.fs = filesystem
}

func (s *apiScaffolder) Scaffold() error {
	fmt.Println("Writing scaffold for you to edit...")

	bpFilePath := filepath.Join(s.config.Repo, hack.DefaultBoilerplatePath)
	boilerplate, err := afero.ReadFile(s.fs.FS, bpFilePath)
	if err != nil {
		return err
	}

	scaffold := machinery.NewScaffold(s.fs,
		machinery.WithResource(&s.resource),
		machinery.WithConfig(&s.config),
		machinery.WithBoilerplate(string(boilerplate)),
	)

	if err := scaffold.Execute(
		&csi.Controller{Force: s.force},
		&csi.Driver{Force: s.force},
		&csi.Identity{Force: s.force},
		&csi.Node{Force: s.force},
		&csi.Version{Force: s.force},
		&internal.Main{Force: s.force},
		&internal.Dockerfile{Force: s.force},

		&deploy.DaemonSetYaml{Force: s.force},
		&deploy.StatefulSetYaml{Force: s.force},
		&deploy.ClusterRoleYaml{Force: s.force},
		&deploy.ClusterRoleBindingYaml{Force: s.force},
		&deploy.ServiceAccountYaml{Force: s.force},
		&deploy.CSIDriverYaml{Force: s.force},
	); err != nil {
		return fmt.Errorf("error scaffolding APIs: %v", err)
	}

	return nil
}
