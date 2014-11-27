
// Explaining indexing & array arithmetic to people

#include <stdio.h>
 
int main (){
   int n[10];
   int i,j;

   for ( i = 0; i < 10; i++ ){
      n[ i ] = i + 100;
   }
   
   printf("%d\n", *(n + 1));
 
   return 0;
}