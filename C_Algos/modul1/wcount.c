#include <string.h>
#include <stdio.h>
#include <math.h>
#include <stdlib.h>


int wcount(char *s) {
	int k = 0;
	for (int i = 0; i < strlen(s); i++)
		if (s[i] != ' ' && (s[i + 1] == ' ' || s[i + 1] == '\0'))
			k++;
	return k;
}


int main() {
	char s[1000];
	char *buf;
	buf = gets(s);
	int sum = wcount(s);
	printf("%d", sum);

	return 0;
}