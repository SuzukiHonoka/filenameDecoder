//go:build windows

package rename

var Filter = FileNameFilter{
	'<',
	'>',
	':',
	'"',
	'/',
	'\\',
	'|',
	'?',
	'*',
}
