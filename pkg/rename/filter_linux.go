//go:build linux || unix

package rename

var Filter = FileNameFilter{
	'/',
}
