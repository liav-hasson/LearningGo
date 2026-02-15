/******************************************************************************

                            Online C Compiler.
                Code, Compile, Run and Debug C program online.
Write your code in this editor and press "Run" button to compile and execute it.

*******************************************************************************/

#include <stdio.h>

void swaper(int *num1, int *num2);

int main()
{
	int num1, num2;

	printf("choose 2 numbers to swap:\n");
	printf("choose first number:");
	scanf("%d", &num1);
	printf("choose second number:");
	scanf("%d", &num2);

	swaper(&num1, &num2);

	printf("the first number is now: %d\n", num1);
	printf("the second number is now %d\n", num2);

	return 0;
}

void swaper(int *num1, int *num2)
{
	int temp = *num1;
	*num1 = *num2;
	*num2 = temp;
}