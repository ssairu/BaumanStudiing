#include <stdio.h>

#define scanf scanf_s



void ssort(int *a, int left, int right)
{
	for (int j = right; j > left; j--)
	{
		int m = j;
		for (int i = j - 1; i >= left; i--)
			if (a[i] > a[m])
				m = i;

		int buf = a[j];
		a[j] = a[m];
		a[m] = buf;
	}
}

int part(int *a, int left, int right)
{
	int i = left;

	for (j = left; j < right; j++)
	{
		if (a[j] < a[right])
		{
			int buf = a[i];
			a[i] = a[j];
			a[j] = buf;
			i++;
		}
	}

	int buf = a[i];
	a[i] = a[right];
	a[right] = buf;
	return i;
}

void sort1(int *a, int left, int right, int m)
{
	if (right - left + 1 < m)
		ssort(a, left, right);
	else
	{
		while (left < right)
		{
			int q = part(a, left, right);
			if (q - left < right - q)
			{
				sort1(a, left, q - 1, m);
				left = q + 1;
			}
			else
			{
				sort1(a, q + 1, right, m);
				right = q - 1;
			}
		}
	}
}

void quicksort(int *a, int n, int m)
{
	sort1(a, 0, n - 1, m);
}

int main()
{
	int n, m;
	scanf("%d %d", &n, &m);
	int a[n];
	for (int i = 0; i < n; i++) scanf("%d", &a[i]);
	quicksort(a, n, m);
	for (int i = 0; i < n; i++) printf("%d ", a[i]);
	printf("\n");
	return 0;
}


















