package machinery

// IfExistsAction determines what to do if the scaffold file already exists
type IfExistsAction int

const (
	// SkipFile skips the file and moves to the next one
	SkipFile IfExistsAction = iota

	// Error returns an error and stops processing
	Error

	// OverwriteFile truncates and overwrites the existing file
	OverwriteFile
)

// File describes a file that will be written
type File struct {
	// Path is the file to write
	Path string

	// Contents is the generated output
	Contents string

	// IfExistsAction determines what to do if the file exists
	IfExistsAction IfExistsAction
}
