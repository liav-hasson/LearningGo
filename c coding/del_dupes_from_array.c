#include <stdio.h>
#define MAXSIZE 100

void deleteelement(int arr[], int j, int size);


int main()
{
	int size, num, i, j;
	int array[MAXSIZE];

	printf("choose array size: ");
	scanf("%d", &size);

	for(i = 0; i < size; i++)
	{
		printf("choose array element %d: ", i + 1);
		scanf("%d", &array[i]);
	}

	for(i = 0; i < size - 1; i++)
	{
		for(j = i + 1; j < size; j++)
		{
			if(array[i] == array[j])
			{
				deleteelement(array, j, size);
				size--;
				j--;
			}
		}
	}

	printf("\nthe array without duplicates is: ");

	for(i = 0; i < size; i++)
	{
		printf("%d. ", array[i]);
	}

	return 0;
}

void deleteelement(int arr[], int j, int size)
{
	int i;
	for(i = j; i < size - 1; i++)
	{
		arr[i] = arr[i + 1];
	}
}











