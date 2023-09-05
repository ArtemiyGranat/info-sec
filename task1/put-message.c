#include "utils.h"
#include <getopt.h>
#include <stdio.h>
#include <string.h>

#define USAGE()                                                               \
  printf ("Usage: %s [-m <MESSAGE>] [-s <STEGO>] [-h] -c <CONTAINER>\n\n"     \
          "Options:\n"                                                        \
          "-m, --message <MESSAGE> — the message path,\n"                     \
          "-s, --stego <STEGO> — the stegocontainer path,\n"                  \
          "-c, --container <CONTAINER> — the container path,\n"               \
          "-h, --help — print help\n",                                        \
          argv[0])

typedef struct config_t
{
  char *message;
  char *stego;
  char *container;
} config_t;



void
put_message (config_t *config)
{
  FILE *container_fp = NULL, *message_fp = NULL, *stego_fp = NULL;
  container_fp = fopen (config->container, "r");
  message_fp = (config->message) ? fopen (config->message, "r") : stdin;
  stego_fp = (config->stego) ? fopen (config->stego, "w") : stdout;

  char buff[BUFF_SIZE];
  int bits[BITS_PER_BYTE];
  int idx = BITS_PER_BYTE;
  while (fgets (buff, BUFF_SIZE, container_fp) != NULL)
    {
      buff[strcspn (buff, "\n")] = 0;
      if (idx == BITS_PER_BYTE)
        {
          if (!feof (message_fp)) {
            char_to_bits (fgetc (message_fp), bits);
            idx = 0;
          }
        }

      fputs (buff, stego_fp);
      if (!feof (message_fp) && bits[idx])
        {
          fputc (' ', stego_fp);
        }

      if (!feof (container_fp))
        {
          fputc ('\n', stego_fp);
          // TODO: Something is wrong here
          ++idx;
        }
    }

  fclose (container_fp);
  fclose (message_fp);
  fclose (stego_fp);
}

int
parse_params (config_t *config, int argc, char *argv[])
{
  struct option longopts[] = { { "message", required_argument, 0, 'm' },
                               { "stego", required_argument, 0, 's' },
                               { "container", required_argument, 0, 'c' },
                               { "help", no_argument, 0, 'h' },
                               { 0, 0, 0, 0 } };

  int opt = 0;
  while ((opt = getopt_long (argc, argv, "m:s:c:h", longopts, NULL)) != -1)
    {
      switch (opt)
        {
        case 'm':
          config->message = optarg;
          break;
        case 's':
          config->stego = optarg;
          break;
        case 'c':
          config->container = optarg;
          break;
        case 'h':
          USAGE ();
          return -1;
        default:
          USAGE ();
          return -1;
        }
    }

  if (config->container == NULL)
    {
      fprintf (stderr,
               "Error: the required argument 'container' was not provided");
      return -1;
    }

  return 0;
}

int
main (int argc, char *argv[])
{
  config_t config = {
    .message = NULL,
    .stego = NULL,
    .container = NULL,
  };
  if (parse_params (&config, argc, argv) == -1)
    {
      return EXIT_FAILURE;
    }

  put_message (&config);

  return EXIT_SUCCESS;
}
