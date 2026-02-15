#include <stdio.h>
#define MAXSIZE 100

void printarray(int array[], int arraysize);

int main()
{
	int array1[MAXSIZE];
	int array2[MAXSIZE];
	int size, i;

	int *ptr1 = array1;
	int *ptr2 = array2;

	printf("choose array size:");
	scanf("%d", &size);
	
		if(size > 100 || size < 1)
	{
	    return 1;
	}

	for(i = 0; i < size; i++)
	{
	    printf("choose array value number %d:", i + 1);
		scanf("%d", (ptr1 + i));
	}

	ptr1 = array1;

	for(i = 0; i < size; i++)
	{
		*ptr2 = *ptr1;
		ptr1++;
		ptr2++;
	}

	printf("the original array is:\n");
	printarray(array1, size);

	printf("the copied array is:\n");
	printarray(array2, size);

	return 0;
}

void printarray(int *arr, int arraysize)
{
	int i;

	for(i = 0; i < arraysize; i++)
	{
		printf("%d\n", *(arr + i));
	}
}