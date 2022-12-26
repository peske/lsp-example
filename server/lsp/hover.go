package lsp

import (
	"context"
	"github.com/peske/lsp-srv/lsp/protocol"
)

// hover https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_hover
func (s *Server) hover(_ context.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	//TODO: Figure out how Range field should be used.
	//TODO: Make Markdown example.
	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind:  protocol.PlainText,
			Value: "Lepo!",
		},
		Range: protocol.Range{
			Start: protocol.Position{
				Line:      0,
				Character: 5,
			},
			End: protocol.Position{
				Line:      0,
				Character: 10,
			},
		},
	}, nil
}

func hover_simple(content []byte) *protocol.Hover {
	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind:  protocol.PlainText,
			Value: "Lepo!",
		},
		Range: protocol.Range{
			Start: protocol.Position{
				Line:      0,
				Character: 5,
			},
			End: protocol.Position{
				Line:      0,
				Character: 10,
			},
		},
	}
}

func hover_markdown(content []byte) *protocol.Hover {
	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "# Nice!\n\nVery nice!",
		},
		Range: protocol.Range{
			Start: protocol.Position{
				Line:      0,
				Character: 5,
			},
			End: protocol.Position{
				Line:      0,
				Character: 10,
			},
		},
	}
}

func hover_range(content []byte) *protocol.Hover {
	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind:  protocol.PlainText,
			Value: "Lepo!",
		},
		Range: protocol.Range{
			Start: protocol.Position{
				Line:      0,
				Character: 5,
			},
			End: protocol.Position{
				Line:      0,
				Character: 10,
			},
		},
	}
}
