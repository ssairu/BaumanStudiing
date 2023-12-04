#include <stdio.h>
#include <stdlib.h>
#include <string.h>


typedef struct Elem {
	struct Elem *next;
	char *word;
} elem;


elem *init(char *word) {
	elem *list;
	list = (se*)malloc(sizeof(se));
	list->next = NULL;
	list->word = word;

	return list;
}


int isempty(se *list) {
	return list == NULL;
}


int len(se *list) {
	int len = 0;
	for (se *x = list; x != NULL; x = x->next) {
		len++;
	}
	return len;
}


se *insert(se *l, se *y) {
	se *x = l;
	while (x->next) {
		x = x->next;
	}
	x->next = y;
	return l;
}


void delete(se *l) {
	se* y = l;
	l = y->next;
	y->next = NULL;
}

elem *bsort(elem *list) {
	if (!list || !list->word) {
		return NULL;
	}
	se *element;
	se *t = list;
	while (t->next) {
		t = t->next;
	}
	while (t != list) {
		se *bound = t;
		t = list;
		for (elem = list; elem != bound && elem->next != NULL; elem = elem->next) {
			if (strlen(elem->next->word) < strlen(elem->word)) {
				char *temp = elem->word;
				elem->word = elem->next->word;
				elem->next->word = temp;

				t = elem->next;
			}
		}
	}
	return list;
}


int main() {
	char s[1001], *word;
	scanf("%[^\n]s", s);
	word = strtok(s, " ");
	se *list = init(word), *elem;
	while (1) {
		word = strtok(NULL, " ");
		if (!word) {
			break;
		}
		elem = init(word);
		insert(list, elem);
	}
	list = bsort(list);
	while (list) {
		printf("%s ", list->word);
		elem = list;
		list = list->next;
		free(elem);
	}
}