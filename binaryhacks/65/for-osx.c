#include <mach-o/dyld.h>
#include <stdio.h>

// BINARY HAKS の#65 をOSXで試してみる

int main() {
  int num, i;
  num = _dyld_image_count();
  for (i = 0; i < num; i++) {
    printf("%08lx %s\n",
        _dyld_get_image_vmaddr_slide(i), _dyld_get_image_name(i));
  }
  return 0;
}
