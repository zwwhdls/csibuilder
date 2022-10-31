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

package machinery

import (
	"csibuilder/pkg/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type mockTemplate struct {
	TemplateMixin
	RepositoryMixin
	ProjectNameMixin
	BoilerplateMixin
	ResourceMixin
}

type mockInserter struct {
	// InserterMixin requires a different type because it collides with TemplateMixin
	InserterMixin
}

var _ = Describe("TemplateMixin", func() {
	const (
		path           = "path/to/file.go"
		ifExistsAction = SkipFile
		body           = "content"
	)

	var tmp = mockTemplate{
		TemplateMixin: TemplateMixin{
			PathMixin:           PathMixin{path},
			IfExistsActionMixin: IfExistsActionMixin{ifExistsAction},
			TemplateBody:        body,
		},
	}

	Context("GetPath", func() {
		It("should return the path", func() {
			Expect(tmp.GetPath()).To(Equal(path))
		})
	})

	Context("GetIfExistsAction", func() {
		It("should return the if-exists action", func() {
			Expect(tmp.GetIfExistsAction()).To(Equal(ifExistsAction))
		})
	})

	Context("GetBody", func() {
		It("should return the body", func() {
			Expect(tmp.GetBody()).To(Equal(body))
		})
	})
})

var _ = Describe("InserterMixin", func() {
	const path = "path/to/file.go"

	var tmp = mockInserter{
		InserterMixin: InserterMixin{
			PathMixin: PathMixin{path},
		},
	}

	Context("GetPath", func() {
		It("should return the path", func() {
			Expect(tmp.GetPath()).To(Equal(path))
		})
	})

	Context("GetIfExistsAction", func() {
		It("should return overwrite file always", func() {
			Expect(tmp.GetIfExistsAction()).To(Equal(OverwriteFile))
		})
	})
})

var _ = Describe("RepositoryMixin", func() {
	const repo = "test"

	var tmp = mockTemplate{}

	Context("InjectRepository", func() {
		It("should inject the provided repository", func() {
			tmp.InjectRepository(repo)
			Expect(tmp.Repo).To(Equal(repo))
		})
	})
})

var _ = Describe("ProjectNameMixin", func() {
	const name = "my project"

	var tmp = mockTemplate{}

	Context("InjectProjectName", func() {
		It("should inject the provided project name", func() {
			tmp.InjectProjectName(name)
			Expect(tmp.ProjectName).To(Equal(name))
		})
	})
})

var _ = Describe("BoilerplateMixin", func() {
	const boilerplate = "Copyright"

	var tmp = mockTemplate{}

	Context("InjectBoilerplate", func() {
		It("should inject the provided boilerplate", func() {
			tmp.InjectBoilerplate(boilerplate)
			Expect(tmp.Boilerplate).To(Equal(boilerplate))
		})
	})
})

var _ = Describe("ResourceMixin", func() {
	var res = &model.Resource{
		Path:           "test-csi",
		CSIName:        "abc",
		AttachRequired: false,
	}

	var tmp = mockTemplate{}

	Context("InjectResource", func() {
		It("should inject the provided resource", func() {
			tmp.InjectResource(res)
			Expect(tmp.Resource.Path == res.Path).To(BeTrue())
			Expect(tmp.Resource.CSIName == res.CSIName).To(BeTrue())
			Expect(tmp.Resource.AttachRequired == res.AttachRequired).To(BeTrue())
		})
	})
})
