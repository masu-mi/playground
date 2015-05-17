#include <stdio.h>
#include <stdlib.h>

extern int fee_count, fie_count, foe_count, fom_count;
extern int yylex( void );

int main( int args, char** argv )
{
  yylex();
  printf("%d %d %d %d\n", fee_count, fie_count, foe_count, fom_count);
  exit( 0 );
}
