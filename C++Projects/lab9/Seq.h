#pragma once
#include <vector>


template <typename T>
class Seq
{
private:
	std::vector<T> a;
	int size;
public:
	int getSize();
	Seq(std::vector<T> b);
	Seq<T> operator+(Seq<T> b);
	std::vector<Seq<T>> operator/(int k);
	Seq<T> operator!();
	T operator[](int k);
	void print();
};


template <typename T>
Seq<T>::Seq(std::vector<T> b) {
	this->size =(int) b.size();
	for (int i = 0; i < (int)b.size(); i++) {
		this->a.push_back(b[i]);
	}
}


template <typename T>
Seq<T> Seq<T>::operator+(Seq<T> b) {
	std::vector<T> res(0);
	for (int i = 0; i < this->getSize(); i++) {
		res.push_back(this->a[i]);
	}
	for (int i = 0; i < b.getSize(); i++) {
		res.push_back(b[i]);
	}
	return Seq<T>(res);
}


template <typename T>
int Seq<T>::getSize() {
	return this->size;
}

template <typename T>
std::vector<Seq<T>> Seq<T>::operator/(int k) {
	std::vector<Seq<T>> res;

	int delta = this->getSize() % k;
	int part = 0;

	for (int i = 0; i < k; i++) {
		std::vector<T> buf;
		if (delta > 0) { part = 1; }
		else { part = 0; }
		for (int j = 0; j < this->getSize() / k + part; j++) {
			buf.push_back(this->a[i * (this->getSize() / k) + (this->getSize() % k) - delta + j]);
		}
		Seq<T> temp(buf);
		if (delta > 0) {
			delta--;
		}
		res.push_back(temp);
	}
	return res;
}


template <typename T>
Seq<T> Seq<T>::operator!() {
	std::vector<T> res;
	for (int i = 0; i < this->getSize(); i++) {
		if (std::find(res.begin(), res.end(), this->a[i]) == res.end()) {
			res.push_back(this->a[i]);
		}
	}
	return Seq<T>(res);
}

template <typename T>
T Seq<T>::operator[](int k) {
	return this->a[k];
}

template <typename T>
void Seq<T>::print() {
	for (int i = 0; i < this->getSize(); i++) {
		std::cout << a[i] << " ";
	}
}