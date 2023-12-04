#include <stdio.h>
#include <math.h>

#define scanf scanf_s

int main()
{
	long long x0, n, p, a, k;
	scanf("%lld%lld\n%lld", &n, &x0, &k);
	a = k;
	p = k * n;

	for (int i = 0; i < n; i++)
	{
		scanf("%lld", &k);
		a = a * x0 + k;
		if (i < n - 1)
			p = p * x0 + k * (n - i - 1);
	}

	printf("%lld\n%lld", a, p);

	return 0;
}