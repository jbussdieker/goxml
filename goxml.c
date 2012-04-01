#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <libxml/xmlmemory.h>
#include "mem.h"

void goXmlInit() {
	//fprintf(stderr, "Running xmlMemSetup()...\n");
	xmlMemSetup(
		(xmlFreeFunc)goXmlFree, 
		(xmlMallocFunc)goXmlMalloc, 
		(xmlReallocFunc)goXmlRealloc,
      	(xmlStrdupFunc)goXmlStrDup
	);

	//char *_LIBXML_VERSION = strdup(LIBXML_DOTTED_VERSION);
	//char *_LIBXML_PARSER_VERSION = strdup(xmlParserVersion);
	//fprintf(stderr, "LIBXML_VERSION: %s\n", _LIBXML_VERSION);
	//fprintf(stderr, "LIBXML_PARSER_VERSION: %s\n", _LIBXML_PARSER_VERSION);

#ifdef LIBXML_ICONV_ENABLED
	//fprintf(stderr, "LIBXML_ICONV_ENABLED: %s\n", "true");
#else
	//fprintf(stderr, "LIBXML_ICONV_ENABLED: %s\n", "false");
#endif

	//xmlInitParser();
}

