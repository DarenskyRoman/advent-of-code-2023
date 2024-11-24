#include <stdio.h>
#include <string.h>

int main(){
	char buff[128];
	unsigned int sum = 0;
    FILE *input = fopen("input.txt", "r");

    if (input == NULL) {
        perror("Error opening file");
        return 1;
    }

    while(fgets(buff, 128, input) != NULL){
    	int firstDigit = 0;
     	int lastDigit = -1;

    	for(int i = 0; ; i++){
     		char ch = buff[i];
       		if(ch >= '0' && ch <= '9'){
         		lastDigit = ch - '0';
         		if(firstDigit == 0)
           			firstDigit = ch - '0';
         	}

       		if(ch == '\0')
         		break;

         	if(ch == '\n')
          		sum += (firstDigit * 10 + lastDigit);
     	}
    }

    printf("%d\n", sum);
    return 0;
}
