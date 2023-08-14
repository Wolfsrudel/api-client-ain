package parse

import "github.com/jonaslu/ain/internal/pkg/data"

func parseMethodSection(template []sourceMarker, parsedTemplate *data.ParsedTemplate) *fatalMarker {
	captureResult, captureFatal := captureSection("Method", template, true)
	if captureFatal != nil {
		return captureFatal
	}

	if captureResult.sectionHeaderLine == emptyLine {
		return nil
	}

	methodLines := captureResult.sectionLines

	if len(methodLines) > 1 {
		for _, hostLine := range methodLines {
			return newFatalMarker("Found several lines under [Method]", hostLine)
		}
	}

	parsedTemplate.Method = methodLines[0].lineContents

	return nil
}
