#include <stdio.h>
#include <libxml/parser.h>
#include <libxml/tree.h>
#include <string.h>

void *goXmlRead(char *input) {
	void *doc;
	doc = xmlReadMemory(input, strlen(input), "noname.xml", NULL, 0);
	if (doc == NULL) {
		// TODO: Handle error
	}
	return doc;
}

char *goXmlDumpDoc(void *doc) {
	int size = 0;
	xmlChar *mem;
	xmlDocDumpMemory(doc, &mem, &size);
	return mem;
}

