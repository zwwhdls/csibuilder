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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"csibuilder/pkg/model"
)

type templateBase struct {
	path           string
	ifExistsAction IfExistsAction
}

func (t templateBase) GetPath() string {
	return t.path
}

func (t templateBase) GetIfExistsAction() IfExistsAction {
	return t.ifExistsAction
}

type templateWithRepository struct {
	templateBase
	repository string
}

func (t *templateWithRepository) InjectRepository(repository string) {
	t.repository = repository
}

type templateWithProjectName struct {
	templateBase
	projectName string
}

func (t *templateWithProjectName) InjectProjectName(projectName string) {
	t.projectName = projectName
}

type templateWithBoilerplate struct {
	templateBase
	boilerplate string
}

func (t *templateWithBoilerplate) InjectBoilerplate(boilerplate string) {
	t.boilerplate = boilerplate
}

type templateWithResource struct {
	templateBase
	resource *model.Resource
}

func (t *templateWithResource) InjectResource(res *model.Resource) {
	t.resource = res
}

var _ = Describe("injector", func() {
	var tmp = templateBase{
		path:           "my/path/to/file",
		ifExistsAction: Error,
	}

	Context("injectInto", func() {
		Context("Config", func() {
			var c *model.Config

			BeforeEach(func() {
				c = &model.Config{}
			})

			Context("Repository", func() {
				var template *templateWithRepository

				BeforeEach(func() {
					template = &templateWithRepository{templateBase: tmp}
				})

				It("should not inject anything if the config is nil", func() {
					injector{}.injectInto(template)
					Expect(template.repository).To(Equal(""))
				})

				It("should not inject anything if the config doesn't have a repository set", func() {
					injector{config: c}.injectInto(template)
					Expect(template.repository).To(Equal(""))
				})

				It("should inject if the config has a repository set", func() {
					const repo = "test"
					Expect(c.SetRepository(repo)).To(Succeed())

					injector{config: c}.injectInto(template)
					Expect(template.repository).To(Equal(repo))
				})
			})
		})

		Context("Boilerplate", func() {
			var template *templateWithBoilerplate

			BeforeEach(func() {
				template = &templateWithBoilerplate{templateBase: tmp}
			})

			It("should not inject anything if no boilerplate was set", func() {
				injector{}.injectInto(template)
				Expect(template.boilerplate).To(Equal(""))
			})

			It("should inject if the a boilerplate was set", func() {
				const boilerplate = `Copyright "The Kubernetes Authors"`

				injector{boilerplate: boilerplate}.injectInto(template)
				Expect(template.boilerplate).To(Equal(boilerplate))
			})
		})

		Context("Resource", func() {
			var template *templateWithResource

			BeforeEach(func() {
				template = &templateWithResource{templateBase: tmp}
			})

			It("should not inject anything if the resource is nil", func() {
				injector{}.injectInto(template)
				Expect(template.resource).To(BeNil())
			})

			It("should inject if the config has a domain set", func() {
				var res = &model.Resource{
					Path:           "test-csi",
					CSIName:        "abc",
					AttachRequired: false,
				}

				injector{resource: res}.injectInto(template)
				Expect(template.resource).To(Equal(res))
			})

		})
	})
})
