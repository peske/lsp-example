package lsp

import (
	"fmt"
	"github.com/peske/lsp-srv/span"
	"regexp"
	"strings"

	"github.com/peske/lsp-srv/lsp/protocol"
	"go.uber.org/zap"
)

func (s *Server) checkDiagnosticsThread(uri protocol.DocumentURI) {
	if u := uri.SpanURI(); u.IsFile() {
		go s.checkDiagnostics(u)
	}
}

func (s *Server) checkDiagnostics(uri span.URI) {
	f := s.helper.Cache.GetFile(uri)
	if f == nil {
		s.logger.Error(fmt.Sprintf("checkDiagnostics; '%s' not found.", string(uri)))
	}
	mp := protocol.NewColumnMapper(f.URI(), f.Content)
	prm := &protocol.PublishDiagnosticsParams{
		URI:         protocol.URIFromSpanURI(f.URI()),
		Version:     f.Version,
		Diagnostics: s.getDiagnosticsForWord(mp, "football", protocol.SeverityInformation, "Football is a very nice sport."),
	}
	prm.Diagnostics = append(prm.Diagnostics, s.getDiagnosticsForWord(mp, "basketball", protocol.SeverityWarning, "Basketball is nice, but football is better.")...)
	prm.Diagnostics = append(prm.Diagnostics, s.getDiagnosticsForWord(mp, "soccer", protocol.SeverityError, "It is not \"soccer\" - it is football!")...)
	prm.Diagnostics = append(prm.Diagnostics, s.getDiagnosticsForWord(mp, "beer", protocol.SeverityHint, "Consider calling me for a beer! :)")...)
	if err := s.client.PublishDiagnostics(s.ctx, prm); err != nil {
		s.logger.Error("checkDiagnostics; s.client.PublishDiagnostics", zap.Error(err))
	}
}

func (s *Server) getDiagnosticsForWord(mapper *protocol.ColumnMapper, word string, severity protocol.DiagnosticSeverity,
	message string) []protocol.Diagnostic {
	ds := make([]protocol.Diagnostic, 0)
	rx := regexp.MustCompile(fmt.Sprintf(`(?i)(^|\s)%s($|\s)`, word))
	ms := rx.FindAllIndex(mapper.Content, -1)
	for _, idx := range ms {
		from := idx[0]
		to := idx[1]
		c := mapper.Content[from:to]
		rx = regexp.MustCompile(fmt.Sprintf(`(?i)%s`, word))
		m := rx.FindAllIndex(c, -1)
		from += m[0][0]
		to = from + m[0][1] - m[0][0]
		if strings.ToLower(string(mapper.Content[from:to])) != strings.ToLower(word) {
			s.logger.Error(fmt.Sprintf("getDiagnosticsForWord; expected: '%s'; got: '%s'", word, string(mapper.Content[from:to])))
		}
		p, err := mapper.OffsetPosition(from)
		if err != nil {
			s.logger.Error("getDiagnosticsForWord; mapper.OffsetPosition", zap.Error(err))
			continue
		}
		ds = append(ds, protocol.Diagnostic{
			Range: protocol.Range{
				Start: p,
				End: protocol.Position{
					Line:      p.Line,
					Character: p.Character + uint32(to-from),
				},
			},
			Severity: severity,
			Source:   "Fat Dragon",
			Message:  message,
		})
	}
	return ds
}
