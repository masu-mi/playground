
root_dir="$(dirname $0)/.."
cd $root_dir
echo "[INFO][PWD]: $root_dir"

# fetch standard wasm glue code
# wget https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.js
# fetch tinygo glue code
wget https://raw.githubusercontent.com/tinygo-org/tinygo/master/targets/wasm_exec.js
echo vim ./wasm_exec.js "# execute it in other terminal to add functions to env"
read -p "Hit enter: "
# fetch example html file with load wasm file
wget https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.html -O index.html
{
  # Buidl wasm with tinygo
  pushd ..
  tinygo build -o test.wasm -target wasm ./test-wasm/
  popd
}
mv ../test.wasm ./
echo vim ./index.html "# to add debug code"
read -p "Hit enter: "
echo python3 -m http.server 8080
