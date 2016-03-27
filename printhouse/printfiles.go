// PrintFiles
//
// Contains communication logic to handle
// printfiles negotiations with the server.
package printhouse

// Lists all the print files created for
// the developer's account.
//
// Usage:
//	1: ph.GetPrintFiles()
func (ph *Printhouse) GetPrintFiles() []PrintFile {
	return nil
}

// Returns the information of a single print file.
//
// Usage:
//	1: ph.GetPrintFile("54808653b7017c7f795a6ac6")
func (ph *Printhouse) GetPrintFile(id string) *PrintFile {
	return nil
}

// Creates a print file
//
// If positive, the PrintFile will have a new Id.
// Usage:
//	1: ph.PostPrintFile(&PrintFile{Product: "0001", FileUrl: "http://dom.myFile.host"})
func (ph *Printhouse) PostPrintFile(printfile *PrintFile) *PrintFile {
	return nil
}


