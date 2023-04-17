#include <iostream>
#include "MyString.h"

int main()
{
	setlocale(LC_ALL, "Russian");
	char a[3] = { 'a', 'b', 'c' };
	MyString* axe = new MyString(a, 3);

	for (int i = 0; i < axe->len(); i++) {
		std::cout << axe->access(i) << " ";
	}

	std::cout << "\nСоздали строку из символов a b c\n";

	int count = 0;
	while (!axe->polyndrom() && count < 5) {
		axe->setChar(0, axe->access(0) + 1);
		for (int i = 0; i < axe->len(); i++) {
			std::cout << axe->access(i) << " ";
		}
		count++;
		std::cout << "\n";
	}

	std::cout << "Заменяли первый символ на следущий в \n";
	std::cout << "ASCII пока не получится полиндром\n\n";
	std::cout << "добавим элементов в строку:\n";
	axe->pushChar(0, 'a');
	axe->pushChar(1, 'b');
	for (int i = 0; i < axe->len(); i++) {
		std::cout << axe->access(i) << " ";
	}
	std::cout << "\n\nСоздадим второй класс и скопируем";
	std::cout << "\nданные из первого, выведем его:";


	MyString* diamond = new MyString();
	diamond->copy(axe);
	for (int i = 0; i < diamond->len(); i++) {
		std::cout << diamond->access(i) << " ";
	}
	std::cout << "\n";

	return 0;
}

