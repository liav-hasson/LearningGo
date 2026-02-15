#include <stdio.h>

void EvenOdd(int a);

int main()
{
    int num;
    printf("pick any real number: ");
    scanf("%d", &num);
    
    EvenOdd(num);

    return 0;
}

void EvenOdd(int a)
{
    if(a % 2 == 0)
    {
        printf("%d is even!", a);
    } else
    {
        printf("%d is odd!", a);
    }
}
