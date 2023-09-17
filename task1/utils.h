#ifndef UTILS_H
#define UTILS_H

#include <stdlib.h>

#define BUFF_SIZE (16384 * 100)
#define BITS_PER_BYTE 8

void
char_to_bits (char ch, int *bits)
{
  for (size_t i = 0; i < BITS_PER_BYTE; ++i)
    {
      bits[i] = (ch >> (7 - i)) & 1;
    }
}

char
bits_to_char (char ch, int *bits)
{
  for (size_t i = 0; i < BITS_PER_BYTE; ++i)
    {
      ch = (ch << 1) | (bits[i] & 1);
    }

  return ch;
}
#endif // UTILS_H
