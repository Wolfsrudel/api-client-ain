package parse

import (
	"github.com/jonaslu/ain/internal/pkg/data"
)

func parseQuerySection(template []sourceMarker, parsedTemplate *data.ParsedTemplate) *fatalMarker {
	captureResult, captureFatal := captureSection("Query", template, true)
	if captureFatal != nil {
		return captureFatal
	}

	if captureResult.sectionHeaderLine == emptyLine {
		return nil
	}

	queryLines := captureResult.sectionLines

	query := []string{}
	for _, queryLine := range queryLines {
		query = append(query, queryLine.lineContents)
	}

	parsedTemplate.Query = query

	return nil
}
