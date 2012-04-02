package libxml

import "testing"

func BenchmarkHtmlParseDoc(b *testing.B) {
	for i := 0; i < 100; i++ {
		doc := HtmlParseDoc("<html><head><title>Test Doc</title></head><body class=\"styled\"></body></html>")
		doc.Free()
	}
}

