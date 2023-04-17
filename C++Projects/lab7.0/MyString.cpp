#include "MyString.h"


MyString::MyString() {
	str = new char[0];
	length = 0;
}

MyString::MyString(char *a, int n) {
	this->str = new char[n];
	this->length = n;
	for (int i = 0; i < n; i++) {
		this->str[i] = a[i];
	}
}

char* MyString::getStr() { return str; }

int MyString::len() { return length; }

char& MyString::operator[](int i) { return this->str[i]; }

char& MyString::access(int i) { return this->str[i]; }

void MyString::setChar(int i, char x) { str[i] = x; }

void MyString::pushChar(int i, char x) { 
	char* s = new char[length + 1];
	this->length++;
	for (int j = 0; j < length; j++) {
		if (j < i) { s[j] = str[j]; }
		if (j == i) { s[j] = x; }
		if (j > i) { s[j] = str[j - 1]; }
	}
	delete str;
	this->str = s;
}

void MyString::copy(MyString* s) {
	char* buf = new char[s->len()];
	this->length = s->len();
	for (int i = 0; i < length; i++) {
		buf[i] = s->access(i);
	}
	this->str = buf;
}

bool MyString::polyndrom() {
	bool res = true;
	for (int i = 0; 2 * i < length; i++) {
		if (str[i] != str[length - i - 1])
			res = false;
	}
	return res;
}
