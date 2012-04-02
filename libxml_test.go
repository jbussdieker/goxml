package libxml

import "testing"
import "fmt"

func init() {
	fmt.Println("Testing libxml ", LIBXML_VERSION, "using parser version", LIBXML_PARSER_VERSION)
}

func blankXmlDoc(t *testing.T) Document {
	doc := XmlNewDoc("1.0")
	if doc == nil {
		t.Fatal("XmlNewDoc returned nil")
	}
	return doc
}

func blankHtmlDoc(t *testing.T) Document {
	doc := HtmlNewDoc("", "")
	if doc == nil {
		t.Fatal("HtmlNewDoc returned nil")
	}
	return doc
}

func parseHtmlDoc(t *testing.T, buf string) Document {
	doc := HtmlParseDoc(buf)
	if doc == nil {
		t.Fatal("HtmlParseDoc returned nil")
	}
	return doc
}

func parseXmlDoc(t *testing.T, buf string) Document {
	doc := XmlParseDoc(buf)
	if doc == nil {
		t.Fatal("HtmlParseDoc returned nil")
	}
	return doc
}

func checkMemory(t *testing.T) {
	if AllocSize() != 0 {
		t.Fatal(AllocSize(), "remaining unfreed objects")
	}
}

