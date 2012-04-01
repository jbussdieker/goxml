#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <time.h>

#include "_cgo_export.h"

//#define TRACE_MEM
//#define CUSTOM_GC

long alloc_count = 0;

#pragma pack(push)
#pragma pack(1)
typedef struct go_xml_allocation {
	size_t size;
	struct timespec timestamp;
	void *p;
} go_xml_allocation;
#pragma pack(pop)

unsigned long goXmlAllocSize() {
	return alloc_count;
}

void goXmlFree(void *p) {
	alloc_count--;
#ifdef CUSTOM_GC
	go_xml_allocation *gxa = (go_xml_allocation *)(p - sizeof(go_xml_allocation));
	fprintf(stderr, "Freeing %lu bytes @ %p created at: %lu\n", gxa->size, gxa->p, gxa->timestamp.tv_nsec);
	return free(gxa);
#else
#ifdef TRACE_MEM
	fprintf(stderr, "%08lu Free %p\n", alloc_count, p);
#endif
	return free(p);
#endif
}

void *goXmlMalloc(int size) {
	alloc_count++;
#ifdef CUSTOM_GC
	go_xml_allocation *gxa = (go_xml_allocation *)malloc(size + sizeof(go_xml_allocation));
	gxa->p = (void *)gxa + sizeof(go_xml_allocation);
	gxa->size = size;
	clock_gettime(CLOCK_REALTIME, &(gxa->timestamp));
	fprintf(stderr, "Allocated %lu bytes @ %p timestamp: %lu\n", gxa->size, gxa->p, gxa->timestamp.tv_nsec);
	return gxa->p;
#else
#ifdef TRACE_MEM
	fprintf(stderr, "%08lu Malloc %d\n", alloc_count, size);
#endif
	return malloc(size);
#endif
}

void *goXmlRealloc(void *p, int size) {
#ifdef TRACE_MEM
	fprintf(stderr, "Realloc %p, %d\n", p, size);
#endif
	return realloc(p, size);
}

void *goXmlStrDup(void *p) {
	alloc_count++;
#ifdef TRACE_MEM
	fprintf(stderr, "%08lu StrDup %p\n", alloc_count, p);
#endif
	return strdup(p);
}
