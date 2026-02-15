
#define MAXSIZE 100
#include <stdio.h>

int main()
{
	int i, size, sum;
	int arr[MAXSIZE];
	int *ptr;
	
	ptr = arr;
	
	printf("choose array size: ");
	scanf("%d", &size);

	for(i = 0; i < size; i++)
	{
		printf("choose array element %d: ", i + 1);
		scanf("%d", &arr[i]);
	}
	
	for(i = 0; i < size; i++)
	{
		sum += *ptr;
		ptr++;
		
	}
	
	printf("the sum is: %d", sum);
	
    return 0;
}