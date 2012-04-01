#include <libxml/tree.h>
#include <libxml/xpath.h>

xmlNodePtr goXmlChild(xmlNodePtr node) {
	xmlNodePtr child = node->children;
	return child;
}

xmlNodeSetPtr goXmlChildren(xmlNodePtr node) {
	xmlNodePtr child = node->children;
	xmlNodeSetPtr set = xmlXPathNodeSetCreate(child);

	if(!child) return set;

	child = child->next;
	while(NULL != child) {
		xmlXPathNodeSetAdd(set, child);
		child = child->next;
	}

	return set;
}
