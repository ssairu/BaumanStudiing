#include <stdio.h>
#include <math.h>

#include <stdlib.h>


#define scanf scanf_s


int pow2(char a) {
	int x = 1;
	for (int i = 0; i < a; i++)
		x *= 2;
	return x;
}

int main()
{
	int a, n = 0;
	scanf("%d", &a);
	if (a < 0)
		a = -a;
	unsigned char * del = (unsigned char *)malloc((a - 1) / 8 + 1);
	for (int i = 0; i < (a - 1) / 8; i++) {
		del[i] = 1;
		for (int j = 0; j < 8; j++)
			del[i] = del[i] * 2 + 1;
	}
	del[(a - 1) / 8] = 0;
	for (int i = 0; i < (a - 1) % 8; i++) {
		del[(a - 1) / 8] = del[(a - 1) / 8] + pow2(7 - i);
	}

	for (int i = 0; i < a / 8 + 1; i++)
		printf("%hhu  ", del[i]);


	for (int i = 0; (i * 8 + 2) * (i * 8 + 2) < a + 1; i++)
	{
		for (int k = 7; k > -1; k--) {
			if ((del[i] / pow2(k)) % 2 == 1)
				for (int j = ((i + 1) * 8 - k + 1) * ((i + 1) * 8 - k + 1); j <= a; j += ((i + 1) * 8 - k + 1))
					del[(j - 2) / 8] -= pow2(7 - (j - 2) % 8);
		}
	}

	for (int i = 0; i < a / 8 + 1; i++)
		printf("%hhu  ", del[i]);

	/*for (int i = 0; i < a - 1; i++)
		if (del[i] == 1)
			n++;

	int *pr = (int *)malloc(n * sizeof(pr[0])), x = 0;

	for (int i = 2; i < a + 1; i++)
		if (del[i - 2] == 1)
		{
			pr[x] = i;
			x++;
		}

	while (a % pr[x] != 0)
		x--;

	printf("%d", pr[x]);

	free(pr);
	free(del);

	/*char a = 255;
	int b = 256;
	printf("%hhd", a / b);*/

	return 0;
}