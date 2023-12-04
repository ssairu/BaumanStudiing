#include <stdio.h>

int compare_int(void* x1, void* x2)
{
	return (*(int*)x1 - *(int*)x2);
}

int compare_ch(void* x1, void* x2)
{
	return (*(char*)x1 - *(char*)x2);
}


int maxarray(void *base, size_t nel, size_t width, int(*compare)(void *a, void*b)) {
	void *max = base;
	for (int i = 1; i < nel; i++) {
		int x = compare(max, (void *)((char *)base + i * width));
		if (x < 0)
			max = (void *)((char *)base + i * width);
	}
	return ((char *)max - (char *)base) / width;
}
