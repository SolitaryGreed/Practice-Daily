#include <ostream>
#include <bitset>
#include <iostream>

using namespace std;

int main(int argc, char* argv[]) {
   
	for (;;) {

		string str;
		cin >> str;

		try {
			bitset<16> bitset(str);
			cout << bitset.to_ulong() << '\n';
		}
		catch (const std::exception&) {
			cout << "Enter valid data" << '\n';
		}
	}
}
