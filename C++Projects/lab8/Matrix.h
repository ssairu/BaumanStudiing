#pragma once
#include <iostream>
#include <vector>
#include <math.h>
#include <string>

template <typename T, int N>
class Matrix
{
private:
	std::vector<std::vector<T>> a;
	long long a1;
public:
	Matrix(T x);
	bool isBoolSmall;
	void print();
	void printMatrix();
	void setValue(int a, int b, T v);
	T getValue(int a, int b);
	Matrix<T, N - 1>* minor(int a, int b);
};

template <typename T, int N>
void Matrix<T, N>::print()
{
	std::cout << N << "\n";
}


template <typename T, int N>
Matrix<T, N>::Matrix(T x) {
	this->isBoolSmall = false;
	this->a1 = 0;
	this->a.resize(N);
	for (int i = 0; i < N; i++) {
		this->a[i].resize(N);
		for (int j = 0; j < N; j++) {
			this->a[i][j] = x;
		}
	}
}


template <typename T, int N>
void Matrix<T, N>::printMatrix() {
	if (this->isBoolSmall) {
		for (int i = 0; i < N; i++) {
			for (int j = 0; j < N; j++) {
				std::cout << (long long)(a1 / pow(2, (8 * i + j))) % 2 << " ";
			}
			std::cout << "\n";
		}
	}
	else {
		for (int i = 0; i < N; i++) {
			for (int j = 0; j < N; j++) {
				std::cout << a[i][j] << " ";
			}
			std::cout << "\n";
		}
	}
}


template <typename T, int N>
void Matrix<T, N>::setValue(int a, int b, T v) {
	if (this->isBoolSmall) {
		a1 |= (long long)pow(2, 8 * a + b);
		if (!v) {
			a1 /= (long long)pow(2, 8 * a + b);
		}
	}
	else {
		this->a[a][b] = v;
	}
}


template <typename T, int N>
T Matrix<T, N>::getValue(int a, int b) {
	if (this->isBoolSmall) {
		return (long long)(a1 / pow(2, (8 * a + b))) % 2 == 1;
	}
	else {
		return this->a[a][b];
	}
}


template <typename T, int N>
Matrix<T, N - 1>* Matrix<T, N>::minor(int x, int y) {
	Matrix<T, N - 1>* b = new Matrix<T, N - 1>(this->a[0][0]);
	if (b->isBoolSmall) {
		int m = 0, l = 0;
		for (int i = 0; i < N; i++) {
			if (i == x) { continue; }
			else { m++; }
			for (int j = 0; j < N; j++) {
				a1 *= (long long)pow(2, 8 * m + l);
			}
			std::cout << "\n";
		}
	}
	else {
		int m = 0, l = 0;
		for (int i = 0; i < N; i++) {
			if (i == x) { continue; }
			for (int j = 0; j < N; j++) {
				if (j == y) { continue; }
				b->setValue(m, l, a[i][j]);
				l++;
			}
			m++;
			l = 0;
		}
	}
	return b;
}