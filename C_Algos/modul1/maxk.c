#include <stdio.h>
#include <math.h>

#define scanf scanf_s


int main()
{
	int n, k, s = 0, l;
	scanf("%d", &n);
	int a[1000000];
	for (int i = 0; i < n; i++)
		scanf("%d", a + i);
	scanf("%d", &k);

	for (int i = 0; i < k; i++)
		s += a[i];
	l = s;

	for (int i = k; i < n; i++)
	{
		l = l - a[i - k] + a[i];
		if (l > s)
			s = l;
	}

	printf("%d", s);

	return 0;
}