#include "utils.h"
#include <getopt.h>
#include <stdio.h>
#include <string.h>

#define USAGE()                                                               \
  printf ("Usage: %s [-m <MESSAGE>] [-s <STEGO>] [-h]\n\n"                    \
          "Options:\n"                                                        \
          "-m, --message <MESSAGE> — the message path,\n"                     \
          "-s, --stego <STEGO> — the stegocontainer path,\n"                  \
          "-h, --help — print help\n",                                        \
          argv[0])

typedef struct config_t
{
  char *message;
  char *stego;
} config_t;

void
get_message (config_t *config)
{
  FILE *stego_fp = NULL, *message_fp = NULL;
  stego_fp = (config->stego) ? fopen (config->stego, "r") : stdin;
  message_fp = (config->message) ? fopen (config->message, "w") : stdout;

  char buff[BUFF_SIZE];
  int bits[BITS_PER_BYTE];
  int idx = 0;
  char ch;
  while (fgets (buff, BUFF_SIZE, stego_fp) != NULL)
    {
      printf ("%d", buff[strcspn (buff, "\n") - 1] == ' ');
      bits[idx] = buff[strcspn (buff, "\n") - 1] == ' ';
      ++idx;

      if (idx == BITS_PER_BYTE)
        {
          ch = bits_to_char (ch, bits);
          printf ("%c", ch);
          if (ch)
            {
              idx = 0;
              fputc (ch, message_fp);
            }
        }
    }
}

int
parse_params (config_t *config, int argc, char *argv[])
{
  struct option longopts[] = { { "message", required_argument, 0, 'm' },
                               { "stego", required_argument, 0, 's' },
                               { "help", no_argument, 0, 'h' },
                               { 0, 0, 0, 0 } };

  int opt = 0;
  while ((opt = getopt_long (argc, argv, "m:s:h", longopts, NULL)) != -1)
    {
      switch (opt)
        {
        case 'm':
          config->message = optarg;
          break;
        case 's':
          config->stego = optarg;
          break;
        case 'h':
          USAGE ();
          return -1;
        default:
          USAGE ();
          return -1;
        }
    }

  return 0;
}

int
main (int argc, char *argv[])
{
  config_t config = {
    .message = NULL,
    .stego = NULL,
  };
  if (parse_params (&config, argc, argv) == -1)
    {
      return EXIT_FAILURE;
    }

  get_message (&config);

  return EXIT_SUCCESS;
}
