#include <stdlib.h>
#include <stdio.h>

void bubblesort(unsigned long nel,
	int(*compare)(unsigned long i, unsigned long j),
	void(*swap)(unsigned long i, unsigned long j))
{
	int l = 0, r = nel - 1, count = 1;
	int l1 = 0, r1 = nel - 1;
	while (count > 0)
	{
		count = 0;
		for (int i = l; i < r; i++)
		{
			if (compare(i, i + 1) > 0)
			{
				swap(i, i + 1);
				count++;
				r1 = i;
			}
		}

		r = r1;
		if (count == 0) 
			break;

		count = 0;
		for (int i = r; i > l; i--)
		{
			if (compare(i - 1, i) > 0)
			{
				swap(i - 1, i);
				count++;
				l1 = i;
			}
		}
		l = l1;
	}
}