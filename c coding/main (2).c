/******************************************************************************

                            Online C Compiler.
                Code, Compile, Run and Debug C program online.
Write your code in this editor and press "Run" button to compile and execute it.

*******************************************************************************/

#include <stdio.h>
#define MAX_SIZE 100

int main()
{

	int size, i;
	int array[MAX_SIZE];
	int *ptr = array;

	printf("choose array size:");
	scanf("%d", &size);


	for (i = 1; i < size + 1; i++)
	{
		printf("choose array element %d:", i);
		scanf("%d", ptr);
		ptr++;
	}
	
	ptr = array;
	printf("the array elements are:\n");
	
		for (i = 0; i < size; i++)
	{
		printf("%d\n", *ptr);
		ptr++;
	}
	
	return 0;
}