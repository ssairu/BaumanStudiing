#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define scanf scanf_s




int main(int argc, char * argv[])
{
	if (argc != 4)
		printf("Usage: frame <height> <width> <text>");
	else {
		int a = atoi(argv[2]), b = atoi(argv[1]);
		if (strlen(argv[3]) > a - 2 || b < 3)
			printf("ERROR");
		else {
			for (int i = 0; i < a; i++)
				printf("*");
			printf("\n");
			for (int i = 0; i < b - 2; i++) {
				if (i == (b - 2) / 2 + (b - 2) % 2 - 1) {
					printf("*");
					for (int j = 0; j < (a - 2 - strlen(argv[3])) / 2; j++)
						printf(" ");
					printf("%s", s);
					for (int j = 0; j < (a - 2 - strlen(argv[3])) / 2 + (a - 2 - strlen(argv[3])) % 2; j++)
						printf(" ");
					printf("*\n");
				}
				else {
					printf("*");
					for (int j = 0; j < a - 2; j++)
						printf(" ");
					printf("*\n");
				}
			}
			for (int i = 0; i < a; i++)
				printf("*");
			printf("\n");
		}
	}
	return 0;
}
