#include <stdio.h>
#include <math.h>

#define scanf scanf_s

int main()
{
	long long a, x, q, w, e, z;
	int i = 1;
	scanf("%lld", &a);
	x = a;
	while (x > 0)
	{
		printf("1");
		w = 1;
		e = 2;
		while (e <= x)
		{
			q = e;
			e = e + w;
			w = q;
			i++;
		}
		q = w;
		w = e - w;
		e = q;
		x -= e;
		while (w > x)
		{
			if (x != 0)
				printf("0");
			q = w;
			w = e - w;
			e = q;
		}
	}

	return 0;
}