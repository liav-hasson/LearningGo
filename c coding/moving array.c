#include <stdio.h>
#define MAXSIZE 10

void arrayleft(int arr[], int SIZE, int N);
void arrayright(int arr[], int SIZE, int N);

int main()
{
	int array[MAXSIZE];
	int size, i, n, direction;

	printf("choose array size: ");
	scanf("%d", &size);

	if(size > 10 || size < 2)
	{
		printf("invalid size!");
		return 1;
	}

	for(i = 0; i < size; i++)
	{
		printf("choose array value %d: ", i + 1);
		scanf("%d", &array[i]);
	}

	printf("\nby how much to rotate the numbers? ");
	scanf("%d", &n);

	printf("\nto rotate right -> 0, to rotate left -> 1: ");
	scanf("%d", &direction);

	if(direction != 0 && direction != 1)
	{
		printf("\ninvalid direction!");
		printf("\nto rotate right -> 0, to rotate left -> 1: ");
		scanf("%d", &direction);
	}

	printf("your array before is: ");
	for(i = 0; i < size; i++)
	{
		printf("%d. ", array[i]);
	}

	if(direction == 0)
	{
		arrayright(array, size, n);
	} else
	{
		arrayleft(array, size, n);
	}

	printf("\nyour modified array is: ");
	for(i = 0; i < size; i++)
	{
		printf("%d. ", array[i]);
	}

	return 0;
}

void arrayleft(int arr[], int SIZE, int N)
{
	int TEMP, i, j;

	for(i = 0; i < N; i++)
	{
		TEMP = arr[0];
		for(j = 0; j < SIZE - 1; j++)
		{
			arr[j] = arr[j + 1];
		}
		arr[SIZE - 1] = TEMP;
	}

}

void arrayright(int arr[], int SIZE, int N)
{
	int TEMP, i, j;

	for(i = 0; i < N; i++)
	{
		TEMP = arr[SIZE - 1];
		for(j = SIZE - 1; j > 0; j--)
		{
			arr[j] = arr[j - 1];
		}
		arr[0] = TEMP;
	}
}










