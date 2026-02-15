#include <stdio.h>
#define MAXSIZE 100

float calcavrage(int array[], int size);

float main()
{
	int arr[MAXSIZE];
	int arrsize;
	float average;
	
	printf("choose how many numbers: ");
	scanf("%d", &arrsize);

	for(int i = 0; i < arrsize; i++)
	{
		printf("choose number %d: ", i + 1);
		scanf("%d", &arr[i]);
	}

	average = calcavrage(arr, arrsize);
	printf("the average of the numbers is %.3f", average);

	return 0;
}

float calcavrage(int array[], int size)
{
	float sum = 0;

	for(int i = 0; i < size; i++)
	{
		sum += array[i];
	}

	return sum / size;
}