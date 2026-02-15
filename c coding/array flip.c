#include <stdio.h>
#define SIZE 100

int main()
{
	int i, size, temp;
	int array[SIZE];

	printf("choose array size: ");
	scanf("%d", &size);

	for(i = 0; i < size; i++)
	{
		printf("choose array element %d: ", i + 1);
		scanf("%d", &array[i]);
	}
	
	printf("\nthe array before flipping is: ");

	for(i = 0; i < size; i++)
	{
		printf("%d. ", array[i]);
	}

	for(i = 0; i < (size / 2); i++)
	{
		temp = array[i];
		array[i] = array[size - i - 1];
		array[size - i - 1] = temp;
	}
	
	printf("\nthe flipped array is: ");

	for(i = 0; i < size; i++)
	{
		printf("%d. ", array[i]);
	}

	return 0;
}


