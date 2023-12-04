#include <stdio.h>
#include <math.h>

#define scanf scanf_s


void out(unsigned int x, unsigned int y)
{
	for (int i = 0; i < 32; i++)
	{
		if ((x % 2) == (y % 2) && (y % 2) != 0)
			printf("%d ", i);
		x /= 2;
		y /= 2;
	}
}

int main()
{
	int na, nb, l;
	unsigned int x = 0, y = 0;
	scanf("%d", &na);
	for (int i = 0; i < na; i++)
	{
		scanf("%d", &l);
		unsigned int k = 1;
		for (int j = 0; j < l; j++)
		{
			k *= 2;
		}
		x += k;
	}

	scanf("%d", &nb);
	for (int i = 0; i < nb; i++)
	{
		scanf("%d", &l);
		unsigned int k = 1;
		for (int j = 0; j < l; j++)
		{
			k *= 2;
		}
		y += k;
	}


	out(x, y);

	return 0;
}