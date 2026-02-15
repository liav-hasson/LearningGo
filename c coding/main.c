/******************************************************************************

                            Online C Compiler.
                Code, Compile, Run and Debug C program online.
Write your code in this editor and press "Run" button to compile and execute it.

*******************************************************************************/

#include <stdio.h>

float main()
{
    float num1, num2, sum, difference, product, quotient;
    float *ptr1, *ptr2;
    
    ptr1 = &num1;
    ptr2 = &num2;
    
    printf("choose 2 numbers:"); 
    scanf("%f%f", ptr1, ptr2);
    
    sum = *ptr1 + *ptr2;
    difference = *ptr1 - *ptr2;
    product = *ptr1 * *ptr2;
    quotient = *ptr1 / *ptr2;
    
    printf("the sum is %.2f\n", sum);
    printf("the difference is %.2f\n", difference);
    printf("the product is %.2f\n", product);
    printf("the quotient is %.2f\n", quotient);

    return 0;
}