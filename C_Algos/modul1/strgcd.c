#include <stdio.h>
#include <math.h>
#include <stdlib.h>

int gcd(int a, int b) {
	if (a < b) {
		int k = a;
		a = b;
		b = k;
	}
	while (a % b != 0) {
		int x = a;
		a = b;
		b = x % b;
	}
	return b;
}

void compute_lgithms(int m, int size, int *lg)
{
	int j = 0;
	for (int i = 1; i <= m; i++)
		for (; j < (1 << i); j++)
			lg[j] = i - 1;
}


int sparse_table_query(int **ST, int l, int r, int *lg)
{
	int j = lg[r - l + 1];
	return gcd(ST[l][j], ST[r - (1 << j) + 1][j]);
}

void sparse_table_build(int *v, int n, int *lg, int ***pST)
{
	int m = lg[n] + 1,
		**ST = (int**)malloc(sizeof(int *) * n);
	for (int i = 0; i < n; i++)
		ST[i] = (int*)malloc(sizeof(int) * m);

	for (int i = 0; i < n; i++) ST[i][0] = v[i];

	for (int i = 1; i < m; i++)
		for (int j = 0; j <= n - (1 << i); j++)
			ST[j][i] = gcd(ST[j][i - 1], ST[j + (1 << (i - 1))][i - 1]);
	*pST = ST;
}

int main() {
	int n, m;
	scanf("%d", &n);
	int *s = (int *)malloc(sizeof(int) * n);
	for (int i = 0; i < n; ++i)
		scanf("%d", &s[i]);
	scanf("%d", &m);

	int c = log2(n) + 1,
		lg_size = (1 << (m + 1)),
		lg[lg_size];

	compute_lgithms(m, lg_size, lg);
	int **ST;
	sparse_table_build(nums, n, lg, &ST);
	  
	int m;
	scanf("%i", &m);
	for (int i = 0; i < m; i++)
	{
		int l, r;
		scanf("%i %i", &l, &r);
		printf("%d\n", sparse_table_query(ST, l, r, lg));
	}

	for (int i = 0; i < n; i++)
		free(ST[i]);
	free(ST);
	return 0;
}


int main()
{
	int n, m;
	scanf("%d", &n);
	int *s = (int *)malloc(sizeof(int) * n);
	for (int i = 0; i < n; ++i)
		scanf("%d", &s[i]);
	scanf("%d", &m);
	for (int i = 0; i < m; i++) {
		int l, r;
		scanf("%d%d", &l, &r);
		int x = abs(s[l]);
		for (int i = l; i < r; i++) {
			if (x > gcd(x, abs(s[i + 1])))
				x = gcd(x, abs(s[i + 1]));
		}
		printf("%d\n", x);
	}

	free(s);

	return 0;
}
