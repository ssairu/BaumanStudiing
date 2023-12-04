#include <stdio.h>
#include <math.h>

#define scanf scanf_s


int main()
{
	int a[8], b[8], s = 1;

	for (int i = 0; i < 8; i++)
	{
		scanf("%d", a + i);
	}

	for (int i = 0; i < 8; i++)
	{
		scanf("%d", b + i);
	}

	for (int i = 0; i < 8; i++)
	{
		for (int j = i; j < 8; j++)
		{
			if (a[i] == b[j])
			{
				int s = b[i];
				b[i] = b[j];
				b[j] = s;
			}
		}
	}

	for (int i = 0; i < 8; i++)
	{
		if (a[i] != b[i])
			s = 0;
	}

	if (s == 1)
		printf("yes");
	else
		printf("no");

	return 0;
}