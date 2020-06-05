#include "libsdpobj.h"
#include <iostream>

using namespace std;

int main(){
//    auto sdpStr = "v=0\no=jdoe 2890844526 2890842807 IN IP4 10.47.16.5\ns=SDP Seminar\n";
    auto result = c_parse("v=0\no=jdoe 2890844526 2890842807 IN IP4 10.47.16.5\ns=SDP Seminar\n");
//    std::string word(result.r0, result.r1);
//    cout << result.r1 << endl;
//    cout << result.r0 << endl;
    string word = result.r1;
    cout << word << endl;
    return 0;
}