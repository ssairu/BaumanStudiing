#include <stdio.h>
#include <math.h>

#define scanf scanf_s


int main()
{
	int a, n = 0;
	scanf("%d", &a);
	a > 0 ? ; a = -a;
	char del[a - 1];
	for (int i = 0; i < a - 1; i++)
		del[i] = 1;

	for (int i = 2; i * i < a + 1; i++)
	{
		if (del[i - 2] == 1)
			for (int j = i * i; j <= a; j += i)
				del[j - 2] = 0;
	}

	for (int i = 0; i < a - 1; i++)
		if (del[i] == 1)
			n++;

	int pr[n], x = 0;

	for (int i = 2; i < a + 1; i++)
		if (del[i - 2] == 1)
		{
			pr[x] = i;
			x++;
		}

	while (a % pr[x] != 0)
		x--;

	printf("%d", pr[x]);

	return 0;
}