#include <stdio.h>
#include <stdlib.h>

int mystrcpy(const char *src, char *dest)
{
    if(src == NULL || dest == NULL){
        return -1;
    }

    while(*src)
        *dest++ = *src++;//address move

    *dest = '\0';
    return 0;
}

int main (int argc, const char *argv[])
{
    int len;
    char * src;
    char * dest;

    printf("input string len >");
    scanf("%[^\n]", &len);
    getchar();

    printf("%d \n", len);

}