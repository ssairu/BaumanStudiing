#include <stdio.h>
#include <math.h>

#define scanf scanf_s

int main()
{
	long long p[64], a, b, r, m;
	scanf("%lld%lld%lld", &a, &b, &m);

	for (int i = 0; i < 64; i++)
	{
		p[i] = b % 2;
		b /= 2;
	}

	r = a * p[63];

	for (int i = 62; i >= 0; i--)
	{
		r = (r * (2 % m)) % m;
		r = (r + ((a * p[i]) % m)) % m;
	}

	printf("%lld\n", r);

	return 0;
}