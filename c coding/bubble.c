#include <stdio.h>

void NumberFlip(int *num1, int *num2);

int main()
{

	int array[10];
	int i, j, swapped;

	printf("choose 10 numbers to sort!\n");

	for(i = 0; i < 10; i++)
	{
		printf("choose number %d:", i + 1);
		scanf("%d", &array[i]);
	}
	
	printf("\nthe array after 0 iteration:");

		for(i = 0; i < 10; i++)
		{
			printf("%d. ", array[i]);
		}

	for(j = 0; j < 10; j++)
	{
	    swapped = 0;
	    
		for(i = 0; i < 9; i++)
		{
			if(array[i] > array[i + 1])
			{
				NumberFlip(&array[i], &array[i + 1]);
				swapped = 1;
			}
		}

		if(swapped == 0)
		{
		    printf("\nthe array is sorted!");
		    break;

		}
		
		printf("\nthe array after %d iteration: ", j + 1);

		for(i = 0; i < 10; i++)
		{
			printf("%d. ", array[i]);
		}
		
	}

	return 0;
}

void NumberFlip(int *num1, int *num2)
{
	int temp = *num1;
	*num1 = *num2;
	*num2 = temp;
}

















