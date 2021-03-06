
root_dir="$(dirname $0)/.."
cd $root_dir
echo "[INFO][PWD]: $root_dir"

# fetch standard wasm glue code
# wget https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.js
# fetch tinygo glue code
if [ ! -f ./static/js/wasm_exec.js ]; then
  wget https://raw.githubusercontent.com/tinygo-org/tinygo/master/targets/wasm_exec.js ./static/js/
  echo vim ./static/js/wasm_exec.js "# execute it in other terminal to add functions to env"
else
  echo ./static/js/wasm_exec.js exists
fi
read -p "Hit enter: "

# fetch example html file with load wasm file
if [ ! -f index.html ]; then
  wget https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.html -O index.html
  echo vim index.html "# to load main.wasm and use /static/js/wasm_exec.js"
else
  echo ./index.html exists
fi
read -p "Hit enter: "
tinygo build -o main.wasm -target wasm ./main/

read -p "Hit enter: "
python3 -m http.server 8080
