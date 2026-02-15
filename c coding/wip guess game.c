
#include <stdio.h>
#include <time.h>
#include <stdlib.h>

int main()
{
	int num, guess;
	char check;

	printf("I chose a number between 1 - 100, and you try to guess it. ");

	srand(time(0));
	num = (rand() % 100) + 1;

	while(num != guess)
	{
		while(scanf("%d%c", &guess, &check) != 2 || check != '\n' || guess > 100 || guess < 1)
		{
			printf("Invalid input, try again! ");
		}
		
		if(num > guess)
		{
			printf("my number is higher. try again! ");
		}
		else
		{
			printf("my number is lower. try again! ");
		}
	}

	printf("\nyou got it right!");

	return 0;
}