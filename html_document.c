#include <stdio.h>
#include <libxml/HTMLparser.h>
#include <libxml/tree.h>
#include <string.h>

htmlDocPtr goHtmlRead(xmlChar *input) {
	htmlDocPtr doc;
	xmlChar *url = "noname.html";
	doc = htmlReadDoc(input, url, "UTF-8", 0);
	return doc;
}

xmlChar *goHtmlDumpDoc(htmlDocPtr doc) {
	int size = 0;
	xmlChar *mem;
	xmlDocDumpMemory(doc, &mem, &size);
	return mem;
}

