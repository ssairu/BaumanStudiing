#pragma once

class MyString
{
private:
	char* str;
	int length;
public:
	MyString(char *a, int n);
	MyString();
	char* getStr();
	int len();
	char& operator[](int i);
	char& access(int i);
	void setChar(int i, char x);
	void pushChar(int i, char x);
	bool polyndrom();
	~MyString();
	void copy(MyString* s);
};