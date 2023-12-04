#include <stdio.h>

unsigned long binsearch(unsigned long nel, int(*compare)(unsigned long i)) {
	unsigned long r = nel - 1;
	unsigned long l = 0;
	while (l != r) {
		unsigned long x = (r + l) / 2;
		if (compare(x) == 0)
			return x;
		if (compare(x) > 0)
			r = x + 1;
		if (compare(x) < 0)
			l = x;
	}
	if (compare(l) == 0)
		return l;
	else
		return nel;
}


int main()
{
	int a[5] = { 1, 2, 3, 4, 5 };
	char b[5] = { 12, 2, 3, 4, 6 };
	long long c[5] = { 1, 2, 3, 4 };

}