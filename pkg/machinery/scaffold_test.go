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
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
	"os"
)

var _ = Describe("Scaffold", func() {
	Describe("NewScaffold", func() {
		It("should succeed for no option", func() {
			s := NewScaffold(Filesystem{FS: afero.NewMemMapFs()})
			Expect(s.fs).NotTo(BeNil())
			Expect(s.dirPerm).To(Equal(defaultDirectoryPermission))
			Expect(s.filePerm).To(Equal(defaultFilePermission))
			Expect(s.injector.config).To(BeNil())
			Expect(s.injector.boilerplate).To(Equal(""))
			Expect(s.injector.resource).To(BeNil())
		})

		It("should succeed with directory permissions option", func() {
			const dirPermissions os.FileMode = 0755

			s := NewScaffold(Filesystem{FS: afero.NewMemMapFs()}, WithDirectoryPermissions(dirPermissions))
			Expect(s.fs).NotTo(BeNil())
			Expect(s.dirPerm).To(Equal(dirPermissions))
			Expect(s.filePerm).To(Equal(defaultFilePermission))
			Expect(s.injector.config).To(BeNil())
			Expect(s.injector.boilerplate).To(Equal(""))
			Expect(s.injector.resource).To(BeNil())
		})

		It("should succeed with file permissions option", func() {
			const filePermissions os.FileMode = 0755

			s := NewScaffold(Filesystem{FS: afero.NewMemMapFs()}, WithFilePermissions(filePermissions))
			Expect(s.fs).NotTo(BeNil())
			Expect(s.dirPerm).To(Equal(defaultDirectoryPermission))
			Expect(s.filePerm).To(Equal(filePermissions))
			Expect(s.injector.config).To(BeNil())
			Expect(s.injector.boilerplate).To(Equal(""))
			Expect(s.injector.resource).To(BeNil())
		})

		It("should succeed with config option", func() {
			cfg := &model.Config{}

			s := NewScaffold(Filesystem{FS: afero.NewMemMapFs()}, WithConfig(cfg))
			Expect(s.fs).NotTo(BeNil())
			Expect(s.dirPerm).To(Equal(defaultDirectoryPermission))
			Expect(s.filePerm).To(Equal(defaultFilePermission))
			Expect(s.injector.config).NotTo(BeNil())
			Expect(s.injector.boilerplate).To(Equal(""))
			Expect(s.injector.resource).To(BeNil())
		})

		It("should succeed with boilerplate option", func() {
			const boilerplate = "Copyright"

			s := NewScaffold(Filesystem{FS: afero.NewMemMapFs()}, WithBoilerplate(boilerplate))
			Expect(s.fs).NotTo(BeNil())
			Expect(s.dirPerm).To(Equal(defaultDirectoryPermission))
			Expect(s.filePerm).To(Equal(defaultFilePermission))
			Expect(s.injector.config).To(BeNil())
			Expect(s.injector.boilerplate).To(Equal(boilerplate))
			Expect(s.injector.resource).To(BeNil())
		})

		It("should succeed with resource option", func() {
			var res = &model.Resource{
				CSIName: "abc",
			}

			s := NewScaffold(Filesystem{FS: afero.NewMemMapFs()}, WithResource(res))
			Expect(s.fs).NotTo(BeNil())
			Expect(s.dirPerm).To(Equal(defaultDirectoryPermission))
			Expect(s.filePerm).To(Equal(defaultFilePermission))
			Expect(s.injector.config).To(BeNil())
			Expect(s.injector.boilerplate).To(Equal(""))
			Expect(s.injector.resource).NotTo(BeNil())
		})
	})

	Describe("Scaffold.Execute", func() {
		const (
			path     = "filename"
			pathGo   = path + ".go"
			pathYaml = path + ".yaml"
			content  = "Hello world!"
		)

		var (
			testErr = errors.New("error text")

			s *Scaffold
		)

		BeforeEach(func() {
			s = &Scaffold{fs: afero.NewMemMapFs()}
		})

		DescribeTable("successes",
			func(path, expected string, files ...Builder) {
				Expect(s.Execute(files...)).To(Succeed())

				b, err := afero.ReadFile(s.fs, path)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(b)).To(Equal(expected))
			},
			Entry("should write the file",
				path, content,
				fakeTemplate{fakeBuilder: fakeBuilder{path: path}, body: content},
			),
			Entry("should skip optional models if already have one",
				path, content,
				fakeTemplate{fakeBuilder: fakeBuilder{path: path}, body: content},
				fakeTemplate{fakeBuilder: fakeBuilder{path: path}},
			),
			Entry("should overwrite required models if already have one",
				path, content,
				fakeTemplate{fakeBuilder: fakeBuilder{path: path}},
				fakeTemplate{fakeBuilder: fakeBuilder{path: path, ifExistsAction: OverwriteFile}, body: content},
			),
			Entry("should format a go file",
				pathGo, "package file\n",
				fakeTemplate{fakeBuilder: fakeBuilder{path: pathGo}, body: "package    file"},
			),
		)

		DescribeTable("file builders related errors",
			func(errType interface{}, files ...Builder) {
				err := s.Execute(files...)
				Expect(err).To(HaveOccurred())
				Expect(errors.As(err, errType)).To(BeTrue())
			},
			Entry("should fail if unable to set default values for a template",
				&SetTemplateDefaultsError{},
				fakeTemplate{err: testErr},
			),
			Entry("should fail if an unexpected previous model is found",
				&ModelAlreadyExistsError{},
				fakeTemplate{fakeBuilder: fakeBuilder{path: path}},
				fakeTemplate{fakeBuilder: fakeBuilder{path: path, ifExistsAction: Error}},
			),
			Entry("should fail if behavior if-exists-action is not defined",
				&UnknownIfExistsActionError{},
				fakeTemplate{fakeBuilder: fakeBuilder{path: path}},
				fakeTemplate{fakeBuilder: fakeBuilder{path: path, ifExistsAction: -1}},
			),
		)

		// Following errors are unwrapped, so we need to check for substrings
		DescribeTable("template related errors",
			func(errMsg string, files ...Builder) {
				err := s.Execute(files...)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(errMsg))
			},
			Entry("should fail if a template is broken",
				"template: ",
				fakeTemplate{body: "{{ .Field }"},
			),
			Entry("should fail if a template params aren't provided",
				"template: ",
				fakeTemplate{body: "{{ .Field }}"},
			),
			Entry("should fail if unable to format a go file",
				"expected 'package', found ",
				fakeTemplate{fakeBuilder: fakeBuilder{path: pathGo}, body: content},
			),
		)

		DescribeTable("insert strings related errors",
			func(errType interface{}, files ...Builder) {
				Expect(afero.WriteFile(s.fs, path, []byte{}, 0666)).To(Succeed())

				err := s.Execute(files...)
				Expect(err).To(HaveOccurred())
				Expect(errors.As(err, errType)).To(BeTrue())
			},
			Entry("should fail if inserting into a model that fails when a file exists and it does exist",
				&FileAlreadyExistsError{},
				fakeTemplate{fakeBuilder: fakeBuilder{path: "filename", ifExistsAction: Error}},
				fakeInserter{fakeBuilder: fakeBuilder{path: "filename"}},
			),
			Entry("should fail if inserting into a model with unknown behavior if the file exists and it does exist",
				&UnknownIfExistsActionError{},
				fakeTemplate{fakeBuilder: fakeBuilder{path: "filename", ifExistsAction: -1}},
				fakeInserter{fakeBuilder: fakeBuilder{path: "filename"}},
			),
		)

		Context("write when the file already exists", func() {
			BeforeEach(func() {
				_ = afero.WriteFile(s.fs, path, []byte{}, 0666)
			})

			It("should skip the file by default", func() {
				Expect(s.Execute(fakeTemplate{
					fakeBuilder: fakeBuilder{path: path},
					body:        content,
				})).To(Succeed())

				b, err := afero.ReadFile(s.fs, path)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(b)).To(BeEmpty())
			})

			It("should write the file if configured to do so", func() {
				Expect(s.Execute(fakeTemplate{
					fakeBuilder: fakeBuilder{path: path, ifExistsAction: OverwriteFile},
					body:        content,
				})).To(Succeed())

				b, err := afero.ReadFile(s.fs, path)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(b)).To(Equal(content))
			})

			It("should error if configured to do so", func() {
				err := s.Execute(fakeTemplate{
					fakeBuilder: fakeBuilder{path: path, ifExistsAction: Error},
					body:        content,
				})
				Expect(err).To(HaveOccurred())
				Expect(errors.As(err, &FileAlreadyExistsError{})).To(BeTrue())
			})
		})
	})
})

var _ Builder = fakeBuilder{}

// fakeBuilder is used to mock a Builder
type fakeBuilder struct {
	path           string
	ifExistsAction IfExistsAction
}

// GetPath implements Builder
func (f fakeBuilder) GetPath() string {
	return f.path
}

// GetIfExistsAction implements Builder
func (f fakeBuilder) GetIfExistsAction() IfExistsAction {
	return f.ifExistsAction
}

var _ Template = fakeTemplate{}

// fakeTemplate is used to mock a File in order to test Scaffold
type fakeTemplate struct {
	fakeBuilder

	body string
	err  error
}

func (f fakeTemplate) InjectTemplatePath(templatePath string) {
	f.path = templatePath
}

// GetBody implements Template
func (f fakeTemplate) GetBody() string {
	return f.body
}

// SetTemplateDefaults implements Template
func (f fakeTemplate) SetTemplateDefaults() error {
	if f.err != nil {
		return f.err
	}

	return nil
}

type fakeInserter struct {
	fakeBuilder
}
