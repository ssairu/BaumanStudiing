#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define scanf scanf_s


char *concat(char **s, int n) {
	size_t length = 1;
	for (int i = 0; i < n; i++) {
		length += strlen(s[i]);
	}
	char *res = (char *)malloc(length);
	sum = 0;
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < strlen(s[i]); j++) {
			res[i + sum] = s[i][j];
			sum += strlen(s[i]);
		}
	}
	res[length - 1] = '\0';

	return (char *)res;
}


int main()
{
	char * a[3];
	char * buf;
	
	scanf("%5s", a[0]);
	scanf("%5s", a[1]);
	scanf("%5s", a[2]);


	printf("%s", concat(a, 3));


	return 0;
}