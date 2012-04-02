package libxml

import "testing"

func BenchmarkXmlParseDoc(b *testing.B) {
	for i := 0; i < 10000; i++ {
		doc := NewXmlFromString("<html><head><title>Test Doc</title></head><body class=\"styled\"></body></html>")
		doc.Free()
	}
}

