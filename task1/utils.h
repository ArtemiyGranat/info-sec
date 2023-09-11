#ifndef COMMON_H
#define COMMON_H

#include <stdlib.h>

#define BUFF_SIZE 16384
#define BITS_PER_BYTE 8

void
char_to_bits (char ch, int *bits)
{
  for (size_t i = 0; i < BITS_PER_BYTE; ++i)
    {
      bits[i] = (ch >> (7 - i)) & 1;
    }
}

void
bits_to_char (char *ch, int *bits)
{
  for (size_t i = 0; i < BITS_PER_BYTE; ++i)
    {
      *ch = (*ch << 1) | (bits[i] & 1);
    }
}

#endif // COMMON_H
