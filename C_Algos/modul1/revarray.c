#include <stdio.h>

void revarray(void *base, size_t nel, size_t width) {
	for (int i = 0; i < nel / 2; i++) {
		for (int j = 0; j < width; j++) {
			char* a = (char*)base;
			char x = *(a + width * i + j);
			*(a + width * i + j) = *(a + (nel - 1 - i) * width + j);
			*(a + (nel - 1 - i) * width + j) = x;
		}
	}
}


int main()
{
	int a[5] = { 1, 2, 3, 4, 5 };
	char b[5] = { 1, 2, 3, 4, 5 };
	long long c[5] = { 1, 2, 3, 4, 5 };
	revarray(a, sizeof(a) / sizeof(a[0]), sizeof(a[0]));
	revarray(b, sizeof(b) / sizeof(b[0]), sizeof(b[0]));
	revarray(c, sizeof(c) / sizeof(c[0]), sizeof(c[0]));
	for (int i = 0; i < sizeof(a) / sizeof(a[0]); i++)
		printf("%d ", a[i]);
	for (int i = 0; i < sizeof(b) / sizeof(b[0]); i++)
		printf("%hhd ", b[i]);
	for (int i = 0; i < sizeof(c) / sizeof(c[0]); i++)
		printf("%lld ", c[i]);
}

