#include "Seq.h"
#include <iostream>

int main()
{
    std::cout << "Create some char and int sequances\n";
	std::cout << "and print them:\n";
	std::vector<char> a(7, 'A');
	std::vector<char> b(4, 'B');
	Seq<char> a1(a);
	a1.print();
	std::cout << "\n";
	Seq<char> b1(b);
	b1.print();
	std::cout << "\n";

	std::vector<int> ai(15, 0);
	std::vector<int> bi(12, 1);
	Seq<int> a1i(ai);
	a1i.print();
	std::cout << "\n";
	Seq<int> b1i(bi);
	b1i.print();
	std::cout << "\n\n";

	std::cout << "find a sum of int and char:\n";
	a1 = a1 + b1;
	a1.print();
	std::cout << "\n";
	a1i = a1i + b1i;
	a1i.print();
	std::cout << "\n\n";


	std::cout << "check operators '/' and '!':\n";
	std::cout << "char seq devides by 3\n";
	std::vector<Seq<char>> div = a1 / 3;
	std::cout << "{ ";
	for (Seq<char> x : div) {
		std::cout << "[ ";
		x.print();
		std::cout << "] ";
	}
	std::cout << "}\n";

	for (int i = 0; i < div.size(); i++) {
		div[i] = !div[i];
	}

	std::cout << "{ ";
	for (Seq<char> x : div) {
		std::cout << "[ ";
		x.print();
		std::cout << "] ";
	}
	std::cout << "}\n\n";

	std::cout << "int seq devides by 6\n";
	std::vector<Seq<int>> divi = a1i / 6;
	std::cout << "{ ";
	for (Seq<int> x : divi) {
		std::cout << "[ ";
		x.print();
		std::cout << "] ";
	}
	std::cout << "}\n";

	for (int i = 0; i < divi.size(); i++) {
		divi[i] = !divi[i];
	}

	std::cout << "{ ";
	for (Seq<int> x : divi) {
		std::cout << "[ ";
		x.print();
		std::cout << "] ";
	}
	std::cout << "}\n";


}
