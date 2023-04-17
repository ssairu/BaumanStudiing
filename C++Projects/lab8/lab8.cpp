#include <iostream>
#include "Matrix.h"
#include <math.h>
#include <string>
#include <vector>


int main()
{
	std::cout << "define new int matrix with dim 10:\n";
	Matrix<int, 10>* a = new Matrix<int, 10>(0);
	a->printMatrix();

	std::cout << "\n\n set some values in matrix: \n";

	a->setValue(1, 3, 12);
	a->setValue(2, 4, 1);
	a->setValue(4, 0, 3);
	a->setValue(8, 5, -4);
	a->setValue(5, 9, 55);

	a->printMatrix();

	std::cout << "\n\n element from position(1, 3): " << a->getValue(1, 3);

	Matrix<int, 9>* b = a->minor(1, 1);

	std::cout << "\n\n minor from (1, 1):\n";
	b->printMatrix();

	std::cout << "\n\n string matrix:\n";
	Matrix<std::string, 8>* c = new Matrix<std::string, 8>("WOW");
	c->printMatrix();
}