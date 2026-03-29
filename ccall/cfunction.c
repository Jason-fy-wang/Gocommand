#include <stdio.h>
#include "ccall.h"



int main(int argc, char* argv[]) {
    PrintMessageFromgo();

    GoInt x = 12;
    GoInt y = 34;

    GoInt p = Multiply(x, y);
    printf("Multiply(%lld, %lld) = %lld\n", x, y, p);
    printf("it worked!\n");
    return 0;
}

