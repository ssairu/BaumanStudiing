#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Elem {
	struct Elem *next;
	char *word;
} elem;


void swap(elem *a, elem *b)
{
	char* t = a->word;
	a->word = b->word;
	b->word = t;
}

elem *bsort(struct elem *xs) {
	for (elem *now = NULL; now != xs; now = i) {
		elem *i = xs;
		while (i->next != now) {
			if (strlen(i->next->word) < strlen(i->word)) {
				swap(i->next, i);
				i = i->next;
			}
		}
	}
	return xs;
}

signed main() {
	elem *xs = NULL;

	char* w[2000];

	int count = 0;
	while (scanf("%ms", &w[count]) != EOF) {
		count++;
	}


	while (cnt > 0) {
		count--;
		elem *newelem = (elem *)malloc(sizeof(elem *));
		newelem->word = w[count];
		newelem->next = xs;
		xs = newelem;
	}

	xs = bsort(xs);

	while (xs) {
		elem *next = xs->next;
		printf("%s ", xs->word);
		free(xs->word);
		free(xs);
		xs = next;
	}

	return 0;
}