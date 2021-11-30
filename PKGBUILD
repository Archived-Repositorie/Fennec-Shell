# Maintainer: JustFoxxo <justafoxxo@outlook.com>
pkgname=fennec-shell
pkgver=1.0.r3.637588f
pkgrel=1
pkgdesc="A Shell written in Go"
arch=(x86_64)
url="https://github.com/JustFoxx/fennec-shell.git"
license=('GPL')
depends=(go)
makedepends=(go)
optdepends=(nodejs python lua bash php php-cgi ruby perl powershell)
source=("git+$url")
md5sums=('SKIP')


pkgver() {
  cd "${_pkgname}"
  printf "1.0.r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"
}

build() {
    cd "${srcdir}/fennec-shell"
    mkdir out
    go build -o out/
}

package() {
    mkdir -p $pkgdir/usr/bin
    chmod +x ${srcdir}/fennec-shell/out/fs
    cp "${srcdir}/fennec-shell/out/fs" "$pkgdir/usr/bin"
}
