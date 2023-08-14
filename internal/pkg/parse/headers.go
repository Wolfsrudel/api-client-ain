package parse

import "github.com/jonaslu/ain/internal/pkg/data"

func parseHeadersSection(template []sourceMarker, parsedTemplate *data.ParsedTemplate) *fatalMarker {
	captureResult, captureFatal := captureSection("Headers", template, true)
	if captureFatal != nil {
		return captureFatal
	}

	if captureResult.sectionHeaderLine == emptyLine {
		return nil
	}

	headerLines := captureResult.sectionLines

	headers := []string{}
	for _, headerLine := range headerLines {
		headers = append(headers, headerLine.lineContents)
	}

	parsedTemplate.Headers = headers

	return nil
}
