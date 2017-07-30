#include <stdio.h>
#include <unistd.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <openssl/ssl.h>
#include <openssl/err.h>

#define BUFFER_SIZE 1024

#define CERT_FILE_PATH "files/cert.pem"
#define KEY_FILE_PATH  "files/key.pem"

int main(int argc, char **argv)
{
  int sock, port;
  char buf[BUFFER_SIZE];

  SSL_CTX *ctx;
  const SSL_METHOD *method;

  if (argc < 2) {
    printf("Usage: %s [port]\n", argv[0]);
    exit(EXIT_FAILURE);
  }

  // init openssl lib.
  SSL_library_init(); // call OpenSSL_add_ssl_algorithms();
  ERR_load_BIO_strings();
  SSL_load_error_strings();

  // setup tls context.
  method = TLSv1_2_server_method(); // SSLv23_server_method();
  ctx = SSL_CTX_new(method);
  if (!ctx) {
    perror("Unable to create SSL context");
    ERR_print_errors_fp(stderr);
    exit(EXIT_FAILURE);
  }
  /* Set the key and cert */
  if (SSL_CTX_use_certificate_file(ctx, CERT_FILE_PATH, SSL_FILETYPE_PEM) <= 0) {
    ERR_print_errors_fp(stderr);
    exit(EXIT_FAILURE);
  }
  if (SSL_CTX_use_PrivateKey_file(ctx, KEY_FILE_PATH, SSL_FILETYPE_PEM) <= 0 ) {
    ERR_print_errors_fp(stderr);
    exit(EXIT_FAILURE);
  }
  SSL_CTX_set_cipher_list(ctx, "ECDHE-RSA-AES256-GCM-SHA384");
  SSL_CTX_set_ecdh_auto(ctx, 1);

  port = atoi(argv[1]);
  { // setup addr
    struct sockaddr_in addr;

    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    addr.sin_addr.s_addr = htonl(INADDR_ANY);

    sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock < 0) {
      perror("Unable to create socket");
      exit(EXIT_FAILURE);
    }
    if (bind(sock, (struct sockaddr*)&addr, sizeof(addr)) < 0) {
      perror("Unable to bind");
      exit(EXIT_FAILURE);
    }
  }
  if (listen(sock, 1) < 0) {
    perror("Unable to listen");
    exit(EXIT_FAILURE);
  }

  /* Handle connections */
  printf("Server is listening on %d\n", port);
  while(1) {
    int con;
    struct sockaddr_in peer_addr;
    uint size = sizeof(peer_addr);
    SSL *ssl;

    // accept(tcp)
    con = accept(sock, (struct sockaddr*)&peer_addr, &size);
    if (con < 0) {
      perror("Unable to accept");
      exit(EXIT_FAILURE);
    }
    // create ssl session
    ssl = SSL_new(ctx);
    SSL_set_fd(ssl, con);
    if (SSL_accept(ssl) <= 0) {
      ERR_print_errors_fp(stderr);
      close(con);
      continue;
    }
    fprintf(stderr, "Connection 1 established\n");
    while (1) { // send/recv
      int len, err;
      len = SSL_read(ssl, buf, BUFFER_SIZE);
      if (!len) break;
      if (len < 0) {
        fprintf(stderr, "error SSL_read");
        break;
      }
      err = SSL_write(ssl, buf, len);
      if (err < 0) {
        fprintf(stderr, "error SSL_write");
        break;
      }
    }
    SSL_free(ssl);
    close(con);
  }

  close(sock);
  SSL_CTX_free(ctx);
}
