#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>

#define BUFFER_SIZE 1024

int main (int argc, char *argv[]) {

  int sock, port;
  char buf[BUFFER_SIZE];

  if (argc < 2) {
    printf("Usage: %s [port]\n", argv[0]);
    exit(EXIT_FAILURE);
  }

  port = atoi(argv[1]);
  { // sestup addr
    int opt_val, err;
    struct sockaddr_in addr;

    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    addr.sin_addr.s_addr = htonl(INADDR_ANY);

    sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock < 0) {
      perror("Could not create socket");
      exit(EXIT_FAILURE);
    }
    opt_val = 1;
    setsockopt(sock, SOL_SOCKET, SO_REUSEADDR, &opt_val, sizeof opt_val);

    err = bind(sock, (struct sockaddr *) &addr, sizeof(addr));
    if (err < 0) {
      perror("Could not bind socket");
      exit(EXIT_FAILURE);
    }
  }
  if (listen(sock, 1) < 0) {
    perror("Unable to listen");
    exit(EXIT_FAILURE);
  }

  /* Handle connections */
  printf("Server is listening on %d\n", port);
  while (1) {
    int con;
    struct sockaddr_in peer_addr;
    socklen_t client_len = sizeof(peer_addr);
    con = accept(sock, (struct sockaddr *) &peer_addr, &client_len);

    if (con < 0) {
      perror("Could not establish new connection");
      exit(EXIT_FAILURE);
    }
    fprintf(stderr, "Connection 1 established\n");
    while (1) {
      int len, err;

      len = recv(con, buf, BUFFER_SIZE, 0);
      if (!len) break; // done reading
      if (len < 0) {
        perror("Client read failed");
        close(con);
        break;
      }
      err = send(con, buf, len, 0);
      if (err < 0) {
        perror("Client write failed");
        close(con);
        break;
      }
    }
    close(con);
  }
  close(sock);
  return 0;
}
