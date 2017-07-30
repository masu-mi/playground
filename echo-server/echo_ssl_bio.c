#include <stdio.h>
#include <openssl/bio.h>
#include <openssl/ssl.h>
#include <openssl/err.h>

#define BUFFER_SIZE 1024

#define CERT_FILE_PATH "files/cert.pem"
#define KEY_FILE_PATH  "files/key.pem"

int main(int argc, char **argv)
{
  char buf[BUFFER_SIZE];

  SSL_CTX *ctx;
  const SSL_METHOD *method;
  BIO *abio, *cbio;

  if (argc < 2) {
    fprintf(stderr, "Usage: %s [port]\n", argv[0]);
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

  // create BIO(high level API); for accept
  abio = BIO_new_accept(argv[1]);
  if(BIO_do_accept(abio) <= 0) {
    // first BIO_do_accept() is as to init BIO.
    fprintf(stderr, "Error setting up accept\n");
    ERR_print_errors_fp(stderr);
    exit(0);
  }
  // configuration for I/O operation over ssl/tls.
  BIO_set_accept_bios(abio, BIO_new_ssl(ctx, 0));

  // listen loop
  printf("Server is listening on %d\n", atoi(argv[1]));
  while (1) {
    if(BIO_do_accept(abio) <= 0) {
      fprintf(stderr, "Error accepting connection\n");
      ERR_print_errors_fp(stderr);
      exit(0);
    }
    fprintf(stderr, "Connection 1 established\n");
    /* Retrieve BIO for connection */
    cbio = BIO_pop(abio);
    while (1) {
      int len, err;
      len = BIO_read(cbio, buf, BUFFER_SIZE);
      if (!len) break;
      if (len < 0) {
        fprintf(stderr, "error SSL_read");
        break;
      }
      err = BIO_write(cbio, buf, len);
      if (err < 0) {
        fprintf(stderr, "error SSL_write");
        break;
      }
    }
    BIO_free(cbio);
  }
}
