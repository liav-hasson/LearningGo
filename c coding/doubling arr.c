#include <stdio.h>
#include <stdlib.h>

void doublearray(int *arr, int n);

int main()
{
	int size;

	printf("choose array size: ");
	scanf("%d", &size);

	int *ptr = (int *)malloc(sizeof(int) * size);
	if(ptr == NULL)
	{
	    printf("Allocation Failed");
	    return 1;
	}

	for(int i = 0; i < size; i++)
	{
		printf("choose element %d: ", i + 1);
		scanf("%d", &ptr[i]);
	}
	
	printf("the array before is: ");

	for(int i = 0; i < size; i++)
	{
		printf("%d. ", ptr[i]);
	}
	
	printf("\nthe new array is: "); 

	doublearray(ptr, size);
	
	for(int i = 0; i < size; i++)
	{
	    printf("%d. ", ptr[i]);
	}
	
	free(ptr);
	return 0;
}

void doublearray(int *p, int n)
{
    for(int i = 0; i < n; i++)
    {
        p[i] *= 2;
    }
}